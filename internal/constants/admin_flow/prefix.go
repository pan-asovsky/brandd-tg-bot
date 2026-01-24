package admin_flow

const (
	AdminBackPrefix = AdminPrefix + "BACK::"
	PrefixBack      = "BACK::"
	AdminPrefix     = "ADMIN::"
	FlowPrefix      = "FLOW::"
	MenuPrefix      = "MENU::"

	PrefixBooking    = "BOOKING::"
	PrefixStatistics = "STAT::"
	PrefixSettings   = "STNG::"
	PrefixReject     = "REJECT:"

	PrefixComplete           = "CMP:"
	PrefixPreCompleteBooking = PrefixComplete + "1:"
	PrefixCompleteBooking    = PrefixComplete + "2:"

	PrefixNoShow           = "NS:"
	PrefixPreNoShowBooking = PrefixNoShow + "1:"
	PrefixNoShowBooking    = PrefixNoShow + "2:"
)
