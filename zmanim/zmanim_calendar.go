package zmanim

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/zmanim/calculator"
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
	"time"
)

/*
The ZmanimCalendar is a specialized calendar that can calculate sunrise, sunset and Jewish zmanim
(religious times) for prayers and other Jewish religious duties. This class contains the main functionality of the
Zmanim library.
For a much more extensive list of zmanim, use the ComplexZmanimCalendar that
extends this class.
See documentation for the ComplexZmanimCalendar and AstronomicalCalendar for simple examples on using the API.
Elevation based zmanim (even sunrise and sunset) should not be used lekula without the guidance of a posek.
According to Rabbi Dovid Yehudah Bursztyn in his [Zmanim Kehilchasam, 7th edition]: https://www.worldcat.org/oclc/1158574217 chapter 2, section 7 (pages 181-182) - and section 9 (pages 186-187), no zmanim besides sunrise and sunset should use elevation.
However, Rabbi Yechiel
Avrahom Zilber in the <a href="https://hebrewbooks.org/51654">Birur Halacha Vol. 6</a> Ch. 58 Pages
[34]: ]https://hebrewbooks.org/pdfpager.aspx?req=51654&amp;pgnum=42 and
[42]: https://hebrewbooks.org/pdfpager.aspx?req=51654&amp;pgnum=50 is of the opinion that elevation should be
accounted for in <em>zmanim</em> calculations. Related to this, Rabbi Yaakov Karp in
[Shimush Zekeinim]: https://www.worldcat.org/oclc/919472094, Ch. 1, page 17 states that obstructing horizons should
be factored into <em>zmanim</em> calculations. The setting defaults to false (elevation will not be used for
zmanim calculations besides sunrise and sunset), unless the setting is changed to true in
SetUseElevation. This will impact sunrise and sunset-based zmanim such as AstronomicalCalendar.Sunrise,
AstronomicalCalendar.Sunset, SofZmanShmaGRA, alos-based zmanim such as SofZmanShmaMGA
that are based on a fixed offset of sunrise or sunset and zmanim based on a percentage of the day such as
ComplexZmanimCalendar.SofZmanShmaMGA90MinutesZmanis that are based on sunrise and sunset. Even when set to
true it will not impact zmanim that are a degree-based offset of sunrise and sunset, such as
ComplexZmanimCalendar.SofZmanShmaMGA16Point1Degrees or ComplexZmanimCalendar.SofZmanShmaBaalHatanya since
these zmanim are not linked to sunrise or sunset times (the calculations are based on the astronomical definition of
sunrise and sunset calculated in a vacuum with the solar radius above the horizon), and are therefore not impacted by the use
of elevation.
For additional information on the halachic impact of elevation on zmanim see:
- [Zmanei Halacha Lema'aseh]: https://www.nli.org.il/en/books/NNL_ALEPH002542826/NLI 4th edition by
[Rabbi Yedidya Manat]: http://beinenu.com/rabbis/%D7%94%D7%A8%D7%91-%D7%99%D7%93%D7%99%D7%93%D7%99%D7%94-%D7%9E%D7%A0%D7%AA.
- See section 1, pages 11-12 for a very concise write-up, with details in section 2, pages 37 - 63 and 133 - 151.
- [Zmanim Kehilchasam]: https://www.worldcat.org/oclc/1158574217 7th edition, by Rabbi Dovid Yehuda Burstein,  vol 1, chapter 2, pages 95 - 188.
- [Hazmanim Bahalacha]: ]https://www.worldcat.org/oclc/36089452 by Rabbi Chaim Banish , perek 7, pages 53 - 63.
Note: It is important to read the technical notes on top of the calculator.AstronomicalCalculator documentation before using this code.
I would like to thank [Rabbi Yaakov Shakow]: https://www.worldcat.org/search?q=au%3AShakow%2C+Yaakov,
the author of Luach Ikvei Hayom who spent a considerable amount of time reviewing, correcting and making suggestions on the
documentation in this library.
Disclaimer: I did my best to get accurate results, but please double-check before relying on these
zmanim for halacha lema'aseh.
*/

