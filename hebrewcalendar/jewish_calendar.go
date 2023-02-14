package hebrewcalendar

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/parsha"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"math"
	"time"
)

/*
The JewishCalendar provides jewish calendar methods.
*/
type JewishCalendar interface {
	// JewishDate and other getters
	//
	JewishDate() JewishDate
	YomTov() YomTovIndex
	IsYomTov() bool
	DayOfChanukah() jdt.JDay
	IsRoshChodesh() bool
	IsMacharChodesh() bool
	DayOfOmer() jdt.JDay
	Parshah() parsha.Parsha
	SpecialShabbos() parsha.Parsha
	IsInIsrael() bool
	IsUseModernHolidays() bool
	IsAssurBemelacha() bool
	IsBirkasHachamah() bool
	HasCandleLighting() bool
	IsTomorrowShabbosOrYomTov() bool
	IsErevYomTovSheni() bool
	IsAseresYemeiTeshuva() bool
	IsCholHamoed() bool
	IsCholHamoedPesach() bool
	IsCholHamoedSuccot() bool
	IsErevYomTov() bool
	IsErevRoshChodesh() bool
	IsYomKippurKatan() bool
	IsBeHaB() bool
	IsTaanis() bool
	IsTaanisBechoros() bool
	IsShabbosMevorchim() bool
	IsIsruChag() bool
	// SofZmanKidushLevanaBetweenMoldos and other getters returning time.Time
	//
	SofZmanKidushLevanaBetweenMoldos() time.Time
	SofZmanKidushLevana15Days() time.Time
	TchilasZmanKidushLevana3Days() time.Time
	MoladAsDate() time.Time
	TchilasZmanKidushLevana7Days() time.Time
	// TekufasTishreiElapsedDays and getters based on TekufasTishreiElapsedDays
	TekufasTishreiElapsedDays() int32
	IsVeseinTalUmatarStartDate() bool
	IsVeseinTalUmatarStartingTonight() bool
	IsVeseinTalUmatarRecited() bool
	IsVeseinBerachaRecited() bool
	// IsMashivHaruachStartDate and other IsMashiv* getters
	IsMashivHaruachStartDate() bool
	IsMashivHaruachEndDate() bool
	IsMashivHaruachRecited() bool
	IsMoridHatalRecited() bool
	// SetInIsrael and other setters
	//
	SetInIsrael(inIsrael bool)
	SetUseModernHolidays(useModernHolidays bool)
}

type jewishCalendar struct {
	jewishDate JewishDate

	/*
	 inIsrael is the calendar set to Israel, where some holidays have different rules.
	 Default is false.
	*/
	inIsrael bool

	/*
		useModernHolidays Is the calendar set to use modern Israeli holidays such as Yom Haatzmaut.
		By default, this value is false.
		The holidays are YomHashoah, YomHazikaron, YomHaatzmaut and YomYerushalayim.
	*/
	useModernHolidays bool
}

func newJewishCalendar() *jewishCalendar {
	return &jewishCalendar{inIsrael: false, useModernHolidays: false}
}

func NewJewishCalendar(jewishDate JewishDate) JewishCalendar {
	t := newJewishCalendar()

	t.jewishDate = jewishDate

	return t
}

func (t *jewishCalendar) IsUseModernHolidays() bool {
	return t.useModernHolidays
}

func (t *jewishCalendar) SetUseModernHolidays(useModernHolidays bool) {
	t.useModernHolidays = useModernHolidays
}

func (t *jewishCalendar) IsInIsrael() bool {
	return t.inIsrael
}

func (t *jewishCalendar) SetInIsrael(inIsrael bool) {
	t.inIsrael = inIsrael
}

func (t *jewishCalendar) JewishDate() JewishDate {
	return t.jewishDate
}

/*
IsBirkasHachamah [Birkas Hachamah]: https://en.wikipedia.org/wiki/Birkat_Hachama" is recited every 28 years based on
Tekufas Shmuel (Julian years) that a year is 365.25 days.
The [Rambam]: https://en.wikipedia.org/wiki/Maimonides
[Hilchos Kiddush Hachodesh 9:3]: http://hebrewbooks.org/pdfpager.aspx?req=14278&amp;st=&amp;pgnum=323"
states that tekufas Nissan of year 1 was 7 days + 9 hours before molad Nissan.
This is calculated as every 10,227 days (28 * 365.25).
*/
func (t *jewishCalendar) IsBirkasHachamah() bool {
	elapsedDays := t.jewishDate.JYear().JewishCalendarElapsedDays()  //elapsed days since molad ToHu
	elapsedDays = elapsedDays + t.jewishDate.DaysSinceStartOfJYear() //elapsed days to the current calendar date

	/* Molad Nissan year 1 was 177 days after molad tohu of Tishrei. We multiply 29.5 days * 6 months from Tishrei
	 * to Nissan = 177. Subtract 7 days since tekufas Nissan was 7 days and 9 hours before the molad as stated in the Rambam
	 * and we are now at 170 days. Because GetJewishCalendarElapsedDays and getDaysSinceStartOfJewishYear use the value for
	 * Rosh Hashana as 1, we have to add 1 day for a total of 171. To this add a day since the tekufah is on a Tuesday,
	 * night, and we push off the bracha to Wednesday AM, resulting in the 172 used in the calculation.
	 */
	if elapsedDays%(28*365.25) == 172 { // 28 years of 365.25 days + the offset from molad tohu mentioned above
		return true
	}

	return false
}

