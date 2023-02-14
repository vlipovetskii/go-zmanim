package gdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/helper"
	"time"
)

type GHour int
type GHourH64 float64
type GMinute int
type GMinuteF64 float64
type GSecond int
type GMillisecond int
type GNanosecond int

func (t GHour) Validate() {
	if t < 0 || t > 23 {
		helper.Panic(fmt.Sprintf("Hour < 0 or > 23 can't be set. %d is invalid.", t))
	}
}

func (t GMinute) Validate() {
	if t < 0 || t > 59 {
		helper.Panic(fmt.Sprintf("Minutes < 0 or > 59 can't be set. %d is invalid.", t))
	}
}

func (t GSecond) Validate() {
	if t < 0 || t > 59 {
		helper.Panic(fmt.Sprintf("Second < 0 or > 59 can't be set. %d is invalid.", t))
	}
}

func (t GNanosecond) Validate() {
	if t < 0 || t > 1000000000 {
		helper.Panic(fmt.Sprintf("Nanosecond < 0 or > 1000000000 can't be set. %d is invalid.", t))
	}
}

func (t GHour) ToMilliseconds() GMillisecond {
	return GMillisecond(t * timeutil.HourMillis)
}

func (t GMinute) ToMilliseconds() GMillisecond {
	return GMillisecond(t * timeutil.MinuteMillis)
}

func (t GSecond) ToMilliseconds() GMillisecond {
	return GMillisecond(t * timeutil.SecondMillis)
}

func (t GMinuteF64) ToMilliseconds() GMillisecond {
	return GMillisecond(t * timeutil.MinuteMillis)
}

/*
GTime is an internal structure to track the time (without time.Location) used by different classes
*/
type GTime struct {
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
}

func NewGTime(hour int, minute int, second int, nanosecond int) GTime {
	return GTime{Hour: hour, Minute: minute, Second: second, Nanosecond: nanosecond}
}

func NewGTime0() GTime {
	return NewGTime(0, 0, 0, 0)
}

func NewGTime1(tm time.Time) GTime {
	return NewGTime(tm.Hour(), tm.Minute(), tm.Second(), tm.Nanosecond())
}

func (t GTime) Validate() {
	GHour(t.Hour).Validate()
	GMinute(t.Minute).Validate()
	GSecond(t.Second).Validate()
	GNanosecond(t.Nanosecond).Validate()
}