type ZmanimCalendar interface {
	AstronomicalCalendar
	// IsUseElevation and other getters
	//
	IsUseElevation() bool
	// Hanetz and other ...
	//
	Hanetz() (tm time.Time, ok bool)
	Shkia() (tm time.Time, ok bool)
	Tzais3(degrees dimension.Degrees, offsetMinutes time.Duration, zmanisOffset time.Duration) (tm time.Time, ok bool)
	Tzais() (tm time.Time, ok bool)
	Tzais72() (tm time.Time, ok bool)
	Alos3(degrees dimension.Degrees, offsetMinutes time.Duration, zmanisOffset time.Duration) (tm time.Time, ok bool)
	Alos() (tm time.Time, ok bool)
	Alos72() (tm time.Time, ok bool)
	Chatzos() (tm time.Time, ok bool)
	SofZmanShma(startOfDay time.Time, endOfDay time.Time) time.Time
	SofZmanShmaGRA() (tm time.Time, ok bool)
	SofZmanShmaMGA() (tm time.Time, ok bool)
	SofZmanTfila(startOfDay time.Time, endOfDay time.Time) time.Time
	SofZmanTfilaGRA() (tm time.Time, ok bool)
	SofZmanTfilaMGA() (tm time.Time, ok bool)
	MinchaGedola() (tm time.Time, ok bool)
	MinchaKetana() (tm time.Time, ok bool)
	PlagHamincha() (tm time.Time, ok bool)
	CandleLighting() (tm time.Time, ok bool)
	ShaahZmanis(sunrise time.Time, sunset time.Time) gdt.GMillisecond
	ShaahZmanisGRA() (i gdt.GMillisecond, ok bool)
	ShaahZmanisMGA() (i gdt.GMillisecond, ok bool)
	ShaahZmanisByDegreesAndOffset(degrees dimension.Degrees, offsetMinutes time.Duration) (i gdt.GMillisecond, ok bool)
	IsAssurBemlacha(currentTime time.Time, tzais time.Time, inIsrael bool) bool
	AlosHashachar() (tm time.Time, ok bool)
	// SetUseElevation and other setters
	//
	SetUseElevation(useElevation bool)
}

type zmanimCalendar struct {
	astronomicalCalendar

	/*
		useElevation is elevation factored in for some zmanim.
		Is elevation above sea level calculated for times besides sunrise and sunset. According to Rabbi Dovid Yehuda
		Bursztyn in his [Zmanim Kehilchasam (second edition published in 2007)]: https://www.worldcat.org/oclc/659793988 chapter 2 (pages 186-187)
		no zmanim besides sunrise and sunset should use elevation.
		However, Rabbi Yechiel Avrahom Zilber in the [Birur Halacha Vol. 6]: https://hebrewbooks.org/51654 Ch. 58 Pages
		[34]: https://hebrewbooks.org/pdfpager.aspx?req=51654&amp;pgnum=42 and
		[42]: https://hebrewbooks.org/pdfpager.aspx?req=51654&amp;pgnum=50 is of the opinion that elevation should be
		accounted for in zmanim calculations. Related to this, Rabbi Yaakov Karp in
		[Shimush Zekeinim]: https://www.worldcat.org/oclc/919472094, Ch. 1, page 17 states that obstructing horizons
		should be factored into zmanim calculations. The setting defaults to false (elevation will not be used for
		zmanim calculations), unless the setting is changed to true in SetUseElevation. This will
		impact sunrise and sunset based zmanim such as AstronomicalCalendar.Sunrise, AstronomicalCalendar.Sunset,
		SofZmanShmaGRA, alos based zmanim such as SofZmanShmaMGA that are based on a
		 fixed offset of sunrise or sunset and zmanim based on a percentage of the day such as
		ComplexZmanimCalendar.SofZmanShmaMGA90MinutesZmanis that are based on sunrise and sunset.
		It will not impact zmanim that are a degree based offset of sunrise and sunset, such as
		ComplexZmanimCalendar.SofZmanShmaMGA16Point1Degrees or ComplexZmanimCalendar.SofZmanShmaBaalHatanya.
	*/
	useElevation bool

	/*
		candleLightingOffset is the Shabbos candle lighting offset,
		the offset in minutes before SeaLevelSunset (default is 18 gdt.GMinuteF64).
		It is used in calculating Shabbos candle lighting time.
		The default time used is 18 gdt.GMinuteF64 before SeaLevelSunset.
		Some calendars use gdt.GMinuteF64,
		while the custom in Jerusalem is to use a 40 gdt.GMinuteF64 offset.
		Please check the local custom for candle lighting time.
		see CandleLighting
	*/
	candleLightingOffset gdt.GMinuteF64
}

func newZmanimCalendar() *zmanimCalendar {
	return &zmanimCalendar{candleLightingOffset: 18}
}

