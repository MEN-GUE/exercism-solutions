package booking
import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"

    t, err := time.Parse(layout, date);
    if (err != nil) {
        fmt.Println("Error: ", err)
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date);
	if (err != nil){
        fmt.Println("Error", err)
    }
    return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 1, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if (err != nil) {
        fmt.Println("Error", err)
    }

	hourOfAppointment := t.Hour()
    if (hourOfAppointment >= 12 && hourOfAppointment < 18){
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "2/1/2006 15:04:00"
    t, err := time.Parse(layout, date)
    if(err != nil){
        layout = "1/2/2006 15:04:00"
        t, err = time.Parse(layout, date)
        return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04.")
        fmt.Println("Error", err)
    }

    return "You have an appointment on " + t.Format("Monday, January 1, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    openingDate := time.Date(2000, 9, 15, 12, 00, 0, 0, time.UTC)
    currentDate := time.Now()
    return time.Date(currentDate.Year(), openingDate.Month(), openingDate.Day(), 00, 00, 0, 0, time.UTC)
}
