package jdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"strconv"
)

/*
JDate is an internal structure to track the jewish date Year, Month, Day as a whole
and simplify comparison operations on jewish date Year, Month, Day
E.g. NewJDate1(gDate.ToAbsDate())
*/
type JDate struct {
	Year  JYear
	Month JMonth
	Day   JDay
}

func NewJDate(year JYear, month JMonth, day JDay) JDate {
	return JDate{Year: year, Month: month, Day: day}
}

/*
NewJDate1 creates JDate based on gAbsDate.
gAbsDate is the absolute date (days since January 1, 0001, on the Gregorian calendar)
*/
func NewJDate1(gAbsDate gdt.GDay) JDate {
	// Approximation from below
	jYear := JYear((gAbsDate - JewishEpoch) / 366)

	// Search forward for year from the approximation
	for {
		jDateOfTishrei1 := NewJDate(jYear+1, TISHREI, 1)
		if gAbsDate < jDateOfTishrei1.ToAbsDate() {
			break
		}
		jYear++
	}

	var jMonth JMonth
	// Search forward for month from either Tishri or Nisan.
	jDateOfNissan1 := NewJDate(jYear, Nissan, 1)
	if gAbsDate < jDateOfNissan1.ToAbsDate() {
		jMonth = TISHREI // Start at Tishri
	} else {
		jMonth = Nissan // Start at Nisan
	}

	for {
		jDateOfLastDayInMonth := NewJDate(jYear, jMonth, DaysInJewishMonth(jMonth, jYear))
		if gAbsDate <= jDateOfLastDayInMonth.ToAbsDate() {
			break
		}
		jMonth++
	}

	// Calculate the day by subtraction
	jDate1 := NewJDate(jYear, jMonth, 1)
	jDay := JDay(gAbsDate - jDate1.ToAbsDate() + 1)

	return NewJDate(jYear, jMonth, jDay)
}

func (t *JDate) Validate() {
	t.Month.Validate(t.Year)
	t.Day.Validate()

	// reject dates prior to 18 Teves, 3761 (1/1/1 AD). This restriction can be relaxed if the date coding is
	// changed/corrected
	if (t.Year < 3761) || (t.Year == 3761 && (t.Month >= TISHREI && t.Month < Tevet)) || (t.Year == 3761 && t.Month == Tevet && t.Day < 18) {
		helper.Panic("A Jewish date earlier than 18 Teves, 3761 (1/1/1 Gregorian) can't be set. " + strconv.Itoa(int(t.Year)) + ", " + strconv.Itoa(int(t.Month)) + ", " + strconv.Itoa(int(t.Day)) + " is invalid.")
	}

}

/*
DaysInJewishMonth returns the number of days of a Jewish month for a given month and year.
*/
func DaysInJewishMonth(month JMonth, year JYear) JDay {
	if (month == Iyar) || (month == Tammuz) || (month == Elul) || ((month == Heshvan) && !(year.IsHeshvanLong())) || ((month == KISLEV) && year.IsKislevShort()) || (month == Tevet) || ((month == Adar) && !(year.IsLeapJYear())) || (month == AdarII) {
		return 29
	} else {
		return 30
	}
}

/*
DaysSinceStartOfYear returns the number of days from Rosh Hashana of the date passed in, to the full date passed in.
*/
func (t *JDate) DaysSinceStartOfYear() JDay {
	elapsedDays := t.Day

	// Before Tishrei (from Nissan to Tishrei), add days in prior months
	if t.Month < TISHREI {

		// this year before and after Nisan.
		for m := TISHREI; m <= t.Year.LastMonthOfJYear(); m++ {
			elapsedDays += DaysInJewishMonth(m, t.Year)
		}

		for m := Nissan; m < t.Month; m++ {
			elapsedDays += DaysInJewishMonth(m, t.Year)
		}
	} else { // Add days in prior months this year
		for m := TISHREI; m < t.Month; m++ {
			elapsedDays += DaysInJewishMonth(m, t.Year)
		}
	}

	return elapsedDays
}

/*
addDechiyos adds the 4 dechiyos for molad Tishrei. These are:
Lo ADU Rosh - Rosh Hashana can't fall on a Sunday, Wednesday or Friday. If the molad fell on one of these
- days, Rosh Hashana is delayed to the following day.
- Molad Zaken - If the molad of Tishrei falls after 12 noon, Rosh Hashana is delayed to the following day. If
- the following day is ADU, it will be delayed an additional day.
- GaTRaD - If on a non leap year the molad of Tishrei falls on a Tuesday (Ga) on or after 9 hours (T) and 204
- chalakim (TRaD) it is delayed till Thursday (one day delay, plus one day for Lo ADU Rosh)
- BeTuTaKFoT - if the year following a leap year falls on a Monday (Be) on or after 15 hours (Tu) and 589
- chalakim (TaKFoT) it is delayed till Tuesday
*/
func addDechiyos(year JYear, moladDay int32, moladParts int32) JDay {
	roshHashanaDay := moladDay // if no dechiyos

	if (moladParts >= 19440) || (((moladDay % 7) == 2) && (moladParts >= 9924) && !year.IsLeapJYear()) || (((moladDay % 7) == 1) && (moladParts >= 16789) && ((year - 1).IsLeapJYear())) {
		roshHashanaDay += 1 // Then postpone Rosh HaShanah one day
	}

	if ((roshHashanaDay % 7) == 0) || ((roshHashanaDay % 7) == 3) || ((roshHashanaDay % 7) == 5) {
		roshHashanaDay++ // Then postpone it one (more) day
	}

	return JDay(roshHashanaDay)
}

