package utility

import (
	"time"
)

func ParseStringToTime(stringDate, format string) (t time.Time) {
	t, _ = time.Parse(format, stringDate)
	return
}

func ParseTimeToString(t time.Time, format string) (stringDate string) {
	stringDate = t.Format(format)
	return
}

func ParseTimeCustomToString(t CustomTime, format string) (stringDate string) {
	stringDate = t.Format(format)
	return
}

func FormatDateDDMMYY(inputTime time.Time) string {
	// Use the Format method to format the time according to the "ddmmyy" layout
	return inputTime.Format("020106")
}

func GetTimeInUTCPlus7(timeParam time.Time) time.Time {

	// Add 7 hours to the current UTC time to get UTC+7
	utcPlus7 := timeParam.Add(7 * time.Hour)

	return utcPlus7
}

type CustomTime struct {
	time.Time
}

const (
	// Define possible time formats
	layoutDate         = "2006-01-02"
	layoutDateTime     = "2006-01-02T15:04:05"
	layoutDateTimeZone = "2006-01-02T15:04:05Z07:00"
)

// UnmarshalJSON handles different time formats
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Remove the surrounding quotes
	s := string(b)
	if s == "null" {
		*ct = CustomTime{Time: time.Time{}}
		return nil
	}
	s = s[1 : len(s)-1]

	// Try parsing in different formats
	var t time.Time
	var err error

	// Attempt to parse as full timestamp with timezone
	t, err = time.Parse(layoutDateTimeZone, s)
	if err != nil {
		// Attempt to parse as datetime without timezone
		t, err = time.Parse(layoutDateTime, s)
		if err != nil {
			// Attempt to parse as date only
			t, err = time.Parse(layoutDate, s)
			if err != nil {
				return err
			}
		}
	}

	ct.Time = t
	return nil
}