/*
parshaYearType return the type of year for parsha calculations. The algorithm follows the
[Luach Arba'ah Shearim]: http://hebrewbooks.org/pdfpager.aspx?req=14268&amp;st=&amp;pgnum=222 in the Tur Ohr Hachaim.
*/
func (t *jewishCalendar) parshaYearType() int32 {
	roshHashanaDayOfWeek := time.Weekday(t.jewishDate.JYear().JewishCalendarElapsedDays()+1) % 7 // plus one to the original Rosh Hashana of year 1 to get a week starting on Sunday

	if roshHashanaDayOfWeek == 0 {
		roshHashanaDayOfWeek = 7 // convert 0 to 7 for Shabbos for readability
	}

	if t.jewishDate.IsLeapJYear() {
		switch roshHashanaDayOfWeek + 1 {
		case time.Monday:
			{
				if t.jewishDate.IsKislevShort() { //BaCh
					if t.IsInIsrael() {
						return 14
					}
					return 6
				}
				if t.jewishDate.IsCheshvanLong() { //BaSh
					if t.IsInIsrael() {
						return 15
					}
					return 7
				}
			}
		case time.Tuesday: //Gak
			{
				if t.IsInIsrael() {
					return 15
				}
				return 7
			}
		case time.Thursday:
			{
				if t.jewishDate.IsKislevShort() { //HaCh
					return 8
				}
				if t.jewishDate.IsCheshvanLong() { //HaSh
					return 9
				}
			}

		case time.Saturday:
			if t.jewishDate.IsKislevShort() { //ZaCh
				return 10
			}
			if t.jewishDate.IsCheshvanLong() { //ZaSh
				if t.IsInIsrael() {
					return 16
				}
				return 11
			}
		}
	} else { //not a leap year
		switch roshHashanaDayOfWeek {
		case time.Monday:
			if t.jewishDate.IsKislevShort() { //BaCh
				return 0
			}
			if t.jewishDate.IsCheshvanLong() { //BaSh
				if t.IsInIsrael() {
					return 12
				}
				return 1
			}

		case time.Tuesday: //GaK
			{
				if t.IsInIsrael() {
					return 12
				}
				return 1
			}
		case time.Thursday:
			{
				if t.jewishDate.IsCheshvanLong() { //HaSh
					return 3
				}
				if !t.jewishDate.IsKislevShort() { //Hak
					if t.IsInIsrael() {
						return 13
					}
					return 2
				}
			}

		case time.Saturday:
			{
				if t.jewishDate.IsKislevShort() { //ZaCh
					return 4
				}
				if t.jewishDate.IsCheshvanLong() { //ZaSh
					return 5
				}
			}
		}
	}

	return -1 //keep the compiler happy
}

/*
Parshah returns this week's parsha.Parsha if it is Shabbos. It returns parsha.None if the date
- is a weekday or if there is no parsha that week (for example Yom Tov that falls on a Shabbos).
*/
func (t *jewishCalendar) Parshah() parsha.Parsha {
	if t.jewishDate.DayOfWeek() != jdt.Saturday {
		return parsha.None //
	}

	yearType := t.parshaYearType()

	roshHashanaDayOfWeek := t.jewishDate.JYear().JewishCalendarElapsedDays() % 7
	day := roshHashanaDayOfWeek + t.jewishDate.DaysSinceStartOfJYear()

	if yearType >= 0 { // negative year should be impossible, but let's cover all bases
		return parsha.List[yearType][day/7]
	}
	return parsha.None //keep the compiler happy
}

/*
SpecialShabbos returns a parsha.Parsha if the Shabbos is one of the four parshiyos of
parsha.SHKALIM, parsha.ZACHOR, parsha.PARA, parsha.HACHODESH or parsha.None for a regular Shabbos (or any weekday).
*/
func (t *jewishCalendar) SpecialShabbos() parsha.Parsha {
	if t.jewishDate.DayOfWeek() == jdt.Saturday {
		if (t.jewishDate.JMonth() == jdt.SHEVAT && !t.jewishDate.IsLeapJYear()) || (t.jewishDate.JMonth() == jdt.Adar && t.jewishDate.IsLeapJYear()) {
			if t.jewishDate.JDay() == 25 || t.jewishDate.JDay() == 27 || t.jewishDate.JDay() == 29 {
				return parsha.SHKALIM
			}
		}

		if (t.jewishDate.JMonth() == jdt.Adar && !t.jewishDate.IsLeapJYear()) || t.jewishDate.JMonth() == jdt.AdarII {
			if t.jewishDate.JDay() == 1 {
				return parsha.SHKALIM
			}

			if t.jewishDate.JDay() == 8 || t.jewishDate.JDay() == 9 || t.jewishDate.JDay() == 11 || t.jewishDate.JDay() == 13 {
				return parsha.ZACHOR
			}

			if t.jewishDate.JDay() == 18 || t.jewishDate.JDay() == 20 || t.jewishDate.JDay() == 22 || t.jewishDate.JDay() == 23 {
				return parsha.PARA
			}

			if t.jewishDate.JDay() == 25 || t.jewishDate.JDay() == 27 || t.jewishDate.JDay() == 29 {
				return parsha.HACHODESH
			}
		}

		if t.jewishDate.JMonth() == jdt.Nissan && t.jewishDate.JDay() == 1 {
			return parsha.HACHODESH
		}
	}
	return parsha.None
}

