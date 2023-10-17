package utils

import "time"

const DATE = "2006-01-02"

func ParseInLocation(timeFormat, timeStr string) (time.Time, error) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(timeFormat, timeStr, location)
}
