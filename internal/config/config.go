package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type Config struct {
	BotToken    string
	WebhookURL  string
	WebhookPath string
	HttpAddress string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string

	RedisURL      string
	RedisPassword string
	RedisDB       int
	CacheTTL      time.Duration

	LogLevel    string
	Environment string
}

func Load() (*Config, error) {
	cfg := &Config{
		BotToken:    utils.GetEnvRequired("BOT_TOKEN"),
		WebhookURL:  utils.GetEnvRequired("WEBHOOK_URL"),
		WebhookPath: utils.GetEnvRequired("WEBHOOK_PATH"),
		HttpAddress: utils.GetEnv("HTTP_ADDRESS", "0.0.0.0:8118"),

		DBHost:     utils.GetEnv("DB_HOST", "localhost"),
		DBPort:     utils.GetEnv("DB_PORT", "5432"),
		DBName:     utils.GetEnv("DB_NAME", "brandd"),
		DBUser:     utils.GetEnv("DB_USER", "brandd_user"),
		DBPassword: utils.GetEnvRequired("DB_PASSWORD"),
		DBSSLMode:  utils.GetEnv("DB_SSL_MODE", "disable"),

		RedisURL:      utils.GetEnv("REDIS_URL", "cache:6379"),
		RedisPassword: utils.GetEnvRequired("REDIS_PASSWORD"),
		RedisDB:       utils.GetEnv("REDIS_DB", 0),
		CacheTTL:      utils.GetEnv("CACHE_TTL", time.Duration(180)*time.Second),

		LogLevel:    utils.GetEnv("LOG_LEVEL", "info"),
		Environment: utils.GetEnv("ENVIRONMENT", "production"),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) validate() error {
	var missing []string

	if cfg.BotToken == "" {
		missing = append(missing, "BOT_TOKEN")
	}
	if cfg.DBPassword == "" {
		missing = append(missing, "DB_PASSWORD")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	return nil
}

func (cfg *Config) DBDsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)
}
