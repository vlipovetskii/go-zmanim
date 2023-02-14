package gdt

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"time"
)

/*
GDateTime is an internal structure to aggregate GDate, GTime
*/
type GDateTime struct {
	D GDate
	T GTime
}

func NewGDateTime(gDate GDate, gTime GTime) GDateTime {
	return GDateTime{D: gDate, T: gTime}
}

/*
NewGDateTime1
E.g. NewGDateTime1(time.Now())
*/
func NewGDateTime1(tm time.Time) GDateTime {
	return GDateTime{D: NewGDate1(tm), T: NewGTime1(tm)}
}

func (t GDateTime) ToTime(loc *time.Location) time.Time {
	if loc != nil {
		return time.Date(int(t.D.Year), t.D.Month, int(t.D.Day), t.T.Hour, t.T.Minute, t.T.Second, t.T.Nanosecond, loc)
	} else {
		return time.Date(int(t.D.Year), t.D.Month, int(t.D.Day), t.T.Hour, t.T.Minute, t.T.Second, t.T.Nanosecond, timeutil.GmtTimezoneOrPanic())
	}
}

func (t GDateTime) Validate() {
	t.D.Validate()
	t.T.Validate()
}
