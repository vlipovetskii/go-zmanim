package hebrewcalendar

type YomTovIndex int32

const (
	NoYomTov YomTovIndex = 0
	// ErevPesach The 14th day of Nissan, the day before of Pesach (Passover).
	ErevPesach YomTovIndex = 0 + iota
	// Pesach The holiday of Pesach (Passover) on the 15th (and 16th out of Israel) day of Nissan.
	Pesach
	// CholHamoedPesach Chol Hamoed (interim days) of Pesach (Passover)
	CholHamoedPesach
	// PesachSheni Pesach Sheni, the 14th day of Iyar, a minor holiday.
	PesachSheni
	// ErevShavuos Erev Shavuos (the day before Shavuos), the 5th of Sivan
	ErevShavuos
	// Shavuos Shavuos (Pentecost), the 6th of Sivan
	Shavuos
	// SeventeenOfTammuz The fast of the 17th day of Tamuz
	SeventeenOfTammuz
	// TishaBeav The fast of the 9th of Av
	TishaBeav
	// TuBeav The 15th day of Av, a minor holiday
	TuBeav
	// ErevRoshHashana Erev Rosh Hashana (the day before Rosh Hashana), the 29th of Elul
	ErevRoshHashana
	// RoshHashana Rosh Hashana, the first of Tishrei.
	RoshHashana
	// FastOfGedalyah The fast of Gedalyah, the 3rd of Tishrei.
	FastOfGedalyah
	// ErevYomKippur The 9th day of Tishrei, the day before of Yom Kippur.
	ErevYomKippur
	// YomKippur The holiday of Yom Kippur, the 10th day of Tishrei
	YomKippur
	// ErevSuccos The 14th day of Tishrei, the day before of Succos/Sukkos (Tabernacles).
	ErevSuccos
	// Succot The holiday of Succot (Tabernacles), the 15th (and 16th out of Israel) day of Tishrei
	Succot
	// CholHamoedSuccos Chol Hamoed (interim days) of Succos/Sukkos (Tabernacles)
	CholHamoedSuccos
	// HoshanaRabba Hoshana Rabba, the 7th day of Succos/Sukkos that occurs on the 21st of Tishrei.
	HoshanaRabba
	// SheminiAtzeres Shmini Atzeres, the 8th day of Succos/Sukkos is an independent holiday that occurs on the 22nd of Tishrei.
	SheminiAtzeres
	/*
		SimchasTorah Simchas Torah, the 9th day of Succos/Sukkos, or the second day of Shmini Atzeres that is celebrated
		IsInIsrael out of Israel on the 23rd of Tishrei.
	*/
	SimchasTorah
	// CHANUKAH The holiday of Chanukah. 8 days starting on the 25th day Kislev.
	CHANUKAH
	// TenthOfTeves The fast of the 10th day of Teves.
	TenthOfTeves
	// TuBeshvat Tu Bishvat on the 15th day of Shevat, a minor holiday.
	TuBeshvat
	// FastOfEsther The fast of Esther, usually on the 13th day of Adar (or Adar II on leap years). It is earlier on some years.
	FastOfEsther
	// PURIM The holiday of Purim on the 14th day of Adar (or Adar II on leap years).
	PURIM
	// ShushanPurim The holiday of Shushan Purim on the 15th day of Adar (or Adar II on leap years).
	ShushanPurim
	// PurimKatan The holiday of Purim Katan on the 14th day of Adar I on a leap year when Purim is on Adar II, a minor holiday.
	PurimKatan
	/*
		YomHashoah Yom HaShoah, Holocaust Remembrance Day, usually held on the 27th of Nissan. If it falls on a Friday, it is moved
		 to the 26th, and if it falls on a Sunday it is moved to the 28th. A IsUseModernHolidays modern holiday.
	*/
	YomHashoah
	/*
		YomHazikaron Yom HaZikaron, Israeli Memorial Day, held a day before Yom Ha'atzmaut.  A IsUseModernHolidays modern holiday.
	*/
	YomHazikaron

	/*
		YomHaatzmaut Yom Ha'atzmaut, Israel Independence Day, the 5th of Iyar, but if it occurs on a Friday or Saturday, the holiday is
		 moved back to Thursday, the 3rd of 4th of Iyar, and if it falls on a Monday, it is moved forward to Tuesday the
		 6th of Iyar.  A IsUseModernHolidays modern holiday.
	*/
	YomHaatzmaut
	/*
		YomYerushalayim Yom Yerushalayim or Jerusalem Day, on 28 Iyar. A IsUseModernHolidays modern holiday.
	*/
	YomYerushalayim

	// LagBaomer The 33rd day of the Omer, the 18th of Iyar, a minor holiday.
	LagBaomer

	// ShushanPurimKatan The holiday of Purim Katan on the 15th day of Adar I on a leap year when Purim is on Adar II, a minor holiday.
	ShushanPurimKatan

	// IsruChag The day following the last day of Pesach, Shavuos and Sukkos.
	IsruChag
)
