package hebrewcalendar

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"github.com/vlipovetskii/go-zmanim/helper/test"
	"testing"
	"time"
)

func TestInitWithNoArgs(t *testing.T) {

	tag := helper.CurrentFuncName()

	today := gdt.NewGDateToday()
	subject := NewJewishDate()

	assert.Equal(t, tag, today, subject.GDate())

	jDate := subject.JDate()

	subject.SetGDate(today)

	assert.Equal(t, tag, jDate, subject.JDate())

}

func TestInitWithModernDateArgs(t *testing.T) {

	tag := helper.CurrentFuncName()

	gregorianDate := gdt.NewGDate(2017, 10, 26)
	subject := NewJewishDate2(gregorianDate)

	assert.Equal(t, tag, subject.GDate(), gregorianDate)

}

func TestInitWithPreGregorianDateArgs(t *testing.T) {

	tag := helper.CurrentFuncName()

	gregorianDate := gdt.NewGDate(1550, 10, 1)
	subject := NewJewishDate2(gregorianDate)

	assert.Equal(t, tag, subject.GDate(), gregorianDate)

	assert.Equal(t, tag, jdt.NewJDate(5311, 7, 11), subject.JDate())
}

func TestInitWithJewishModernDateArgs(t *testing.T) {

	tag := helper.CurrentFuncName()

	subject := NewJewishDate1(jdt.NewJDate(5778, 8, 6))

	assert.Equal(t, tag, jdt.NewJDate(5778, 8, 6), subject.JDate())

	assert.Equal(t, tag, subject.GDate(), gdt.NewGDate(2017, time.October, 26))

}

func TestInitWithJewishPreGregorianDateArgs(t *testing.T) {

	tag := helper.CurrentFuncName()

	subject := NewJewishDate1(jdt.NewJDate(5311, 7, 11))

	assert.Equal(t, tag, jdt.NewJDate(5311, 7, 11), subject.JDate())

	assert.Equal(t, tag, subject.GDate(), gdt.NewGDate(1550, time.October, 1))

}

func TestInitWithMoladArgBeforeMidnight(t *testing.T) {

	tag := helper.CurrentFuncName()

	subject := NewJewishDate3(54700170003)

	assert.Equal(t, tag, jdt.NewJDate(5778, 5, 30), subject.JDate())

	assert.Equal(t, tag, subject.GDate(), gdt.NewGDate(2018, time.August, 11))

	assert.Equal(t, tag, subject.MoladTime(), jdt.NewMoladTime(19, 33, 9))
}

func TestInitWithMoladArgAfterMidnight(t *testing.T) {

	tag := helper.CurrentFuncName()

	subject := NewJewishDate3(54692515673)

	assert.Equal(t, tag, jdt.NewJDate(5778, 7, 30), subject.JDate())

	assert.Equal(t, tag, subject.GDate(), gdt.NewGDate(2017, time.October, 20))

	assert.Equal(t, tag, subject.MoladTime(), jdt.NewMoladTime(12, 12, 17))
}

// test_from_date see TestInitWithModernDateArgs
// test_from_jewish_date see TestInitWithJewishModernDateArgs
// test_from_molad see TestInitWithMoladArgAfterMidnight

func TestResetDate(t *testing.T) {

	tag := helper.CurrentFuncName()

	today := timeutil.NewDateOfToday(nil)
	subject := NewJewishDate2(gdt.NewGDate1(today.Add(24 * 3 * time.Hour)))

	// Resets this date to the current local date.
	subject.SetGDate(gdt.NewGDate1(time.Now()))

	assert.Equal(t, tag, subject.GDate(), gdt.NewGDate1(today))

}

func TestDateAssignment(t *testing.T) {

	tag := helper.CurrentFuncName()

	subject := NewJewishDate()
	gregorianDate := gdt.NewGDate(2017, 10, 26)

	subject.SetGDate(gregorianDate)

	assert.Equal(t, tag, subject.GDate(), gregorianDate)
	assert.Equal(t, tag, subject.JDate(), jdt.NewJDate(5778, 8, 6))
	assert.Equal(t, tag, subject.DayOfWeek(), jdt.JWeekday(5))
	subject.SetGDate(gdt.NewGDate1(gregorianDate.ToTime(nil).Add(24 * 2 * time.Hour)))
	assert.Equal(t, tag, subject.DayOfWeek(), jdt.JWeekday(7))
}

