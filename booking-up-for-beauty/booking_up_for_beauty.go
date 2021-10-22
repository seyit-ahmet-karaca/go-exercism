package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	mTime, _ := time.Parse("1/02/2006 15:04:05", date)

	return mTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	mTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return mTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	mTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return mTime.Hour() >= 12 && mTime.Hour() < 18
}
// Description returns a formatted string of the appointment time
func Description(date string) string {
	mTime, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s", mTime.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	mTime, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return mTime
}
