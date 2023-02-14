package timeutil

import (
	"github.com/vlipovetskii/go-zmanim/helper"
	"time"
)

const (
	/*
		--- Why Was UTC Created When There Was GMT? https://www.worldtimeserver.com/learn/why-was-utc-created-when-there-was-gmt/\
		- UTC vs. GMT
		UTC is not a time zone, while GMT is.
		Although there was GMT, a committee at the United Nations officially adopted UTC as a standard. This is because it is more accurate than GMT for setting clocks.
		Coordinated Universal Time and Greenwich Mean Time is used interchangeably;
		that’s why knowing their difference is a must, especially when dealing with time.
		They might be confusing, but all you need to remember is that UTC is the current standard time used as a reference for other time zones.
	*/
	GeoLocationNameGMT = "GMT"

	// SecondMillis constant for milliseconds in a second (1,000)
	SecondMillis = 1000

	// MinuteMillis constant for milliseconds in a minute (60,000)
	MinuteMillis = 60 * SecondMillis

	// HourMillis constant for milliseconds in an hour (3,600,000)
	HourMillis = MinuteMillis * 60
)

func LoadLocationOrPanic(name string) *time.Location {
	location, err := time.LoadLocation(name)
	if err != nil {
		helper.PanicOnError(err)
	}
	return location
}

func GmtTimezoneOrPanic() *time.Location {
	return LoadLocationOrPanic(GeoLocationNameGMT)
}

func NewDate(year int, month time.Month, day int, loc *time.Location) time.Time {
	if loc != nil {
		return time.Date(year, month, day, 0, 0, 0, 0, loc)
	} else {
		return time.Date(year, month, day, 0, 0, 0, 0, GmtTimezoneOrPanic())
	}
}

func NewDateOfToday(loc *time.Location) time.Time {
	now := time.Now()
	return NewDate(now.Year(), now.Month(), now.Day(), loc)
}

/*
	func RequireUtcTime(tm time.Time) {
		if tm.Location().String() != GeoLocationNameGMT {
			helper.Panic(fmt.Sprintf("%v is not UTC", tm))
		}
	}
*/
func WithoutNanoseconds(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), 0, tm.Location())
}