/*
YomTov returns an index of the Jewish holiday or fast day for the current day, or a -1 if there is no holiday for this day.
There are constants in this class representing each Yom Tov.
*/
func (t *jewishCalendar) YomTov() YomTovIndex {
	day := t.jewishDate.JDay()
	dayOfWeek := t.jewishDate.DayOfWeek()

	// check by month (starting from Nissan)
	switch t.jewishDate.JMonth() {
	case jdt.Nissan:
		{
			if day == 14 {
				return ErevPesach
			}
			if day == 15 || day == 21 || (!t.inIsrael && (day == 16 || day == 22)) {
				return Pesach
			}
			if day >= 17 && day <= 20 || (day == 16 && t.inIsrael) {
				return CholHamoedPesach
			}
			if (day == 22 && t.inIsrael) || (day == 23 && !t.inIsrael) {
				return IsruChag
			}
			if t.IsUseModernHolidays() && ((day == 26 && dayOfWeek == jdt.Thursday) || (day == 28 && dayOfWeek == jdt.Monday) || (day == 27 && dayOfWeek != jdt.Sunday && dayOfWeek != jdt.Friday)) {
				return YomHashoah
			}
		}
	case jdt.Iyar:
		{
			if t.IsUseModernHolidays() && ((day == 4 && dayOfWeek == jdt.Tuesday) || ((day == 3 || day == 2) && dayOfWeek == jdt.Wednesday) || (day == 5 && dayOfWeek == jdt.Monday)) {
				return YomHazikaron
			}
			// if 5 Iyar falls on Wed, Yom Haatzmaut is that day. If it falls on Friday or Shabbos, it is moved back to
			// Thursday. If it falls on Monday it is moved to Tuesday
			if t.IsUseModernHolidays() && ((day == 5 && dayOfWeek == jdt.Wednesday) || ((day == 4 || day == 3) && dayOfWeek == jdt.Thursday) || (day == 6 && dayOfWeek == jdt.Tuesday)) {
				return YomHaatzmaut
			}
			if day == 14 {
				return PesachSheni
			}
			if day == 18 {
				return LagBaomer
			}
			if t.IsUseModernHolidays() && day == 28 {
				return YomYerushalayim
			}
		}

	case jdt.Sivan:
		{
			if day == 5 {
				return ErevShavuos
			}
			if day == 6 || (day == 7 && !t.inIsrael) {
				return Shavuos
			}
			if (day == 7 && t.inIsrael) || (day == 8 && !t.inIsrael) {
				return IsruChag
			}
		}
	case jdt.Tammuz:
		// push off the fast day if it falls on Shabbos
		if (day == 17 && dayOfWeek != jdt.Saturday) || (day == 18 && dayOfWeek == jdt.Sunday) {
			return SeventeenOfTammuz
		}

	case jdt.Av:
		{
			// if Tisha B'av falls on Shabbos, push off until Sunday
			if (dayOfWeek == jdt.Sunday && day == 10) || (dayOfWeek != jdt.Saturday && day == 9) {
				return TishaBeav
			}
			if day == 15 {
				return TuBeav
			}
		}
	case jdt.Elul:
		if day == 29 {
			return ErevRoshHashana
		}

	case jdt.TISHREI:
		{
			if day == 1 || day == 2 {
				return RoshHashana
			}
			if (day == 3 && dayOfWeek != jdt.Saturday) || (day == 4 && dayOfWeek == jdt.Sunday) {
				// push off Tzom Gedalia if it falls on Shabbos
				return FastOfGedalyah
			}
			if day == 9 {
				return ErevYomKippur
			}
			if day == 10 {
				return YomKippur
			}
			if day == 14 {
				return ErevSuccos
			}
			if day == 15 || (day == 16 && !t.inIsrael) {
				return Succot
			}
			if day >= 17 && day <= 20 || (day == 16 && t.inIsrael) {
				return CholHamoedSuccos
			}
			if day == 21 {
				return HoshanaRabba
			}
			if day == 22 {
				return SheminiAtzeres
			}
			if day == 23 && !t.inIsrael {
				return SimchasTorah
			}
			if (day == 23 && t.inIsrael) || (day == 24 && !t.inIsrael) {
				return IsruChag
			}
		}
	case jdt.KISLEV: // no yomtov in Heshvan
		// if (day == 24) {
		// return EREV_CHANUKAH;
		// } else
		if day >= 25 {
			return CHANUKAH
		}

	case jdt.Tevet:
		{
			if day == 1 || day == 2 || (day == 3 && t.jewishDate.IsKislevShort()) {
				return CHANUKAH
			}
			if day == 10 {
				return TenthOfTeves
			}
		}

	case jdt.SHEVAT:
		if day == 15 {
			return TuBeshvat
		}

	case jdt.Adar:
		{
			if !t.jewishDate.IsLeapJYear() {
				// if 13th Adar falls on Friday or Shabbos, push back to Thursday
				if ((day == 11 || day == 12) && dayOfWeek == jdt.Thursday) || (day == 13 && !(dayOfWeek == jdt.Friday || dayOfWeek == jdt.Saturday)) {
					return FastOfEsther
				}
				if day == 14 {
					return PURIM
				}
				if day == 15 {
					return ShushanPurim
				}
			} else { // else if a leap year
				if day == 14 {
					return PurimKatan
				}
				if day == 15 {
					return ShushanPurimKatan
				}
			}
		}
	case jdt.AdarII:
		{
			// if 13th Adar falls on Friday or Shabbos, push back to Thursday
			if ((day == 11 || day == 12) && dayOfWeek == jdt.Thursday) || (day == 13 && !(dayOfWeek == jdt.Friday || dayOfWeek == jdt.Saturday)) {
				return FastOfEsther
			}
			if day == 14 {
				return PURIM
			}
			if day == 15 {
				return ShushanPurim
			}
		}
	}

	// if we get to this stage, then there are no holidays for the given date return NoYomTov
	return NoYomTov

}