func NewZmanimCalendar(gDateTime gdt.GDateTime, geoLocation calculator.GeoLocation, astronomicalCalculator calculator.AstronomicalCalculator) ZmanimCalendar {
	t := newZmanimCalendar()

	t.initAstronomicalCalendar(gDateTime, geoLocation, astronomicalCalculator)

	return t
}

func (t *zmanimCalendar) IsUseElevation() bool {
	return t.useElevation
}

func (t *zmanimCalendar) SetUseElevation(useElevation bool) {
	t.useElevation = useElevation
}

/**
 * This method will return {@link #getSeaLevelSunrise() sea level sunrise} if {@link #isUseElevation()} is false (the
 * default), or elevation adjusted {@link AstronomicalCalendar#getSunrise()} if it is true. This allows relevant <em>zmanim</em>
 * in this and extending classes (such as the {@link ComplexZmanimCalendar}) to automatically adjust to the elevation setting.
 *
 * @return {@link #getSeaLevelSunrise()} if {@link #isUseElevation()} is false (the default), or elevation adjusted
 *         {@link AstronomicalCalendar#getSunrise()} if it is true.
 * @see com.kosherjava.zmanim.AstronomicalCalendar#getSunrise()
 */
func (t *zmanimCalendar) elevationAdjustedSunrise() (tm time.Time, ok bool) {
	if t.IsUseElevation() {
		return t.Sunrise()
	}
	return t.SeaLevelSunrise()
}

/**
 * This method will return {@link #getSeaLevelSunrise() sea level sunrise} if {@link #isUseElevation()} is false (the default),
 * or elevation adjusted {@link AstronomicalCalendar#getSunrise()} if it is true. This allows relevant <em>zmanim</em>
 * in this and extending classes (such as the {@link ComplexZmanimCalendar}) to automatically adjust to the elevation setting.
 *
 * @return {@link #getSeaLevelSunset()} if {@link #isUseElevation()} is false (the default), or elevation adjusted
 *         {@link AstronomicalCalendar#getSunset()} if it is true.
 * @see com.kosherjava.zmanim.AstronomicalCalendar#getSunset()
 */
func (t *zmanimCalendar) elevationAdjustedSunset() (tm time.Time, ok bool) {
	if t.IsUseElevation() {
		return t.Sunset()
	}
	return t.SeaLevelSunset()
}

func (t *zmanimCalendar) Tzais3(degrees dimension.Degrees, offsetMinutes time.Duration, zmanisOffset time.Duration) (tm time.Time, ok bool) {
	var sunsetForDegrees time.Time

	if degrees == 0 {
		sunsetForDegrees, ok = t.elevationAdjustedSunset()
		if !ok {
			return time.Time{}, false
		}
	} else {
		sunsetForDegrees, ok = t.SunsetOffsetByDegrees(calculator.GeometricZenith + degrees)
		if !ok {
			return time.Time{}, false
		}
	}

	if zmanisOffset != 0 {
		return t.offsetByMinutesZmanis(sunsetForDegrees, zmanisOffset)
	} else {
		return sunsetForDegrees.Add(offsetMinutes * time.Minute), true
	}

}

/*
Tzais is a method that returns tzais (nightfall), when the sun is zenith8Point5 8.5 deg below the
calculator.GeometricZenith 90 deg after AstronomicalCalendar.Sunset, a time that Rabbi Meir
Posen in his the [Ohr Meir]: https://www.worldcat.org/oclc/29283612 calculated that 3 small
stars are visible, which is later than the required 3 medium stars.
See the zenith8Point5 constant.
the method return the time.Time of nightfall.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle where the sun may not reach
low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
ComplexZmanimCalendar.TzaisGeonim8Point5Degrees that returns an identical time to this generic tzais
*/
func (t *zmanimCalendar) Tzais() (tm time.Time, ok bool) {
	return t.Tzais3(8.5, 0, 0)
}

