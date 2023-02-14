package jdt

type JYear int32

/*
IsLeapJYear returns if the year is a Jewish leap year. Years 3, 6, 8, 11, 14, 17 and 19 in 19 years cycle are leap years.
*/
func (t JYear) IsLeapJYear() bool {
	return ((7*t)+1)%19 < 7
}

/*
LastMonthOfJYear returns the last month of a given Jewish year.
This will be 12 on a non IsLeapJYear or 13 on a leap year.
*/
func (t JYear) LastMonthOfJYear() JMonth {
	if t.IsLeapJYear() {
		return AdarII
	} else {
		return Adar
	}
}

/*
DaysInJYear returns the number of days for a given Jewish year.
*/
func (t JYear) DaysInJYear() JDay {
	return (t + 1).JewishCalendarElapsedDays() - t.JewishCalendarElapsedDays()
}

/*
IsHeshvanLong returns if Heshvan is long in a given Jewish year.
*/
func (t JYear) IsHeshvanLong() bool {
	return t.DaysInJYear()%10 == 5
}

/*
IsKislevShort returns if Kislev is short (29 days VS 30 days) in a given Jewish year. The method name isShort is done since in
a Kesidran (ordered) year Kislev is long. ND+ER
*/
func (t JYear) IsKislevShort() bool {
	return t.DaysInJYear()%10 == 3
}

/*
JewishCalendarElapsedDays returns the number of days elapsed from the Sunday prior to the start of the Jewish calendar to the mean
conjunction of Tishri of the Jewish year.
*/
func (t JYear) JewishCalendarElapsedDays() JDay {
	chalakimSince := ChalakimSinceMoladTohu(t, TISHREI)
	moladDay := chalakimSince / ChalakimPerDay
	moladParts := int32(chalakimSince - moladDay*ChalakimPerDay)

	// delay Rosh Hashana for the 4 dechiyos
	return addDechiyos(t, int32(moladDay), moladParts)
}