/*
IsYomTov returns true if the current day is Yom Tov. The method returns true even for holidays such as CHANUKAH
and minor ones such as TuBeav and PesachSheni.
Erev YomTov with the except of HoshanaRabba,
erev the second days of Pesach returns false, as do IsTaanis fast days besides YomKippur.
Use IsAssurBemelacha to find the days that have a prohibition of work.
*/
func (t *jewishCalendar) IsYomTov() bool {
	yomTov := t.YomTov()
	if (t.IsErevYomTov() && yomTov == CholHamoedPesach && t.jewishDate.JDay() != 20) || (t.IsTaanis() && yomTov != YomKippur) || yomTov == IsruChag {
		return false
	}

	return t.YomTov() != NoYomTov
}

/*
IsYomTovAssurBemelacha returns true if the Yom Tov day has a melacha (work) prohibition.
This method will return false for a non-Yom Tov day, even if it is Shabbos.
*/
func (t *jewishCalendar) IsYomTovAssurBemelacha() bool {
	yomTov := t.YomTov()
	return yomTov == Pesach || yomTov == Shavuos || yomTov == Succot || yomTov == SheminiAtzeres || yomTov == SimchasTorah || yomTov == RoshHashana || yomTov == YomKippur
}

/*
IsAssurBemelacha returns true if it is Shabbos or if it is a Yom Tov day that has a melacha (work) prohibition.
*/
func (t *jewishCalendar) IsAssurBemelacha() bool {
	return t.jewishDate.DayOfWeek() == jdt.Saturday || t.IsYomTovAssurBemelacha()
}

/*
HasCandleLighting returns true if the day has candle lighting.
This will return true on Erev Shabbos, Erev Yom Tov, the first day of Rosh Hashana and the first days of Yom Tov out of Israel.
It is identical to call IsTomorrowShabbosOrYomTov.
*/
func (t *jewishCalendar) HasCandleLighting() bool {
	return t.IsTomorrowShabbosOrYomTov()
}

/*
IsTomorrowShabbosOrYomTov returns true if tomorrow is Shabbos or Yom Tov.
This will return true on Erev Shabbos, Erev Yom Tov, the first day of Rosh Hashana and erev the first days of Yom Tov
out of Israel. It is identical to call HasCandleLighting.
*/
func (t *jewishCalendar) IsTomorrowShabbosOrYomTov() bool {
	return t.jewishDate.DayOfWeek() == jdt.Friday || t.IsErevYomTov() || t.IsErevYomTovSheni()
}

/*
IsErevYomTovSheni returns true if the day is the second day of Yom Tov.
This impacts the second day of Rosh Hashana everywhere and the second days of Yom Tov in chutz laaretz (out of Israel).
*/
func (t *jewishCalendar) IsErevYomTovSheni() bool {
	return t.jewishDate.JMonth() == jdt.TISHREI && (t.jewishDate.JDay() == 1) || (!t.IsInIsrael() && ((t.jewishDate.JMonth() == jdt.Nissan && (t.jewishDate.JDay() == 15 || t.jewishDate.JDay() == 21)) || (t.jewishDate.JMonth() == jdt.TISHREI && (t.jewishDate.JDay() == 15 || t.jewishDate.JDay() == 22)) || (t.jewishDate.JMonth() == jdt.Sivan && t.jewishDate.JDay() == 6)))
}

/*
IsAseresYemeiTeshuva returns true if the current day is Aseres Yemei Teshuva.
*/
func (t *jewishCalendar) IsAseresYemeiTeshuva() bool {
	return t.jewishDate.JMonth() == jdt.TISHREI && t.jewishDate.JDay() <= 10
}

/*
IsCholHamoed returns true if the current day is Chol Hamoed of Pesach or Succos.
*/
func (t *jewishCalendar) IsCholHamoed() bool {
	return t.IsCholHamoedPesach() || t.IsCholHamoedSuccot()
}

/*
IsCholHamoedPesach returns true if the current day is Chol Hamoed of Pesach.
*/
func (t *jewishCalendar) IsCholHamoedPesach() bool {
	return t.YomTov() == CholHamoedPesach
}

/*
IsCholHamoedSuccot returns true if the current day is Chol Hamoed of Succos.
*/
func (t *jewishCalendar) IsCholHamoedSuccot() bool {
	return t.YomTov() == CholHamoedSuccos
}