/*
AlosHashachar returns alos (dawn) based on the time when the sun is zenith16Point1 16.1 deg below the
eastern calculator.GeometricZenith before AstronomicalCalendar.Sunrise. This is based on the
calculation that the time between dawn and sunrise (and sunset to nightfall) is 72 minutes, the time that is
takes to walk 4 mil at 18 minutes a mil [Rambam]: https://en.wikipedia.org/wiki/Maimonides and others.
The sun's position at 72 minutes before AstronomicalCalendar.Sunrise in Jerusalem
on the [around the equinox / - equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/ is 16.1 deg below calculator.GeometricZenith.
see
- ComplexZmanimCalendar.Alos16Point1Degrees
the method	return the time.Time of dawn.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle where the sun may not reach
low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) AlosHashachar() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith16Point1)
}

func (t *zmanimCalendar) Alos3(degrees dimension.Degrees, offsetMinutes time.Duration, zmanisOffset time.Duration) (tm time.Time, ok bool) {
	var sunriseForDegrees time.Time

	if degrees == 0 {
		sunriseForDegrees, ok = t.elevationAdjustedSunrise()
		if !ok {
			return time.Time{}, false
		}
	} else {
		sunriseForDegrees, ok = t.SunriseOffsetByDegrees(calculator.GeometricZenith + degrees)
		if !ok {
			return time.Time{}, false
		}
	}

	if zmanisOffset != 0 {
		return t.offsetByMinutesZmanis(sunriseForDegrees, -zmanisOffset)
	} else {
		return sunriseForDegrees.Add(-offsetMinutes * time.Minute), true
	}

}

func (t *zmanimCalendar) Alos() (tm time.Time, ok bool) {
	return t.Alos3(16.1, 0, 0)
}

/*
Alos72 is the method to return alos (dawn) calculated using 72 minutes before AstronomicalCalendar.Sunrise or
AstronomicalCalendar.SeaLevelSunrise (depending on the IsUseElevation setting).
This time is based on the time to walk the distance of 4 Mil at 18 minutes a Mil.
The 72 minutes time (but not the concept of fixed minutes) is based on the opinion that the time of the Neshef (twilight between
dawn and sunrise) does not vary by the time of year or location but depends on the time it takes to walk the
distance of 4 Mil.
The method return the time.Time representing the time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) Alos72() (tm time.Time, ok bool) {
	return t.Alos3(0, 72, 0)
}

/*
Chatzos returns chatzos (midday) following most opinions that chatzos is the midpoint
between AstronomicalCalendar.SeaLevelSunrise and AstronomicalCalendar.SeaLevelSunset.
A day starting at <em>alos</em> and ending at tzais using the same time or degree offset will also return
the same time. The returned value is identical to AstronomicalCalendar.sunTransit.
In reality due to lengthening or shortening of day,
this is not necessarily the exact midpoint of the day, but it is very close.
The method return the time.Time of chatzos.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) Chatzos() (tm time.Time, ok bool) {
	return t.SunTransit()
}

/*
SofZmanShma is a generic method for calculating the latest zman krias shema (time to recite shema in the morning)
that is 3 * shaos zmaniyos (temporal hours) after the start of the day, calculated using the start and
end of the day passed to this method.
The time from the start of day to the end of day are divided into 12 shaos zmaniyos (temporal hours),
and the latest zman krias shema is calculated as 3 of those shaos zmaniyos after the beginning of
the day.
As an example, passing AstronomicalCalendar.Sunrise and AstronomicalCalendar.Sunset
or
AstronomicalCalendar.SeaLevelSunrise and AstronomicalCalendar.SeaLevelSunset
(depending on the IsUseElevation) to this method will return sof zman krias shema according to the opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
startOfDay is the start of day for calculating zman krias shema.
This can be sunrise or any alos passed to this method.
endOfDay is the end of day for calculating zman krias shema.
This can be sunset or any tzais passed to this method.
The method	return the time.Time of the latest zman shema based on the start and end of day times passed to this
method.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day
a year, where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) SofZmanShma(startOfDay time.Time, endOfDay time.Time) time.Time {
	return t.ShaahZmanisBasedZman(startOfDay, endOfDay, 3)
}

/*
SofZmanShmaGRA eturns the latest zman krias shema (time to recite shema in the morning)
that is 3 * ShaahZmanisGRA (solar hours) after AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the IsUseElevation setting),
according to the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
The day is calculated from AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunset
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the IsUseElevation setting).
see
- SofZmanShma
- ShaahZmanisGRA
- IsUseElevation
- ComplexZmanimCalendar.SofZmanShmaBaalHatanya
The method  return the time.Time of the latest zman shema according to the GRA.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See the detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) SofZmanShmaGRA() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(elevationAdjustedSunrise, elevationAdjustedSunset), true
}

/*
SofZmanShmaMGA returns the latest zman krias shema (time to recite shema in the morning)
that is 3 * ShaahZmanisMGA shaos zmaniyos (solar hours) after Alos72,
according to the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern.
The day is calculated from 72 minutes before AstronomicalCalendar.SeaLevelSunrise to 72 minutes after
AstronomicalCalendar.SeaLevelSunrise
or
from 72 minutes before AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset (depending on the IsUseElevation setting).
The method	return the time.Time of the latest zman shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- SofZmanShma
- ComplexZmanimCalendar.ShaahZmanis72Minutes
- ZmanimCalendar.Alos72
- ComplexZmanimCalendar.SofZmanShmaMGA72Minutes
*/
func (t *zmanimCalendar) SofZmanShmaMGA() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	tzais72, ok := t.Tzais72()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos72, tzais72), true
}

