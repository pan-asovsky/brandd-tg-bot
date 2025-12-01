package utils

import (
	"log"
	"regexp"
)

func ParseZoneCallback(cd string) (zone string, date string) {
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

func ParseTimeCallback(cd string) (time string, zone string, date string) {
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

func ParseServiceCallback(callback string) (pfx string, time string, date string) {
	re := regexp.MustCompile(`^([A-Z_]+):(\d{2}:\d{2}-\d{2}:\d{2}):(\d{4}-\d{2}-\d{2})$`)

	matches := re.FindStringSubmatch(callback)
	if matches == nil {
		log.Printf("[parse_callback] invalid callback format: %s", callback)
		return
	}

	pfx = matches[1]
	time = matches[2]
	date = matches[3]
	//log.Printf("[parse_callback] pfx: %s, time: %s, date: %s", pfx, time, date)

	return pfx, time, date
}
