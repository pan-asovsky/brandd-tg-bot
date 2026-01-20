package callback

type AdminCallbackBuilderService interface {
	StartUser() string
	StartAdmin() string
	Bookings() string
	Statistics() string
	Settings() string
	Back(direction string) string
}
