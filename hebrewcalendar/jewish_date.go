package hebrewcalendar

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"time"
)

/*
JewishDate is the base calendar class, that supports maintenance of a GDate instance along with the corresponding JDate.
*/
type JewishDate interface {
	// GDate and other getters
	//
	GDate() gdt.GDate
	GYear() gdt.GYear
	GMonth() time.Month
	GDay() gdt.GDay
	JDate() jdt.JDate
	JYear() jdt.JYear
	JMonth() jdt.JMonth
	JDay() jdt.JDay
	DayOfWeek() jdt.JWeekday
	Molad() JewishDate
	MoladTime() jdt.MoladTime
	MoladHours() jdt.MoladHours
	MoladMinutes() jdt.MoladMinutes
	MoladChalakim() jdt.MoladChalakim
	GAbsDate() gdt.GDay
	LastMonthOfJYear() jdt.JMonth
	DaysInJYear() jdt.JDay
	DaysInJMonth() jdt.JDay
	DaysSinceStartOfJYear() jdt.JDay
	IsLeapJYear() bool
	IsCheshvanLong() bool
	IsKislevShort() bool
	CheshvanKislevKviah() int32
	ChalakimSinceMoladTohu() jdt.MoladChalakim64
	CompareTo(jewishDate JewishDate) int32
	// SetGDate and other setters
	//
	SetGDate(gDate gdt.GDate)
	SetJDate(jDate jdt.JDate)
	SetMoladTime(moladTime jdt.MoladTime)
	SetMoladTime1(chalakim jdt.MoladChalakim)
	// ForwardJDay and other arithmetic methods
	//
	ForwardJDay(amount jdt.JDay)
	ForwardJMonth(amount jdt.JMonth)
	ForwardJYear(amount jdt.JYear)
	BackJDay(amount jdt.JDay)
	BackJMonth(amount jdt.JMonth)
	BackJYear(amount jdt.JYear)
}

type jewishDate struct {
	jDate jdt.JDate

	moladTime jdt.MoladTime

	gDate gdt.GDate

	dayOfWeek jdt.JWeekday

	// gregorianAbsDate the absolute date (days since January 1, 0001, on the Gregorian calendar).
	gregorianAbsDate gdt.GDay
}

func newJewishDate() *jewishDate {
	return &jewishDate{}
}

// NewJewishDate default constructor will set a default date to the current system date.
func NewJewishDate() JewishDate {
	t := newJewishDate()

	t.SetGDate(gdt.NewGDate1(time.Now()))

	return t
}

func NewJewishDate1(jDate jdt.JDate) JewishDate {
	t := newJewishDate()

	t.SetJDate(jDate)

	return t
}

func NewJewishDate2(gDate gdt.GDate) JewishDate {
	t := newJewishDate()

	t.SetGDate(gDate)

	return t
}

/*
NewJewishDate3 constructor that creates a JewishDate based on a molad passed in. The molad would be the number of chalakim/parts
starting at the beginning of Sunday prior to the molad Tohu BeHaRaD (Be = Monday, Ha= 5 hours and Rad =204
chalakim/parts) - prior to the start of the Jewish calendar. BeHaRaD is 23:11:20 on Sunday night(5 hours 204/1080
chalakim after sunset on Sunday evening).
*/
func NewJewishDate3(molad jdt.MoladChalakim64) JewishDate {

	t := newJewishDate()

	t.setDatesFromGDate(gdt.NewGDate2(molad.ToAbsDate()))

	t.SetMoladTime1(jdt.MoladChalakim(int64(molad) % int64(jdt.ChalakimPerDay)))

	return t
}

// MoladTime returns the molad time.
func (t *jewishDate) MoladTime() jdt.MoladTime {
	return t.moladTime
}

// MoladHours returns the molad hours
func (t *jewishDate) MoladHours() jdt.MoladHours {
	return t.moladTime.Hours
}

// MoladMinutes returns the molad minutes.
func (t *jewishDate) MoladMinutes() jdt.MoladMinutes {
	return t.moladTime.Minutes
}

// MoladChalakim returns the molad chalakim/parts.
func (t *jewishDate) MoladChalakim() jdt.MoladChalakim {
	return t.moladTime.Chalakim
}

// SetMoladTime sets the molad time (hours, minutes and chalakim)
func (t *jewishDate) SetMoladTime(moladTime jdt.MoladTime) {
	t.moladTime = moladTime
}

