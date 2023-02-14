package gdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"time"
)

/*
ValidateMonth validates a Gregorian month for validity.
month the Gregorian month number to validate. It will enforce that the month is between 1 - 12
*/
func ValidateMonth(month time.Month) {
	if month < 1 || month > 12 {
		helper.Panic(fmt.Sprintf("The Gregorian month has to be between 1 - 12. %d is invalid.", month))
	}
}
