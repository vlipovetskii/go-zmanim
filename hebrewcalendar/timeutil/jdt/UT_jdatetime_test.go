package jdt

import (
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"testing"
)

func TestNewJDateTime(t *testing.T) {
	tag := helper.CurrentFuncName()
	jdate := NewJDateTime(NewJDate(5781, Nissan, 1), NewMoladTime0())
	assert.Equal(t, tag, JDay(1), jdate.D.Day)
}
