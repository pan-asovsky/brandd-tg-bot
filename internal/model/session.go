package model

type State string

const (
	DateSelected        State = "DATE"
	ZoneSelected        State = "ZONE"
	TimeSelected        State = "TIME"
	ServiceTypeSelected State = "SERVICE_TYPE"
	RimSizeSelected     State = "RIM_SIZE"
	WheelCountSelected  State = "WHEEL_COUNT"
)

type Session struct {
	ChatID      int64  `json:"chat_id"`
	State       State  `json:"state"`
	Date        string `json:"date"`
	Zone        string `json:"zone"`
	Time        string `json:"time"`
	ServiceType string `json:"service_type"`
	RimSize     int64  `json:"rim_size"`
	WheelCount  int64  `json:"wheel_count"`
	TotalPrice  int64  `json:"total_price"`
}
