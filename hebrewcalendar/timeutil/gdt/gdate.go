package gdt

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"time"
)

/*
GDate is an internal structure to track the date (without time.Location) used by different classes
*/
type GDate struct {
	Year  GYear
	Month time.Month
	Day   GDay
}

func NewGDate(year GYear, month time.Month, day GDay) GDate {
	return GDate{Year: year, Month: month, Day: day}
}

func NewGDate1(tm time.Time) GDate {
	return NewGDate(GYear(tm.Year()), tm.Month(), GDay(tm.Day()))
}

func NewGDateToday() GDate {
	return NewGDate1(time.Now())
}

/*
NewGDate2 creates GDate from gAbsDate.
gAbsDate is the absolute date (days since January 1, 0001, on the Gregorian calendar)
*/
func NewGDate2(gAbsDate GDay) GDate {
	year := GYear(gAbsDate / 366) // Search forward year by year from approximate year

	for gAbsDate >= NewGDate(year+1, 1, 1).ToAbsDate() { // gregorianDateToAbsDate(year+1, 1, 1) {
		year++
	}

	var month time.Month = 1 // Search forward month by month from January
	for gAbsDate > NewGDate(year, month, LastGDayOfGMonth(month, year)).ToAbsDate() {
		month++
	}

	day := gAbsDate - NewGDate(year, month, 1).ToAbsDate() + 1

	return NewGDate(year, month, day)

}

func (t GDate) ToTime(loc *time.Location) time.Time {
	if loc != nil {
		return time.Date(int(t.Year), t.Month, int(t.Day), 0, 0, 0, 0, loc)
	} else {
		return time.Date(int(t.Year), t.Month, int(t.Day), 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	}
}

/*
ToAbsDate computes the absolute date from a Gregorian date. ND+ER
year the Gregorian year
month the Gregorian month. Unlike the Java Calendar where January has the value of 0,
This expects a 1 for January
dayOfMonth the day of the month (1st, 2nd, etc...)
return the absolute Gregorian day
*/
func (t GDate) ToAbsDate() GDay {
	absDate := t.Day

	for m := t.Month - 1; m > 0; m-- {
		absDate += LastGDayOfGMonth(m, t.Year) // days in prior months of the year
	}

	return absDate + 365*GDay(t.Year-1) + GDay(t.Year-1)/4 - GDay(t.Year-1)/100 + GDay(t.Year-1)/400

}

/*
LastGDayOfGMonth returns the number of days in a given month in a given month and year.
*/
func LastGDayOfGMonth(month time.Month, year GYear) GDay {
	switch month {
	case time.February:
		if year.IsLeapGYear() {
			return 29
		} else {
			return 28
		}
	case time.April, time.June, time.September, time.November:
		return 30
	default:
		return 31
	}
}

func (t GDate) Validate() {
	t.Year.Validate()
	ValidateMonth(t.Month)
	t.Day.Validate()
}
