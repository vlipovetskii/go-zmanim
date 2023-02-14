package zmanim

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"github.com/vlipovetskii/go-zmanim/zmanim/calculator"
	"testing"
	"time"
)

func testNewZmanimCalendar(year int, month time.Month, day int, geoLocation calculator.GeoLocation) ZmanimCalendar {
	return NewZmanimCalendar(gdt.NewGDateTime1(timeutil.NewDate(year, month, day, nil)), geoLocation, calculator.NewNOAACalculator())
}

func test2to1[T any](tm T, ok bool) func(funcInvocation string) T {
	return func(funcInvocation string) T {
		if !ok {
			helper.Panic(fmt.Sprintf("%s -> ok is false", funcInvocation))
		}
		return tm
	}
}

func TestHanetz(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testNewZmanimCalendar(2017, 10, 17, calculator.LakewoodGeoLocation())
	assert.Equal(t, tag, test2to1(cal.SeaLevelSunrise())("cal.SeaLevelSunrise()"), test2to1(cal.Hanetz())("cal.Hanetz()"))
}

func TestShkia(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testNewZmanimCalendar(2017, 10, 17, calculator.LakewoodGeoLocation())
	assert.Equal(t, tag, test2to1(cal.SeaLevelSunset())("cal.SeaLevelSunset()"), test2to1(cal.Shkia())("cal.Shkia()"))
}

func TestTzais(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"18:54:29",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Tzais())("cal.Tzais()")
		},
	)
}

func TestTzaisWithCustomDegreeOffset(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:53:34",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Tzais3(19.8, 0, 0))("cal.Tzais3(19.8, 0, 0)")
		},
	)
}

// TestTzaisWithCustomMinuteOffset
func TestTzaisWithCustomMinuteOffset(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:13:58",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Tzais3(0, 60, 0))("cal.Tzais3(0, 60, 0)")
		},
	)
}

// TestTzaisWithCustomZmanisOffset is test_tzais_with_custom_temporal_minute_offset in py-zmanim
func TestTzaisWithCustomZmanisOffset(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:36:59",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Tzais3(0, 0, 90))("cal.Tzais3(0, 0, 90)")
		},
	)
}

func TestTzais72(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:25:58",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Tzais72())("cal.Tzais72()")
		},
	)
}

func TestAlos(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"5:49:30",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Alos())("cal.Alos()")
		},
	)
}

func TestAlossWithCustomDegreeOffset(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"5:30:07",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Alos3(19.8, 0, 0))("cal.Alos3(19.8, 0, 0)")
		},
	)
}

func TestAlosWithCustomMinuteOffset(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"06:09:51",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Alos3(0, 60, 0))("cal.Alos3(0, 60, 0)")
		},
	)
}

// TestAlosWithCustomZmanisOffset is test_alos_with_custom_temporal_minute_offset in py-zmanim
func TestAlosWithCustomZmanisOffset(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"05:46:50",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Alos3(0, 0, 90))("cal.Alos3(0, 0, 90)")
		},
	)
}

func TestAlos72(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"05:57:51",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Alos72())("cal.Alos72()")
		},
	)
}

func TestChatzos(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"12:41:55",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.Chatzos())("cal.Chatzos()")
		},
	)
}

func TestSofZmanShma(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"09:42:10",
		func(cal ZmanimCalendar) time.Time {
			return cal.SofZmanShma(
				test2to1(cal.SunriseOffsetByDegrees(96))("cal.SunriseOffsetByDegrees(96)"),
				test2to1(cal.SunsetOffsetByDegrees(96))("cal.SunsetOffsetByDegrees(96)"),
			)
		},
	)
}

func TestSofZmanShmaGRA(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"09:55:53",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.SofZmanShmaGRA())("cal.SofZmanShmaGRA()")
		},
	)
}

func TestSofZmanShmaMGA(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"09:19:53",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.SofZmanShmaMGA())("cal.SofZmanShmaMGA()")
		},
	)
}

