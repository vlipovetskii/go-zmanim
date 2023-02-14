package jdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
)

type MoladMinutes int32

func (t MoladMinutes) Validate() {
	if t < 0 || t > 59 {
		helper.Panic(fmt.Sprintf("Minutes < 0 or > 59 can't be set. %d is invalid.", t))
	}
}
