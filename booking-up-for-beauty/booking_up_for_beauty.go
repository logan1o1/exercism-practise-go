package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"

	t, _ := time.ParseInLocation(layout, date, time.UTC)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		t, err = time.Parse("2/01/2006 15:04:05", date)
		if err != nil {
			t, err = time.Parse("Monday, 2/01/2006 15:04:05", date)
			if err != nil {
				return false // Or handle the error as appropriate
			}
		}
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	dt, _ := time.Parse(layout, date)
	return dt.Hour() >= 12 && dt.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dt := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", dt.Weekday(), dt.Month().String(), dt.Day(), dt.Year(), dt.Hour(), dt.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	curtime := time.Now().UTC()
	curYr := curtime.Year()
	return time.Date(curYr, time.September, 15, 0, 0, 0, 0, time.UTC)
}