func TestSofZmanTfila(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"10:42:05",
		func(cal ZmanimCalendar) time.Time {
			return cal.SofZmanTfila(
				test2to1(cal.SunriseOffsetByDegrees(96))("cal.SunriseOffsetByDegrees(96)"),
				test2to1(cal.SunsetOffsetByDegrees(96))("cal.SunsetOffsetByDegrees(96)"),
			)
		},
	)
}

func TestSofZmanTfilaGRA(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"10:51:14",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.SofZmanTfilaGRA())("cal.SofZmanTfilaGRA()")
		},
	)
}

func TestSofZmanTfilaMGA(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"10:27:14",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.SofZmanTfilaMGA())("cal.SofZmanTfilaMGA()")
		},
	)
}

func TestMinchaGedola(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"13:09:35",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.MinchaGedola())("cal.MinchaGedola()")
		},
	)
}

func TestMinchaKetana(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"15:55:37",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.MinchaKetana())("cal.MinchaKetana()")
		},
	)
}

func TestPlagHamincha(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"17:04:48",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.PlagHamincha())("cal.PlagHamincha()")
		},
	)
}

func TestCandleLighting(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"17:55:58",
		func(cal ZmanimCalendar) time.Time {
			return test2to1(cal.CandleLighting())("cal.CandleLighting()")
		},
	)
}

func TestShaahZmanis(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(3594499),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			return cal.ShaahZmanis(
				test2to1(cal.SunriseOffsetByDegrees(96))("cal.SunriseOffsetByDegrees(96)"),
				test2to1(cal.SunsetOffsetByDegrees(96))("cal.SunsetOffsetByDegrees(96)"),
			)
		},
	)
}

func TestShaahZmanisGRA(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(3320608),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			return test2to1(cal.ShaahZmanisGRA())("cal.ShaahZmanisGRA()")
		},
	)
}

func TestShaahZmanisMGA(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(4040608),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			return test2to1(cal.ShaahZmanisMGA())("cal.ShaahZmanisMGA()")
		},
	)
}

func TestShaahZmanisByDegreesAndOffsetWithDegrees(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(4040608),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			return test2to1(cal.ShaahZmanisByDegreesAndOffset(0, 72))("cal.ShaahZmanisByDegreesAndOffset(0, 72)")
		},
	)
}

func TestShaahZmanisByDegreesAndOffsetWithBothDegreesAndOffset(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(4314499),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			return test2to1(cal.ShaahZmanisByDegreesAndOffset(6, 72))("cal.ShaahZmanisByDegreesAndOffset(6, 72)")
		},
	)
}

func TestHanetzUsingElevation(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testZmanimCalendar()
	cal.SetUseElevation(true)
	assert.Equal(t, tag, test2to1(cal.Hanetz())("cal.Hanetz()"), test2to1(cal.Sunrise())("cal.Sunrise()"))
}

func TestShkiaUsingElevation(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testZmanimCalendar()
	cal.SetUseElevation(true)
	assert.Equal(t, tag, test2to1(cal.Shkia())("cal.Shkia()"), test2to1(cal.Sunset())("cal.Sunset()"))
}

func TestTzaisUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"18:54:29",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Tzais())("cal.Tzais()")
		},
	)
}

func TestTzaisWithCustomDegreeOffsetUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:53:34",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Tzais3(19.8, 0, 0))("cal.Tzais3(19.8, 0, 0)")
		},
	)
}

func TestTzaisWithCustomMinuteOffsetUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:14:38",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Tzais3(0, 60, 0))("cal.Tzais3(0, 60, 0)")
		},
	)
}

// TestTzaisWithCustomZmanisOffset is test_tzais_with_custom_temporal_minute_offset_using_elevation in py-zmanim
func TestTzaisWithCustomZmanisOffsetUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:37:49",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Tzais3(0, 0, 90))("cal.Tzais3(0, 0, 90)")
		},
	)
}

