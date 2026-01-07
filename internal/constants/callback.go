package constants

const (
	NewBookingCbk       = PrefixBooking + New
	MyBookingsCbk       = PrefixBooking + My
	PreCancelBookingCbk = PrefixBooking + PreCancel
	CancelBookingCbk    = PrefixBooking + Cancel
	NoCancelBookingCbk  = PrefixBooking + NoCancel

	HelpCbk     = PrefixMenu + Help
	CalendarCbk = PrefixMenu + Calendar

	ConfirmBookingCbk = PrefixConfirm + Yes
	RejectCbk         = PrefixConfirm + No
)
