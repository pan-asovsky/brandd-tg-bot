package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/bot"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/config"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler"
	"github.com/pan-asovsky/brandd-tg-bot/internal/infrastructure/db"
	"github.com/pan-asovsky/brandd-tg-bot/internal/infrastructure/httpsrv"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	"github.com/pan-asovsky/brandd-tg-bot/internal/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Context context.Context
	Config  *config.Config
	Cache   *redis.Client
	SqlDB   *sql.DB

	ProviderContainer provider.Container

	BotAPI        *tgapi.BotAPI
	UpdateHandler ihandler.UpdateHandler
	httpServer    *httpsrv.Server
	//httpServer    *http.Server
}

func NewApp(ctx context.Context) *App {
	return &App{Context: ctx}
}

func (a *App) Init() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	a.Config = cfg

	redisClient, err := cache.NewRedis(cfg)
	if err != nil {
		return err
	}
	a.Cache = redisClient

	conn, err := db.NewDBConn(a.Config, a.Context)
	if err != nil {
		return utils.WrapError(err)
	}
	a.SqlDB = conn

	tgbot, err := bot.NewTelegramBot(a.Config.BotToken, a.Config.WebhookURL)
	if err != nil {
		return err
	}
	a.BotAPI = tgbot

	callback := provider.NewCallbackProvider()
	repo := provider.NewRepoProvider(a.SqlDB)
	cachep := provider.NewCacheProvider(a.Cache, a.Config.CacheTTL)
	service := provider.NewServiceProvider(repo, cachep, callback)
	formatter := provider.NewMessageFormatterProvider(service.DateTime())
	keyboard := provider.NewKeyboardProvider(service.DateTime(), callback)
	telegram := provider.NewTelegramProvider(a.BotAPI, service, keyboard, formatter)
	notification := provider.NewNotificationProvider(service.User(), telegram.Common(), formatter)
	statistics := provider.NewStatisticsProvider(repo)

	a.ProviderContainer = *provider.NewContainer(
		repo, service, cachep,
		telegram, callback, formatter,
		keyboard, notification, statistics,
	)
	a.UpdateHandler = handler.NewUpdateHandler(a.ProviderContainer)

	return nil
}

func (a *App) Run() error {
	wh := httpsrv.NewWebhookHandler(a.UpdateHandler, a.Config.WebhookPath)
	a.httpServer = httpsrv.NewServer(a.Config.HttpAddress, wh.Handler())
	log.Printf("Started webhook server on %s, path %s", a.Config.HttpAddress, a.Config.WebhookPath)

	if err := a.httpServer.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	if a.SqlDB != nil {
		err := a.SqlDB.Close()
		if err != nil {
			log.Printf("Error closing db connection: %s", err)
		}
	}
	if a.Cache != nil {
		err := a.Cache.Close()
		if err != nil {
			log.Printf("Error closing cache connection: %s", err)
		}
	}
}
