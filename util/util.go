package util

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func GetCurrentDate() string {
	d := time.Now()

	return d.Format("2006-01-02")
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func FormatAmount(amount float64) string {
	rounded := fmt.Sprintf("%.2f", amount)

	return strings.TrimRight(rounded, ".0")
}

func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}

	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