/*
IsErevYomTov returns true if the current day is Erev Yom Tov.
The method returns true for Erev - Pesach
(first and last days), Shavuos, RoshHashana, YomKippur, Succot and HoshanaRabba.
*/
func (t *jewishCalendar) IsErevYomTov() bool {
	yomTov := t.YomTov()
	return yomTov == ErevPesach || yomTov == ErevShavuos || yomTov == ErevRoshHashana || yomTov == ErevYomKippur || yomTov == ErevSuccos || yomTov == HoshanaRabba || (yomTov == CholHamoedPesach && t.jewishDate.JDay() == 20)
}

/*
IsErevRoshChodesh returns true if the current day is Erev RoshChodesh.
Returns false for Erev RoshHashana.
*/
func (t *jewishCalendar) IsErevRoshChodesh() bool {
	// Erev Rosh Hashana is not Erev Rosh Chodesh.
	return t.jewishDate.JDay() == 29 && t.jewishDate.JMonth() != jdt.Elul
}

/*
IsYomKippurKatan returns true if the current day is Yom Kippur Katan.
returns false for Erev RoshHashana,
Erev RoshChodesh Heshvan, jdt.Teves and Iyyar. If Erev RoshChodesh occurs
on a Friday or Shabbos, YomKippurKatan is moved back to Thursday.
The day before Rosh Chodesh (moved to Thursday if Rosh Chodesh is on a Friday or Shabbos) in most months.
*/
func (t *jewishCalendar) IsYomKippurKatan() bool {
	dayOfWeek := t.jewishDate.DayOfWeek()
	month := t.jewishDate.JMonth()
	day := t.jewishDate.JDay()
	if month == jdt.Elul || month == jdt.TISHREI || month == jdt.KISLEV || month == jdt.Nissan {
		return false
	}

	if day == 29 && dayOfWeek != jdt.Friday && dayOfWeek != jdt.Saturday {
		return true
	}

	if (day == 27 || day == 28) && dayOfWeek == jdt.Thursday {
		return true
	}

	return false
}

/*
IsBeHaB the Monday, Thursday and Monday after the first Shabbos after IsRoshChodesh RoshChodesh
jdt.Heshvan and jdt.Iyar are [BeHaB]: https://outorah.org/p/41334/ days.
If the last Monday of Iyar's BeHaB coincides with PesachSheni,
the method currently considers it both PesachSheni and BeHaB.
- As seen in an Ohr Sameach  article on the subject
[The  - unknown Days: BeHaB Vs. PesachSheni ]: https://ohr.edu/this_week/insights_into_halacha/9340
there are some customs that delay the day to various points in the future.

It is the Monday, Thursday and Monday after the first Shabbos after Rosh Chodesh Cheshvan and Iyarem
are BeHaB days.
*/
func (t *jewishCalendar) IsBeHaB() bool {
	dayOfWeek := t.jewishDate.DayOfWeek()
	month := t.jewishDate.JMonth()
	day := t.jewishDate.JDay()

	if month == jdt.Heshvan || month == jdt.Iyar {
		if (dayOfWeek == jdt.Monday && day > 4 && day < 18) || (dayOfWeek == jdt.Thursday && day > 7 && day < 14) {
			return true
		}
	}

	return false
}

/*
IsTaanis return true if the day is a Taanis (fast day). Return true for 17 of Tammuz, Tisha B'Av,
Yom Kippur, Fast of Gedalyah, 10 of Teves and the Fast of Esther.
*/
func (t *jewishCalendar) IsTaanis() bool {
	yomTov := t.YomTov()
	return yomTov == SeventeenOfTammuz || yomTov == TishaBeav || yomTov == YomKippur || yomTov == FastOfGedalyah || yomTov == TenthOfTeves || yomTov == FastOfEsther
}

/*
IsTaanisBechoros return true if the day is Taanis Bechoros (on Erev Pesach).
It will return true for the 14th of Nissan if it is not on Shabbos, or if the 12th of Nissan occurs on a Thursday.
*/
func (t *jewishCalendar) IsTaanisBechoros() bool {
	day := t.jewishDate.JDay()
	dayOfWeek := t.jewishDate.DayOfWeek()

	// on 14 Nissan unless that is Shabbos where the fast is moved back to Thursday
	return t.jewishDate.JMonth() == jdt.Nissan && ((day == 14 && dayOfWeek != jdt.Saturday) || (day == 12 && dayOfWeek == jdt.Thursday))
}

/*
DayOfChanukah returns the day of Chanukah or -1 if it is not Chanukah.
*/
func (t *jewishCalendar) DayOfChanukah() jdt.JDay {
	day := t.jewishDate.JDay()

	if t.IsChanukah() {
		if t.jewishDate.JMonth() == jdt.KISLEV {
			return day - 24
		} else { // teves
			// return isKislevShort() ? day + 5 : day + 6
			if t.jewishDate.IsKislevShort() {
				return day + 5
			} else {
				return day + 6
			}
		}
	} else {
		return -1
	}
}

/*
IsChanukah returns true if the current day is one of the 8 days of Chanukah.
*/
func (t *jewishCalendar) IsChanukah() bool {
	return t.YomTov() == CHANUKAH
}

/*
IsIsruChag returns true if the current day is Isru Chag. The method returns true for the day following Pesach
Shavuos and Succos. It utilizes {@see #getInIsrael()} to return the proper date.
*/
func (t *jewishCalendar) IsIsruChag() bool {
	return t.YomTov() == IsruChag
}

