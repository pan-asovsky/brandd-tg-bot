package constant

import (
	stat "github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
)

var ServiceNames = map[string]string{
	"TAKE_IT_OUT":        "Съём-установка",
	"BALANCING":          "Балансировка",
	"TIRE_SERVICE":       "Шиномонтаж",
	"COMPLEX":            "Комплекс",
	"TAKE_AND_TIRE":      "Съём-установка и шиномонтаж",
	"TAKE_AND_BALANCING": "Съём-установка и балансировка",
	"TIRE_AND_BALANCING": "Шиномонтаж и балансировка",
}

var Time = map[string]string{
	"09:00": "🕘",
	"10:00": "🕙",
	"11:00": "🕚",
	"12:00": "🕛",
	"13:00": "🕐",
	"14:00": "🕑",
	"15:00": "🕒",
	"16:00": "🕓",
	"17:00": "🕔",
	"18:00": "🕕",
	"19:00": "🕖",
	"20:00": "🕗",
}

var PeriodLabels = map[stat.Label]string{
	stat.Today:     "сегодня",
	stat.Yesterday: "вчера",
	stat.Week:      "неделю",
	stat.Month:     "месяц",
}