func TestTzais72UsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"19:26:38",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Tzais72())("cal.Tzais72()")
		},
	)
}

func TestAlosUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"05:49:30",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Alos())("cal.Alos()")
		},
	)
}

func TestAlosWithCustomDegreeOffsetUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"05:30:07",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Alos3(19.8, 0, 0))("cal.Alos3(19.8, 0, 0)")
		},
	)
}

func TestAlosWithCustomMinuteOffsetUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"06:09:11",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Alos3(0, 60, 0))("cal.Alos3(0, 60, 0)")
		},
	)
}

// TestAlosWithCustomZmanisOffset is test_alos_with_custom_temporal_minute_offset_using_elevation in py-zmanim
func TestAlosWithCustomZmanisOffsetUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"05:46:00",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Alos3(0, 0, 90))("cal.Alos3(0, 0, 90)")
		},
	)
}

func TestAlos72UsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"05:57:11",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Alos72())("cal.Alos72()")
		},
	)
}

func TestChatzosUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"12:41:55",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.Chatzos())("cal.Chatzos()")
		},
	)
}

func TestSofZmanShmaUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"09:42:10",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return cal.SofZmanShma(
				test2to1(cal.SunriseOffsetByDegrees(96))("cal.SunriseOffsetByDegrees(96)"),
				test2to1(cal.SunsetOffsetByDegrees(96))("cal.SunsetOffsetByDegrees(96)"),
			)
		},
	)
}

func TestSofZmanShmaGRAUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"09:55:33",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.SofZmanShmaGRA())("cal.SofZmanShmaGRA()")
		},
	)
}

func TestSofZmanShmaMGAUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"09:19:33",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.SofZmanShmaMGA())("cal.SofZmanShmaMGA()")
		},
	)
}

func TestSofZmanTfilaUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"10:42:05",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return cal.SofZmanTfila(
				test2to1(cal.SunriseOffsetByDegrees(96))("cal.SunriseOffsetByDegrees(96)"),
				test2to1(cal.SunsetOffsetByDegrees(96))("cal.SunsetOffsetByDegrees(96)"),
			)
		},
	)
}

func TestSofZmanTfilaGRAUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"10:51:00",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.SofZmanTfilaGRA())("cal.SofZmanTfilaGRA()")
		},
	)
}

func TestSofZmanTfilaMGAUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"10:27:00",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.SofZmanTfilaMGA())("cal.SofZmanTfilaMGA()")
		},
	)
}

func TestMinchaGedolaUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"13:09:38",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.MinchaGedola())("cal.MinchaGedola()")
		},
	)
}

func TestMinchaKetanaUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"15:56:00",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.MinchaKetana())("cal.MinchaKetana()")
		},
	)
}

func TestPlagHaminchaUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"17:05:19",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.PlagHamincha())("cal.PlagHamincha()")
		},
	)
}

func TestCandleLightingUsingElevation(t *testing.T) {
	testZmanimCalendarTimeResult(t, helper.CurrentFuncName(),
		"17:55:58",
		func(cal ZmanimCalendar) time.Time {
			cal.SetUseElevation(true)
			return test2to1(cal.CandleLighting())("cal.CandleLighting()")
		},
	)
}

func TestShaahZmanisUsingElevation(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(3594499),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			cal.SetUseElevation(true)
			return cal.ShaahZmanis(
				test2to1(cal.SunriseOffsetByDegrees(96))("cal.SunriseOffsetByDegrees(96)"),
				test2to1(cal.SunsetOffsetByDegrees(96))("cal.SunsetOffsetByDegrees(96)"),
			)
		},
	)
}

func TestShaahZmanisGRAUsingElevation(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(3327251),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			cal.SetUseElevation(true)
			return test2to1(cal.ShaahZmanisGRA())("cal.ShaahZmanisGRA()")
		},
	)
}

