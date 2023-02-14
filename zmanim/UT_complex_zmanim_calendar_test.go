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

func testNewComplexZmanimCalendar(year int, month time.Month, day int, geoLocation calculator.GeoLocation) ComplexZmanimCalendar {
	return NewComplexZmanimCalendar(gdt.NewGDateTime1(timeutil.NewDate(year, month, day, nil)), geoLocation, calculator.NewNOAACalculator())
}

func TestShaahZmanis19Point8Degrees(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testNewComplexZmanimCalendar(2017, 10, 17, calculator.LakewoodGeoLocation())
	assert.Equal(t, tag, test2to1(cal.ShaahZmanis19Point8Degrees())("cal.ShaahZmanis19Point8Degrees()"), temporalHour(test2to1(cal.Alos19Point8Degrees())("cal.Alos19Point8Degrees"), test2to1(cal.Tzais19Point8Degrees())("cal.Tzais19Point8Degrees")))
}
