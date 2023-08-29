package utils

import "time"

func ParseDateString(d string) (time.Time, error) {
	date, err := time.Parse("2023-11-22 23:59:00", d)
	if err != nil {
		return time.Now(), nil
	}
	return date, nil
}
