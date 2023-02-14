package jdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
)

type MoladHours int32

func (t MoladHours) Validate() {
	if t < 0 || t > 23 {
		helper.Panic(fmt.Sprintf("Hour < 0 or > 23 can't be set. %d is invalid.", t))
	}
}
