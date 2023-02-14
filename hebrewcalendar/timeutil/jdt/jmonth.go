package jdt

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/helper"
)

// A JMonth specifies a jewish month of the year (Nissan = 1, ...).
type JMonth int32

const (

	/*
		Nissan value of the month field indicating Nissan, the first numeric month of the year in the Jewish calendar. With the
		year starting at TISHREI, it would actually be the 7th (or 8th in a IsLeapJYear leap
		year) month of the year.
	*/
	Nissan JMonth = 1 + iota

	/*
		Iyar value of the month field indicating Iyar, the second numeric month of the year in the Jewish calendar. With the
		year starting at {@link #TISHREI}, it would actually be the 8th (or 9th in a IsLeapJYear leap
		year) month of the year.
	*/
	Iyar

	/*
		Sivan value of the month field indicating Sivan, the third numeric month of the year in the Jewish calendar. With the
		year starting at {@link #TISHREI}, it would actually be the 9th (or 10th in a IsLeapJYear leap
		year) month of the year.
	*/
	Sivan

	/*
		Tammuz value of the month field indicating Tammuz, the fourth numeric month of the year in the Jewish calendar. With the
		year starting at {@link #TISHREI}, it would actually be the 10th (or 11th in a IsLeapJYear leap
		 year) month of the year.
	*/
	Tammuz

	/*
		Av value of the month field indicating Av, the fifth numeric month of the year in the Jewish calendar. With the year
		starting at TISHREI, it would actually be the 11th (or 12th in a IsLeapJYear leap year)
		month of the year.
	*/
	Av

	/*
		Elul value of the month field indicating Elul, the sixth numeric month of the year in the Jewish calendar. With the
		year starting at TISHREI, it would actually be the 12th (or 13th in a IsLeapJYear leap year) month of the year.
	*/
	Elul

	/*
		TISHREI value of the month field indicating Tishrei, the seventh numeric month of the year in the Jewish calendar. With
		the year starting at this month, it would actually be the 1st month of the year.
	*/
	TISHREI

	/*
		Heshvan value of the month field indicating Cheshvan/marcheshvan, the eighth numeric month of the year in the Jewish
		calendar. With the year starting at TISHREI, it would actually be the 2nd month of the year.
	*/
	Heshvan

	/*
		KISLEV value of the month field indicating Kislev, the ninth numeric month of the year in the Jewish calendar. With the
		year starting at TISHREI, it would actually be the 3rd month of the year.
	*/
	KISLEV

	// Tevet
	/*
		Tevet value of the month field indicating Teves, the tenth numeric month of the year in the Jewish calendar. With the
		year starting at TISHREI, it would actually be the 4th month of the year.
	*/
	Tevet

	// SHEVAT
	/*
		SHEVAT value of the month field indicating Shevat, the eleventh numeric month of the year in the Jewish calendar. With
		the year starting at TISHREI, it would actually be the 5th month of the year.
	*/
	SHEVAT

	/*
	 Adar value of the month field indicating Adar (or Adar I in a IsLeapJYear leap year), the twelfth
	 numeric month of the year in the Jewish calendar. With the year starting at TISHREI, it would actually
	 be the 6th month of the year.
	*/
	Adar

	/*
		AdarII value of the month field indicating Adar II, the leap (intercalary or embolismic) thirteenth (Undecimber) numeric
		 month of the year added in Jewish IsLeapJYear leap year. The leap years are years 3, 6, 8, 11,
		 14, 17 and 19 of a 19-year cycle. With the year starting at TISHREI, it would actually be the 7th month
		 of the year.
	*/
	AdarII
)

func (t JMonth) Validate(year JYear) {
	if t < Nissan || t > year.LastMonthOfJYear() {
		helper.Panic(fmt.Sprintf("The Jewish month has to be between 1 and 12 (or 13 on a leap year). %d is invalid for the year %d.", t, year))
	}
}
