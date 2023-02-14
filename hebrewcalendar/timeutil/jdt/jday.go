package jdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
)

type JDay int32

func (t JDay) Validate() {
	if t < 1 || t > 30 {
		helper.Panic(fmt.Sprintf("The Jewish day of month can't be < 1 or > 30. %d is invalid.", t))
	}
}
