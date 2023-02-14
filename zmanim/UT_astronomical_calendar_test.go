package zmanim

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"github.com/vlipovetskii/go-zmanim/zmanim/calculator"
	"testing"
	"time"
)

func TestSunrise(t *testing.T) {
	tm := time.Date(2017, 10, 17, 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	expectedTimes := []time.Time{
		time.Date(tm.Year(), tm.Month(), tm.Day(), 7, 9, 11, 0, basicTestGeoLocations[0].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 6, 39, 32, 0, basicTestGeoLocations[1].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 7, 0, 25, 0, basicTestGeoLocations[2].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 5, 48, 20, 0, basicTestGeoLocations[3].TimeZone()),
		// time.Date(???, basicTestGeoLocations[4].TimeZone()), -> cal.Sunrise() returns nil
		time.Date(tm.Year(), tm.Month(), tm.Day(), 6, 54, 18, 0, basicTestGeoLocations[4].TimeZone()),
	}

	testAstronomicalCalendarTimeResult(t, helper.CurrentFuncName(), expectedTimes,
		func(cal AstronomicalCalendar) time.Time {
			sunrise, ok := cal.Sunrise()
			if !ok {
				helper.Panic("cal.Sunrise() -> ok is false")
			}
			return sunrise
		})
}

func TestSunset(t *testing.T) {
	tm := time.Date(2017, 10, 17, 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	expectedTimes := []time.Time{
		time.Date(tm.Year(), tm.Month(), tm.Day(), 18, 14, 38, 0, basicTestGeoLocations[0].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 18, 8, 46, 0, basicTestGeoLocations[1].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 18, 19, 05, 0, basicTestGeoLocations[2].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 17, 04, 46, 0, basicTestGeoLocations[3].TimeZone()),
		//// time.Date(???, basicTestGeoLocations[4].TimeZone()), -> cal.Sunrise() returns nil
		time.Date(tm.Year(), tm.Month(), tm.Day(), 19, 31, 07, 0, basicTestGeoLocations[4].TimeZone()),
	}

	testAstronomicalCalendarTimeResult(t, helper.CurrentFuncName(), expectedTimes, func(cal AstronomicalCalendar) time.Time {
		sunset, ok := cal.Sunset()
		if !ok {
			helper.Panic("cal.Sunset() -> ok is false")
		}
		return sunset
	})
}

func TestSeaLevelSunrise(t *testing.T) {
	tm := time.Date(2017, 10, 17, 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	expectedTimes := []time.Time{
		time.Date(tm.Year(), tm.Month(), tm.Day(), 7, 9, 51, 0, basicTestGeoLocations[0].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 6, 43, 43, 0, basicTestGeoLocations[1].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 7, 1, 45, 0, basicTestGeoLocations[2].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 5, 49, 21, 0, basicTestGeoLocations[3].TimeZone()),
		//// time.Date(???, basicTestGeoLocations[4].TimeZone()), -> cal.Sunrise() returns nil
		time.Date(tm.Year(), tm.Month(), tm.Day(), 7, 0, 5, 0, basicTestGeoLocations[4].TimeZone()),
	}

	testAstronomicalCalendarTimeResult(t, helper.CurrentFuncName(), expectedTimes, func(cal AstronomicalCalendar) time.Time {
		seaLevelSunrise, ok := cal.SeaLevelSunrise()
		if !ok {
			helper.Panic("cal.SeaLevelSunrise() -> os is false")
		}
		return seaLevelSunrise
	})
}

func TestSeaLevelSunset(t *testing.T) {
	tm := time.Date(2017, 10, 17, 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	expectedTimes := []time.Time{
		time.Date(tm.Year(), tm.Month(), tm.Day(), 18, 13, 58, 0, basicTestGeoLocations[0].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 18, 4, 36, 0, basicTestGeoLocations[1].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 18, 17, 45, 0, basicTestGeoLocations[2].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 17, 3, 45, 0, basicTestGeoLocations[3].TimeZone()),
		//// time.Date(???, basicTestGeoLocations[4].TimeZone()), -> cal.Sunrise() returns nil
		time.Date(tm.Year(), tm.Month(), tm.Day(), 19, 25, 19, 0, basicTestGeoLocations[4].TimeZone()),
	}

	testAstronomicalCalendarTimeResult(t, helper.CurrentFuncName(), expectedTimes, func(cal AstronomicalCalendar) time.Time {
		seaLevelSunset, ok := cal.SeaLevelSunset()
		if !ok {
			helper.Panic("cal.SeaLevelSunset() -> ok is false")
		}
		return seaLevelSunset
	})
}

func TestUtcSunrise(t *testing.T) {
	expectedTimes := []float64{
		11.15327065, 3.65893934, 14.00708152, 20.8057012,
		//None,
		16.90510688,
	}

	testFloat64Result(t, helper.CurrentFuncName(), expectedTimes,
		func(cal AstronomicalCalendar) float64 { return cal.UTCSunrise(90) },
	)
}

func TestUtcSunset(t *testing.T) {
	expectedTimes := []float64{
		22.24410903, 15.14635336, 1.31819979, 8.07962871,
		//None,
		5.51873532,
	}

	testFloat64Result(t, helper.CurrentFuncName(), expectedTimes,
		func(cal AstronomicalCalendar) float64 { return cal.UTCSunset(90) },
	)
}

func TestUtcSeaLevelSunrise(t *testing.T) {
	expectedTimes := []float64{
		11.16434723, 3.72862262, 14.02926518, 20.82268461,
		//None,
		17.00158411,
	}

	testFloat64Result(t, helper.CurrentFuncName(), expectedTimes,
		func(cal AstronomicalCalendar) float64 { return cal.UTCSeaLevelSunrise(90) },
	)
}

func TestUtcSeaLevelSunset(t *testing.T) {
	expectedTimes := []float64{
		22.23304301, 15.07671429, 1.29603174, 8.06265871,
		//None,
		5.42214918,
	}

	testFloat64Result(t, helper.CurrentFuncName(), expectedTimes,
		func(cal AstronomicalCalendar) float64 { return cal.UTCSeaLevelSunset(90) },
	)
}

func TestSunriseOffsetByDegrees(t *testing.T) {
	tm := time.Date(2017, 10, 17, 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	expectedTimes := []time.Time{
		time.Date(tm.Year(), tm.Month(), tm.Day(), 6, 10, 57, 0, basicTestGeoLocations[0].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 5, 50, 43, 0, basicTestGeoLocations[1].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 6, 7, 22, 0, basicTestGeoLocations[2].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 4, 53, 55, 0, basicTestGeoLocations[3].TimeZone()),
		// time.Date(???, basicTestGeoLocations[4].TimeZone()), -> cal.Sunrise() returns nil
		time.Date(tm.Year(), tm.Month(), tm.Day(), 6, 13, 13, 0, basicTestGeoLocations[4].TimeZone()),
	}

	testAstronomicalCalendarTimeResult(t, helper.CurrentFuncName(), expectedTimes, func(cal AstronomicalCalendar) time.Time {
		sunriseOffsetByDegrees, ok := cal.SunriseOffsetByDegrees(102)
		if !ok {
			helper.Panic("cal.SunriseOffsetByDegrees(102) -> ok is false")
		}
		return sunriseOffsetByDegrees
	})
}

func TestSunsetOffsetByDegrees(t *testing.T) {
	tm := time.Date(2017, 10, 17, 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	expectedTimes := []time.Time{
		time.Date(tm.Year(), tm.Month(), tm.Day(), 19, 12, 49, 0, basicTestGeoLocations[0].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 18, 57, 33, 0, basicTestGeoLocations[1].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 19, 12, 5, 0, basicTestGeoLocations[2].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 17, 59, 8, 0, basicTestGeoLocations[3].TimeZone()),
		// time.Date(???, basicTestGeoLocations[4].TimeZone()), -> cal.Sunrise() returns nil
		time.Date(tm.Year(), tm.Month(), tm.Day(), 20, 12, 15, 0, basicTestGeoLocations[4].TimeZone()),
	}

	testAstronomicalCalendarTimeResult(t, helper.CurrentFuncName(), expectedTimes, func(cal AstronomicalCalendar) time.Time {
		sunsetOffsetByDegrees, ok := cal.SunsetOffsetByDegrees(102)
		if !ok {
			helper.Panic("cal.SunsetOffsetByDegrees(102) -> ok is false")
		}
		return sunsetOffsetByDegrees
	})
}

func TestSunriseOffsetByDegreesForArcticTimezoneExtremities(t *testing.T) {
	gregorianDateTime := gdt.NewGDateTime(gdt.NewGDate(2017, 4, 20), gdt.NewGTime0())
	cal := NewAstronomicalCalendar(gregorianDateTime, calculator.DaneborgGeoLocation(), calculator.NewNOAACalculator())

	result, ok := cal.SunriseOffsetByDegrees(94)
	if !ok {
		helper.Panic("cal.SunriseOffsetByDegrees(94) -> ok is false")
	}

	tag := helper.CurrentFuncName()

	// compare without nanoseconds
	expectedTime := time.Date(2017, 4, 19, 23, 54, 23, 0, result.Location())

	assert.Equal(t, tag, expectedTime, time.Date(result.Year(), result.Month(), result.Day(), result.Hour(), result.Minute(), result.Second(), 0, result.Location()))
}

func TestSunsetOffsetByDegreesForArcticTimezoneExtremities(t *testing.T) {
	gregorianDateTime := gdt.NewGDateTime(gdt.NewGDate(2017, 6, 21), gdt.NewGTime0())
	cal := NewAstronomicalCalendar(gregorianDateTime, calculator.HooperBayGeoLocation(), calculator.NewNOAACalculator())

	result, ok := cal.SunsetOffsetByDegrees(94)
	if !ok {
		helper.Panic("cal.SunsetOffsetByDegrees(94) -> ok is false")
	}

	tag := helper.CurrentFuncName()

	// compare without nanoseconds
	expectedTime := time.Date(2017, 6, 22, 2, 0, 16, 0, result.Location())

	assert.Equal(t, tag, expectedTime, time.Date(result.Year(), result.Month(), result.Day(), result.Hour(), result.Minute(), result.Second(), 0, result.Location()))
}

func TestTemporalHour(t *testing.T) {
	expectedTimes := []float64{
		0.92239132,
		0.94567431,
		0.93889721,
		0.936664, // 0.93666451,
		//None,
		1.03504709,
	}

	testFloat64Result(t, helper.CurrentFuncName(), expectedTimes,
		// return geo, (None if result is None else round(result / MathHelper.HOUR_MILLIS, 8))

		func(cal AstronomicalCalendar) float64 {
			temporalHour, ok := cal.TemporalHour()
			if !ok {
				helper.Panic("TemporalHour() -> ok is false")
			}
			return float64(temporalHour) / timeutil.HourMillis
		},
	)
}

func TestSunTransit(t *testing.T) {
	tm := time.Date(2017, 10, 17, 0, 0, 0, 0, timeutil.GmtTimezoneOrPanic())
	basicTestGeoLocations := calculator.BasicTestGeoLocations()

	expectedTimes := []time.Time{
		time.Date(tm.Year(), tm.Month(), tm.Day(), 12, 41, 55, 0, basicTestGeoLocations[0].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 12, 24, 9, 0, basicTestGeoLocations[1].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 12, 39, 45, 0, basicTestGeoLocations[2].TimeZone()),
		time.Date(tm.Year(), tm.Month(), tm.Day(), 11, 26, 33, 0, basicTestGeoLocations[3].TimeZone()),
		// time.Date(???, basicTestGeoLocations[4].TimeZone()), -> cal.Sunrise() returns nil
		time.Date(tm.Year(), tm.Month(), tm.Day(), 13, 12, 42, 0, basicTestGeoLocations[4].TimeZone()),
	}

	testAstronomicalCalendarTimeResult(t, helper.CurrentFuncName(), expectedTimes, func(cal AstronomicalCalendar) time.Time {
		sunTransit, ok := cal.SunTransit()
		if !ok {
			helper.Panic("cal.SunTransit() -> ok is false")
		}
		return sunTransit
	})
}
