package notification

type Type string

const (
	BookingCreated   Type = "booking_created"
	BookingCancelled Type = "booking_cancelled"
	BookingCompleted Type = "booking_completed"
)

type Event struct {
	Type Type
	Data any
}
