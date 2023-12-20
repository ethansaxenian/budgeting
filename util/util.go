package util

import "time"

func GetCurrentDate() string {
	d := time.Now()

	return d.Format("2006-01-02")
}
