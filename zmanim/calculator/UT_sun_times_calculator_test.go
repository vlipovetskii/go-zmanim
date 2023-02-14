package calculator

import (
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"testing"
)

func TestNewSunTimesCalculator(t *testing.T) {
	tag := helper.CurrentFuncName()
	calc := NewSunTimesCalculator()
	calc.CalculatorName()
	assert.Equal(t, tag, "US Naval Almanac Algorithm", calc.CalculatorName())
}
