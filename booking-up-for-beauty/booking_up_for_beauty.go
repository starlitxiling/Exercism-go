package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// panic("Please implement the Schedule function")
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("parsing error: ", err)
		return time.Time{}
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// panic("Please implement the HasPassed function")
	layout := "January 2, 2006 15:04:05"
	// t, err := time.Parse(layout, date)
	t, _ := time.Parse(layout, date)
	// if err != nil {
	// 	fmt.Println("parsing error: ", err)
	// 	return false
	// }
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// panic("Please implement the IsAfternoonAppointment function")
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// panic("Please implement the Description function")
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %02d:%02d.",
		t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// panic("Please implement the AnniversaryDate function")
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
