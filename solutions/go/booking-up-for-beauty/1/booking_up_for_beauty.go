package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout:= "1/2/2006 15:04:05"

    t, _ := time.Parse(layout, date)

    // if err != nil {
    //     fmt.Println(err)
    //     }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
    layout:= "January 2, 2006 15:04:05"
    parsed, _ := time.Parse(layout, date)
	
    return parsed.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout:= "Monday, January 2, 2006 15:04:05"
    parsed, _ := time.Parse(layout, date)
    hour := parsed.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout:= "1/2/2006 15:04:05"
    parsed, _ := time.Parse(layout, date)

    return parsed.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "01/02/2006"
    currentYear := time.Now().Year()
    anniversaryString := fmt.Sprintf("09/15/%d", currentYear)
    parsed, _ := time.Parse(layout, anniversaryString)
    return parsed
}