/*
IsRoshChodesh returns if the day is RoshChodesh. RoshHashana will return false
Rosh Chodesh, the new moon on the first day of the Jewish month,
and the 30th day of the previous month in the case of a month with 30 days.
*/
func (t *jewishCalendar) IsRoshChodesh() bool {
	// Rosh Hashana is not rosh chodesh. Elul never has 30 days
	return (t.jewishDate.JDay() == 1 && t.jewishDate.JMonth() != jdt.TISHREI) || t.jewishDate.JDay() == 30
}

/*
IsMacharChodesh returns if the day is Shabbos and Sunday is RoshChodesh.
*/
func (t *jewishCalendar) IsMacharChodesh() bool {
	return t.jewishDate.DayOfWeek() == jdt.Saturday && (t.jewishDate.JDay() == 30 || t.jewishDate.JDay() == 29)
}

/*
IsShabbosMevorchim returns if the day is Shabbos Mevorchim.
*/
func (t *jewishCalendar) IsShabbosMevorchim() bool {
	return t.jewishDate.DayOfWeek() == jdt.Saturday && t.jewishDate.JDay() >= 23 && t.jewishDate.JDay() <= 29 && t.jewishDate.JMonth() != jdt.Elul
}

/*
DayOfOmer returns the int value of the Omer day or -1 if the day is not in the Omer.
*/
func (t *jewishCalendar) DayOfOmer() jdt.JDay {
	var omer jdt.JDay = -1 // not a day of the Omer
	month := t.jewishDate.JMonth()
	day := t.jewishDate.JDay()

	// if Nissan and second day of Pesach and on
	if month == jdt.Nissan && day >= 16 {
		omer = day - 15
		// if Iyar
	} else if month == jdt.Iyar {
		omer = day + 15
		// if Sivan and before Shavuos
	} else if month == jdt.Sivan && day < 6 {
		omer = day + 44
	}

	return omer
}

/*
MoladAsDate returns the molad in Standard Time in Yerushalayim as a Date. The traditional calculation uses local time.
This method subtracts 20.94 minutes (20 minutes and 56.496 seconds) from the local time (of Har Habayis
with a longitude of 35.2354&deg; is 5.2354&deg; away from the %15 timezone longitude) to get to standard time. This
method intentionally uses standard time and not daylight savings time. Java will implicitly format the time to the
default (or set) Timezone.
*/
func (t *jewishCalendar) MoladAsDate() time.Time {
	/*	molad := t.Molad()
		locationName := "Jerusalem, Israel"

		latitude := 31.778 // Har Habayis latitude
		longitude := 35.2354 // Har Habayis longitude

		// The raw molad Date (point in time) must be generated using standard time. Using "Asia/Jerusalem" timezone will result in the time
		// being incorrectly off by an hour in the summer due to DST. Proper adjustment for the actual time in DST will be done by the date
		// formatter class used to display the Date.
		// TimeZone yerushalayimStandardTZ = TimeZone.getTimeZone("GMT+2");
		yerushalayimStandardTZ, err := time.LoadLocation("GMT+2")
		if err != nil {
			panic(err)
		}

		geo := util.NewGeoLocation1(locationName, latitude, longitude, yerushalayimStandardTZ)
	*/

	/*
		Calendar cal = Calendar.getInstance(geo.getTimeZone());

		cal.clear()

		double moladSeconds = molad.getMoladChalakim() * 10 / (double) 3;

		cal.set(molad.getGregorianYear(), molad.getGregorianMonth(), molad.getGregorianDayOfMonth(), molad.getMoladHours(), molad.getMoladMinutes(), (int) moladSeconds);
		cal.set(Calendar.MILLISECOND, (int) (1000 * (moladSeconds - (int) moladSeconds)));
		// subtract local time difference of 20.94 minutes (20 minutes and 56.496 seconds) to get to Standard time
		cal.add(Calendar.MILLISECOND, -1 * (int) geo.getLocalMeanTimeOffset());

		return cal.getTime();
	*/
	panic("implement me")
}

/*
TchilasZmanKidushLevana3Days returns the earliest time of Kiddush Levana calculated as 3 days after the molad. This method returns the time
even if it is during the day when Kiddush Levana can't be said. Callers of this method should consider
displaying the next tzais if the zman is between alos and tzais.
*/
func (t *jewishCalendar) TchilasZmanKidushLevana3Days() time.Time {
	/*Date molad = getMoladAsDate();
	  Calendar cal = Calendar.getInstance();
	  cal.setTime(molad);
	  cal.add(Calendar.HOUR, 72); // 3 days after the molad
	  return cal.getTime();
	*/
	panic("implement me")
}

/*
TchilasZmanKidushLevana7Days returns the earliest time of Kiddush Levana calculated as 7 days after the molad as mentioned
by the [Mechaber]: http://en.wikipedia.org/wiki/Yosef_Karo".
See the [Bach's]: http://en.wikipedia.org/wiki/Yoel_Sirkis" opinion on this time.
This method returns the time
- even if it is during the day when Kiddush Levana can't be said. Callers of this method should consider
- displaying the next tzais if the zman is between alos and tzais.
*/
func (t *jewishCalendar) TchilasZmanKidushLevana7Days() time.Time {
	/*Date molad = getMoladAsDate();
	  Calendar cal = Calendar.getInstance();
	  cal.setTime(molad);
	  cal.add(Calendar.HOUR, 168); // 7 days after the molad
	  return cal.getTime();
	*/
	panic("implement me")
}