func TestDateAssignmentWithPriorMolad(t *testing.T) {

	tag := helper.CurrentFuncName()

	subject := NewJewishDate3(54692515673)

	subject.SetGDate(gdt.NewGDate(2017, 10, 26))
	assert.Equal(t, tag, subject.MoladTime(), jdt.NewMoladTime(0, 0, 0))
}

func TestSetGregorianDateWithInvalidYear(t *testing.T) {

	// tag := helper.CurrentFuncName()

	subject := NewJewishDate()

	defer assert.Raises(t, helper.CurrentFuncName())()
	subject.SetGDate(gdt.NewGDate(-5, 11, 5))

}

func TestSetGregorianDateWithInvalidMonth(t *testing.T) {

	// tag := helper.CurrentFuncName()

	subject := NewJewishDate()

	defer assert.Raises(t, helper.CurrentFuncName())()
	subject.SetGDate(gdt.NewGDate(2000, 13, 5))

}

func TestSetGregorianDateWithInvalidDay(t *testing.T) {

	// tag := helper.CurrentFuncName()

	subject := NewJewishDate()

	defer assert.Raises(t, helper.CurrentFuncName())()
	subject.SetGDate(gdt.NewGDate(2000, 11, 32))

}

func TestSetJewishDate(t *testing.T) {

	tag := helper.CurrentFuncName()

	subject := NewJewishDate()
	subject.SetJDate(jdt.NewJDate(5778, 8, 6))

	assert.Equal(t, tag, subject.JDate(), jdt.NewJDate(5778, 8, 6))
	assert.Equal(t, tag, subject.GDate(), gdt.NewGDate(2017, 10, 26))

}

func TestSetJewishDateWithInvalidYear(t *testing.T) {

	subject := NewJewishDate()

	defer assert.Raises(t, helper.CurrentFuncName())()
	subject.SetJDate(jdt.NewJDate(3660, 11, 23))

}

func TestSetJewishDateWithInvalidMonth(t *testing.T) {

	subject := NewJewishDate()

	defer assert.Raises(t, helper.CurrentFuncName())()
	subject.SetJDate(jdt.NewJDate(5778, 14, 23))

}

func TestSetJewishDateWithInvalidDay(t *testing.T) {

	subject := NewJewishDate()

	defer assert.Raises(t, helper.CurrentFuncName())()
	subject.SetJDate(jdt.NewJDate(5778, 11, 31))

}

// test_set_jewish_date_resets_month_to_max_month_in_year
// this functionality is not supported by go-zmanim

func TestSetJewishDateWithPriorMolad(t *testing.T) {

	subject := NewJewishDate3(54692515673)

	subject.SetJDate(jdt.NewJDate(5778, 8, 15))
	assert.Equal(t, helper.CurrentFuncName(), subject.MoladTime(), jdt.NewMoladTime0())
}

func TestSetJewishDateWithPassingMolad(t *testing.T) {

	subject := NewJewishDate()

	subject.SetJDate(jdt.NewJDate(5778, 8, 15))
	subject.SetMoladTime(jdt.NewMoladTime(4, 5, 6))
	assert.Equal(t, helper.CurrentFuncName(), subject.MoladTime(), jdt.NewMoladTime(4, 5, 6))
}

// test_forward_with_no_args
// this functionality is not supported by go-zmanim

func TestForwardWithAnIncrementInSameMonth(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 10, 15))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	subject.ForwardJDay(5)
	assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), jdt.NewJDate(5778, 10, 20))
	assert.Equal(t, helper.CurrentFuncName(), subject.DayOfWeek(), jdt.JWeekday(1))
}

func TestForwardWithAnIncrementIntoNextMonth(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 10, 28))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	subject.ForwardJDay(5)
	assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), jdt.NewJDate(5778, 11, 4))
}

func TestForwardWithAnIncrementIntoNextYear(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 6, 28))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	subject.ForwardJDay(5)
	assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), jdt.NewJDate(5779, 7, 4))
}

func TestForwardWithAnIncrementIntoFirstJewishMonth(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5779, 13, 29))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	subject.ForwardJDay(5)
	assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), jdt.NewJDate(5779, 1, 5))
}

