package gdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
)

type GYear int32

/*
IsLeapGYear returns is the year passed in is a [Gregorian leap year]: https://en.wikipedia.org/wiki/Leap_year#Gregorian_calendar
year: the Gregorian year
*/
func (t GYear) IsLeapGYear() bool {
	return (t%4 == 0 && t%100 != 0) || (t%400 == 0)
}

func (t GYear) DaysInGYear() GDay {
	if t.IsLeapGYear() {
		return 366
	} else {
		return 365
	}
}

/*
Validate validates a Gregorian year for validity.
year the Gregorian year to validate. It will reject any year < 1.
*/
func (t GYear) Validate() {
	if t < 1 {
		helper.Panic(fmt.Sprintf("Years < 1 can't be calculated. %d is invalid.", t))
	}
}
