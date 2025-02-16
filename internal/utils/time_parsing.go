package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseStringToDuration(durationString string) (time.Duration, bool) {

	if strings.HasSuffix(durationString, "d") {
		daysStr := strings.TrimSuffix(durationString, "d")
		days, err := strconv.Atoi(daysStr)

		if err != nil {
			return 0, false
		}

		// return number of days * 24 hours * duration of an hour
		return time.Duration(days*24) * time.Hour, true
	}

	duration, err := time.ParseDuration(durationString)

	if err != nil {
		return 0, false
	}

	return duration, true
}