func TestForwardWithALargeIncrement(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 6, 28))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	subject.ForwardJDay(505)
	assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), jdt.NewJDate(5780, 10, 29))
}

// test_back_with_no_args
// this functionality is not supported by go-zmanim

func TestBackWithADecrementInSameMonth(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 10, 15))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	// To implement back(self, decrement: int = 1)
	//subject.Back(5)
	//assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), timeutil.NewJDate(5778, 10, 20))
	//assert.Equal(t, helper.CurrentFuncName(), subject.DayOfWeek(), JWeekday(1))
}

func TestBackWithADecrementIntoPreviousMonth(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 10, 15))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	// To implement back(self, decrement: int = 1)
	//subject.Back(5)
	//assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), timeutil.NewJDate(5778, 10, 20))
	//assert.Equal(t, helper.CurrentFuncName(), subject.DayOfWeek(), JWeekday(1))
}

func TestBackWithADecrementIntoPreviousYear(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 10, 15))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	// To implement back(self, decrement: int = 1)
	//subject.Back(5)
	//assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), timeutil.NewJDate(5778, 10, 20))
	//assert.Equal(t, helper.CurrentFuncName(), subject.DayOfWeek(), JWeekday(1))
}

func TestBackWithADecrementIntoLastJewishMonth(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 10, 15))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	// To implement back(self, decrement: int = 1)
	//subject.Back(5)
	//assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), timeutil.NewJDate(5778, 10, 20))
	//assert.Equal(t, helper.CurrentFuncName(), subject.DayOfWeek(), JWeekday(1))
}

func TestBackWithALargeDecrement(t *testing.T) {

	subject := NewJewishDate1(jdt.NewJDate(5778, 10, 15))
	initialGregorian := subject.GDate()

	assert.Equal(t, helper.CurrentFuncName(), subject.GDate(), initialGregorian)

	// To implement back(self, decrement: int = 1)
	//subject.Back(5)
	//assert.Equal(t, helper.CurrentFuncName(), subject.JDate(), timeutil.NewJDate(5778, 10, 20))
	//assert.Equal(t, helper.CurrentFuncName(), subject.DayOfWeek(), JWeekday(1))
}

/*
py-zmanim: test_addition_with_an_integer,
py-zmanim: test_addition_with_a_timedelta,
py-zmanim: test_subtraction_with_an_integer,
py-zmanim: test_subtraction_with_a_timedelta
test overload operators +/-, which is not supported in go-zmanim,
See TestForwardWithAnIncrementInSameMonth, ... which test ForwardJDay instead
*/

/*
py-zmanim: test_subtraction_with_another_jewish_date,
py-zmanim: test_subtraction_with_a_gregorian_date
test overload operators +/-, which is not supported in go-zmanim,
and leverage combinations like return JewishDate(self.gregorian_date - subtrahend) under hood
*/

func TestComparisonWithJewishDate(t *testing.T) {

	date1 := NewJewishDate1(jdt.NewJDate(5778, 8, 5))
	date2 := NewJewishDate1(jdt.NewJDate(5778, 9, 5))
	date3 := NewJewishDate1(jdt.NewJDate(5778, 6, 5))
	date2Copy := NewJewishDate1(jdt.NewJDate(5778, 9, 5))

	tag := helper.CurrentFuncName()
	assert.Equal(t, tag, date2, date2Copy)
	assert.True(t, tag, date2 != date1)
	assert.True(t, tag, date2.CompareTo(date1) > 0)
	assert.True(t, tag, date2.CompareTo(date1) >= 0)
	assert.True(t, tag, date2.CompareTo(date2Copy) >= 0)
	assert.True(t, tag, date2.CompareTo(date3) < 0)
	assert.True(t, tag, date2.CompareTo(date3) <= 0)
	assert.True(t, tag, date2.CompareTo(date2Copy) <= 0)

}

/*
py-zmanim: test_gregorian_year_assignment: subject.gregorian_year = 2016,
py-zmanim: test_gregorian_month_assignment ...,
py-zmanim: test_gregorian_day_assignment ...,
test syntax sugar over self.set_gregorian_date(...) or subject.SetGDate(today)
*/

