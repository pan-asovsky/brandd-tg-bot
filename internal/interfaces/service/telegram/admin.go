package telegram

type TelegramAdminService interface {
	StartMenu(chatID int64) error
}
