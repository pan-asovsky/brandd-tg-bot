package user_flow

const (
	NewBookingCbk       = UserPrefix + PrefixBooking + New
	MyBookingsCbk       = UserPrefix + PrefixBooking + My
	PreCancelBookingCbk = UserPrefix + PrefixBooking + PreCancel
	CancelBookingCbk    = UserPrefix + PrefixBooking + Cancel
	NoCancelBookingCbk  = UserPrefix + PrefixBooking + NoCancel

	ConfirmBookingCbk = UserPrefix + PrefixConfirm + Yes
	RejectCbk         = UserPrefix + PrefixConfirm + No
)
