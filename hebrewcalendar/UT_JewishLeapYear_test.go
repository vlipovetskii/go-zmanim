package hebrewcalendar

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"testing"
)

/**
 * Verify correct calculations of when a Hebrew leap year occurs.
 */

func TestIsLeapYear(t *testing.T) {

	shouldBeLeapYear(t, 5160)
	shouldNotBeLeapYear(t, 5536)

	shouldNotBeLeapYear(t, 5770)
	shouldBeLeapYear(t, 5771)
	shouldNotBeLeapYear(t, 5772)
	shouldNotBeLeapYear(t, 5773)
	shouldBeLeapYear(t, 5774)
	shouldNotBeLeapYear(t, 5775)
	shouldBeLeapYear(t, 5776)
	shouldNotBeLeapYear(t, 5777)
	shouldNotBeLeapYear(t, 5778)
	shouldBeLeapYear(t, 5779)
	shouldNotBeLeapYear(t, 5780)
	shouldNotBeLeapYear(t, 5781)
	shouldBeLeapYear(t, 5782)
	shouldNotBeLeapYear(t, 5783)
	shouldBeLeapYear(t, 5784)
	shouldNotBeLeapYear(t, 5785)
	shouldNotBeLeapYear(t, 5786)
	shouldBeLeapYear(t, 5787)
	shouldNotBeLeapYear(t, 5788)
	shouldNotBeLeapYear(t, 5789)
	shouldBeLeapYear(t, 5790)
	shouldNotBeLeapYear(t, 5791)
	shouldNotBeLeapYear(t, 5792)
	shouldBeLeapYear(t, 5793)
	shouldNotBeLeapYear(t, 5794)
	shouldBeLeapYear(t, 5795)
}

func shouldBeLeapYear(t *testing.T, year jdt.JYear) {

	tag := "shouldBeLeapYear"

	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	assert.True(t, tag, jewishDate.IsLeapJYear())
}

func shouldNotBeLeapYear(t *testing.T, year jdt.JYear) {

	tag := "shouldNotBeLeapYear"

	jewishDate := NewJewishDate()
	jewishDate.SetJDate(jdt.NewJDate(year, jewishDate.JMonth(), jewishDate.JDay()))

	assert.False(t, tag, jewishDate.IsLeapJYear())
}
