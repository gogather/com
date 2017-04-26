package time

import (
	"fmt"
	"time"
)

// Tonight get tonight, this day 00:00:00
func Tonight() time.Time {
	now := time.Now()
	date := fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day())
	pt, _ := time.Parse("2006-01-02", date)
	return pt
}

// Days get d days period
func Days(d int) (time.Duration, error) {
	return time.ParseDuration(fmt.Sprintf("%dh", d*24))
}