/*
SofZmanKidushLevanaBetweenMoldos returns the latest time of Kiddush Levana according to the
[Maharil's]: http://en.wikipedia.org/wiki/Yaakov_ben_Moshe_Levi_Moelin opinion that it is calculated as
halfway between molad and molad. This adds half the 29 days, 12 hours and 793 chalakim
time between molad and molad (14 days, 18 hours, 22 minutes and 666 milliseconds) to the month's
molad. This method returns the time even if it is during the day when Kiddush Levana can't be
recited. Callers of this method should consider displaying alos before this time if the zman is
between alos and tzais.
*/
func (t *jewishCalendar) SofZmanKidushLevanaBetweenMoldos() time.Time {
	/*Date molad = getMoladAsDate();
	  Calendar cal = Calendar.getInstance();
	  cal.setTime(molad);
	  // add half the time between molad and molad (half of 29 days, 12 hours and 793 chalakim (44 minutes, 3.3
	  // seconds), or 14 days, 18 hours, 22 minutes and 666 milliseconds). Add it as hours, not days, to avoid
	  // DST/ST crossover issues.
	  cal.add(Calendar.HOUR, (24 * 14) + 18);
	  cal.add(Calendar.MINUTE, 22);
	  cal.add(Calendar.SECOND, 1);
	  cal.add(Calendar.MILLISECOND, 666);
	  return cal.getTime();
	*/
	panic("implement me")
}

/*
SofZmanKidushLevana15Days returns the latest time of Kiddush Levana calculated as 15 days after the molad. This is the
opinion brought down in the Shulchan Aruch (Orach Chaim 426). It should be noted that some opinions hold that
the [Rema]: http://en.wikipedia.org/wiki/Moses_Isserles" who brings down the the
[Maharil's]: http://en.wikipedia.org/wiki/Yaakov_ben_Moshe_Levi_Moelin opinion of calculating it as
SofZmanKidushLevanaBetweenMoldos(), half-way between molad and molad is of the
opinion of the Mechaber as well. Also see the Aruch Hashulchan. For additional details on the subject, See Rabbi
Dovid Heber's very detailed writeup in Siman Daled (chapter 4) of
[Shaarei Zmanim]: http://www.worldcat.org/oclc/461326125". This method returns the time even if it is during
the day when Kiddush Levana can't be said. Callers of this method should consider displaying alos
before this time if the zman is between alos and tzais.
*/
func (t *jewishCalendar) SofZmanKidushLevana15Days() time.Time {
	/*Date molad = getMoladAsDate();
	  Calendar cal = Calendar.getInstance();
	  cal.setTime(molad);
	  cal.add(Calendar.HOUR, 24 * 15); //15 days after the molad. Add it as hours, not days, to avoid DST/ST crossover issues.
	  return cal.getTime();
	*/
	panic("implement me")
}

/*
TekufasTishreiElapsedDays returns the elapsed days since Tekufas Tishrei. This uses Tekufas Shmuel (identical to the
[Julian Year]: https://en.wikipedia.org/wiki/Julian_year_(astronomy) with a solar year length of 365.25 days).
The notation used below is D = days, H = hours and C = chalakim.
[Molad]: "https://en.wikipedia.org/wiki/Molad"
BaHaRad was 2D,5H,204C or 5H,204C from the start of Rosh Hashana year 1. For molad
Nissan add 177D, 4H and 438C (6 * 29D, 12H and 793C), or 177D,9H,642C after Rosh Hashana year 1.
Tekufas Nissan was 7D, 9H and 642C before molad Nissan according to the Rambam, or 170D, 0H and
0C after Rosh Hashana year 1. Tekufas Tishrei was 182D and 3H (365.25 / 2) before tekufas
Nissan, or 12D and 15H before Rosh Hashana of year 1. Outside of Israel we start reciting Tal
Umatar in Birkas Hashanim from 60 days after tekufas Tishrei. The 60 days include the day of
the tekufah and the day we start reciting Tal Umatar. 60 days from the tekufah == 47D and 9H
from Rosh Hashana year 1.
*/
func (t *jewishCalendar) TekufasTishreiElapsedDays() int32 {
	// Days since Rosh Hashana year 1. Add 1/2 day as the first tekufas tishrei was 9 hours into the day. This allows all
	// 4 years of the secular leap year cycle to share 47 days. Truncate 47D and 9H to 47D for simplicity.
	days := float64(t.jewishDate.JYear().JewishCalendarElapsedDays()) + float64(t.jewishDate.DaysSinceStartOfJYear()-1) + 0.5
	// days of completed solar years
	solar := float64(t.jewishDate.JYear()-1) * 365.25

	return int32(math.Floor(days - solar))
}

