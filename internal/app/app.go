package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/bot"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/config"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/provider"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Context  context.Context
	Config   *config.Config
	Cache    *redis.Client
	Postgres *sql.DB

	RepoProvider     iprovider.RepoProvider
	ServiceProvider  iprovider.ServiceProvider
	CacheProvider    iprovider.CacheProvider
	TelegramProvider iprovider.TelegramProvider
	CallbackProvider iprovider.CallbackProvider
	MsgFmtProvider   iprovider.MessageFormatterProvider

	BotAPI        *tgapi.BotAPI
	UpdateHandler ihandler.UpdateHandler
	httpServer    *http.Server
}

func NewApp(ctx context.Context) *App {
	return &App{Context: ctx}
}

func (a *App) Init() error {
	// configuration load
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	a.Config = cfg

	// cache client
	redisClient, err := cache.NewRedis(cfg)
	if err != nil {
		return err
	}
	a.Cache = redisClient

	// postgres client
	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		return err
	}
	a.Postgres = db

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode,
	)
	postgres.RunMigrations(dbURL)

	tgbot, err := bot.NewTelegramBot(a.Config.BotToken, a.Config.WebhookURL)
	if err != nil {
		return err
	}
	a.BotAPI = tgbot

	// provider
	a.CallbackProvider = provider.NewCallbackProvider()
	a.RepoProvider = provider.NewRepoProvider(a.Postgres)
	a.CacheProvider = provider.NewCacheProvider(a.Cache, a.Config.CacheTTL)
	a.ServiceProvider = provider.NewServiceProvider(a.RepoProvider, a.CacheProvider, a.CallbackProvider)
	a.MsgFmtProvider = provider.NewMessageFormatterProvider(a.ServiceProvider.DateTime())
	a.TelegramProvider = provider.NewTelegramProvider(a.BotAPI, a.ServiceProvider, a.MsgFmtProvider)

	// handler
	a.UpdateHandler = handler.NewUpdateHandler(a.BotAPI, a.ServiceProvider, a.RepoProvider, a.CacheProvider, a.CallbackProvider, a.TelegramProvider)

	return nil
}

func (a *App) Run() error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	_ = router.SetTrustedProxies([]string{"127.0.0.1"})

	router.POST(a.Config.WebhookPath, func(c *gin.Context) {
		var update tgapi.Update
		if err := c.ShouldBindJSON(&update); err != nil {
			c.Status(400)
			log.Printf("Invalid update payload: %v", err)
			return
		}
		if err := a.UpdateHandler.Handle(&update); err != nil {
			log.Printf("[app] error handling update: %v", err)
		}
		c.Status(200)
	})

	a.httpServer = &http.Server{
		Addr:    a.Config.HttpAddress,
		Handler: router,
	}
	log.Printf("Started webhook server on %s, path %s", a.Config.HttpAddress, a.Config.WebhookPath)

	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

func (a *App) Close() {
	if a.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := a.httpServer.Shutdown(ctx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
	}

	if a.Postgres != nil {
		err := a.Postgres.Close()
		if err != nil {
			log.Printf("Error closing postgres connection: %s", err)
		}
	}
	if a.Cache != nil {
		err := a.Cache.Close()
		if err != nil {
			log.Printf("Error closing cache connection: %s", err)
		}
	}
}
