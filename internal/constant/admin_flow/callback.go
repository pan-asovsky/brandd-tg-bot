package admin_flow

import consts "github.com/pan-asovsky/brandd-tg-bot/internal/constant"

const (
	BookingsCbk     = AdminPrefix + MenuPrefix + Bookings
	StatisticsCbk   = AdminPrefix + MenuPrefix + Statistics
	SettingsCbk     = AdminPrefix + MenuPrefix + Settings
	UserFlowCbk     = FlowPrefix + consts.USER
	AdminFlowCbk    = FlowPrefix + consts.ADMIN
	RejectActionCbk = AdminPrefix + PrefixReject
)
