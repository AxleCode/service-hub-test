package utility

import "time"

func DateStringToDateOnly(date string) (time.Time, error) {
	return time.Parse(time.DateOnly, date)
}

func ParseStringDateToRFC3339(date string) time.Time {
	rDate, _ := time.Parse(time.RFC3339, date+"T00:00:00Z")
	return rDate
}

func ParseStringDateRFC3339ToTime(date string) time.Time {
	rDate, _ := time.Parse(time.RFC3339, date)
	return rDate
}

func ParseStringDatetimeToRFC3339(date, times string) time.Time {
	rDate, _ := time.Parse(time.RFC3339, date+"T"+times+"Z")
	return rDate
}

func ParseStringTimeToRFC3339(times string) time.Time {
	date := time.Now().Format("2006-01-02")
	rDate, _ := time.Parse(time.RFC3339, date+"T"+times+"Z")
	return rDate
}

func ParseMonthToFirstDayOfMonth(month, year string) string {

	if len(month) < 1 {
		month = "0" + month
	}

	return year + "-" + month + "-01"
}

func ParseMonthToLastDayOfMonth(month, year string) string {
	firstDate := ParseMonthToFirstDayOfMonth(month, year)

	t := ParseStringToTime(firstDate, "2006-01-02")
	y, m, _ := t.Date()
	lastday := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC)

	return lastday.Format("2006-01-02")
}

func ParseStringStartEndDate(startDate, endDate string) (time.Time, time.Time) {
	start, _ := time.Parse(time.RFC3339, startDate+"T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, endDate+"T23:59:59Z")

	return start, end
}

func GetNowFirstLastDayOfMonth() (firstDay, lastDay string) {
	t, _ := time.Parse("2006-01-02", time.Now().UTC().Format("2006-01-02"))
	firstDay = t.Format("2006-01-02")

	y, m, _ := t.Date()
	lastday := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC)
	lastDay = lastday.Format("2006-01-02")

	return
}
