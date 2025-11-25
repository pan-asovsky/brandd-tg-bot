package bot

import (
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewTelegramBot(token string, webhookURL string) (*tg.BotAPI, error) {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create telegram bot: %w", err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	if err = setupWebhook(bot, webhookURL); err != nil {
		return nil, fmt.Errorf("failed to setup webhook: %w", err)
	}

	return bot, nil
}

func setupWebhook(api *tg.BotAPI, webhookURL string) error {
	whInfo, _ := api.GetWebhookInfo()
	if whInfo.URL == webhookURL {
		return nil
	}

	webhook, err := tg.NewWebhook(webhookURL)
	if err != nil {
		return fmt.Errorf("failed to create webhook: %w", err)
	}

	if _, err := api.Request(webhook); err != nil {
		return fmt.Errorf("failed to set webhook: %w", err)
	}

	log.Println("Webhook set:", webhookURL)
	return nil
}
