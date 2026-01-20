package admin_flow

import consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"

const (
	BookingsCbk   = AdminPrefix + MenuPrefix + Bookings
	StatisticsCbk = AdminPrefix + MenuPrefix + Statistics
	SettingsCbk   = AdminPrefix + MenuPrefix + Settings
	UserFlowCbk   = FlowPrefix + consts.USER
	AdminFlowCbk  = FlowPrefix + consts.ADMIN
)
