package zmanim

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"github.com/vlipovetskii/go-zmanim/zmanim/calculator"
	"math"
	"testing"
	"time"
)

func testAstronomicalCalendarTimeResult(t *testing.T, tag string, expectedTimes []time.Time, testFunc func(cal AstronomicalCalendar) time.Time) {

	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	gregorianDateTime := gdt.NewGDateTime(gdt.NewGDate(2017, 10, 17), gdt.NewGTime0())

	for i := 0; i < len(expectedTimes); i++ {

		cal := NewAstronomicalCalendar(gregorianDateTime, basicTestGeoLocations[i], calculator.NewNOAACalculator())

		expectedTime := expectedTimes[i]

		result := testFunc(cal)
		// compare without nanoseconds
		assert.Equal(t, fmt.Sprintf("%s[%d]", tag, i), expectedTime, timeutil.WithoutNanoseconds(result))
	}
}

func testZmanimCalendar() ZmanimCalendar {
	return NewZmanimCalendar(gdt.NewGDateTime1(timeutil.NewDate(2017, 10, 17, nil)), calculator.LakewoodGeoLocation(), calculator.NewNOAACalculator())
}

func wantTime(wantHour int, wantMinute int, wantSecond int) time.Time {
	return time.Date(2017, 10, 17, wantHour, wantMinute, wantSecond, 0, calculator.LakewoodGeoLocation().TimeZone())
}

func wantTime2(wantTimeStr string) time.Time {
	tm, err := time.Parse("15:04:05", wantTimeStr)
	if err != nil {
		helper.PanicOnError(err)
	}
	return wantTime(tm.Hour(), tm.Minute(), tm.Second())
}

func testZmanimCalendarTimeResult(t *testing.T, tag string, wantTimeStr string, testFunc func(cal ZmanimCalendar) time.Time) {
	cal := testZmanimCalendar()
	want := wantTime2(wantTimeStr)
	result := testFunc(cal)
	assert.Equal(t, tag, want, timeutil.WithoutNanoseconds(result))
}

func testZmanimCalendarGMillisecondResult(t *testing.T, tag string, want gdt.GMillisecond, testFunc func(cal ZmanimCalendar) gdt.GMillisecond) {
	cal := testZmanimCalendar()
	result := testFunc(cal)
	assert.Equal(t, tag, want, result)
}

/*
roundFloat
--- Round float to any precision in Go https://gosamples.dev/round-float/
--- How can we truncate float64 type to a particular precision? https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision
*/
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func testFloat64Result(t *testing.T, tag string, expectedTimes []float64, testFunc func(cal AstronomicalCalendar) float64) {
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	gregorianDateTime := gdt.NewGDateTime(gdt.NewGDate(2017, 10, 17), gdt.NewGTime0())

	for i := 0; i < len(expectedTimes); i++ {

		cal := NewAstronomicalCalendar(gregorianDateTime, basicTestGeoLocations[i], calculator.NewNOAACalculator())

		expectedTime := expectedTimes[i]

		result := testFunc(cal)
		assert.Equal(t, fmt.Sprintf("%s[%d]", tag, i), roundFloat(expectedTime, 6), roundFloat(result, 6))
	}
}