/*
SetMoladTime1 sets the molad time (hours, minutes and chalakim)
based on the number of chalakim since the start of the day.
*/
func (t *jewishDate) SetMoladTime1(chalakim jdt.MoladChalakim) {
	adjustedChalakim := chalakim

	moladHours := jdt.MoladHours(adjustedChalakim / jdt.ChalakimPerHour)

	// molad hours start at 18:00, which means that
	// we cross a secular date boundary if hours are 6 or greater
	if moladHours >= 6 {
		t.ForwardJDay(1)
	}

	t.moladTime.Hours = (moladHours + 18) % 24

	adjustedChalakim = adjustedChalakim - (jdt.MoladChalakim(moladHours) * jdt.ChalakimPerHour)
	t.moladTime.Minutes = jdt.MoladMinutes(adjustedChalakim / jdt.ChalakimPerMinute)
	t.moladTime.Chalakim = adjustedChalakim - jdt.MoladChalakim(t.moladTime.Minutes)*jdt.ChalakimPerMinute
}

// GAbsDate returns the absolute date (days since January 1, 0001, on the Gregorian calendar).
func (t *jewishDate) GAbsDate() gdt.GDay {
	return t.gregorianAbsDate
}

/*
IsLeapJYear returns if the year the calendar is set to is a Jewish leap year.
Years 3, 6, 8, 11, 14, 17 and 19 in the 19 years cycle are leap years.
*/
func (t *jewishDate) IsLeapJYear() bool {
	return t.JYear().IsLeapJYear()
}

/*
ChalakimSinceMoladTohu returns the number of chalakim (parts - 1080 to the hour) from the original hypothetical Molad Tohu to the Jewish
year and month that this Object is set to.
*/
func (t *jewishDate) ChalakimSinceMoladTohu() jdt.MoladChalakim64 {
	return jdt.ChalakimSinceMoladTohu(t.jDate.Year, t.jDate.Month)
}

// DaysInJYear returns the number of days for the current year that the calendar is set to.
func (t *jewishDate) DaysInJYear() jdt.JDay {
	return t.JYear().DaysInJYear()
}

/*
IsCheshvanLong returns if Cheshvan is long (30 days VS 29 days) for the current year that the calendar is set to. The method
name isLong is done since in a Kesidran (ordered) year Cheshvan is short.
*/
func (t *jewishDate) IsCheshvanLong() bool {
	return t.JYear().IsHeshvanLong()
}

/*
IsKislevShort returns if the Kislev is short for the year that this class is set to. The method name isShort is done since in a
Kesidran (ordered) year Kislev is long.
*/
func (t *jewishDate) IsKislevShort() bool {
	return t.JYear().IsKislevShort()
}

/*
CheshvanKislevKviah returns the Cheshvan and Kislev kviah (whether a Jewish year is short, regular or long). It will return
if both cheshvan and kislev are 30 days, {@link #Kesidran} if Cheshvan is 29 days and Kislev is 30 days and {@link #Chaserim} if both are 29 days.
*/
func (t *jewishDate) CheshvanKislevKviah() int32 {
	if t.IsCheshvanLong() && !t.IsKislevShort() {
		return jdt.Shelaimim
	} else if !t.IsCheshvanLong() && t.IsKislevShort() {
		return jdt.Chaserim
	} else {
		return jdt.Kesidran
	}
}

// DaysInJMonth returns the number of days of the Jewish month that the calendar is currently set to.
func (t *jewishDate) DaysInJMonth() jdt.JDay {
	return jdt.DaysInJewishMonth(t.JMonth(), t.JYear())
}

/*
Molad returns the molad for a given year and month. Returns a JewishDate set to the date of the molad
with the MoladHours(), MoladMinutes() and MoladChalakim() set.
In the current implementation, it sets the molad time based on a midnight date rollover. This
means that Rosh Chodesh Adar II, 5771 with a molad of 7 chalakim past midnight on Shabbos 29 Adar I / March 5,
2011 12:00 AM and 7 chalakim, will have the following values: hours: 0, minutes: 0, Chalakim: 7.
*/
func (t *jewishDate) Molad() JewishDate {
	// JewishDate moladDate = new JewishDate(getChalakimSinceMoladTohu());
	moladDate := NewJewishDate3(t.ChalakimSinceMoladTohu())
	return moladDate
}

// DaysSinceStartOfJYear see jdt.DaysSinceStartOfYear
func (t *jewishDate) DaysSinceStartOfJYear() jdt.JDay {
	return t.jDate.DaysSinceStartOfYear()
}

func (t *jewishDate) SetGDate(gDate gdt.GDate) {
	gDate.Validate()
	t.setDatesFromGDate(gDate)

	t.SetMoladTime(jdt.NewMoladTime0())
}

/*
SetJewishDate2 sets the Jewish Date and updates the Gregorian date accordingly.
*/
func (t *jewishDate) SetJewishDate2(jDate jdt.JDate, moladTime jdt.MoladTime) {
	jDate.Validate()
	t.setDatesFromGDate(gdt.NewGDate2(jDate.ToAbsDate()))

	moladTime.Validate()
	t.SetMoladTime(moladTime)
}