/*
Tzais72 returns the tzais (nightfall) based on the opinion of Rabbeinu Tam that
tzais hakochavim is calculated as 72 minutes, the time it takes to walk 4 Mil at 18 minutes
a Mil.
According to the [Machtzis Hashekel]: https://en.wikipedia.org/wiki/Samuel_Loew in
Orach Chaim 235:3, the [Pri Megadim]: https://en.wikipedia.org/wiki/Joseph_ben_Meir_Teomim in Orach
Chaim 261:2 (see the Biur Halacha) and others (see Hazmanim Bahalacha 17:3 and 17:5) the 72 minutes are standard
clock minutes any time of the year in any location.
Depending on the IsUseElevation setting a 72 minutes offset from
either AstronomicalCalendar.Sunset or AstronomicalCalendar.SeaLevelSunset is used.
see
- ComplexZmanimCalendar.Tzais16Point1Degrees
The method return the time.Time representing 72 minutes after sunset.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise,and one where it does not set, ok is false will be returned
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) Tzais72() (tm time.Time, ok bool) {
	return t.Tzais3(0, 72, 0)
}

/*
CandleLighting is a method to return candle lighting time,
calculated as CandleLightingOffset minutes before AstronomicalCalendar.SeaLevelSunset.
This will return the time for any day of the week, since it can be
used to calculate candle lighting time for Yom Tov (mid-week holidays) as well.
Elevation adjustments are intentionally not performed by this method,
but you can calculate it by passing the elevation adjusted sunset to zmanim.timeOffset.
The method return candle lighting time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) CandleLighting() (tm time.Time, ok bool) {
	seaLevelSunset, ok := t.SeaLevelSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(seaLevelSunset, -t.CandleLightingOffset().ToMilliseconds()), true
}

/*
SofZmanTfila is a generic method for calculating the latest zman tfilah
(time to recite the morning prayers) that is 4 * shaos zmaniyos (temporal hours)
after the start of the day, calculated using the start and end of the day passed to this method.
The time from the start of day to the end of day are divided into 12 shaos zmaniyos (temporal hours),
sof zman tfila is calculated as 4 of those shaos zmaniyos after the beginning of the day.
As an example, passing AstronomicalCalendar.Sunrise and AstronomicalCalendar.Sunset
or
AstronomicalCalendar.SeaLevelSunrise and AstronomicalCalendar.SeaLevelSunset (depending on the IsUseElevation setting) to this method will return
zman tfilah according to the opinion of the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
startOfDay is the start of day for calculating zman tfilah.
This can be sunrise or any alos passed to this method.
endOfDay is the end of day for calculating zman tfilah.
This can be sunset or any tzais passed to this method.
The method return the time.Time of the latest zman tfilah based on the start and end of day times passed
to this method.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the {@link AstronomicalCalendar} documentation.
*/
func (t *zmanimCalendar) SofZmanTfila(startOfDay time.Time, endOfDay time.Time) time.Time {
	return t.ShaahZmanisBasedZman(startOfDay, endOfDay, 4)
}

/*
SofZmanTfilaGRA returns the latest zman tfila (time to recite shema in the morning)
that is 4 * ShaahZmanisGRA shaos zmaniyos (solar hours)
after AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the IsUseElevation setting),
according to the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
The day is calculated from AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunrise
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the IsUseElevation setting).
see
- SofZmanTfila
- ShaahZmanisGRA
- ComplexZmanimCalendar.SofZmanTfilaBaalHatanya
The method return the time.Time of the latest zman tfilah.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) SofZmanTfilaGRA() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(elevationAdjustedSunrise, elevationAdjustedSunset), true
}

/*
SofZmanTfilaMGA method returns the latest zman tfila (time to recite shema in the morning)
that is 4 * ShaahZmanisMGA shaos zmaniyos (solar hours) after Alos72,
according to the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern.
The day is calculated from 72 minutes before AstronomicalCalendar.SeaLevelSunrise to 72 minutes after AstronomicalCalendar.SeaLevelSunset
or
from 72 minutes before AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the IsUseElevation setting).
return the time.Time of the latest zman tfila.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- SofZmanTfila
- ShaahZmanisMGA
- Alos72
*/
func (t *zmanimCalendar) SofZmanTfilaMGA() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	tzais72, ok := t.Tzais72()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos72, tzais72), true
}

