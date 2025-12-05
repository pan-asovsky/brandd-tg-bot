package utils

import (
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	t "github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

func FmtRimMsg(data t.UserSessionInfo) string {
	return fmt.Sprintf(consts.PreConfirmMsg, data.Date, consts.Time[data.Time], data.Time, consts.ServiceNames[data.Service], data.Radius)
}
