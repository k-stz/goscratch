package main

import (
	"fmt"
	"time"
)

func main() {
	Schedule("7/25/2019 13:45:00")
	// => 2019-07-25 13:45:00 +0000 UTC
	fmt.Println("haspassed() => true", HasPassed("October 3, 2019 20:32:00"))
	// => true
	fmt.Println("IsAfternoon()=>must true; is=", IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	// => true)
	fmt.Println("Should return: You have an appointment on Thursday, July 25, 2019, at 13:45.")
	fmt.Println("Desciption():", Description("7/25/2019 13:45:00"))
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
	fmt.Println("Anniversary:", AnniversaryDate())

}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return 12 <= t.Hour() && t.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
