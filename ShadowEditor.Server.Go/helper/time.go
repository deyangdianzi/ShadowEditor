package helper

import (
	"strconv"
	"strings"
	"time"
)

// TimeToString convert a time to string.
func TimeToString(time time.Time, format string) string {
	year := strconv.Itoa(time.Year())
	month := strconv.Itoa(int(time.Month()))
	day := strconv.Itoa(time.Day())

	if len(month) < 2 {
		month = "0" + month
	}
	if len(day) < 2 {
		day = "0" + day
	}

	hour := strconv.Itoa(time.Hour())
	minute := strconv.Itoa(time.Minute())
	second := strconv.Itoa(time.Second())

	if len(hour) < 2 {
		hour = "0" + hour
	}

	if len(minute) < 2 {
		minute = "0" + minute
	}

	if len(second) < 2 {
		second = "0" + second
	}

	result := strings.ReplaceAll(format, "yyyy", year)
	result = strings.ReplaceAll(result, "MM", month)
	result = strings.ReplaceAll(result, "dd", day)
	result = strings.ReplaceAll(result, "HH", hour)
	result = strings.ReplaceAll(result, "mm", minute)
	result = strings.ReplaceAll(result, "ss", second)

	return result
}
