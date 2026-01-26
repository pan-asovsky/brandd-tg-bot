package model

type SlotLockInfo struct {
	Key  string
	UUID string
}

type UserSessionInfo struct {
	ChatID                               int64
	Date, Zone, Time, RimRadius, Service string
	SelectedServices                     map[string]bool
	TotalPrice                           int64
}
