package utils

import "time"

const (
	fpDate      = "2006-01-02"
	fpTimestamp = "2006-01-02T00:00:00Z"
	fpTime      = "15:04"
)

func GetFormattedDate(t time.Time) string {
	return t.Format(fpDate)
}

func GetFormattedTime(t time.Time) string {
	return t.Format(fpTime)
}

func ParseDateFromString(date string) (time.Time, error) {
	return time.Parse(fpDate, date)
}

func ParseDateFromTimestampString(date string) (time.Time, error) {
	return time.Parse(fpTimestamp, date)
}