/*
setDatesFromGDate sets the hidden internal representation of the Gregorian date,
and updates the Jewish date accordingly.
*/
func (t *jewishDate) setDatesFromGDate(gDate gdt.GDate) {
	t.gDate = gDate

	t.gregorianAbsDate = t.gDate.ToAbsDate()
	t.jDate = jdt.NewJDate1(t.gregorianAbsDate)
	t.setDayOfWeek()
}

/*
setDatesFromJDate sets the hidden internal representation of the Jewish date,
and updates the Gregorian date accordingly.
*/
func (t *jewishDate) setDatesFromJDate(jDate jdt.JDate) {
	t.jDate = jDate

	t.gregorianAbsDate = t.jDate.ToAbsDate()
	t.gDate = gdt.NewGDate2(t.gregorianAbsDate)
	t.setDayOfWeek()
}

func (t *jewishDate) setDayOfWeek() {
	t.dayOfWeek = jdt.JWeekday(t.gregorianAbsDate%7) + 1 // set day of week
}

func (t *jewishDate) String() string {
	return fmt.Sprintf("%v-%v-%v", t.gDate, t.jDate, t.moladTime)
}

/*
ForwardJDay see jdt.ForwardDay
It modifies both the Gregorian and Jewish dates accordingly.
*/
func (t *jewishDate) ForwardJDay(amount jdt.JDay) {
	t.jDate.ForwardDay(amount)
	t.setDatesFromJDate(t.jDate)
}

/*
ForwardJMonth see jdt.ForwardMonth
*/
func (t *jewishDate) ForwardJMonth(amount jdt.JMonth) {
	t.jDate.ForwardMonth(amount)
	t.setDatesFromJDate(t.jDate)
}

func (t *jewishDate) ForwardJYear(amount jdt.JYear) {
	t.jDate.ForwardYear(amount)
	t.setDatesFromJDate(t.jDate)
}

// BackJDay see jdt.BackDay
func (t *jewishDate) BackJDay(amount jdt.JDay) {
	t.jDate.BackDay(amount)
	t.setDatesFromJDate(t.jDate)
}

// BackJMonth see jdt.BackMonth
func (t *jewishDate) BackJMonth(amount jdt.JMonth) {
	t.jDate.BackMonth(amount)
	t.setDatesFromJDate(t.jDate)
}

// BackJYear see jdt.BackYear
func (t *jewishDate) BackJYear(amount jdt.JYear) {
	t.jDate.BackYear(amount)
	t.setDatesFromJDate(t.jDate)
}

/*
Equals indicates whether some other object is "equal to" this one.
*/
func (t *jewishDate) Equals(jewishDate JewishDate) bool {
	return t.gregorianAbsDate == jewishDate.GAbsDate()
}

/*
CompareTo compares two dates as per the compareTo() method in the Comparable interface. Returns a value less than 0 if this
date is "less than" (before) the date, greater than 0 if this date is "greater than" (after) the date, or 0 if
hey are equal.
*/
func (t *jewishDate) CompareTo(jewishDate JewishDate) int32 {
	return int32(t.GAbsDate() - jewishDate.GAbsDate())
}

func (t *jewishDate) GDate() gdt.GDate {
	return t.gDate
}

/*
GMonth Returns the Gregorian month
*/
func (t *jewishDate) GMonth() time.Month {
	return t.gDate.Month
}

/*
GDay returns the Gregorian day of the month.
*/
func (t *jewishDate) GDay() gdt.GDay {
	return t.gDate.Day
}

/*
GYear returns the Gregotian year.
*/
func (t *jewishDate) GYear() gdt.GYear {
	return t.gDate.Year
}

func (t *jewishDate) JDate() jdt.JDate {
	return t.jDate
}

/*
JMonth returns the Jewish month 1-12 (or 13 years in a leap year). The month count starts with 1 for Nisan and goes to
13 for Adar II
*/
func (t *jewishDate) JMonth() jdt.JMonth {
	return t.jDate.Month
}

/*
JDay returns the Jewish day of month.
*/
func (t *jewishDate) JDay() jdt.JDay {
	return t.jDate.Day
}

/*
JYear returns the Jewish year.
*/
func (t *jewishDate) JYear() jdt.JYear {
	return t.jDate.Year
}

/*
DayOfWeek returns the day of the week.
*/
func (t *jewishDate) DayOfWeek() jdt.JWeekday {
	return t.dayOfWeek
}

func (t *jewishDate) SetJDate(jDate jdt.JDate) {
	t.SetJewishDate2(jDate, jdt.NewMoladTime0())
}

func (t *jewishDate) LastMonthOfJYear() jdt.JMonth {
	return t.JYear().LastMonthOfJYear()
}
