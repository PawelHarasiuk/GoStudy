package helpers

import "time"

var dateFormat = "02.01.2006 15:04:05"

func StringToDate(stringDate string) (time.Time, error) {
	birthday, err := time.Parse(dateFormat, stringDate)
	return birthday, err
}

func DateToString(date time.Time) string {
	return time.Time.Format(date, dateFormat)
}