/*
MinchaGedola2 is a generic method for calculating the latest mincha gedola (the earliest time to recite the mincha  prayers)
that is 6.5 * shaos zmaniyos (temporal hours) after the start of the day,
calculated using the start and end of the day passed to this method.
The time from the start of day to the end of day are divided into 12 shaos zmaniyos (temporal hours),
and mincha gedola is calculated as 6.5 of those shaos zmaniyos after the beginning of the day.
As an example, passing AstronomicalCalendar.Sunrise and AstronomicalCalendar.Sunset
or
AstronomicalCalendar.SeaLevelSunrise and AstronomicalCalendar.SeaLevelSunset
(depending on the IsUseElevation setting) to this method will return mincha gedola according to the opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
startOfDay is the start of day for calculating Mincha gedola</em>.
This can be sunrise or any alos passed to this method.
endOfDay is the end of day for calculating Mincha gedola.
This can be sunset or any tzais passed to this method.
The method return the time.Time of the time of Mincha gedola based on the start and end of day times
passed to this method.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) MinchaGedola2(startOfDay time.Time, endOfDay time.Time) time.Time {
	return t.ShaahZmanisBasedZman(startOfDay, endOfDay, 6.5)
}

/*
MinchaGedola returns the latest mincha gedola, the earliest time one can pray mincha that is
6.5 * ShaahZmanisGRA shaos zmaniyos (solar hours) after AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the IsUseElevation setting), according to the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
Mincha gedola is the earliest time one can pray mincha.
The Ramba"m is of the opinion that it is better to delay mincha until MinchaKetana while the
Ra"sh, Tur, GRA and others are of the opinion that mincha can be prayed lechatchila starting at mincha gedola.
The day is calculated from AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunset
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the IsUseElevation setting).
see
- MinchaGedola
- ShaahZmanisGRA
- MinchaKetana
- ComplexZmanimCalendar.MinchaGedolaBaalHatanya
the method return the time.Time of the time of mincha gedola.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok mis false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) MinchaGedola() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaGedola2(elevationAdjustedSunrise, elevationAdjustedSunset), true
}

/*
SamuchLeMinchaKetana is a generic method for calculating samuch lemincha ketana,
/ near mincha ketana time that is half an hour before MinchaKetana or 9 * shaos zmaniyos (temporal hours)
after the start of the day, calculated using the start and end of the day passed to this method.
The time from the start of day to the end of day are divided into 12 shaos zmaniyos (temporal hours),
and samuch lemincha ketana is calculated as 9 of those shaos zmaniyos after the beginning of the day.
For example, passing AstronomicalCalendar.Sunrise and AstronomicalCalendar.Sunset or AstronomicalCalendar.SeaLevelSunrise
and AstronomicalCalendar.SeaLevelSunset (depending on the IsUseElevation setting) to this method will return
samuch lemincha ketana according to the opinion of the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
startOfDay is the start of day for calculating samuch lemincha ketana.
This can be sunrise or any alos passed to this method.
endOfDay the end of day for calculating samuch lemincha ketana.
This can be sunset or any tzais passed to this method.
The method return the time.Time of the time of Mincha ketana based on the start and end of day times
passed to this method.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- ComplexZmanimCalendar.SamuchLeMinchaKetanaGRA
- ComplexZmanimCalendar.SamuchLeMinchaKetana16Point1Degrees
- ComplexZmanimCalendar.SamuchLeMinchaKetana72Minutes
*/
func (t *zmanimCalendar) SamuchLeMinchaKetana(startOfDay time.Time, endOfDay time.Time) time.Time {
	return t.ShaahZmanisBasedZman(startOfDay, endOfDay, 9)
}

