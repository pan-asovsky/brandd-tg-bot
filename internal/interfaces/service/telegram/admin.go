package tg

type TelegramAdminService interface {
	ChoiceMenu(chatID int64) error
	StartMenu(chatID int64) error
}
