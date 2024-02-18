package util

import "time"

func GetCurrentMonthStr() string {
	d := time.Now()

	return d.Format("2006-01")
}

func GetDateMonth(date time.Time) string {
	return date.Format("2006-01")
}

func GetCurrentDate() string {
	d := time.Now()

	return d.Format(time.DateOnly)
}

func FormatDate(date time.Time) string {
	return date.Format(time.DateOnly)
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse(time.DateOnly, date)
}
