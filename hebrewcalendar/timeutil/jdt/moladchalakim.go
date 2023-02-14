package jdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
)

type MoladChalakim int32
type MoladChalakim64 int64

func (t MoladChalakim) Validate() {
	if t < 0 || t > 17 {
		helper.Panic(fmt.Sprintf("Chalakim/parts < 0 or > 17 can't be set. %d is invalid.", t))
	}
}

/*
ToAbsDate returns the number of days from the Jewish epoch from the number of chalakim from the epoch passed in.
chalakim is the number of chalakim since the beginning of Sunday prior to BaHaRaD
*/
func (t MoladChalakim64) ToAbsDate() gdt.GDay {
	return gdt.GDay(int32(int64(t)/int64(ChalakimPerDay))) + JewishEpoch
}
