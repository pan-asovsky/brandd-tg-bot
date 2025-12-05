package service

type LockService interface {
	Toggle(chatID int64, date, time string) error
}
