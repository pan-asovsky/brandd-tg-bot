package types

type UserSessionInfo struct {
	ChatID                            int64
	Date, Zone, Time, Radius, Service string
	SelectedServices                  map[string]bool
	TotalPrice                        int64
}
