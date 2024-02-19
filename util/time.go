package util

import "time"

func CurrentYear() int {
	return time.Now().Year()
}

func CurrentMonth() time.Month {
	return time.Now().Month()
}

func CurrentDate() (int, time.Month, int) {
	return time.Now().Date()
}

func FormatDate(date time.Time) string {
	return date.Format(time.DateOnly)
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse(time.DateOnly, date)
}
