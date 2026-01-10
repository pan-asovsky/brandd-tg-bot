package service

type SlotLocking interface {
	Lock(date, time string) (uid string, ok bool, err error)
	Unlock(date, uid string) error
	IsLocked(date, time string) (bool, error)
	RefreshTTL(key string) error
	AreLocked(keys ...string) (map[string]bool, error)
	FormatKey(date, time string) string
}
