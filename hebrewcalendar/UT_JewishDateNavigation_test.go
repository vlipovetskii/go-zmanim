package hebrewcalendar

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"testing"
	"time"
)

func TestJewishForwardMonthToMonth(t *testing.T) {

	jewishDate := NewJewishDate1(jdt.NewJDate(5771, 1, 1))

	tag := "TestJewishForwardMonthToMonth"

	assert.Equal(t, tag, gdt.GDay(5), jewishDate.GDay())
	assert.Equal(t, tag, time.April, jewishDate.GMonth())
	assert.Equal(t, tag, gdt.GYear(2011), jewishDate.GYear())
}

func TestComputeRoshHashana5771(t *testing.T) {

	// At one point, this test was failing as the JewishDate class spun through a never-ending loop...

	jewishDate := NewJewishDate1(jdt.NewJDate(5771, 7, 1))

	tag := "TestComputeRoshHashana5771"

	assert.Equal(t, tag, gdt.GDay(9), jewishDate.GDay())
	assert.Equal(t, tag, time.September, jewishDate.GMonth())
	assert.Equal(t, tag, gdt.GYear(2010), jewishDate.GYear())
}
