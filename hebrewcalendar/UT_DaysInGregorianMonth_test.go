package hebrewcalendar

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"testing"
)

/**
 * Verify the calculation of the number of days in a month. Not too hard...just the rules about when February
 *  has 28 or 29 days...
 */

func TestDaysInMonth(t *testing.T) {
	assertDaysInMonth(t, false, 2011)
}

func TestDestDaysInMonthLeapYear(t *testing.T) {
	assertDaysInMonth(t, true, 2012)
}

func TestDaysInMonth100Year(t *testing.T) {
	assertDaysInMonth(t, false, 2100)
}

func TestDaysInMonth400Year(t *testing.T) {
	assertDaysInMonth(t, true, 2000)
}

func assertDaysInMonth(t *testing.T, febIsLeap bool, gregorianYear gdt.GYear) {

	tag := "assertDaysInMonth"

	assert.Equal(t, tag, gdt.GDay(31), gdt.LastGDayOfGMonth(1, gregorianYear))

	if febIsLeap {
		assert.Equal(t, tag, gdt.GDay(29), gdt.LastGDayOfGMonth(2, gregorianYear))
	} else {
		assert.Equal(t, tag, gdt.GDay(28), gdt.LastGDayOfGMonth(2, gregorianYear))
	}

	assert.Equal(t, tag, gdt.GDay(31), gdt.LastGDayOfGMonth(3, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(30), gdt.LastGDayOfGMonth(4, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(31), gdt.LastGDayOfGMonth(5, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(30), gdt.LastGDayOfGMonth(6, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(31), gdt.LastGDayOfGMonth(7, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(31), gdt.LastGDayOfGMonth(8, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(30), gdt.LastGDayOfGMonth(9, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(31), gdt.LastGDayOfGMonth(10, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(30), gdt.LastGDayOfGMonth(11, gregorianYear))
	assert.Equal(t, tag, gdt.GDay(31), gdt.LastGDayOfGMonth(12, gregorianYear))
}