/*
ToAbsDate returns the absolute date of Jewish date.
*/
func (t *JDate) ToAbsDate() gdt.GDay {
	elapsed := t.DaysSinceStartOfYear()

	// add elapsed days this year + Days in prior years + Days elapsed before absolute year 1
	return gdt.GDay(elapsed) + gdt.GDay(t.Year.JewishCalendarElapsedDays()) + JewishEpoch
}

/*
ForwardDay rolls the date, month or year forward by the number of jdt.JDay passed in.
*/
func (t *JDate) ForwardDay(amount JDay) {

	if amount < 1 {
		helper.Panic(fmt.Sprintf("the amount of months %d to forward has to be greater than zero.", amount))
	}

	for i := JDay(1); i <= amount; i++ {

		// Change the Jewish Date
		if t.Day == DaysInJewishMonth(t.Month, t.Year) {
			// if it last day of elul (i.e. last day of Jewish year)
			if t.Month == Elul {
				t.Year++
				t.Month++
				t.Day = 1
			} else if t.Month == t.Year.LastMonthOfJYear() {
				// if it is the last day of Adar, or Adar II as case may be
				t.Month = Nissan
				t.Day = 1
			} else {
				t.Month++
				t.Day = 1
			}
		} else { // if not last date of month
			t.Day++
		}

	}

}

/*
adjustDayToLastDayOfMonth
if 30 is passed for a jewishMonth that only has 29 days (for example by rolling the jewishMonth from a jewishMonth that had 30
days to a jewishMonth that only has 29) set the date to 29th
*/
func (t *JDate) adjustDayToLastDayOfMonth() {
	daysInJewishMonth := DaysInJewishMonth(t.Month, t.Year)
	if t.Day > daysInJewishMonth {
		t.Day = daysInJewishMonth
	}
}

/*
ForwardMonth Forward the Jewish date by the number of months passed in
*/
func (t *JDate) ForwardMonth(amount JMonth) {
	if amount < 1 {
		helper.Panic(fmt.Sprintf("The amount of months %d to forward has to be greater than zero.", amount))
	}

	for i := JMonth(1); i <= amount; i++ {
		if t.Month == Elul {
			t.Month = TISHREI
			t.Year++
		} else if (!t.Year.IsLeapJYear() && t.Month == Adar) || (t.Year.IsLeapJYear() && t.Month == AdarII) {
			t.Month = Nissan
		} else {
			t.Month++
		}
	}

	t.adjustDayToLastDayOfMonth()

}

func (t *JDate) ForwardYear(amount JYear) {
	if amount < 1 {
		helper.Panic(fmt.Sprintf("The amount of years %d to forward has to be greater than zero.", amount))
	}
	t.Year += amount

	t.adjustDayToLastDayOfMonth()
}

/*
BackDay Rolls the date back by 1 day.
*/
func (t *JDate) BackDay(amount JDay) {

	if amount < 1 {
		helper.Panic(fmt.Sprintf("the amount of days %d to backward has to be greater than zero.", amount))
	}

	for i := JDay(1); i <= amount; i++ {

		// change Jewish date
		if t.Day == 1 { // if first day of the Jewish month
			if t.Month == Nissan {
				t.Month = t.Year.LastMonthOfJYear()
			} else if t.Month == TISHREI { // if Rosh Hashana
				t.Year--
				t.Month--
			} else {
				t.Month--
			}
			t.Day = DaysInJewishMonth(t.Month, t.Year)
		} else {
			t.Day--
		}

	}
}

func (t *JDate) BackMonth(amount JMonth) {

	if amount < 1 {
		helper.Panic(fmt.Sprintf("The amount of months %d to backward has to be greater than zero.", amount))
	}

	for i := JMonth(1); i <= amount; i++ {
		if t.Month == TISHREI {
			t.Month = Elul
			t.Year--
			//} else if (!t.Year.IsLeapJYear() && t.Month == Adar) || (t.Year.IsLeapJYear() && t.Month == AdarII) {
			//	t.Month = Nissan
		} else if t.Month == Nissan {
			if t.Year.IsLeapJYear() {
				t.Month = AdarII
			} else {
				t.Month = Adar
			}
		} else {
			t.Month--
		}
	}

	t.adjustDayToLastDayOfMonth()

}

func (t *JDate) BackYear(amount JYear) {
	if amount < 1 {
		helper.Panic(fmt.Sprintf("The amount of years %d to forward has to be greater than zero.", amount))
	}
	t.Year -= amount

	t.adjustDayToLastDayOfMonth()
}

/*
JewishMonthOfYear converts the Nissan based constants used by this class to numeric month starting from
TISHREI. This is required for Molad claculations.
*/
func JewishMonthOfYear(year JYear, month JMonth) JYear {
	isLeapYear := year.IsLeapJYear()

	var m1, m2 int32
	if isLeapYear {
		m1, m2 = 6, 13
	} else {
		m1, m2 = 5, 12
	}
	return JYear((int32(month)+m1)%m2 + 1)
}

/*
ChalakimSinceMoladTohu returns the number of chalakim (parts - 1080 to the hour) from the original hypothetical Molad Tohu to the year
and month passed in.
*/
func ChalakimSinceMoladTohu(year JYear, month JMonth) MoladChalakim64 {
	// Jewish lunar month = 29 days, 12 hours and 793 chalakim.
	// chalakim since Molad Tohu BeHaRaD - 1 day, 5 hours and 204 chalakim.
	monthOfYear := JewishMonthOfYear(year, month)

	monthsElapsed := (235 * ((year - 1) / 19)) + (12 * ((year - 1) % 19)) + ((7*((year-1)%19) + 1) / 19) + (monthOfYear - 1)

	// return chalakim prior to BeHaRaD + number of chalakim since
	return ChalakimMoladTohu + ChalakimPerMonth*MoladChalakim64(monthsElapsed)
}
