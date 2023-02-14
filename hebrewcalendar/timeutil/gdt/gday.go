package gdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
)

type GDay int32

/*
Validate validates a Gregorian day of month for validity.
dayOfMonth the day of the Gregorian month to validate. It will reject any value < 1 and > 31
*/
func (t GDay) Validate() {
	if t < 1 || t > 31 {
		helper.Panic(fmt.Sprintf("The day of month can't be less than 1 or bigger than 31. %d is invalid.", t))
	}
}