func TestDaysInGregorianYearForStandardYear(t *testing.T) {
	tag := helper.CurrentFuncName()
	assert.Equal(t, tag, gdt.GYear(2010).DaysInGYear(), gdt.GDay(365))
}

func TestDaysInGregorianYearForLeapYear(t *testing.T) {
	tag := helper.CurrentFuncName()
	assert.Equal(t, tag, gdt.GYear(2012).DaysInGYear(), gdt.GDay(366))
}

/*
py-zmanim: test_days_in_gregorian_year_defaults_to_current_year,
test syntax sugar over days_in_gregorian_year(...) or gdt.DaysInGYear(...)
*/

// ...

func TestIsCheshvanLong(t *testing.T) {
	tag := helper.CurrentFuncName()

	assert.False(t, tag, test.StandardMondayChaseirim.IsHeshvanLong())
	assert.False(t, tag, test.StandardTuesdayKesidran.IsHeshvanLong())
	assert.True(t, tag, test.StandardShabbosShelaimim.IsHeshvanLong())
	assert.False(t, tag, test.LeapThursdayChaseirim.IsHeshvanLong())
	assert.False(t, tag, test.LeapTuesdayKesidran.IsHeshvanLong())
	assert.True(t, tag, test.LeapMondayShelaimim.IsHeshvanLong())
}

/*
py-zmanim: test_is_cheshvan_long_defaults_to_current_year
py-zmanim: test_is_cheshvan_short
py-zmanim: test_is_cheshvan_short_defaults_to_current_year
test syntax sugar over is_cheshvan_long(...) or jdt.IsHeshvanLong(...)
*/

/*
py-zmanim: test_is_kislev_long
py-zmanim: test_is_kislev_long_defaults_to_current_year
py-zmanim: test_is_kislev_short_defaults_to_current_year
test syntax sugar over is_kislev_long(...) or jdt.IsKislevShort(...)
*/

/*
TestIsKislevShort is based on py-zmanim: test_is_kislev_short
*/
func TestIsKislevShort(t *testing.T) {
	tag := helper.CurrentFuncName()

	assert.True(t, tag, test.StandardMondayChaseirim.IsKislevShort())
	assert.False(t, tag, test.StandardTuesdayKesidran.IsKislevShort())
	assert.False(t, tag, test.StandardShabbosShelaimim.IsKislevShort())
	assert.True(t, tag, test.LeapThursdayChaseirim.IsKislevShort())
	assert.False(t, tag, test.LeapTuesdayKesidran.IsKislevShort())
	assert.False(t, tag, test.LeapMondayShelaimim.IsKislevShort())
}

/*
py-zmanim: test_cheshvan_kislev_kviah_defaults_to_current_year
py-zmanim: test_cheshvan_kislev_kviah_defaults_to_current_year
test cheshvan_kislev_kviah, which return str,
while func (t *jewishDate) CheshvanKislevKviah() int32 return int32
*/

/*
py-zmanim: test_kviah
py-zmanim: test_kviah_defaults_to_current_year
test kviah, which return tuple of (cheshvan_kislev_kviah(), ...)
while func (t *jewishDate) CheshvanKislevKviah() int32 return int32
*/

/*
py-zmanim: molad
test _chalakim_since_molad_tohu under hood
*/
func TestChalakimSinceMoladTohu(t *testing.T) {
	tag := helper.CurrentFuncName()

	subject := NewJewishDate3(jdt.ChalakimSinceMoladTohu(5778, 5))

	assert.Equal(t, tag, subject.JDate(), jdt.NewJDate(5778, 5, 1))
	assert.Equal(t, tag, subject.MoladTime(), jdt.NewMoladTime(6, 49, 8))
}

/*
py-zmanim: test_molad_defaults_to_current_month
*/
func TestChalakimSinceMoladTohu1(t *testing.T) {
	tag := helper.CurrentFuncName()

	subject := NewJewishDate1(jdt.NewJDate(5778, 5, 10))

	molad := subject.Molad()

	assert.Equal(t, tag, molad.JDate(), jdt.NewJDate(5778, 5, 1))
	assert.Equal(t, tag, molad.MoladTime(), jdt.NewMoladTime(6, 49, 8))
}
