package util

import (
	"fmt"
	"strings"
	"unicode"
)

func FormatAmount(amount float64) string {
	rounded := fmt.Sprintf("%.2f", amount)

	if before, ok := strings.CutSuffix(rounded, ".00"); ok {
		return before
	}

	if before, ok := strings.CutSuffix(rounded, ".0"); ok {
		return before
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
