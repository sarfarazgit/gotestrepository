package servicepackage

import "time"

func CurrentDate() string {
	current_date := time.Now().Local()
	return current_date.Format("2-01-2006")
}

func CurrenTime() string {
	current_date := time.Now().Local()
	return current_date.Format("15:04:05")
}
