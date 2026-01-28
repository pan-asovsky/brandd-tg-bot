package notification

type Channel interface {
	Send(chatID int64, msg string) error
}
