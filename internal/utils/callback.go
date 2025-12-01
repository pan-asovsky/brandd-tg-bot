package utils

import (
	"log"
	"regexp"
	"time"
)

func ParseZoneCallback(cd string) (zone, date string) {
	re := regexp.MustCompile(`^([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}):([0-9]{4}-[0-9]{2}-[0-9]{2})$`)

	matches := re.FindStringSubmatch(cd)
	if len(matches) != 3 {
		log.Printf("[parse_callback] invalid callback format: %s", cd)
		return
	}

	zone = matches[1]
	date = matches[2]
	//log.Printf("[parse_callback] zone: %s, date: %s", zone, date)

	return zone, date
}

func ParseTimeCallback(cd string) (time, zone, date string) {
	re := regexp.MustCompile(`^(\d{2}:\d{2}-\d{2}:\d{2}):(\d{2}:\d{2}-\d{2}:\d{2}):(\d{4}-\d{2}-\d{2})$`)

	matches := re.FindStringSubmatch(cd)
	if matches == nil {
		log.Printf("[parse_callback] invalid callback format: %s", cd)
		return
	}

	time = matches[1]
	zone = matches[2]
	date = matches[3]
	//log.Printf("[parse_callback] time: %s, zone: %s, date: %s", time, zone, date)

	return time, zone, date
}

func ParseServiceCallback(callback string) (svc, time, date string) {
	re := regexp.MustCompile(`^([A-Z_]+):(\d{2}:\d{2}-\d{2}:\d{2}):(\d{4}-\d{2}-\d{2})$`)

	matches := re.FindStringSubmatch(callback)
	if matches == nil {
		log.Printf("[parse_callback] invalid callback format: %s", callback)
		return
	}

	svc = matches[1]
	time = matches[2]
	date = matches[3]
	//log.Printf("[parse_callback] svc: %s, time: %s, date: %s", svc, time, date)

	return svc, time, date
}

func ParseRimCallback(callback string) (r, svc, time, date string) {
	re := regexp.MustCompile(`^(\d+):([A-Z_]+):(\d{2}:\d{2}-\d{2}:\d{2}):(\d{4}-\d{2}-\d{2})$`)

	matches := re.FindStringSubmatch(callback)
	if matches == nil {
		log.Printf("[parse_callback] invalid callback format: %s", callback)
		return
	}

	r = matches[1]
	svc = matches[2]
	time = matches[3]
	date = matches[4]
	log.Printf("[parse_callback] r: %s, svc: %s, time: %s, date: %s", r, svc, time, date)

	return r, svc, time, reformatDate(date)
}

func reformatDate(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("[reformat_date] failed: %s", date)
		return ""
	}
	return t.Format("02.01.2006")
}