/*
MinchaKetana2 is a generic method for calculating mincha ketana,
(the preferred time to recite the mincha prayers in the opinion of the [Rambam]: https://en.wikipedia.org/wiki/Maimonides and others) that is
9.5 * shaos zmaniyos (temporal hours) after the start of the day, calculated using the start and end of the day passed to this method.
The time from the start of day to the end of day are divided into 12 shaos zmaniyos (temporal hours),
and mincha ketana is calculated as 9.5 of those shaos zmaniyos after the beginning of the day.
As an example, passing
AstronomicalCalendar.Sunrise and AstronomicalCalendar.Sunset
or
AstronomicalCalendar.SeaLevelSunrise and AstronomicalCalendar.SeaLevelSunset
(depending on the IsUseElevation setting) to this method will return mincha ketana according to the opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
startOfDay is the start of day for calculating Mincha ketana.
This can be sunrise or any alos passed to this method.
endOfDay is the end of day for calculating Mincha ketana.
This can be sunset or any tzais passed to this method.
The method return the time.Time of the time of Mincha ketana based on the start and end of day times
passed to this method.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) MinchaKetana2(startOfDay time.Time, endOfDay time.Time) time.Time {
	return t.ShaahZmanisBasedZman(startOfDay, endOfDay, 9.5)
}

/*
MinchaKetana method returns mincha ketana, preferred the earliest time to pray mincha in the
opinion of the [Rambam]: https://en.wikipedia.org/wiki/Maimonides and others,
that is 9.5 ShaahZmanisGRA shaos zmaniyos (solar hours) after
AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the IsUseElevation setting),
according to the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
For more information on this see the documentation on MinchaGedola.
The day is calculated from
AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunset
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the IsUseElevation setting).
see
  - MinchaKetana
  - ShaahZmanisGRA
  - MinchaGedola
  - ComplexZmanimCalendar.MinchaKetanaBaalHatanya

The method return the time.Time of the time of mincha ketana.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) MinchaKetana() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaKetana2(elevationAdjustedSunrise, elevationAdjustedSunset), true
}

/*
PlagHamincha2 is a generic method for calculating plag hamincha (the earliest time that Shabbos can be started)
that is 10.75 hours after the start of the day, (or 1.25 hours before the end of the day) based on the start and end of
the day passed to the method.
The time from the start of day to the end of day are divided into 12 shaos zmaniyos (temporal hours),
and plag hamincha is calculated as 10.75 of those shaos zmaniyos after the beginning of the day.
As an example, passing
AstronomicalCalendar.Sunrise and AstronomicalCalendar.Sunset
or
AstronomicalCalendar.SeaLevelSunrise and AstronomicalCalendar.SeaLevelSunset
(depending on the IsUseElevation setting)
to this method will return plag mincha according to the opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
startOfDay is the start of day for calculating plag.
This can be sunrise or any alos passed to this method.
endOfDay is the end of day for calculating plag. This can be sunset or any tzais passed to this method.
The method return the time.Time of the time of plag hamincha based on the start and end of day times
passed to this method.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) PlagHamincha2(startOfDay time.Time, endOfDay time.Time) time.Time {
	return t.ShaahZmanisBasedZman(startOfDay, endOfDay, 10.75)
}

/*
PlagHamincha method returns plag hamincha, that is 10.75 * ShaahZmanisGRA shaos zmaniyos (solar hours) after
AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the IsUseElevation setting),
according to the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon
Plag hamincha is the earliest time that Shabbos can be started.
The day is calculated from
AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunset
or
AstronomicalCalendar.Sunrise() to AstronomicalCalendar.Sunset
(depending on the IsUseElevation)
see
- PlagHamincha
  - ComplexZmanimCalendar.PlagHaminchaBaalHatanya

The method return the time.Time of the time of plag hamincha.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) PlagHamincha() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(elevationAdjustedSunrise, elevationAdjustedSunset), true
}

/*
ShaahZmanisGRA is a method that returns a shaah zmanis temporalHour according to
the opinion of the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
This calculation divides the day based on the opinion of the GRA that the day runs from
AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunset
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the IsUseElevation setting).
The day is split into 12 equal parts with each one being a shaah zmanis.
This method is similar to temporalHour, but can account for elevation.
The method return the gdt.GMillisecond length of a shaah zmanis calculated from sunrise to sunset.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - temporalHour
  - AstronomicalCalendar.SeaLevelSunrise
  - AstronomicalCalendar.SeaLevelSunset
  - ComplexZmanimCalendar.ShaahZmanisBaalHatanya
*/
func (t *zmanimCalendar) ShaahZmanisGRA() (i gdt.GMillisecond, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return 0, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return 0, false
	}
	return temporalHour(elevationAdjustedSunrise, elevationAdjustedSunset), true
}

