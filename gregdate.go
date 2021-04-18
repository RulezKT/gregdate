package gregdate

import "fmt"

type GregDate struct {
	Year    int `json:"Year"`
	Month   int `json:"Month"`
	Day     int `json:"Day"`
	Hour    int
	Minutes int
	Seconds int
}

func (gd GregDate) String() string {

	var monthsFormatted string
	if gd.Month > 9 {
		monthsFormatted = "" + fmt.Sprint(gd.Month)

	} else {
		monthsFormatted = "0" + fmt.Sprint(gd.Month)
	}

	var daysFormatted string
	if gd.Day > 9 {
		daysFormatted = "" + fmt.Sprint(gd.Day)

	} else {
		daysFormatted = "0" + fmt.Sprint(gd.Day)
	}

	var hoursFormatted string
	if gd.Hour > 9 {
		hoursFormatted = "" + fmt.Sprint(gd.Hour)

	} else {
		hoursFormatted = "0" + fmt.Sprint(gd.Hour)
	}

	var minutesFormatted string
	if gd.Minutes > 9 {
		minutesFormatted = "" + fmt.Sprint(gd.Minutes)

	} else {
		minutesFormatted = "0" + fmt.Sprint(gd.Minutes)
	}

	var secondsFormatted string
	if gd.Seconds > 9 {
		secondsFormatted = "" + fmt.Sprint(gd.Seconds)

	} else {
		secondsFormatted = "0" + fmt.Sprint(gd.Seconds)
	}

	return "Date:  " + daysFormatted + "." + monthsFormatted + "." + fmt.Sprint(gd.Year) +
		"  Time: " + hoursFormatted + ":" + minutesFormatted + ":" + secondsFormatted
}
