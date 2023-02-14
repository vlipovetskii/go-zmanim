package hebrewcalendar

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"testing"
)

/**
 * Validate the days in a Hebrew month (in various types of years) are correct.
 */

func TestDaysInMonthsInHaserYear(t *testing.T) {

	assertHaser(t, 5773)
	assertHaser(t, 5777)
	assertHaser(t, 5781)

	assertHaserLeap(t, 5784)
	assertHaserLeap(t, 5790)
	assertHaserLeap(t, 5793)
}

func TestDaysInMonthsInQesidrahYear(t *testing.T) {

	assertQesidrah(t, 5769)
	assertQesidrah(t, 5772)
	assertQesidrah(t, 5778)
	assertQesidrah(t, 5786)
	assertQesidrah(t, 5789)
	assertQesidrah(t, 5792)

	assertQesidrahLeap(t, 5782)
}

func TestDaysInMonthsInShalemYear(t *testing.T) {

	assertShalem(t, 5770)
	assertShalem(t, 5780)
	assertShalem(t, 5783)
	assertShalem(t, 5785)
	assertShalem(t, 5788)
	assertShalem(t, 5791)
	assertShalem(t, 5794)

	assertShalemLeap(t, 5771)
	assertShalemLeap(t, 5774)
	assertShalemLeap(t, 5776)
	assertShalemLeap(t, 5779)
	assertShalemLeap(t, 5787)
	assertShalemLeap(t, 5795)
}

func assertHaser(t *testing.T, year jdt.JYear) {
	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	tag := "assertHaser"

	assert.False(t, tag, jewishDate.IsCheshvanLong())
	assert.True(t, tag, jewishDate.IsKislevShort())
}

func assertHaserLeap(t *testing.T, year jdt.JYear) {
	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	tag := "assertHaserLeap"

	assert.True(t, tag, jewishDate.IsLeapJYear())
}

func assertQesidrah(t *testing.T, year jdt.JYear) {
	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	tag := "assertQesidrah"

	assert.False(t, tag, jewishDate.IsCheshvanLong())
	assert.False(t, tag, jewishDate.IsKislevShort())
}

func assertQesidrahLeap(t *testing.T, year jdt.JYear) {
	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	tag := "assertQesidrahLeap"

	assertQesidrah(t, year)
	assert.True(t, tag, jewishDate.IsLeapJYear())
}

func assertShalem(t *testing.T, year jdt.JYear) {
	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	tag := "assertShalem"

	assert.True(t, tag, jewishDate.IsCheshvanLong())
	assert.False(t, tag, jewishDate.IsKislevShort())
}

func assertShalemLeap(t *testing.T, year jdt.JYear) {
	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	tag := "assertShalemLeap"

	assertShalem(t, year)
	assert.True(t, tag, jewishDate.IsLeapJYear())
}
