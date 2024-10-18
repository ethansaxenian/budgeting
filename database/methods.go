package database

import (
	"fmt"
	"time"
)

func (m Month) FormatStr() string {
	return fmt.Sprintf("%d-%02d", m.Year, m.Month)
}

func (m Month) HasDate(date time.Time) bool {
	return m.Month == date.Month() && m.Year == date.Year()
}

func (m Month) StartEndDates() (time.Time, time.Time) {
	start := time.Date(m.Year, m.Month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, -1)
	return start, end
}

func (m Month) Date() (time.Time, error) {
	return time.Parse("2006-01", m.FormatStr())
}
