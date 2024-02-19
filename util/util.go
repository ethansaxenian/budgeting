package util

import (
	"fmt"
	"strings"
	"unicode"
)

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

func FormatAmountWithDollarSign(amount float64) string {
	str := fmt.Sprintf("%.2f", amount)

	str = strings.TrimSuffix(str, ".00")

	str = strings.TrimSuffix(str, ".0")

	if !strings.Contains(str, "-") {
		return "$" + str
	}

	return "-" + "$" + strings.TrimLeft(str, "-")
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
