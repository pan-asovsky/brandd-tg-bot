package cache

type ServiceTypeCache interface {
	Toggle(chatID int64, clickedService string) (map[string]bool, error)
	Clean(chatID int64) error
}
