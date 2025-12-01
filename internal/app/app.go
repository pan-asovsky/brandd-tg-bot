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
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/bot"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
	"github.com/pan-asovsky/brandd-tg-bot/internal/config"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler"
	"github.com/pan-asovsky/brandd-tg-bot/internal/postgres"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	rd "github.com/pan-asovsky/brandd-tg-bot/internal/repository/redis"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	kb "github.com/pan-asovsky/brandd-tg-bot/internal/service/keyboard"
	slot "github.com/pan-asovsky/brandd-tg-bot/internal/service/slot"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Context       context.Context
	Config        *config.Config
	Cache         *redis.Client
	Postgres      *sql.DB
	SessionRepo   *rd.SessionRepo
	PriceRepo     pg.PriceRepo
	BookingRepo   pg.BookingRepo
	SlotRepo      pg.SlotRepo
	ServiceRepo   pg.ServiceRepo
	TelegramBot   *tg.BotAPI
	Slot          slot.SlotService
	LockService   *service.LockService
	Keyboard      kb.KeyboardService
	UpdateHandler *handler.UpdateHandler
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

	sl, err := locker.NewSlotLocker(a.Cache, a.Config.SlotLockTTL)
	if err != nil {
		return err
	}
	a.LockService = service.NewLockService(sl, cache.NewLockCache(a.Cache, a.Config.SlotLockTTL))

	// repositories
	a.SessionRepo = rd.NewSessionRepo(a.Cache, a.Config.SessionTTL)
	a.BookingRepo = pg.NewBookingRepo(a.Postgres)
	a.SlotRepo = pg.NewSlotRepo(a.Postgres)
	a.ServiceRepo = pg.NewServiceRepo(a.Postgres)
	a.PriceRepo = pg.NewPriceRepo(a.Postgres)

	tgbot, err := bot.NewTelegramBot(a.Config.BotToken, a.Config.WebhookURL)
	if err != nil {
		return err
	}
	a.TelegramBot = tgbot

	// service + handler
	a.Keyboard = kb.NewKeyboard()
	a.Slot = slot.NewSlot(a.SlotRepo, sl)
	a.UpdateHandler = handler.NewUpdateHandler(a.TelegramBot, a.Keyboard, a.Slot, *a.LockService, a.ServiceRepo, a.PriceRepo)

	return nil
}

func (a *App) Run() error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	_ = router.SetTrustedProxies([]string{"127.0.0.1"})

	router.POST(a.Config.WebhookPath, func(c *gin.Context) {
		var update tg.Update
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
		return fmt.Errorf("server error: %v", err)
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
