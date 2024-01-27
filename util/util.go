package util

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func GetCurrentMonth() string {
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

func FormatAmount(amount float64) string {
	rounded := fmt.Sprintf("%.2f", amount)

	if strings.HasSuffix(rounded, ".00") {
		return strings.TrimSuffix(rounded, ".00")
	}

	if strings.HasSuffix(rounded, ".0") {
		return strings.TrimSuffix(rounded, ".0")
	}

	return rounded
}

func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}

	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Includes[T comparable](arr []T, item T) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}

	return false
}