/*
ShaahZmanisMGA returns a shaah zmanis (temporal hour) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern
based on 72 minutes alos and tzais.
This calculation divides the day that runs from dawn to dusk (for sof zman krias shema
and tfila).
Dawn for this calculation is 72 minutes before AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the IsUseElevation elevation setting)
and
dusk is 72 minutes after AstronomicalCalendar.Sunset or AstronomicalCalendar.SeaLevelSunset
(depending on the IsUseElevation setting).
This day is split into 12 equal parts with each part being a shaah zmanis.
Alternate methods of calculating a shaah zmanis according to the Magen Avraham (MGA)
are available in the subclass ComplexZmanimCalendar.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) ShaahZmanisMGA() (i gdt.GMillisecond, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return 0, false
	}
	tzais72, ok := t.Tzais72()
	if !ok {
		return 0, false
	}
	return temporalHour(alos72, tzais72), true
}

func (t *zmanimCalendar) ShaahZmanisByDegreesAndOffset(degrees dimension.Degrees, offsetMinutes time.Duration) (i gdt.GMillisecond, ok bool) {
	alos3, ok := t.Alos3(degrees, offsetMinutes, 0)
	if !ok {
		return 0, false
	}
	tzais3, ok := t.Tzais3(degrees, offsetMinutes, 0)
	if !ok {
		return 0, false
	}
	return temporalHour(alos3, tzais3), true
}

func (t *zmanimCalendar) CandleLightingOffset() gdt.GMinuteF64 {
	return t.candleLightingOffset
}

/*
IsAssurBemlacha is a utility method to determine if the current Date (date-time) passed in
has a melacha (work) prohibition.
Since there are many opinions on the time of tzais,
the tzais for the current day has to be passed to this class.
Sunset is the classes current day's elevationAdjustedSunset that observes the
IsUseElevation settings.
The JewishCalendar.IsInIsrael will be set by the inIsrael parameter.
currentTime is the current time
tzais the time of tzais
inIsrael whether to use Israel holiday scheme or not
The method return true if melacha is prohibited or false if it is not.
see
- JewishCalendar.IsAssurBemelacha
- JewishCalendar.HasCandleLighting
*/
func (t *zmanimCalendar) IsAssurBemlacha(currentTime time.Time, tzais time.Time, inIsrael bool) bool {
	jewishCalendar := hebrewcalendar.NewJewishCalendar(hebrewcalendar.NewJewishDate2(t.gDateTime.D))
	jewishCalendar.SetInIsrael(inIsrael)

	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return false
	}

	if jewishCalendar.HasCandleLighting() && currentTime.Sub(elevationAdjustedSunset) >= 0 { //erev shabbos, YT or YT sheni and after shkiah
		return true
	}

	if jewishCalendar.IsAssurBemelacha() && currentTime.Sub(tzais) <= 0 { //is shabbos or YT and it is before tzais
		return true
	}

	return false
}

/*
ShaahZmanisBasedZman ia a generic utility method for calculating any shaah zmanis temporalHour based zman
with the day defined as the start and end of day (or night) and the number of shaahos zmaniyos passed to the
method.
This simplifies the code in other methods such as PlagHamincha and cuts down on code replication.
As an example, passing Sunrise and Sunset or SeaLevelSunrise and SeaLevelSunset (depending on the
IsUseElevation) and 10.75 hours to this method will return plag mincha
according to the opinion of the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
startOfDay is the start of day for calculating the zman. This can be sunrise or any alos passed to this method.
endOfDay is the end of day for calculating the zman. This can be sunset or any tzais passed to this method.
hours the number of shaahos zmaniyos (temporal hours) to offset from the start of day
return the time.Time of the time of zman with the shaahos zmaniyos (temporal hours)
in the day offset from the start of day passed to this method.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *zmanimCalendar) ShaahZmanisBasedZman(startOfDay time.Time, endOfDay time.Time, hours gdt.GHourH64) time.Time {
	shaahZmanis := temporalHour(startOfDay, endOfDay)
	return timeOffset(startOfDay, gdt.GMillisecond(float64(shaahZmanis)*float64(hours)))
}

func (t *zmanimCalendar) ShaahZmanis(sunrise time.Time, sunset time.Time) gdt.GMillisecond {
	return temporalHour(sunrise, sunset)
}

func (t *zmanimCalendar) Hanetz() (tm time.Time, ok bool) {
	return t.elevationAdjustedSunrise()
}

func (t *zmanimCalendar) Shkia() (tm time.Time, ok bool) {
	return t.elevationAdjustedSunset()
}

func (t *zmanimCalendar) offsetByMinutesZmanis(tm time.Time, minutes time.Duration) (tmr time.Time, ok bool) {
	shaahZmanisGRA, ok := t.ShaahZmanisGRA()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanisSkew := float64(shaahZmanisGRA) / float64(timeutil.HourMillis)
	return tm.Add(time.Duration(float64(minutes*time.Minute) * shaahZmanisSkew)), true
}
