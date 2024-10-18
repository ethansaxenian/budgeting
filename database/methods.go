package database

import (
	"fmt"
	"time"
)

func (m Month) FormatStr() string {
	return fmt.Sprintf("%d-%02d", m.Year, m.Month)
}

func (m Month) Date() (time.Time, error) {
	return time.Parse("2006-01", m.FormatStr())
}
