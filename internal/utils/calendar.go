package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func GenerateICS(startDate, endDate time.Time) string {
	var buffer bytes.Buffer

	buffer.WriteString("BEGIN:VCALENDAR\r\n")
	buffer.WriteString("VERSION:2.0\r\n")
	buffer.WriteString("PRODID:-//Bandd//Calendar Event//EN\r\n")
	buffer.WriteString("CALSCALE:GREGORIAN\r\n")

	buffer.WriteString("BEGIN:VEVENT\r\n")
	buffer.WriteString(fmt.Sprintf("UID:%s\r\n", generateUID()))
	buffer.WriteString(fmt.Sprintf("DTSTAMP:%s\r\n", startDate.UTC().Format("20060102T150405Z")))

	buffer.WriteString(fmt.Sprintf("DTSTART:%s\r\n", startDate.UTC().Format("20060102T150405Z")))
	buffer.WriteString(fmt.Sprintf("DTEND:%s\r\n", endDate.UTC().Format("20060102T150405Z")))

	buffer.WriteString(fmt.Sprintf("SUMMARY:%s\r\n", "Шиномонтаж"))
	buffer.WriteString(fmt.Sprintf("DESCRIPTION:%s\r\n", "Запись на шиномонтаж в Bandd"))
	buffer.WriteString(fmt.Sprintf("LOCATION:%s\r\n", "Bandd, Казань, ГК Халева 1, 21"))

	buffer.WriteString("END:VEVENT\r\n")
	buffer.WriteString("END:VCALENDAR\r\n")

	return buffer.String()
}

func ParseDateTimeInMSKZone(dateStr, timeStr string) (time.Time, error) {
	mscZone := time.FixedZone("MSK", 3*60*60)
	return time.ParseInLocation("2006-01-02 15:04", dateStr+" "+timeStr, mscZone)
}

func ParseDate(dateStr string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Now(), WrapError(err)
	}
	return date, nil
}

func generateUID() string {
	return fmt.Sprintf("%s-%d@bandd.com",
		time.Now().Format("20060102150405"),
		rand.Intn(917332))
}
