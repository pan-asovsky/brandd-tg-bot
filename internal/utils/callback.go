package utils

import (
	"log"
	"regexp"
)

func ParseCallback(cb string) (zone string, date string, error error) {
	re := regexp.MustCompile(`^([0-9]{2}:[0-9]{2}-[0-9]{2}:[0-9]{2}):([0-9]{4}-[0-9]{2}-[0-9]{2})$`)

	matches := re.FindStringSubmatch(cb)
	if len(matches) != 3 {
		log.Println("[parse_callback] invalid callback format")
		return
	}

	zone = matches[1]
	date = matches[2]
	log.Printf("[parse_callback] zone: %s, date: %s", zone, date)

	return zone, date, nil
}
