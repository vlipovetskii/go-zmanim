package hebrewcalendar

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"testing"
	"time"
)

/**
 * Checks that we can roll forward & backward the gregorian dates...
 */

func TestGregorianForwardMonthToMonth(t *testing.T) {

	tag := helper.CurrentFuncName()

	hebrewDate := NewJewishDate()

	gYear := gdt.GYear(2011)

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.January, 31))

	assert.Equal(t, tag, jdt.NewJDate(5771, jdt.JMonth(11), 26), hebrewDate.JDate())

	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.February, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(11), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(27), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.February, 28))
	assert.Equal(t, tag, time.February, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(28), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(12), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(24), hebrewDate.JDay())

	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.March, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(12), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(25), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.March, 31))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.April, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(13), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(26), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.April, 30))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.May, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(1), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(27), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.May, 31))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.June, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(2), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(28), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.June, 30))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.July, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(3), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(29), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.July, 31))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.August, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(5), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(1), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.August, 31))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.September, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(6), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(2), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.September, 30))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.October, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(7), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(3), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.October, 31))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.November, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JYear(5772), hebrewDate.JYear())
	assert.Equal(t, tag, jdt.JMonth(8), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(4), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.November, 30))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, time.December, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(9), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(5), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.December, 31))
	hebrewDate.ForwardJDay(1)
	assert.Equal(t, tag, gdt.GYear(2012), hebrewDate.GYear())
	assert.Equal(t, tag, time.January, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(1), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(10), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(6), hebrewDate.JDay())
}

func TestGregorianBackwardMonthToMonth(t *testing.T) {

	tag := "TestGregorianBackwardMonthToMonth"

	hebrewDate := NewJewishDate()

	gYear := gdt.GYear(2010)

	hebrewDate.SetGDate(gdt.NewGDate(2011, time.January, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, gdt.GYear(2010), hebrewDate.GYear())
	assert.Equal(t, tag, time.December, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(31), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(10), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(24), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.December, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.November, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(30), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(9), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(23), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.November, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.October, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(31), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(8), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(23), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.October, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.September, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(30), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(7), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(22), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.September, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.August, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(31), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JYear(5770), hebrewDate.JYear())
	assert.Equal(t, tag, jdt.JMonth(6), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(21), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.August, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.July, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(31), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(5), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(20), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.July, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.June, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(30), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(4), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(18), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.June, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.May, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(31), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(3), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(18), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.May, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.April, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(30), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(2), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(16), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.April, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.March, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(31), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(1), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(16), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.March, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.February, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(28), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(12), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(14), hebrewDate.JDay())

	hebrewDate.SetGDate(gdt.NewGDate(gYear, time.February, 1))
	hebrewDate.BackJDay(1)
	assert.Equal(t, tag, time.January, hebrewDate.GMonth())
	assert.Equal(t, tag, gdt.GDay(31), hebrewDate.GDay())
	assert.Equal(t, tag, jdt.JMonth(11), hebrewDate.JMonth())
	assert.Equal(t, tag, jdt.JDay(16), hebrewDate.JDay())

}