/*
IsVeseinTalUmatarStartDate returns if it is the Jewish day (starting the evening before) to start reciting Vesein Tal Umatar
Livracha (Sheailas Geshamim). In Israel this is the 7th day of Marcheshvan. Outside
Israel recitation starts on the evening of December 4th (or 5th if it is the year before a civil leap year)
in the 21st century and shifts a day forward every century not evenly divisible by 400. This method will
return true if vesein tal umatar on the current Jewish date that starts on the previous night, so
Dec 5/6 will be returned by this method in the 21st century. vesein tal umatar is not recited on
Shabbos and the start date will be delayed a day when the start day is on a Shabbos (this
can only occur out of Israel).
*/
func (t *jewishCalendar) IsVeseinTalUmatarStartDate() bool {
	if t.inIsrael {
		// The 7th Cheshvan can't occur on Shabbos, so always return true for 7 Cheshvan
		if t.jewishDate.JMonth() == jdt.Heshvan && t.jewishDate.JDay() == 7 {
			return true
		}
	} else {
		if t.jewishDate.DayOfWeek() == jdt.Saturday { //Not recited on Friday night
			return false
		}
		tekufasTishreiElapsedDays := t.TekufasTishreiElapsedDays()
		if t.jewishDate.DayOfWeek() == jdt.Sunday { // When starting on Sunday, it can be the start date or delayed from Shabbos
			return tekufasTishreiElapsedDays == 48 || tekufasTishreiElapsedDays == 47
		} else {
			return tekufasTishreiElapsedDays == 47
		}
	}

	return false // keep the compiler happy
}

/*
IsVeseinTalUmatarStartingTonight returns true if tonight is the first night to start reciting Vesein Tal Umatar Livracha (
Sheailas Geshamim). In Israel this is the 7th day of Marcheshvan (so the 6th will return
true). Outside Israel recitation starts on the evening of December 4th (or 5th if it is the year before a
civil leap year) in the 21st century and shifts a day forward every century not evenly divisible by 400.
Vesein tal umatar is not recited on Shabbos and the start date will be delayed a day when
the start day is on a Shabbos (this can only occur out of Israel).
*/
func (t *jewishCalendar) IsVeseinTalUmatarStartingTonight() bool {
	if t.inIsrael {
		// The 7th Cheshvan can't occur on Shabbos, so always return true for 6 Cheshvan
		if t.jewishDate.JMonth() == jdt.Heshvan && t.jewishDate.JDay() == 6 {
			return true
		}
	} else {
		if t.jewishDate.DayOfWeek() == jdt.Friday { //Not recited on Friday night
			return false
		}
		tekufasTishreiElapsedDays := t.TekufasTishreiElapsedDays()
		if t.jewishDate.DayOfWeek() == jdt.Saturday { // When starting on motzai Shabbos, it can be the start date or delayed from Friday night
			return tekufasTishreiElapsedDays == 47 || tekufasTishreiElapsedDays == 46
		} else {
			return tekufasTishreiElapsedDays == 46
		}
	}
	return false
}

/*
IsVeseinTalUmatarRecited returns if Vesein Tal Umatar Livracha (Sheailas Geshamim) is recited. This will return
true for the entire season, even on Shabbos when it is not recited.
*/
func (t *jewishCalendar) IsVeseinTalUmatarRecited() bool {
	if t.jewishDate.JMonth() == jdt.Nissan && t.jewishDate.JDay() < 15 {
		return true
	}
	if t.jewishDate.JMonth() < jdt.Heshvan {
		return false
	}
	if t.inIsrael {
		return t.jewishDate.JMonth() != jdt.Heshvan || t.jewishDate.JDay() >= 7
	} else {
		return t.TekufasTishreiElapsedDays() >= 47
	}
}

/*
IsVeseinBerachaRecited returns if Vesein Beracha is recited. It is recited from 15 Nissan to the point that
IsVeseinTalUmatarRecited vesein tal umatar is recited.
*/
func (t *jewishCalendar) IsVeseinBerachaRecited() bool {
	return !t.IsVeseinTalUmatarRecited()
}

/*
IsMashivHaruachStartDate returns if the date is the start date for reciting Mashiv Haruach Umorid Hageshem. The date is 22 Tishrei.
*/
func (t *jewishCalendar) IsMashivHaruachStartDate() bool {
	return t.jewishDate.JMonth() == jdt.TISHREI && t.jewishDate.JDay() == 22
}

/*
IsMashivHaruachEndDate returns if the date is the end date for reciting Mashiv Haruach Umorid Hageshem. The date is 15 Nissan.
*/
func (t *jewishCalendar) IsMashivHaruachEndDate() bool {
	return t.jewishDate.JMonth() == jdt.Nissan && t.jewishDate.JDay() == 15
}

/*
IsMashivHaruachRecited returns if Mashiv Haruach Umorid Hageshem is recited. This period starts on 22 Tishrei and ends
on the 15th day of Nissan.
*/
func (t *jewishCalendar) IsMashivHaruachRecited() bool {
	startDate := NewJewishDate1(jdt.NewJDate(t.jewishDate.JYear(), jdt.TISHREI, 22))
	endDate := NewJewishDate1(jdt.NewJDate(t.jewishDate.JYear(), jdt.Nissan, 15))

	return t.jewishDate.CompareTo(startDate) > 0 && t.jewishDate.CompareTo(endDate) < 0
}

/*
IsMoridHatalRecited returns if Morid Hatal (or the lack of reciting Mashiv Haruach following nussach Ashkenaz) is recited.
This period starts on 22 Tishrei and ends on the 15th day of Nissan.
*/
func (t *jewishCalendar) IsMoridHatalRecited() bool {
	return !t.IsMashivHaruachRecited() || t.IsMashivHaruachStartDate() || t.IsMashivHaruachEndDate()
}