func TestShaahZmanisMGAUsingElevation(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(4047251),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			cal.SetUseElevation(true)
			return test2to1(cal.ShaahZmanisMGA())("cal.ShaahZmanisMGA()")
		},
	)
}

func TestShaahZmanisByDegreesAndOffsetWithDegreesUsingElevation(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(4047251),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			cal.SetUseElevation(true)
			return test2to1(cal.ShaahZmanisByDegreesAndOffset(0, 72))("cal.ShaahZmanisByDegreesAndOffset(0, 72)")
		},
	)
}

func TestShaahZmanisByDegreesAndOffsetWithBothDegreesAndOffsetUsingElevation(t *testing.T) {
	testZmanimCalendarGMillisecondResult(t, helper.CurrentFuncName(),
		gdt.GMillisecond(4314499),
		func(cal ZmanimCalendar) gdt.GMillisecond {
			cal.SetUseElevation(true)
			return test2to1(cal.ShaahZmanisByDegreesAndOffset(6, 72))("cal.ShaahZmanisByDegreesAndOffset(6, 72)")
		},
	)
}

func TestAssurBemelachaForStandardDay(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testZmanimCalendar()
	assert.False(t, tag, cal.IsAssurBemlacha(test2to1(cal.Shkia())("cal.Shkia()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.False(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.False(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
}

func TestAssurBemelachaForMelachadDay(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testNewZmanimCalendar(2017, 10, 21, calculator.LakewoodGeoLocation())
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Shkia())("cal.Shkia()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.False(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
}

func TestAssurBemelachaForMelachadDayWithCustomTzaisTime(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testNewZmanimCalendar(2017, 10, 21, calculator.LakewoodGeoLocation())
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.False(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
}

func TestAssurBemelachaForMelachadDayWithCustomTzaisRule(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testNewZmanimCalendar(2017, 10, 21, calculator.LakewoodGeoLocation())
	tzais := test2to1(cal.Tzais3(11.5, 0, 0))("cal.Tzais3(11.5, 0, 0)")
	assert.True(t, tag, cal.IsAssurBemlacha(tzais.Add(-2*time.Second), tzais, false))
	assert.False(t, tag, cal.IsAssurBemlacha(tzais.Add(2*time.Second), tzais, false))
}

func TestAssurBemelachaPriorToIssurMelachadDay(t *testing.T) {
	tag := helper.CurrentFuncName()
	cal := testNewZmanimCalendar(2017, 10, 20, calculator.LakewoodGeoLocation())
	assert.False(t, tag, cal.IsAssurBemlacha(test2to1(cal.Shkia())("cal.Shkia()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
}

func TestAssurBemelachaOnFirstOfTwoIssurMelachaDays(t *testing.T) {
	tag := helper.CurrentFuncName()

	// first day of pesach
	cal := testNewZmanimCalendar(2018, 03, 31, calculator.LakewoodGeoLocation())

	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Shkia())("cal.Shkia()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), false))
}

func TestAssurBemelachaOnFirstOfSingleIssurMelachaInIsrael(t *testing.T) {
	tag := helper.CurrentFuncName()

	// first day of pesach
	cal := testNewZmanimCalendar(2018, 03, 31, calculator.LakewoodGeoLocation())

	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), true))
	assert.False(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), true))
}

func TestAssurBemelachaOnFirstOfTwoIssurMelachaDaysInIsrael(t *testing.T) {
	tag := helper.CurrentFuncName()

	// Shabbos before Shavuos
	cal := testNewZmanimCalendar(2018, 05, 19, calculator.LakewoodGeoLocation())

	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(-2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), true))
	assert.True(t, tag, cal.IsAssurBemlacha(test2to1(cal.Tzais())("cal.Tzais()").Add(2*time.Second), test2to1(cal.Tzais())("cal.Tzais()"), true))
}
