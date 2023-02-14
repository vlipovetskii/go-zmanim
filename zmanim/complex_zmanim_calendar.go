package zmanim

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/zmanim/calculator"
	"time"
)

/*
ComplexZmanimCalendar extends ZmanimCalendar and provides many more zmanim than available in the ZmanimCalendar.
The basis for most zmanim in this class are from the sefer [Yisroel Vehazmanim]: https://hebrewbooks.org/9765 by
[Rabbi Yisrael Dovid Harfenes]: https://en.wikipedia.org/wiki/Yisroel_Dovid_Harfenes.
As an example of the number of different zmanim made available by this class, there are methods to return
- 18 different calculations for alos (dawn),
- 18 for plag hamincha
- and 29 for tzais available in this API.
The real power of this API is the ease in calculating zmanim that are not part of the library.
The methods for zmanim calculations not present in this class, or it's superclass ZmanimCalendar are contained in the
AstronomicalCalendar, the base class of the calendars in our API since they are generic methods for calculating
time based on degrees or time before or after AstronomicalCalendar.Sunrise and AstronomicalCalendar.Sunset and are of interest
for calculation beyond zmanim calculations. Here are some examples.

TO-DO port comments with examples in java to golang
*/
type ComplexZmanimCalendar interface {
	ZmanimCalendar
	// ShaahZmanis19Point8Degrees and other ShaahZmanis*
	//
	ShaahZmanis19Point8Degrees() (i gdt.GMillisecond, ok bool)
	// Alos19Point8Degrees and other Alos*
	//
	Alos19Point8Degrees() (tm time.Time, ok bool)
	// Misheyakir11Point5Degrees and other Misheyakir*
	//
	Misheyakir11Point5Degrees() (tm time.Time, ok bool)
	// AteretTorahSunsetOffset and other AteretTorah*
	//
	AteretTorahSunsetOffset() gdt.GMinuteF64
	// SofZmanShmaMGA90MinutesZmanis and other SofZmanShma*
	//
	SofZmanShmaMGA90MinutesZmanis() (tm time.Time, ok bool)
	SofZmanShmaMGA16Point1Degrees() (tm time.Time, ok bool)
	// SofZmanShmaBaalHatanya and other SofZmanShmaBaal*
	//
	SofZmanShmaBaalHatanya() (tm time.Time, ok bool)
	// Tzais19Point8Degrees and other Tzais*
	//
	Tzais19Point8Degrees() (tm time.Time, ok bool)
}

type complexZmanimCalendar struct {
	zmanimCalendar

	/*
		ateretTorahSunsetOffset is the offset in gdt.GMinuteF64 (defaults to 40) after sunset used to calculate tzais
		based on calculations of Chacham Yosef Harari-Raful of Yeshivat Ateret Torah.
		see
		- TzaisAteretTorah
		- AteretTorahSunsetOffset
		- SetAteretTorahSunsetOffset
	*/
	ateretTorahSunsetOffset gdt.GMinuteF64
}

func newComplexZmanimCalendar() *complexZmanimCalendar {
	return &complexZmanimCalendar{ateretTorahSunsetOffset: 40}
}

func NewComplexZmanimCalendar(gDateTime gdt.GDateTime, geoLocation calculator.GeoLocation, astronomicalCalculator calculator.AstronomicalCalculator) ComplexZmanimCalendar {
	t := newComplexZmanimCalendar()

	t.initAstronomicalCalendar(gDateTime, geoLocation, astronomicalCalculator)

	return t
}

/*
ShaahZmanis19Point8Degrees is the ethod to return a shaah zmanis (temporal hour) calculated using a 19.8 deg dip.
This calculation divides the day based on the opinion
of the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern that the day runs from dawn to dusk.
Dawn for this calculation is when the sun is 19.8 deg below the eastern geometric horizon before sunrise.
Dusk for this is when the sun is 19.8 deg below the western geometric horizon after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle
where the sun may not reach low enough below the horizon for this calculation, a {@link Long#MIN_VALUE}
will be returned. See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) ShaahZmanis19Point8Degrees() (i gdt.GMillisecond, ok bool) {
	alos19Point8Degrees, ok := t.Alos19Point8Degrees()
	if !ok {
		return 0, false
	}
	tzais19Point8Degrees, ok := t.Tzais19Point8Degrees()
	if !ok {
		return 0, false
	}

	return temporalHour(alos19Point8Degrees, tzais19Point8Degrees), true
}

/*
ShaahZmanis18Degrees return a shaah zmanis (temporal hour) calculated using 18 deg dip.
This calculation divides the day based on the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern,
that the day runs from dawn to dusk.
Dawn for this calculation is when the sun is 18 deg;
below the eastern geometric horizon before sunrise.
Dusk for this is when the sun is 18 deg; below the western geometric horizon after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle
and north of the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) ShaahZmanis18Degrees() (i gdt.GMillisecond, ok bool) {
	alos18Degrees, ok := t.Alos18Degrees()
	if !ok {
		return 0, false
	}
	tzais18Degrees, ok := t.Tzais18Degrees()
	if !ok {
		return 0, false
	}
	return temporalHour(alos18Degrees, tzais18Degrees), true
}

/*
ShaahZmanis26Degrees return shaah zmanis (temporal hour) calculated using a dip of 26 deg.
This calculation divides the day based on the opinion of the
[">Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern,
that the day runs from dawn to dusk.
Dawn for this calculation is when the sun is Alos26Degrees 26 deg below the eastern geometric horizon before sunrise.
Dusk for this is when the sun is Tzais26Degrees 26 deg below the western geometric horizon after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
Since zmanim that use this method are extremely late or early and at a point when the sky is a long time past the 18 deg point,
where the darkest point is reached, zmanim that use this should only be used lechumra,
such as delaying the start of nighttime mitzvos.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as northern and
southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- ShaahZmanis120Minutes
*/
func (t *complexZmanimCalendar) ShaahZmanis26Degrees() (i gdt.GMillisecond, ok bool) {
	alos26Degrees, ok := t.Alos26Degrees()
	if !ok {
		return 0, false
	}
	tzais26Degrees, ok := t.Tzais26Degrees()
	if !ok {
		return 0, false
	}
	return temporalHour(alos26Degrees, tzais26Degrees), true
}

/*
ShaahZmanis16Point1Degrees return a shaah zmanis (temporal hour) calculated using a dip of 16.1 deg.
This calculation divides the day based on the opinion that the day runs from dawn to dusk.
Dawn for this calculation is when the sun is 16.1 deg; below the eastern geometric horizon before sunrise
and dusk is when the sun is 16.1 deg below the western geometric horizon after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- Alos16Point1Degrees
- Tzais16Point1Degrees
- SofZmanShmaMGA16Point1Degrees
- SofZmanTfilaMGA16Point1Degrees
- MinchaGedola16Point1Degrees
- MinchaKetana16Point1Degrees
- PlagHamincha16Point1Degrees
*/
func (t *complexZmanimCalendar) ShaahZmanis16Point1Degrees() (i gdt.GMillisecond, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return 0, false
	}
	tzais16Point1Degrees, ok := t.Tzais16Point1Degrees()
	if !ok {
		return 0, false
	}
	return temporalHour(alos16Point1Degrees, tzais16Point1Degrees), true
}

/*
ShaahZmanis60Minutes return a shaah zmanis (solar hour) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern.
This calculation divides the day based on the opinion of the MGA that the day runs from dawn to dusk.
Dawn for this calculation is 60 minutes before sunrise and dusk is 60 minutes after sunset.
This day is split into 12 equal parts with each part being shaah zmanis.
Alternate methods of calculating a shaah zmanis are available in ComplexZmanimCalendar.
The method return the gdt.GMillisecond length of shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- Alos60
- Tzais60
- PlagHamincha60Minutes
*/
func (t *complexZmanimCalendar) ShaahZmanis60Minutes() (i gdt.GMillisecond, ok bool) {
	alos60, ok := t.Alos60()
	if !ok {
		return 0, false
	}
	tzais60, ok := t.Tzais60()
	if !ok {
		return 0, false
	}
	return temporalHour(alos60, tzais60), true
}

/*
ShaahZmanis72Minutes return a shaah zmanis (solar hour) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern.
This calculation divides the day based on the opinion of the MGA that the day runs from dawn to dusk.
Dawn for this calculation is 72 minutes before sunrise and dusk is 72 minutes after sunset.
This day is split into 12 equal parts with each part being shaah zmanis.
Alternate methods of calculating a shaah zmanis are available in ComplexZmanimCalendar.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) ShaahZmanis72Minutes() (i gdt.GMillisecond, ok bool) {
	return t.ShaahZmanisMGA()
}

/*
ShaahZmanis72MinutesZmanis return shaah zmanis (temporal hour) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos72Zmanis 72 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This calculation divides the day based on the opinion of the MGA,
that the day runs from dawn to dusk.
Dawn for this calculation is 72 minutes zmaniyos before sunrise and dusk is 72 minutes zmaniyos after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
This is identical to 1/10th of the day from AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle,
where there is at least one day a year, where the sun does not rise, and one where it does not set,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- Alos72Zmanis
- Tzais72Zmanis
*/
func (t *complexZmanimCalendar) ShaahZmanis72MinutesZmanis() (i gdt.GMillisecond, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return 0, false
	}
	tzais72Zmanis, ok := t.Tzais72Zmanis()
	if !ok {
		return 0, false
	}
	return temporalHour(alos72Zmanis, tzais72Zmanis), true
}

/*
ShaahZmanis90Minutes return a shaah zmanis (temporal hour) calculated using a dip of 90 minutes.
This calculation divides the day based on the opinion of the
[Magen Avraham (MGA)]: ]https://en.wikipedia.org/wiki/Avraham_Gombinern,
that the day runs from dawn to dusk.
Dawn for this calculation is 90 minutes before sunrise and dusk is 90 minutes after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) ShaahZmanis90Minutes() (i gdt.GMillisecond, ok bool) {
	alos90, ok := t.Alos90()
	if !ok {
		return 0, false
	}
	tzais90, ok := t.Tzais90()
	if !ok {
		return 0, false
	}
	return temporalHour(alos90, tzais90), true
}

/*
ShaahZmanis90MinutesZmanis return a shaah zmanis (temporal hour) according to the opinion of the
[Magen Avraham (MGA): https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos being
Alos90Zmanis 90 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This calculation divides the day based on the opinion of the MGA that the day runs from dawn to dusk.
Dawn for this calculation is 90 minutes zmaniyos before sunrise and dusk is 90 minutes zmaniyos after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
This is 1/8th of the day from AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- Alos90Zmanis
- Tzais90Zmanis
*/
func (t *complexZmanimCalendar) ShaahZmanis90MinutesZmanis() (i gdt.GMillisecond, ok bool) {
	alos90Zmanis, ok := t.Alos90Zmanis()
	if !ok {
		return 0, false
	}
	tzais90Zmanis, ok := t.Tzais90Zmanis()
	if !ok {
		return 0, false
	}
	return temporalHour(alos90Zmanis, tzais90Zmanis), true
}

/*
ShaahZmanis96MinutesZmanis return a shaah zmanis (temporal hour) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos being
Alos96Zmanis 96 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This calculation divides the day based on the opinion of the MGA that the day runs from dawn to dusk.
Dawn for this calculation is 96 minutes zmaniyos before sunrise and dusk is 96 minutes zmaniyos after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
This is identical to 1/7.5th of the day from AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset.
The method return the gdt.GMillisecond length of a shaah zmanis.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
- Alos96Zmanis
- Tzais96Zmanis
*/
func (t *complexZmanimCalendar) ShaahZmanis96MinutesZmanis() (i gdt.GMillisecond, ok bool) {
	alos96Zmanis, ok := t.Alos96Zmanis()
	if !ok {
		return 0, false
	}
	tzais96Zmanis, ok := t.Tzais96Zmanis()
	if !ok {
		return 0, false
	}
	return temporalHour(alos96Zmanis, tzais96Zmanis), true
}

/*
ShaahZmanisAteretTorah return a shaah zmanis (temporal hour) according to the opinion of the
Chacham Yosef Harari-Raful of Yeshivat Ateret Torah calculated with alos being 1/10th
of sunrise to sunset day, or Alos72Zmanis 72 minutes zmaniyos of such a day before
AstronomicalCalendar.Sunrise, and tzais is usually calculated as TzaisAteretTorah 40 minutes
(configurable to any offset via ComplexZmanimCalendar.ateretTorahSunsetOffset after AstronomicalCalendar.Sunset).
This day is split into 12 equal parts with each part being a shaah zmanis.
Note that with this system, chatzos (midday) will not be the point that the sun is AstronomicalCalendar.sunTransit halfway across
the sky.
the method return the gdt.GMillisecond length of a shaah zmanis.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos72Zmanis
  - TzaisAteretTorah
  - AteretTorahSunsetOffset
  - ateretTorahSunsetOffset
*/
func (t *complexZmanimCalendar) ShaahZmanisAteretTorah() (i gdt.GMillisecond, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return 0, false
	}
	tzaisAteretTorah, ok := t.TzaisAteretTorah()
	if !ok {
		return 0, false
	}
	return temporalHour(alos72Zmanis, tzaisAteretTorah), true
}

/*
ShaahZmanisAlos16Point1ToTzais3Point8 return a shaah zmanis (temporal hour) used by some zmanim according to the opinion of
[Rabbi Yaakov Moshe Hillel]: https://en.wikipedia.org/wiki/Yaakov_Moshe_Hillel
as published in the luach of the Bais Horaah of Yeshivat Chevrat Ahavat Shalom,
that is based on a day starting 72 minutes before sunrise in degrees Alos16Point1Degrees alos 16.1 deg
and ending 14 minutes after sunset in degrees TzaisGeonim3Point8Degrees tzais 3.8 deg.
This day is split into 12 equal parts with each part being a shaah zmanis.
Note that with this system, chatzos (midday) will not be the point that the sun is AstronomicalCalendar.sunTransit halfway across the sky.
These shaos zmaniyos are used for Mincha Ketana and Plag Hamincha.
The 14 minutes are based on 3/4 of an 18 minutes mil, with half a minute added for Rav Yosi.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - MinchaKetanaAhavatShalom
  - PlagAhavatShalom
*/
func (t *complexZmanimCalendar) ShaahZmanisAlos16Point1ToTzais3Point8() (i gdt.GMillisecond, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return 0, false
	}
	tzaisGeonim3Point8Degrees, ok := t.TzaisGeonim3Point8Degrees()
	if !ok {
		return 0, false
	}
	return temporalHour(alos16Point1Degrees, tzaisGeonim3Point8Degrees), true
}

/*
ShaahZmanisAlos16Point1ToTzais3Point7 return a shaah zmanis (temporal hour) used by some zmanim according to the opinion of
[Rabbi Yaakov Moshe Hillel]: https://en.wikipedia.org/wiki/Yaakov_Moshe_Hillel as published in the
luach of the Bais Horaah of Yeshivat Chevrat Ahavat Shalom that is based on a day starting 72 minutes before
sunrise in degrees Alos16Point1Degrees alos 16.1 deg and ending 13.5 minutes after sunset in
degrees TzaisGeonim3Point7Degrees tzais 3.7 deg.
This day is split into 12 equal parts with each part being a shaah zmanis.
Note that with this system, chatzos (midday) will not be the point that the sun is AstronomicalCalenadr.sunTransit halfway across the sky.
These shaos zmaniyos are used for Mincha Gedola calculation.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- MinchaGedolaAhavatShalom
*/
func (t *complexZmanimCalendar) ShaahZmanisAlos16Point1ToTzais3Point7() (i gdt.GMillisecond, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return 0, false
	}
	tzaisGeonim3Point7Degrees, ok := t.TzaisGeonim3Point7Degrees()
	if !ok {
		return 0, false
	}
	return temporalHour(alos16Point1Degrees, tzaisGeonim3Point7Degrees), true
}

/*
ShaahZmanis96Minutes return a shaah zmanis (temporal hour) calculated using a dip of 96 minutes.
This calculation divides the day based on the opinion of the
[Magen * Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern that the day runs from dawn to dusk.
Dawn for this calculation is 96 minutes before sunrise and dusk is 96 minutes after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) ShaahZmanis96Minutes() (i gdt.GMillisecond, ok bool) {
	alos96, ok := t.Alos96()
	if !ok {
		return 0, false
	}
	tzais96, ok := t.Tzais96()
	if !ok {
		return 0, false
	}
	return temporalHour(alos96, tzais96), true
}

/*
ShaahZmanis120Minutes return a shaah zmanis (temporal hour) calculated using a dip of 120 minutes.
This calculation divides the day based on the opinion of the
[Magen Avraham (MGA)]: ]https://en.wikipedia.org/wiki/Avraham_Gombinern,
that the day runs from dawn to dusk.
Dawn for this calculation is 120 minutes before sunrise and dusk is 120 minutes after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
Since zmanim that use this method are extremely late or early and at a point when the sky is a long time
past the 18deg point where the darkest point is reached, zmanim that use this should only be used
lechumra only, such as delaying the start of nighttime mitzvos.
The method return the gdt.GMillisecond length of a shaah zmanis.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis26Degrees
*/
func (t *complexZmanimCalendar) ShaahZmanis120Minutes() (i gdt.GMillisecond, ok bool) {
	alos120, ok := t.Alos120()
	if !ok {
		return 0, false
	}
	tzais120, ok := t.Tzais120()
	if !ok {
		return 0, false
	}
	return temporalHour(alos120, tzais120), true
}

/*
ShaahZmanis120MinutesZmanis return a shaah zmanis (temporal hour) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos being
Alos120Zmanis 120 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This calculation divides the day based on the opinion of the MGA that the day runs from dawn to dusk.
Dawn for this calculation is 120 minutes zmaniyos before sunrise and dusk is 120 minutes zmaniyos after sunset.
This day is split into 12 equal parts with each part being a shaah zmanis.
This is identical to 1/6th of the day from AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset.
Since zmanim that use this method are extremely late or early and at a point,
when the sky is a long time past the 18 deg point where the darkest point is reached,
zmanim that use this should only be used lechumra such as delaying the start of nighttime mitzvos.

The method return the gdt.GMillisecond length of a shaah zmanis.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos120Zmanis
  - Tzais120Zmanis
*/
func (t *complexZmanimCalendar) ShaahZmanis120MinutesZmanis() (i gdt.GMillisecond, ok bool) {
	alos120Zmanis, ok := t.Alos120Zmanis()
	if !ok {
		return 0, false
	}
	tzais120Zmanis, ok := t.Tzais120Zmanis()
	if !ok {
		return 0, false
	}
	return temporalHour(alos120Zmanis, tzais120Zmanis), true
}

/*
Alos60 return alos (dawn) calculated as 60 minutes before sunrise.
This is the time to walk the distance of 4 Mil at 15 minutes a Mil.
This seems to be the opinion of the [Chavas Yair]: https://en.wikipedia.org/wiki/Yair_Bacharach in the Mekor Chaim,
Orach Chaim Ch. 90, though the Mekor Chaim in Ch. 58 and in the
[Chut Hashani Cha 97]: ]https://hebrewbooks.org/pdfpager.aspx?req=45193&pgnum=214 states that,
a person walks 3 and a 1/3 mil in an hour, or an 18-minute mil.
Also see the [Divrei Malkiel]: https://he.wikipedia.org/wiki/%D7%9E%D7%9C%D7%9B%D7%99%D7%90%D7%9C_%D7%A6%D7%91%D7%99_%D7%98%D7%A0%D7%A0%D7%91%D7%95%D7%99%D7%9D"
[Vol. 4, Ch. 20, page 34]: ]https://hebrewbooks.org/pdfpager.aspx?req=803&pgnum=33) who
  - mentions the 15 minutes mil lechumra by baking matzos.

Also see the [Maharik]: https://en.wikipedia.org/wiki/Joseph_Colon_Trabotto
[Ch. 173]: https://hebrewbooks.org/pdfpager.aspx?req=1142&pgnum=216,
where the questioner quoting the [Ra'avan]: ]https://en.wikipedia.org/wiki/Eliezer_ben_Nathan is of the opinion,
that the time to walk a mil is 15 minutes (5 mil in a little over an hour).
There are many who believe that there is a ta'us sofer (scribe's error) in the Ra'avan,
and it should 4 mil in a little over an hour, or an 18-minute mil.
Time based offset calculations are based on the opinion of the [Rishonim]: https://en.wikipedia.org/wiki/Rishonim
who stated that the time of the neshef (time between dawn and sunrise) does not vary by the time of year or location,
but purely depends on the time it takes to walk the distance of 4* mil.
TzaisGeonim9Point75Degrees is a related zman that is a degree-based calculation based on 60 minutes.
The method return the time.Time representing the time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Tzais60
  - PlagHamincha60Minutes
  - ShaahZmanis60Minutes
*/
func (t *complexZmanimCalendar) Alos60() (tm time.Time, ok bool) {
	if sunrise, ok := t.Sunrise(); ok {
		return timeOffset(sunrise, -gdt.GMinute(60).ToMilliseconds()), true
	} else {
		return time.Time{}, false
	}
}

/*
Alos72Zmanis return alos (dawn) calculated using 72 minutes zmaniyos or 1/10th of the day before sunrise.
This is based on an 18-minute Mil so the time for 4 Mil is 72 minutes which is 1/10th of a day (12 * 60 = 720)
based on a day being from
AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunrise
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the ZmanimCalendar.IsUseElevation setting).
The actual calculation is AstronomicalCalendar.SeaLevelSunrise - ZmanimCalendar.ShaahZmanisGRA * 1.2.
This calculation is used in the calendars published by the
[Hisachdus Harabanim D'Artzos Habris Ve'Canada]: https://en.wikipedia.org/wiki/Central_Rabbinical_Congress.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where,
the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ZmanimCalendar.ShaahZmanisGRA
*/
func (t *complexZmanimCalendar) Alos72Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(-1.2)
}

/*
Alos96 return alos (dawn) calculated using 96 minutes before
AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the ZmanimCalendar.IsUseElevation setting)
that is based on the time to walk the distance of 4 Mil at 24 minutes a Mil.
Time based offset calculations for alos are based on the opinion of the
[Rishonim]: https://en.wikipedia.org/wiki/Rishonim,
who stated that the time of the Neshef (time between dawn and sunrise) does not vary by the time of year or location,
but purely depends on the time it takes to walk the distance of 4 Mil.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) Alos96() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunrise, -gdt.GMinute(96).ToMilliseconds()), true
}

/*
Alos90Zmanis return alos (dawn) calculated using 90 minutes zmaniyos or 1/8th of the day before
AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the ZmanimCalendar.IsUseElevation setting).
This is based on a 22.5-minute Mil, so the time for 4 Mil is 90 minutes which is 1/8th of a day (12 * 60) / 8 = 90
The day is calculated from AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunrise
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the ZmanimCalendar.IsUseElevation).
The actual calculation used is AstronomicalCalendar.Sunrise - ZmanimCalendar.ShaahZmanisGRA * 1.5.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ZmanimCalendar.ShaahZmanisGRA
*/
func (t *complexZmanimCalendar) Alos90Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(-1.5)
}

/*
Alos96Zmanis returns alos (dawn) calculated using 96 minutes zmaniyos or 1/7.5th of the day before
AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise
(depending on the ZmanimCalendar.IsUseElevation setting).
This is based on a 24-minute Mil, so the time for 4 Mil is 96 minutes which is 1/7.5th of a day (12 * 60 / 7.5 = 96).
The day is calculated from AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunrise
or
AstronomicalCalendar.Sunrise to AstronomicalCalendar.Sunset
(depending on the ZmanimCalendar.IsUseElevation).
The actual calculation used is AstronomicalCalendar.Sunrise - ZmanimCalendar.ShaahZmanisGRA * 1.6.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ZmanimCalendar.ShaahZmanisGRA
*/
func (t *complexZmanimCalendar) Alos96Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(-1.6)
}

/*
Alos90 return alos (dawn) calculated using 90 minutes before AstronomicalCalendar.SeaLevelSunrise,
based on the time to walk the distance of 4 Mil at 22.5 minutes a Mil.
Time based offset calculations for alos are based on the opinion of the [Rishonim]: https://en.wikipedia.org/wiki/Rishonim,
who stated that the time of the Neshef (time between dawn and sunrise) does not vary by the time of year or location but purely depends on the time it
takes to walk the distance of 4 Mil.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar	documentation.
*/
func (t *complexZmanimCalendar) Alos90() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunrise, -gdt.GMinute(90).ToMilliseconds()), true
}

/*
Alos18Degrees return alos (dawn) calculated when the sun is calculator.AstronomicalZenith 18 deg below the
eastern geometric horizon before sunrise.

The method return the time.Time representing alos.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - calculator.AstronomicalZenith
*/
func (t *complexZmanimCalendar) Alos18Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(calculator.AstronomicalZenith)
}

/*
Alos19Degrees return alos (dawn) calculated when the sun is zenith19Degrees 19 deg below the
eastern geometric horizon before sunrise. This is the [Rambam]: https://en.wikipedia.org/wiki/Maimonides
alos according to Rabbi Moshe Kosower's [Maaglei Tzedek]: https://www.worldcat.org/oclc/145454098, page 88,
[Ayeles Hashachar Vol. I, page 12]: https://hebrewbooks.org/pdfpager.aspx?req=33464&pgnum=13,
[Yom Valayla Shel Torah, Ch. 34, p. 222]: https://hebrewbooks.org/pdfpager.aspx?req=55960&pgnum=258
and Rabbi Yaakov Shakow's [Luach Ikvei Hayom]: https://www.worldcat.org/oclc/1043573513.

The method return the time.Time representing alos.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - calculator.AstronomicalZenith
*/
func (t *complexZmanimCalendar) Alos19Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith19Degrees)
}

/*
Alos19Point8Degrees return alos (dawn) calculated when the sun is calculator.zenith19Point8 19.8 deg below the
eastern geometric horizon before sunrise.
This calculation is based on the same calculation of Alos90 90 minutes,
but uses a degree-based calculation instead of 90 exact minutes.
This calculation is based on the position of the sun 90 minutes before sunrise in Jerusalem
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
which calculates to 19.8 deg below calculator.GeometricZenith.

The method return the time.Time representing alos.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - calculator.zenith19Point8
  - Alos90
*/
func (t *complexZmanimCalendar) Alos19Point8Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith19Point8)
}

/*
Alos16Point1Degrees return alos (dawn), calculated when the sun is calculator.zenith16Point1 16.1 deg below the
eastern geometric horizon before sunrise.
This calculation is based on the same calculation of ZmanimCalendar.Alos72 72 minutes,
but uses a degree-based calculation instead of 72 exact minutes.
This calculation is based on the position of the sun 72 minutes before sunrise in Jerusalem
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
which calculates to 16.1 deg below {calculator.GeometricZenith}.

The method return the time.Time representing alos.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - calculator.zenith16Point1
  - ZmanimCalendar.Alos72
*/
func (t *complexZmanimCalendar) Alos16Point1Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith16Point1)
}

/*
Misheyakir11Point5Degrees returns misheyakir based on the position of the sun when it is calculator.zenith11Degrees
11.5 deg below calculator.GeometricZenith (90 deg).
This calculation is used for calculating misheyakir according to some opinions.
This calculation is based on the position of the sun 52 minutes
before AstronomicalCalendar.Sunrise in Jerusalem
[around the equinox / equilux]: "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
which calculates to 11.5 deg below calculator.GeometricZenith.
TO_DO recalculate.

The method return the time.Time of misheyakir.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
- calculator.zenith11Degrees
*/
func (t *complexZmanimCalendar) Misheyakir11Point5Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith11Point5)
}

/*
Misheyakir11Degrees returns misheyakir based on the position of the sun when it is calculator.zenith11Degrees
11 deg below calculator.GeometricZenith (90 deg).
This calculation is used for calculating misheyakir according to some opinions.
This calculation is based on the position of the sun 48 minutes before AstronomicalCalendar.Sunrise in Jerusalem
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
which calculates to 11 deg below calculator.GeometricZenith.

The method return time.Time
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) Misheyakir11Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith11Degrees)
}

/*
Misheyakir10Point2Degrees returns misheyakir based on the position of the sun when it is zenith10Point2
10.2 deg below calculator.GeometricZenith (90 deg).
This calculation is used for calculating misheyakir according to some opinions.
This calculation is based on the position of the sun 45 minutes before AstronomicalCalendar.Sunrise in Jerusalem
[around the equinox]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/ which calculates
to 10.2 deg below calculator.GeometricZenith.

The method return the time.Time of misheyakir.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) Misheyakir10Point2Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith10Point2)
}

/*
Misheyakir7Point65Degrees returns misheyakir based on the position of the sun when it is zenith7Point65
7.65 deg below {calculator.GeometricZenith} (90 deg).
The degrees are based on a 35/36 minute zman
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
when the neshef (twilight) is the shortest.
This time is based on [Rabbi Moshe Feinstein]: https://en.wikipedia.org/wiki/Moshe_Feinstein who writes in
[Ohr Hachaim Vol. 4, Ch. 6]: "https://hebrewbooks.org/pdfpager.aspx?req=14677&pgnum=7
that misheyakir in New York is 35-40 minutes before sunset, something that is a drop less than 8 deg.
[Rabbi Yisroel Taplin]: https://en.wikipedia.org/wiki/Yisroel_Taplin in
[Zmanei Yisrael]: https://www.worldcat.org/oclc/889556744 (page 117) notes that
[Rabbi Yaakov Kamenetsky]: https://en.wikipedia.org/wiki/Yaakov_Kamenetsky stated that it is not less than 36
minutes before sunrise (maybe it is 40 minutes). Sefer Yisrael Vehazmanim (p. 7) quotes the Tamar Yifrach
in the name of the
[Satmar Rov]: https://en.wikipedia.org/wiki/Joel_Teitelbaum that one should be stringent
not consider misheyakir before 36 minutes. This is also the accepted
[minhag]: https://en.wikipedia.org/wiki/Minhag in
[Lakewood]: https://en.wikipedia.org/wiki/Lakewood_Township,_New_Jersey that is used in the
[Yeshiva]: https://en.wikipedia.org/wiki/Beth_Medrash_Govoha. This follows the opinion of
[Rabbi Shmuel Kamenetsky]: https://en.wikipedia.org/wiki/Shmuel_Kamenetsky who provided the time of 35/36 minutes,
but did not provide a degree-based time. Since this zman depends on the level of light, Rabbi Yaakov Shakow
presented this degree-based calculations to Rabbi Kamenetsky who agreed to them.

The method return the time.Time of misheyakir.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Misheyakir9Point5Degrees
*/
func (t *complexZmanimCalendar) Misheyakir7Point65Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith7Point65)
}

/*
Misheyakir9Point5Degrees returns misheyakir based on the position of the sun when it is zenith9Point5 9.5 deg
below {calculator.GeometricZenith} (90 deg).
This calculation is based on Rabbi Dovid Kronglass's Calculation of 45 minutes in Baltimore as mentioned in
[Divrei Chachamim No. 24]: https://hebrewbooks.org/pdfpager.aspx?req=20287&pgnum=29 brought down by the
[Birur Halacha, Tinyana, Ch. 18]: https://hebrewbooks.org/pdfpager.aspx?req=50535&pgnum=87.
This calculates to 9.5 deg.
Also see [Rabbi Yaakov Yitzchok Neiman]: https://en.wikipedia.org/wiki/Jacob_Isaac_Neiman in Kovetz
Eitz Chaim Vol. 9, p. 202 that the Vya'an Yosef did not want to rely on times earlier than 45 minutes in New York. This
zman is also used in the calendars published by Rabbi Hershel Edelstein. As mentioned in Yisroel Vehazmanim,
Rabbi Edelstein who was given the 45 minute zman by Rabbi Bick. The calendars published by the
[Edot Hamizrach]: https://en.wikipedia.org/wiki/Mizrahi_Jews communities also use this zman.
This also follows the opinion of
[Rabbi Shmuel Kamenetsky]: https://en.wikipedia.org/wiki/Shmuel_Kamenetsky who provided
the time of 36 and 45 minutes, but did not provide a degree-based time. Since this zman depends on the level of
light, Rabbi Yaakov Shakow presented these degree-based times to Rabbi Shmuel Kamenetsky who agreed to them.

The method return the time.Time of misheyakir.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle where,
the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- Misheyakir7Point65Degrees
*/
func (t *complexZmanimCalendar) Misheyakir9Point5Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith9Point5)
}

/*
SofZmanShmaMGA19Point8Degrees returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos19Point8Degrees 19.8 deg before AstronomicalCalendar.Sunrise.
This time is 3 ShaahZmanis19Point8Degrees shaos zmaniyos (solar hours) after Alos19Point8Degrees dawn
based on the opinion of the MGA that the day is calculated from dawn to nightfall
with both being 19.8 deg below sunrise or sunset. This returns the time of 3 * ShaahZmanis19Point8Degrees after Alos19Point8Degrees dawn.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanis19Point8Degrees
  - Alos19Point8Degrees
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA19Point8Degrees() (tm time.Time, ok bool) {
	alos19Point8Degrees, ok := t.Alos19Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais19Point8Degrees, ok := t.Tzais19Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos19Point8Degrees, tzais19Point8Degrees), true
}

/*
SofZmanShmaMGA16Point1Degrees returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based
on alos being Alos16Point1Degrees 16.1 deg before AstronomicalCalendar.Sunrise. This time is 3 ShaahZmanis16Point1Degrees shaos zmaniyos (solar hours) after
Alos16Point1Degrees dawn based on the opinion of the MGA that the day is calculated from
dawn to nightfall with both being 16.1 deg below sunrise or sunset.
This returns the time of 3 * ShaahZmanis16Point1Degrees after Alos16Point1Degrees dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis16Point1Degrees()
  - Alos16Point1Degrees()
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA16Point1Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais16Point1Degrees, ok := t.Tzais16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos16Point1Degrees, tzais16Point1Degrees), true
}

/*
SofZmanShmaMGA18Degrees
This method returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the []: https://en.wikipedia.org/wiki/Avraham_Gombinern">Magen Avraham (MGA)</a> based
on alos being {@link Alos18Degrees() 18 deg} before {AstronomicalCalendar.Sunrise. This time is 3
{@link ShaahZmanis18Degrees() shaos zmaniyos} (solar hours) after {@link Alos18Degrees() dawn}
based on the opinion of the MGA that the day is calculated from dawn to nightfall with both being 18 deg
below sunrise or sunset. This returns the time of 3 * {@link ShaahZmanis18Degrees()} after
{@link Alos18Degrees() dawn}.

The method return the time.Time of the latest zman krias shema. If the calculation can't be computed such

	as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle
	where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
	See detailed explanation on top of the AstronomicalCalendar documentation.

- ShaahZmanis18Degrees()
- Alos18Degrees()
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA18Degrees() (tm time.Time, ok bool) {
	alos18Degrees, ok := t.Alos18Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais18Degrees, ok := t.Tzais18Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos18Degrees, tzais18Degrees), true
}

/*
SofZmanShmaMGA72Minutes returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being  ZmanimCalendar.Alos72 minutes before AstronomicalCalendar.Sunrise.
This time is 3 ShaahZmanis72Minutesshaos zmaniyos (solar hours) after ZmanimCalendar.Alos72 based on the opinion
of the MGA that the day is calculated from a ZmanimCalendar.Alos72 of 72 minutes before sunrise to
ZmanimCalendar.Tzais72 nightfall of 72 minutes after sunset. This returns the time of 3 * ShaahZmanis72Minutes
after ZmanimCalendar.Alos72. This class returns an identical time to ZmanimCalendar.SofZmanShmaMGA and is repeated here for clarity.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis72Minutes
  - ZmanimCalendar.Alos72
  - ZmanimCalendar.SofZmanShmaMGA
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA72Minutes() (tm time.Time, ok bool) {
	return t.SofZmanShmaMGA()
}

/*
SofZmanShmaMGA72MinutesZmanis returns the latest zman krias shema (time to recite Shema in the morning) according
to the opinion of the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based
on alos being Alos72Zmanis 72 minutes zmaniyos, or 1/10th of the day before
AstronomicalCalendar.Sunrise. This time is 3 ShaahZmanis90MinutesZmanis shaos zmaniyos
(solar hours) after Alos72Zmanis dawn based on the opinion of the MGA that the day is calculated
from a Alos72Zmanis dawn of 72 minutes zmaniyos, or 1/10th of the day before
AstronomicalCalendar.SeaLevelSunrise to Tzais72Zmanis nightfall of 72 minutes
zmaniyos after AstronomicalCalendar.SeaLevelSunset.
This returns the time of 3 * ShaahZmanis72MinutesZmanis after Alos72Zmanis dawn.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis72MinutesZmanis
  - Alos72Zmanis
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA72MinutesZmanis() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais72Zmanis, ok := t.Tzais72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos72Zmanis, tzais72Zmanis), true
}

/*
SofZmanShmaMGA90Minutes returns the latest zman krias shema (time to recite Shema in the morning) according
to the opinion of the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos90 90 minutes before AstronomicalCalendar.Sunrise. This time is 3 ShaahZmanis90Minutes shaos zmaniyos (solar hours) after Alos90 dawn based on
the opinion of the MGA that the day is calculated from Alos90 dawn of 90 minutes before sunrise to
ZmanimCalendar.Tzais90 nightfall of 90 minutes after sunset. This returns the time of 3 * ShaahZmanis90Minutes after Alos90 dawn.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis90Minutes
  - Alos90
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA90Minutes() (tm time.Time, ok bool) {
	alos90, ok := t.Alos90()
	if !ok {
		return time.Time{}, false
	}
	tzais90, ok := t.Tzais90()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos90, tzais90), true
}

/*
SofZmanShmaMGA90MinutesZmanis returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based
on alos being Alos90Zmanis 90 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This time is 3 ShaahZmanis90MinutesZmanis shaos zmaniyos (solar hours) after
Alos90Zmanis dawn based on the opinion of the MGA that the day is calculated from a Alos90Zmanis dawn of 90 minutes zmaniyos
before sunrise to Tzais90Zmanis nightfall of 90 minutes zmaniyos after sunset.
This returns the time of 3 * ShaahZmanis90MinutesZmanis after Alos90Zmanis dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis90MinutesZmanis()
  - Alos90Zmanis()
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA90MinutesZmanis() (tm time.Time, ok bool) {
	alos90Zmanis, ok := t.Alos90Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais90Zmanis, ok := t.Tzais90Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos90Zmanis, tzais90Zmanis), true
}

/*
SofZmanShmaMGA96Minutes returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based
on alos being Alos96 96 minutes before AstronomicalCalendar.Sunrise.
This time is 3 * ShaahZmanis96Minutes shaos zmaniyos (solar hours) after Alos96 dawn based on
the opinion of the MGA that the day is calculated from Alos96 dawn of 96 minutes before
sunrise to Tzais96 nightfall of 96 minutes after sunset.
This returns the time of 3 * ShaahZmanis96Minutes after Alos96 dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis96Minutes
  - Alos96
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA96Minutes() (tm time.Time, ok bool) {
	alos96, ok := t.Alos96()
	if !ok {
		return time.Time{}, false
	}
	tzais96, ok := t.Tzais96()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos96, tzais96), true
}

/*
SofZmanShmaMGA96MinutesZmanis returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based
on alos being Alos90Zmanis 96 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This time is 3 * ShaahZmanis96MinutesZmanis shaos zmaniyos (solar hours) after
Alos96Zmanis dawn based on the opinion of the MGA that the day is calculated from a Alos96Zmanis dawn of 96 minutes zmaniyos
before sunrise to Tzais90Zmanis nightfall of 96 minutes zmaniyos after sunset.
This returns the time of 3 * ShaahZmanis96MinutesZmanis after Alos96Zmanis dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis96MinutesZmanis
  - Alos96Zmanis
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA96MinutesZmanis() (tm time.Time, ok bool) {
	alos96Zmanis, ok := t.Alos96Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais96Zmanis, ok := t.Tzais96Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos96Zmanis, tzais96Zmanis), true
}

/*
SofZmanShma3HoursBeforeChatzos returns the latest zman krias shema (time to recite Shema in the morning) calculated
as 3 hours (regular clock hours and not sha'os zmaniyos) before ZmanimCalendar#getChatzos.
Generally known as part of the "Komarno" zmanim after
[Rav Yitzchak Eizik of Komarno]: https://en.wikipedia.org/wiki/Komarno_(Hasidic_dynasty)#Rabbi_Yitzchak_Eisik_Safrin,
a proponent of this calculation, it actually predates him a lot. It is the opinion of the
Shach in the Nekudas Hakesef (Yoreh Deah 184),
[Rav Moshe Lifshitz]: https://hebrewbooks.org/pdfpager.aspx?req=21638&st=&pgnum=30 in his commentary
[Lechem Mishneh on Brachos 1:2]: https://hebrewbooks.org/pdfpager.aspx?req=21638&st=&pgnum=50.
It is next brought down about 100 years later by the [Yaavetz]: https://en.wikipedia.org/wiki/Jacob_Emden
(in his siddur, [Mor Uktziah Orach Chaim 1]: https://hebrewbooks.org/pdfpager.aspx?req=7920&st=&pgnum=6,
[Lechem Shamayim, Brachos 1:2]: https://hebrewbooks.org/pdfpager.aspx?req=22309&st=&pgnum=30
and [She'elos Yaavetz vol. 1 no. 40]: https://hebrewbooks.org/pdfpager.aspx?req=1408&st=&pgnum=69),
Rav Yitzchak Eizik of Komarno in the Ma'aseh Oreg on Mishnayos Brachos 11:2, Shevus Yaakov, Chasan Sofer and others.
See Yisrael Vehazmanim [vol. 1 7:3, page 55 - 62]: https://hebrewbooks.org/pdfpager.aspx?req=9765&st=&pgnum=83.
A variant of this calculation SofZmanShmaFixedLocal uses FixedLocalChatzos fixed local chatzos for calculating this type of zman.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ZmanimCalendar.Chatzos
  - SofZmanShmaFixedLocal
  - SofZmanTfila2HoursBeforeChatzos
*/
func (t *complexZmanimCalendar) SofZmanShma3HoursBeforeChatzos() (tm time.Time, ok bool) {
	chatzos, ok := t.Chatzos()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(chatzos, -gdt.GMinute(180).ToMilliseconds()), true
}

/*
SofZmanShmaMGA120Minutes returns the latest zman krias shema (time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos being Alos120 120 minutes
or 1/6th of the day before AstronomicalCalendar.Sunrise.
This time is 3 * ShaahZmanis120Minutes shaos zmaniyos (solar hours) after Alos120 dawn based on the opinion of the MGA,
that the day is calculated from Alos120 dawn of 120 minutes before sunrise to Tzais120 nightfall of 120 minutes after sunset.
This returns the time of 3 * link ShaahZmanis120Minutes after Alos120 dawn.
This is an extremely early zman that is very much a chumra.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis120Minutes()
  - Alos120()
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA120Minutes() (tm time.Time, ok bool) {
	alos120, ok := t.Alos120()
	if !ok {
		return time.Time{}, false
	}
	tzais120, ok := t.Tzais120()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos120, tzais120), true
}

/*
SofZmanShmaAlos16Point1ToSunset returns the latest zman krias shema (time to recite Shema in the morning) based
on the opinion that the day starts at Alos16Point1Degrees alos 16.1 deg and ends at AstronomicalCalendar.SeaLevelSunset.
This is the opinion of the
[\u05D7\u05D9\u05D3\u05D5\u05E9\u05D9\u05D5\u05DB\u05DC\u05DC\u05D5\u05EA \u05D4\u05E8\u05D6\u05F4\u05D4]: https://hebrewbooks.org/40357
and the
[\u05DE\u05E0\u05D5\u05E8\u05D4 \u05D4\u05D8\u05D4\u05D5\u05E8\u05D4]: https://hebrewbooks.org/14799 as
mentioned by Yisrael Vehazmanim
[vol 1, sec. 7, ch. 3 no. 16]: https://hebrewbooks.org/pdfpager.aspx?req=9765&pgnum=81.
Three shaos zmaniyos are calculated based on this day and added to Alos16Point1Degrees alos to reach this time.
This time is 3 shaos zmaniyos (solar hours) after Alos16Point1Degrees dawn based on the opinion that the day is calculated
from a Alos16Point1Degrees alos 16.1 deg to AstronomicalCalendar.SeaLevelSunset.
Note: Based on this calculation chatzos will not be at midday.

The method return the time.Time of the latest zman krias shema based on this day.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the
Antarctic Circle where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Alos16Point1Degrees
  - AstronomicalCalendar.SeaLevelSunset
*/
func (t *complexZmanimCalendar) SofZmanShmaAlos16Point1ToSunset() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos16Point1Degrees, elevationAdjustedSunset), true
}

/*
SofZmanShmaAlos16Point1ToTzaisGeonim7Point083Degrees returns the latest zman krias shema (time to recite Shema in the morning) based on the
opinion that the day starts at Alos16Point1Degrees alos 16.1 deg and ends at TzaisGeonim7Point083Degrees tzais 7.083 deg.
3 shaos zmaniyos are calculated based on this day and added to Alos16Point1Degrees alos to reach this time.
This time is 3 shaos zmaniyos (temporal hours) after Alos16Point1Degrees alos 16.1 deg based on the opinion that the day is calculated
from a Alos16Point1Degrees alos 16.1 deg to
TzaisGeonim7Point083Degrees tzais 7.083 deg.
Note: Based on this calculation chatzos will not be at midday.

The method return the time.Time of the latest zman krias shema based on this calculation.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned. See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Alos16Point1Degrees
  - TzaisGeonim7Point083Degrees
*/
func (t *complexZmanimCalendar) SofZmanShmaAlos16Point1ToTzaisGeonim7Point083Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzaisGeonim7Point083Degrees, ok := t.TzaisGeonim7Point083Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos16Point1Degrees, tzaisGeonim7Point083Degrees), true
}

/*
SofZmanTfilaMGA19Point8Degrees returns the latest zman tfila (time to recite the morning prayers) according to the opinion
of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos19Point8Degrees 19.8 deg before AstronomicalCalendar.Sunrise.
This time is 4 ShaahZmanis19Point8Degrees shaos zmaniyos (solar hours) after Alos19Point8Degrees dawn
based on the opinion of the MGA that the day is calculated from dawn to nightfall with both being 19.8 deg below sunrise or sunset.
This returns the time of 4 * ShaahZmanis19Point8Degrees after Alos19Point8Degrees dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis19Point8Degrees()
  - Alos19Point8Degrees()
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA19Point8Degrees() (tm time.Time, ok bool) {
	alos19Point8Degrees, ok := t.Alos19Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais19Point8Degrees, ok := t.Tzais19Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos19Point8Degrees, tzais19Point8Degrees), true
}

/*
SofZmanTfilaMGA16Point1Degrees returns the latest zman tfila (time to recite the morning prayers) according to the opinion
of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos16Point1Degrees 16.1 deg before AstronomicalCalendar.Sunrise.
This time is 4 ShaahZmanis16Point1Degrees shaos zmaniyos (solar hours) after Alos16Point1Degrees dawn
based on the opinion of the MGA that the day is calculated from dawn to nightfall with both being 16.1 deg below sunrise or sunset.
This returns the time of 4 * ShaahZmanis16Point1Degrees after Alos16Point1Degrees dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis16Point1Degrees()
  - Alos16Point1Degrees()
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA16Point1Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais16Point1Degrees, ok := t.Tzais16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos16Point1Degrees, tzais16Point1Degrees), true
}

/*
SofZmanTfilaMGA18Degrees returns the latest zman tfila (time to recite the morning prayers) according to the opinion
of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos18Degrees 18 deg before AstronomicalCalendar.Sunrise.
This time is 4 ShaahZmanis18Degrees shaos zmaniyos (solar hours) after Alos18Degrees dawn
based on the opinion of the MGA that the day is calculated from dawn to nightfall with both being 18 deg
below sunrise or sunset.
This returns the time of 4 * ShaahZmanis18Degrees after Alos18Degrees dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis18Degrees()
  - Alos18Degrees()
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA18Degrees() (tm time.Time, ok bool) {
	alos18Degrees, ok := t.Alos18Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais18Degrees, ok := t.Tzais18Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos18Degrees, tzais18Degrees), true
}

/*
SofZmanTfilaMGA72Minutes returns the latest zman tfila (time to recite the morning prayers) according to the opinion
of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being  ZmanimCalendar.Alos72 minutes before AstronomicalCalendar.Sunrise.
This time is 4 * ShaahZmanis72Minutes shaos zmaniyos (solar hours) after ZmanimCalendar.Alos72 based on
the opinion of the MGA that the day is calculated from a ZmanimCalendar.Alos72 of 72 minutes before
sunrise to ZmanimCalendar.Tzais72 nightfall of 72 minutes after sunset. This returns the time of 4 *
ShaahZmanis72Minutes after ZmanimCalendar.Alos72. This class returns an identical time to
ZmanimCalendar.SofZmanTfilaMGA and is repeated here for clarity.

The method return the time.Time of the latest zman tfila.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis72Minutes()
  - ZmanimCalendar.Alos72
  - ZmanimCalendar.SofZmanShmaMGA
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA72Minutes() (tm time.Time, ok bool) {
	return t.SofZmanTfilaMGA()
}

/*
SofZmanTfilaMGA72MinutesZmanis returns the latest zman tfila (time to the morning prayers) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos
being Alos72Zmanis 72 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This time is 4 * ShaahZmanis72MinutesZmanis shaos zmaniyos (solar hours) after Alos72Zmanis dawn
based on the opinion of the MGA that the day is calculated from a Alos72Zmanis dawn of 72 minutes zmaniyos before sunrise
to Tzais72Zmanis nightfall of 72 minutes zmaniyos after sunset.
This returns the time of 4 * ShaahZmanis72MinutesZmanis after Alos72Zmanis dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis72MinutesZmanis()
  - Alos72Zmanis()
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA72MinutesZmanis() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais72Zmanis, ok := t.Tzais72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos72Zmanis, tzais72Zmanis), true
}

/*
SofZmanTfilaMGA90Minutes returns the latest zman tfila (time to recite the morning prayers) according to the opinion
of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos90 90 minutes before AstronomicalCalendar.Sunrise.
This time is 4 * ShaahZmanis90Minutes shaos zmaniyos (solar hours) after Alos90 dawn based on
the opinion of the MGA that the day is calculated from Alos90 dawn of 90 minutes before sunrise to
Tzais90 nightfall of 90 minutes after sunset.
This returns the time of 4 * ShaahZmanis90Minutes after Alos90 dawn.

The method return the time.Time of the latest zman tfila.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis90Minutes()
  - Alos90
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA90Minutes() (tm time.Time, ok bool) {
	alos90, ok := t.Alos90()
	if !ok {
		return time.Time{}, false
	}
	tzais90, ok := t.Tzais90()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos90, tzais90), true
}

/*
SofZmanTfilaMGA90MinutesZmanis returns the latest zman tfila (time to the morning prayers) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos
being Alos90Zmanis 90 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This time is 4 * ShaahZmanis90MinutesZmanis shaos zmaniyos (solar hours) after Alos90Zmanis dawn
based on the opinion of the MGA that the day is calculated from a Alos90Zmanis dawn
of 90 minutes zmaniyos before sunrise to Tzais90Zmanis nightfall of 90 minutes
zmaniyos after sunset. This returns the time of 4 * ShaahZmanis90MinutesZmanis after Alos90Zmanis dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis90MinutesZmanis()
  - Alos90Zmanis()
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA90MinutesZmanis() (tm time.Time, ok bool) {
	alos90Zmanis, ok := t.Alos90Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais90Zmanis, ok := t.Tzais90Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos90Zmanis, tzais90Zmanis), true
}

/*
SofZmanTfilaMGA96Minutes returns the latest zman tfila (time to recite the morning prayers) according to the opinion
of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos96 96 minutes before AstronomicalCalendar.Sunrise.
This time is 4 * ShaahZmanis96Minutes shaos zmaniyos (solar hours) after Alos96 dawn based on
the opinion of the MGA that the day is calculated from Alos96 dawn of 96 minutes before
sunrise to Tzais96 nightfall of 96 minutes after sunset.
This returns the time of 4 * ShaahZmanis96Minutes after Alos96 dawn.

The method return the time.Time of the latest zman tfila.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis96Minutes()
  - Alos96()
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA96Minutes() (tm time.Time, ok bool) {
	alos96, ok := t.Alos96()
	if !ok {
		return time.Time{}, false
	}
	tzais96, ok := t.Tzais96()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos96, tzais96), true
}

/*
SofZmanTfilaMGA96MinutesZmanis returns the latest zman tfila (time to the morning prayers) according to the opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos
being Alos96Zmanis 96 minutes zmaniyos before AstronomicalCalendar.Sunrise.
This time is 4 * ShaahZmanis96MinutesZmanis shaos zmaniyos (solar hours) after Alos96Zmanis dawn
based on the opinion of the MGA that the day is calculated from a Alos96Zmanis dawn
of 96 minutes zmaniyos before sunrise to Tzais96Zmanis nightfall of 96 minutes
zmaniyos after sunset.
This returns the time of 4 * ShaahZmanis96MinutesZmanis after Alos96Zmanis dawn.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle,
where there is at least one day a year, where the sun does not rise, and one where it does not set,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis90MinutesZmanis()
  - Alos90Zmanis()
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA96MinutesZmanis() (tm time.Time, ok bool) {
	alos96Zmanis, ok := t.Alos96Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais96Zmanis, ok := t.Tzais96Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos96Zmanis, tzais96Zmanis), true
}

/*
SofZmanTfilaMGA120Minutes returns the latest zman tfila (time to recite the morning prayers) according to the opinion
of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on
alos being Alos120 120 minutes before AstronomicalCalendar.Sunrise.
This time is 4 * ShaahZmanis120Minutes shaos zmaniyos (solar hours) after Alos120 dawn
based on the opinion of the MGA that the day is calculated from Alos120 dawn of 120 minutes
before sunrise to Tzais120 nightfall of 120 minutes after sunset.
This returns the time of 4 * ShaahZmanis120Minutes after Alos120 dawn.
This is an extremely early zman that is very much a chumra.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis120Minutes
  - Alos120
*/
func (t *complexZmanimCalendar) SofZmanTfilaMGA120Minutes() (tm time.Time, ok bool) {
	alos120, ok := t.Alos120()
	if !ok {
		return time.Time{}, false
	}
	tzais120, ok := t.Tzais120()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos120, tzais120), true
}

/*
SofZmanTfila2HoursBeforeChatzos returns the latest zman tfila (time to recite the morning prayers) calculated as 2 hours
before ZmanimCalendar.Chatzos.
This is based on the opinions that calculate sof zman krias shema as SofZmanShma3HoursBeforeChatzos.
This returns the time of 2 hours before ZmanimCalendar.Chatzos.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ZmanimCalendar#getChatzos
  - SofZmanShma3HoursBeforeChatzos
*/
func (t *complexZmanimCalendar) SofZmanTfila2HoursBeforeChatzos() (tm time.Time, ok bool) {
	chatzos, ok := t.Chatzos()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(chatzos, -gdt.GMinute(120).ToMilliseconds()), true
}

/*
MinchaGedola30Minutes returns mincha gedola calculated as 30 minutes after ZmanimCalendar.Chatzos
and not 1/2 of a {ZmanimCalendar.ShaahZmanisGRA shaah zmanis} after ZmanimCalendar.Chatzos as
calculated by MinchaGedola. Some use this time to delay the start of mincha in the winter when
1/2 of a ZmanimCalendar.ShaahZmanisGRA shaah zmanis is less than 30 minutes.
See MinchaGedolaGreaterThan30 for a convenience method that returns the latter of the 2 calculations.
One should not use this time to start mincha before the standard ZmanimCalendar.MinchaGedola.
See Shulchan Aruch [Orach Chayim 234:1]: https://hebrewbooks.org/pdfpager.aspx?req=49624&st=&pgnum=291 and
the Shaar Hatziyon seif katan ches.

The method return the time.Time of 30 gdt.GMinute after chatzos.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ZmanimCalendar.MinchaGedola
  - MinchaGedolaGreaterThan30
*/
func (t *complexZmanimCalendar) MinchaGedola30Minutes() (tm time.Time, ok bool) {
	chatzos, ok := t.Chatzos()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(chatzos, gdt.GMinute(30).ToMilliseconds()), true
}

/*
MinchaGedola72Minutes returns the time of mincha gedola according to the Magen Avraham with the day starting 72
minutes before sunrise and ending 72 minutes after sunset.
This is the earliest time to pray mincha.
For more information on this see the documentation on MinchaGedola. This is
calculated as 6.5 temporalHour solar hours after alos.
The calculation used is 6.5 * ShaahZmanis72Minutes after ZmanimCalendar.Alos72 alos.
see
  - ZmanimCalendar.Alos72
  - ZmanimCalendar.MinchaGedola
  - ZmanimCalendar.MinchaKetana

The method return the time.Time of the time of mincha gedola.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaGedola72Minutes() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	tzais72, ok := t.Tzais72()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaGedola2(alos72, tzais72), true
}

/*
MinchaGedola16Point1Degrees returns the time of mincha gedola according to the Magen Avraham with the day starting and
ending 16.1 deg below the horizon.
This is the earliest time to pray mincha.
For more information on this see the documentation on ZmanimCalendar.MinchaGedola.
This is calculated as 6.5 temporalHour solar hours after alos.
The calculation used is 6.5 * ShaahZmanis16Point1Degrees after Alos16Point1Degrees alos.
see
  - ShaahZmanis16Point1Degrees
  - ZmanimCalendar.MinchaGedola()
  - ZmanimCalendar.MinchaKetana()

The method return the time.Time of the time of mincha gedola.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun  may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaGedola16Point1Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais16Point1Degrees, ok := t.Tzais16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaGedola2(alos16Point1Degrees, tzais16Point1Degrees), true
}

/*
MinchaGedolaAhavatShalom returns the time of mincha gedola based on the opinion of
[Rabbi Yaakov Moshe Hillel]: https://en.wikipedia.org/wiki/Yaakov_Moshe_Hillel as published in the luach
of the Bais Horaah of Yeshivat Chevrat Ahavat Shalom that mincha gedola is calculated as half a shaah
zmanis after chatzos with shaos zmaniyos calculated based on a day starting 72 minutes befoe sunrise
Alos16Point1Degrees alos 16.1 deg and ending 13.5 minutes after sunset TzaisGeonim3Point7Degrees tzais 3.7 deg.
Mincha gedola is the earliest time to pray mincha.
The latter of this time or 30 clock minutes after chatzos is returned. See MinchaGedolaGreaterThan30
(though that calculation is based on mincha gedola GRA).
For more information about mincha gedola see the documentation on ZmanimCalendar.MinchaGedola mincha gedola.

The method return the time.Time of the mincha gedola.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Alos16Point1Degrees
  - TzaisGeonim3Point7Degrees
  - ShaahZmanisAlos16Point1ToTzais3Point7()
  - MinchaGedolaGreaterThan30
*/
func (t *complexZmanimCalendar) MinchaGedolaAhavatShalom() (tm time.Time, ok bool) {
	minchaGedola30Minutes, ok := t.MinchaGedola30Minutes()
	if !ok {
		return time.Time{}, false
	}
	chatzos, ok := t.Chatzos()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanisAlos16Point1ToTzais3Point7, ok := t.ShaahZmanisAlos16Point1ToTzais3Point7()
	if !ok {
		return time.Time{}, false
	}

	if minchaGedola30Minutes.UnixMilli()-timeOffset(chatzos, shaahZmanisAlos16Point1ToTzais3Point7/2).UnixMilli() > 0 {
		return minchaGedola30Minutes, true
	} else {
		return timeOffset(chatzos, shaahZmanisAlos16Point1ToTzais3Point7/2), true
	}
}

/*
MinchaGedolaGreaterThan30  is a convenience method that returns the latter of ZmanimCalendar.MinchaGedola and
MinchaGedola30Minutes. In the winter when 1/2 of a ZmanimCalendar.ShaahZmanisGRA shaah zmanis is
less than 30 minutes MinchaGedola30Minutes will be returned, otherwise ZmanimCalendar.MinchaGedola.
will be returned.

The method return the time.Time of the latter of ZmanimCalendar.MinchaGedola and MinchaGedola30Minutes.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaGedolaGreaterThan30() (tm time.Time, ok bool) {
	minchaGedola30Minutes, ok := t.MinchaGedola30Minutes()
	if !ok {
		return time.Time{}, false
	}
	minchaGedola, ok := t.MinchaGedola()
	if !ok {
		return time.Time{}, false
	}
	if minchaGedola30Minutes.UnixMilli()-minchaGedola.UnixMilli() > 0 {
		return minchaGedola30Minutes, true
	} else {
		return minchaGedola, true
	}
}

/*
MinchaKetana16Point1Degrees returns the time of mincha ketana according to the Magen Avraham with the day starting and
ending 16.1 deg below the horizon. This is the preferred the earliest time to pray mincha according to the
opinion of the [Rambam]: https://en.wikipedia.org/wiki/Maimonides and others. For more information on
this see the documentation on ZmanimCalendar.MinchaGedola. This is calculated as 9.5
temporalHour solar hours after alos. The calculation used is 9.5 *
ShaahZmanis16Point1Degrees after Alos16Point1Degrees.
see
  - ShaahZmanis16Point1Degrees
  - ZmanimCalendar.MinchaGedola
  - ZmanimCalendar.MinchaKetana

The method return the time.Time of the time of mincha ketana.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaKetana16Point1Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais16Point1Degrees, ok := t.Tzais16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaKetana2(alos16Point1Degrees, tzais16Point1Degrees), false
}

/*
MinchaKetanaAhavatShalom method returns the time of mincha ketana based on the opinion of
[Rabbi Yaakov Moshe Hillel]: https://en.wikipedia.org/wiki/Yaakov_Moshe_Hillel as published in the luach
of the Bais Horaah of Yeshivat Chevrat Ahavat Shalom that mincha ketana is calculated as 2.5 shaos zmaniyos
before TzaisGeonim3Point8Degrees tzais 3.8 deg with shaos zmaniyos calculated based on a day starting at
Alos16Point1Degrees alos 16.1 deg and ending at tzais 3.8 deg.
Mincha ketana is the preferred the earliest time to pray mincha according to the opinion of the
[Rambam]: https://en.wikipedia.org/wiki/Maimonides and others.
For more information on this see the documentation on ZmanimCalendar.MinchaKetana.

The method return the time.Time of the time of mincha ketana.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle
and north of the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see:
- ShaahZmanisAlos16Point1ToTzais3Point8
- MinchaGedolaAhavatShalom
- PlagAhavatShalom
*/
func (t *complexZmanimCalendar) MinchaKetanaAhavatShalom() (tm time.Time, ok bool) {
	tzaisGeonim3Point8Degrees, ok := t.TzaisGeonim3Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanisAlos16Point1ToTzais3Point8, ok := t.ShaahZmanisAlos16Point1ToTzais3Point8()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(tzaisGeonim3Point8Degrees, -gdt.GMillisecond(float64(shaahZmanisAlos16Point1ToTzais3Point8)*2.5)), true
}

/*
MinchaKetana72Minutes returns the time of mincha ketana according to the Magen Avraham with the day
starting 72 minutes before sunrise and ending 72 minutes after sunset.
This is the preferred the earliest time to pray mincha according to the opinion of the
[Rambam]: https://en.wikipedia.org/wiki/Maimonides
and others. For more information on this see the documentation on ZmanimCalendar.MinchaGedola.
This is calculated as 9.5 ShaahZmanis72Minutes after alos.
The calculation used is 9.5 * ShaahZmanis72Minutes after ZmanimCalendar.Alos72 alos.
see
  - ShaahZmanis16Point1Degrees
  - ZmanimCalendar.MinchaGedola
  - ZmanimCalendar.MinchaKetana

The method return the time.Time of the time of mincha ketana.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar	documentation.
*/
func (t *complexZmanimCalendar) MinchaKetana72Minutes() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	tzais72, ok := t.Tzais72()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaKetana2(alos72, tzais72), true
}

/*
PlagHamincha60Minutes returns the time of plag hamincha according to the Magen Avraham with the day starting 60
minutes before sunrise and ending 60 minutes after sunset.
This is calculated as 10.75 hours after Alos60 dawn.
The formula used is 10.75 ShaahZmanis60Minutes after Alos60.

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ShaahZmanis60Minutes
  - Alos60
  - Tzais60
*/
func (t *complexZmanimCalendar) PlagHamincha60Minutes() (tm time.Time, ok bool) {
	alos60, ok := t.Alos60()
	if !ok {
		return time.Time{}, false
	}
	tzais60, ok := t.Tzais60()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos60, tzais60), true
}

/*
PlagAlos16Point1ToTzaisGeonim7Point083Degrees returns the time of plag hamincha based on the opinion that the day starts at
Alos16Point1Degrees alos 16.1 deg and ends at TzaisGeonim7Point083Degrees tzais.
10.75 shaos zmaniyos are calculated based on this day and added to Alos16Point1Degrees alos to reach this time.
This time is 10.75 shaos zmaniyos (temporal hours) after Alos16Point1Degrees dawn based on the opinion,
that the day is calculated from a Alos16Point1Degrees dawn of 16.1 degrees before sunrise to TzaisGeonim7Point083Degrees tzais.
This returns the time of 10.75 * the calculated shaah zmanis after Alos16Point1Degrees dawn.

The method return the time.Time of the plag.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Alos16Point1Degrees
  - TzaisGeonim7Point083Degrees
*/
func (t *complexZmanimCalendar) PlagAlos16Point1ToTzaisGeonim7Point083Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzaisGeonim7Point083Degrees, ok := t.TzaisGeonim7Point083Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos16Point1Degrees, tzaisGeonim7Point083Degrees), true
}

/*
PlagAhavatShalom returns the time of plag hamincha (the earliest time that Shabbos can be started) based on the
opinion of [Rabbi Yaakov Moshe Hillel]: https://en.wikipedia.org/wiki/Yaakov_Moshe_Hillel as published in
the luach of the Bais Horaah of Yeshivat Chevrat Ahavat Shalom that plag hamincha is calculated
as 1.25 shaos zmaniyos before TzaisGeonim3Point8Degrees tzais 3.8 deg with shaos
zmaniyos calculated based on a day starting at Alos16Point1Degrees and ending at tzais 3.8 deg.

The method return the time.Time of the plag.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanisAlos16Point1ToTzais3Point8
  - MinchaGedolaAhavatShalom
  - MinchaKetanaAhavatShalom
*/
func (t *complexZmanimCalendar) PlagAhavatShalom() (tm time.Time, ok bool) {
	tzaisGeonim3Point8Degrees, ok := t.TzaisGeonim3Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanisAlos16Point1ToTzais3Point8, ok := t.ShaahZmanisAlos16Point1ToTzais3Point8()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(tzaisGeonim3Point8Degrees, -gdt.GMillisecond(float64(shaahZmanisAlos16Point1ToTzais3Point8)*1.25)), true
}

/*
BainHasmashosRT13Point24Degrees the beginning of bain hashmashos of Rabbeinu Tam calculated when the sun is
zenith13Point24 13.24 deg below the western calculator.GeometricZenith (90 deg)
after sunset.
This calculation is based on the same calculation of BainHasmashosRT58Point5Minutes
bain hashmashos Rabbeinu Tam 58.5 minutes, but uses a degree-based calculation instead of 58.5 exact
minutes.
This calculation is based on the position of the sun 58.5 minutes after sunset in Jerusalem
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
which calculates to 13.24 deg below {calculator.GeometricZenith}.
NOTE: As per Yisrael Vehazmanim Vol. III page 1028, No. 50, a dip of slightly less than 13 deg should be used.
Calculations show that the proper dip to be 13.2456 deg (truncated to 13.24 that provides about 1.5 second
earlier (lechumra) time) below the horizon at that time. This makes a difference of 1 minute and 10
seconds in Jerusalem during the Equinox, and 1 minute 29 seconds during the solstice as compared to the proper
13.24 deg versus 13 deg. For NY during the solstice, the difference is 1 minute 56 seconds.
TO-DO recalculate the above based on equilux/equinox calculations.

The method return the time.Time of the sun being 13.24 deg below calculator.GeometricZenith (90 deg).
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- zenith13Point24
- BainHasmashosRT58Point5Minutes
*/
func (t *complexZmanimCalendar) BainHasmashosRT13Point24Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith13Point24)
}

/*
BainHasmashosRT58Point5Minutes returns the beginning of Bain hashmashos of Rabbeinu Tam calculated as a 58.5
minute offset after sunset. bain hashmashos is 3/4 of a Mil before tzais or 3 1/4
Mil after sunset. With a Mil calculated as 18 minutes, 3.25 * 18 = 58.5 minutes.

The method return the time.Time of 58.5 minutes after sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) BainHasmashosRT58Point5Minutes() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, gdt.GMinuteF64(58.5).ToMilliseconds()), true
}

/*
BainHasmashosRT13Point5MinutesBefore7Point083Degrees returns the beginning of bain hashmashos based on the calculation of 13.5 minutes (3/4 of an
18-minute Mil) before shkiah calculated as TzaisGeonim7Point083Degrees 7.083 deg.

The method return the time.Time of the bain hashmashos of Rabbeinu Tam in this calculation.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - TzaisGeonim7Point083Degrees
*/
func (t *complexZmanimCalendar) BainHasmashosRT13Point5MinutesBefore7Point083Degrees() (tm time.Time, ok bool) {
	sunsetOffsetByDegrees, ok := t.SunsetOffsetByDegrees(zenith7Point83)
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(sunsetOffsetByDegrees, -gdt.GMinuteF64(13.5).ToMilliseconds()), true
}

/*
BainHasmashosRT2Stars returns the beginning of bain hashmashos of Rabbeinu Tam calculated according to the
opinion of the Divrei Yosef (see Yisrael Vehazmanim) calculated 5/18th (27.77%) of the time between
alos (calculated as 19.8 deg before sunrise) and sunrise. This is added to sunset to arrive at the time
for bain hashmashos of Rabbeinu Tam.

The method return the time.Time of bain hashmashos of Rabbeinu Tam for this calculation.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) BainHasmashosRT2Stars() (tm time.Time, ok bool) {

	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}

	alos19Point8, ok := t.Alos19Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	sunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}

	return timeOffset(elevationAdjustedSunset, gdt.GMillisecond(float64(sunrise.UnixMilli()-alos19Point8.UnixMilli())*(5/float64(180)))), true // (5 / 18d))
}

/*
BainHasmashosYereim18Minutes returns the beginning of bain hashmashos (twilight) according to the
[Yereim (Rabbi Eliezer of Metz)]: https://en.wikipedia.org/wiki/Eliezer_ben_Samuel calculated as 18 minutes
or 3/4 of a 24-minute Mil before sunset. According to the Yereim, bain hashmashos starts 3/4
of a Mil before sunset and tzais or nightfall starts at sunset.

The method return the time.Time of 18 minutes before sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- BainHasmashosYereim3Point05Degrees
*/
func (t *complexZmanimCalendar) BainHasmashosYereim18Minutes() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, -gdt.GMinute(18).ToMilliseconds()), true
}

/*
BainHasmashosYereim3Point05Degrees returns the beginning of bain hashmashos (twilight) according to the
[Yereim (Rabbi Eliezer of Metz)]: https://en.wikipedia.org/wiki/Eliezer_ben_Samuel calculated as the sun's
position 3.05 deg above the horizon
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
its position 18 minutes or 3/4 of an 24-minute mil before sunset. According to the Yereim, bain
hashmashos starts 3/4 of a Mil before sunset and tzais or nightfall starts at sunset.
Note that lechumra (of about 14 seconds) a refraction value of 0.5166 deg as opposed to the traditional
0.566 deg is used. This is more inline with the actual refraction in Eretz Yisrael and is brought down
by
[Rabbi Yedidya Manet]: http://beinenu.com/rabbis/%D7%94%D7%A8%D7%91-%D7%99%D7%93%D7%99%D7%93%D7%99%D7%94-%D7%9E%D7%A0%D7%AA in his
[Zmanei Halacha Lema’aseh]: https://www.nli.org.il/en/books/NNL_ALEPH002542826/NLI (p. 11).
That is the first source that I am aware of that calculates degree-based Yereim zmanim.
The 0.5166 deg refraction is also used by the
[Luach Itim Lebinah]: https://zmanim.online/. Calculating the Yereim's bain hashmashos using 18-minute based degrees is also suggested
in the upcoming 8th edition of the zmanim Kehilchasam. For more details, see the article
[The Yereim’s Bein Hashmashos]: https://kosherjava.com/2020/12/07/the-yereims-bein-hashmashos/.

TO-TO recalculate based on equinox/equilux

The method return the time.Time of the sun's position 3.05 deg minutes before sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - zenithMinus3Point05
  - BainHasmashosYereim18Minutes
  - BainHasmashosYereim2Point8Degrees
  - BainHasmashosYereim2Point1Degrees
*/
func (t *complexZmanimCalendar) BainHasmashosYereim3Point05Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenithMinus3Point05)
}

/*
BainHasmashosYereim16Point875Minutes returns the beginning of bain hashmashos (twilight) according to the
[Yereim (Rabbi Eliezer of Metz)]: https://en.wikipedia.org/wiki/Eliezer_ben_Samuel calculated as 16.875
minutes or 3/4 of a 22.5-minute Mil before sunset. According to the Yereim, bain hashmashos
starts 3/4 of a Mil before sunset and tzais or nightfall starts at sunset.

The method return the time.Time of 16.875 minutes before sunset.

If the calculation can't be computed such as in the	Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - BainHasmashosYereim2Point8Degrees
*/
func (t *complexZmanimCalendar) BainHasmashosYereim16Point875Minutes() (tm time.Time, oki bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, -gdt.GMinuteF64(16.875).ToMilliseconds()), true
}

/*
BainHasmashosYereim2Point8Degrees returns the beginning of bain hashmashos (twilight) according to the
[Yereim (Rabbi Eliezer of Metz)]: https://en.wikipedia.org/wiki/Eliezer_ben_Samuel calculated as the sun's
position 2.8 deg above the horizon
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
its position 16.875 minutes or 3/4 of an 18-minute Mil before sunset. According to the Yereim, bain
hashmashos starts 3/4 of a Mil before sunset and tzais or nightfall starts at sunset.
Details, including how the degrees were calculated can be seen in the documentation of BainHasmashosYereim3Point05Degrees.

The method return the time.Time of the sun's position 2.8 deg minutes before sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - zenithMinus2Point8
  - BainHasmashosYereim16Point875Minutes
  - BainHasmashosYereim3Point05Degrees
  - BainHasmashosYereim2Point1Degrees
*/
func (t *complexZmanimCalendar) BainHasmashosYereim2Point8Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenithMinus2Point8)
}

/*
BainHasmashosYereim13Point5Minutes returns the beginning of bain hashmashos (twilight) according to the
[Yereim (Rabbi Eliezer of Metz)]: https://en.wikipedia.org/wiki/Eliezer_ben_Samuel calculated as 13.5 minutes
or 3/4 of an 18-minute Mil before sunset. According to the Yereim, bain hashmashos starts 3/4 of a Mil before sunset and tzais or nightfall starts at sunset.

The method return the time.Time of 13.5 minutes before sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - BainHasmashosYereim2Point1Degrees
*/
func (t *complexZmanimCalendar) BainHasmashosYereim13Point5Minutes() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, -gdt.GMinuteF64(13.5).ToMilliseconds()), true
}

/*
BainHasmashosYereim2Point1Degrees returns the beginning of bain hashmashos according to the
[Yereim (Rabbi Eliezer of Metz)]: https://en.wikipedia.org/wiki/Eliezer_ben_Samuel calculated as the sun's
position 2.1 deg above the horizon
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/ in
Yerushalayim, its position 13.5 minutes or 3/4 of an 18-minute Mil before sunset. According to the Yereim,
bain hashmashos starts 3/4 of a mil before sunset and tzais or nightfall starts at sunset.
Details, including how the degrees were calculated can be seen in the documentation of
BainHasmashosYereim3Point05Degrees.

The method return the time.Time of the sun's position 2.1 deg minutes before sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
- zenithMinus2Point1
- BainHasmashosYereim13Point5Minutes
- BainHasmashosYereim2Point8Degrees
- BainHasmashosYereim3Point05Degrees
*/
func (t *complexZmanimCalendar) BainHasmashosYereim2Point1Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenithMinus2Point1)
}

func (t *complexZmanimCalendar) AteretTorahSunsetOffset() gdt.GMinuteF64 {
	return t.ateretTorahSunsetOffset
}

/*
SofZmanShmaAteretTorah returns the latest zman krias shema (time to recite Shema in the morning) based on the
calculation of Chacham Yosef Harari-Raful of Yeshivat Ateret Torah, that the day starts
Alos72Zmanis 1/10th of the day before sunrise and is usually calculated as ending
TzaisAteretTorah 40 minutes after sunset (configurable to any offset via ateretTorahSunsetOffset).
shaos zmaniyos are calculated based on this day and added
to Alos72Zmanis alos to reach this time.
This time is 3 ShaahZmanisAteretTorah shaos zmaniyos (temporal hours) after
Alos72Zmanis alos 72 zmaniyos.
Note: Based on this calculation chatzos will not be at midday.

The method return the time.Time of the latest zman krias shema based on this calculation.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos72Zmanis
  - TzaisAteretTorah
  - AteretTorahSunsetOffset
  - ShaahZmanisAteretTorah
*/
func (t *complexZmanimCalendar) SofZmanShmaAteretTorah() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzaisAteretTorah, ok := t.TzaisAteretTorah()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(alos72Zmanis, tzaisAteretTorah), true
}

/*
SofZmanTfilahAteretTorah returns the latest zman tfila (time to recite the morning prayers) based on the calculation
of Chacham Yosef Harari-Raful of Yeshivat Ateret Torah, that the day starts Alos72Zmanis
1/10th of the day before sunrise and is usually calculated as ending TzaisAteretTorah 40 minutes
after sunset (configurable to any offset via ateretTorahSunsetOffset). shaos zmaniyos
are calculated based on this day and added to Alos72Zmanis alos to reach this time.
This time is 4 * ShaahZmanisAteretTorah shaos zmaniyos (temporal hours) after
Alos72Zmanis alos 72 zmaniyos.
Note: Based on this calculation chatzos will not be at midday.

The method return the time.Time of the latest zman krias shema based on this calculation.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos72Zmanis
  - TzaisAteretTorah
  - ShaahZmanisAteretTorah
  - ateretTorahSunsetOffset
*/
func (t *complexZmanimCalendar) SofZmanTfilahAteretTorah() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzaisAteretTorah, ok := t.TzaisAteretTorah()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(alos72Zmanis, tzaisAteretTorah), true
}

/*
MinchaGedolaAteretTorah returns the time of mincha gedola based on the calculation of Chacham Yosef
Harari-Raful of Yeshivat Ateret Torah, that the day starts Alos72Zmanis 1/10th of the day
before sunrise and is usually calculated as ending TzaisAteretTorah 40 minutes after sunset
(configurable to any offset via ateretTorahSunsetOffset). This is the preferred the earliest
time to pray mincha according to the opinion of the
[Rambam]: https://en.wikipedia.org/wiki/Maimonides and others.
For more information on this see the documentation on ZmanimCalendar.MinchaGedola.
This is calculated as 6.5 ShaahZmanisAteretTorah solar hours after alos.
The calculation used is 6.5 * ShaahZmanisAteretTorah after Alos72Zmanis alos.

see
  - Alos72Zmanis
  - TzaisAteretTorah()
  - ShaahZmanisAteretTorah()
  - ZmanimCalendar.MinchaGedola
  - MinchaKetanaAteretTorah
  - ZmanimCalendar.MinchaGedola
  - AteretTorahSunsetOffset

The method return the time.Time of the time of mincha gedola.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaGedolaAteretTorah() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzaisAteretTorah, ok := t.TzaisAteretTorah()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaGedola2(alos72Zmanis, tzaisAteretTorah), true
}

/*
MinchaKetanaAteretTorah returns the time of mincha ketana based on the calculation of
Chacham Yosef Harari-Raful of Yeshivat Ateret Torah, that the day starts
Alos72Zmanis 1/10th of the day before sunrise and is usually calculated as ending
TzaisAteretTorah 40 minutes after sunset (configurable to any offset via
ateretTorahSunsetOffset). This is the preferred the earliest time to pray mincha
according to the opinion of the [Rambam]: https://en.wikipedia.org/wiki/Maimonides and others.
For more information on this see the documentation on ZmanimCalendar.MinchaGedola.
This is calculated as 9.5 ShaahZmanisAteretTorah solar hours after Alos72Zmanis alos.
The calculation used is 9.5 * ShaahZmanisAteretTorah after Alos72Zmanis alos.
see
  - Alos72Zmanis()
  - TzaisAteretTorah
  - ShaahZmanisAteretTorah()
  - AteretTorahSunsetOffset
  - ZmanimCalendar.MinchaGedola
  - ZmanimCalendar.MinchaKetana

The method return the time.Time of the time of mincha ketana.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaKetanaAteretTorah() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzaisAteretTorah, ok := t.TzaisAteretTorah()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaKetana2(alos72Zmanis, tzaisAteretTorah), true
}

/*
PlagHaminchaAteretTorah returns the time of plag hamincha based on the calculation of Chacham Yosef Harari-Raful
of Yeshivat Ateret Torah, that the day starts Alos72Zmanis 1/10th of the day before sunrise and is
usually calculated as ending TzaisAteretTorah 40 minutes after sunset (configurable to any offset
via ateretTorahSunsetOffset). shaos zmaniyos are calculated based on this day and
added to Alos72Zmanis alos to reach this time.
This time is 10.75 ShaahZmanisAteretTorah shaos zmaniyos (temporal hours) after Alos72Zmanis dawn.

The method return the time.Time of the plag.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos72Zmanis
  - TzaisAteretTorah
  - ShaahZmanisAteretTorah
  - AteretTorahSunsetOffset
*/
func (t *complexZmanimCalendar) PlagHaminchaAteretTorah() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzaisAteretTorah, ok := t.TzaisAteretTorah()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos72Zmanis, tzaisAteretTorah), true
}

/*
TzaisGeonim3Point7Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated at the
sun's position at zenith3Point7 3.7 deg below the western horizon.

The method return the time.Time representing the time when the sun is 3.7 deg below sea level.
*/
func (t *complexZmanimCalendar) TzaisGeonim3Point7Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith3Point7)
}

/*
TzaisGeonim3Point8Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated at the
sun's position at zenith3Point8 3.8 deg below the western horizon.

The method return the time.Time representing the time when the sun is 3.8 deg below sea level.
*/
func (t *complexZmanimCalendar) TzaisGeonim3Point8Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith3Point8)
}

/*
TzaisGeonim5Point95Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated at the
sun's position at zenith5Point95 5.95 deg below the western horizon.

The method return the time.Time representing the time when the sun is 5.95 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim5Point95Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith5Point95)
}

/*
TzaisGeonim3Point65Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated as 3/4
of a [Mil]: https://en.wikipedia.org/wiki/Biblical_and_Talmudic_units_of_measurement based on an 18
minutes Mil, or 13.5 minutes. It is the sun's position at zenith3Point65 3.65 deg below the western
horizon. This is a very early zman and should not be relied on without Rabbinical guidance.

The method return the time.Time representing the time when the sun is 3.65 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim3Point65Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith3Point65)
}

/*
TzaisGeonim3Point676Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated as 3/4
of a [Mil]: https://en.wikipedia.org/wiki/Biblical_and_Talmudic_units_of_measurement based on 18-minute Mil, or 13.5 minutes. It is the sun's position at zenith3Point676 3.676 deg below the western
horizon based on the calculations of Stanley Fishkind. This is a very early zman and should not be
relied on without Rabbinical guidance.

The method return the time.Time representing the time when the sun is 3.676 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim3Point676Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith3Point676)
}

/*
TzaisGeonim4Point61Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated as 3/4
of a [mil]: https://en.wikipedia.org/wiki/Biblical_and_Talmudic_units_of_measurement based
on a 24-minute Mil, or 18 minutes. It is the sun's position at zenith4Point61 4.61 deg below the
western horizon. This is a very early zman and should not be relied on without Rabbinical guidance.

The method return the time.Time representing the time when the sun is 4.61 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim4Point61Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith4Point61)
}

/*
TzaisGeonim4Point37Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated as 3/4
of a [Mil]: https://en.wikipedia.org/wiki/Biblical_and_Talmudic_units_of_measurement, based on a 22.5
minute Mil, or 16 7/8 minutes. It is the sun's position at zenith4Point37 4.37 deg below the western
horizon. This is a very early zman and should not be relied on without Rabbinical guidance.

The method return the time.Time representing the time when the sun is 4.37 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim4Point37Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith4Point37)
}

/*
TzaisGeonim5Point88Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated as 3/4
of a 24-minute [Mil]: https://en.wikipedia.org/wiki/Biblical_and_Talmudic_units_of_measurement,
based on a Mil being 24 minutes, and is calculated as 18 + 2 + 4 for a total of 24 minutes. It is the
sun's position at zenith5Point88 5.88 deg below the western horizon. This is a very early
zman and should not be relied on without Rabbinical guidance.

TO-DO Additional detailed documentation needed.
The method return the time.Time representing the time when the sun is 5.88 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim5Point88Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith5Point88)
}

/*
TzaisGeonim4Point8Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated as 3/4
of a [Mil]: "https://en.wikipedia.org/wiki/Biblical_and_Talmudic_units_of_measurement based on the
sun's position at zenith4Point8 4.8 deg below the western horizon. This is based on Rabbi Leo Levi's
calculations. This is a very early zman and should not be relied on without Rabbinical guidance.
TO-DO Additional documentation needed.

The method return the time.Time representing the time when the sun is 4.8 deg below sea level.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
oks is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim4Point8Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith4Point8)
}

/*
TzaisGeonim6Point45Degrees returns the tzais (nightfall) based on the opinion of the Geonim as calculated by
[Rabbi Yechiel Michel Tucazinsky]: https://en.wikipedia.org/wiki/Yechiel_Michel_Tucazinsky.
It is based on of the position of the sun no later than TzaisGeonim6Point45Degrees 31 minutes after sunset
in Jerusalem the height of the summer solstice and is 28 minutes after shkiah
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/.
This computes to 6.45 deg below the western horizon.
TO-DO Additional documentation details needed.

The method return the time.Time representing the time when the sun is 6.45 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim6Point45Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith6Point45)
}

/*
TzaisGeonim7Point083Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated when the
sun's position zenith7Point83 7.083 deg (or 7 deg 5\u2032) below the western horizon.
This is often referred to as 7 deg5 or 7 deg and 5 minutes.
This calculation is based on the observation of 3 medium-sized stars by Dr. Baruch (Berthold) Cohn in his luach
[Tabellen enthaltend die Zeitangaben für den Beginn der Nacht und des Tages für die Breitengrade + 66 bis -38]: https://sammlungen.ub.uni-frankfurt.de/freimann/content/titleinfo/983088 published in Strasbourg, France in 1899.
This calendar was very popular in Europe, and many other calendars based their time on it.
[Rav Dovid Tzvi Hoffman]: https://en.wikipedia.org/wiki/David_Zvi_Hoffmann in his
[Sh"Ut Melamed Leho'il]: https://hebrewbooks.org/1053 in an exchange of letters with Baruch Cohn in
[Orach Chaim 30]: https://hebrewbooks.org/pdfpager.aspx?req=1053&st=&pgnum=37 agreed to this zman (page 36),
as did the Sh"Ut Bnei Tziyon and the Tenuvas Sadeh. It is very close to the time of the
[Mekor Chesed]: https://hebrewbooks.org/22044 of the Sefer chasidim.
It is close to the position of the sun 30 minutes after sunset in Jerusalem
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
but not Exactly. The actual position of the sun 30 minutes after sunset in Jerusalem at the equilux is 7.205 deg and 7.199 deg
at the equinox. See Hazmanim Bahalacha vol 2, pages 520-521 for more details.

The method return the time.Time representing the time when the sun is 7.083 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned. See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim7Point083Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith7Point83)
}

/*
TzaisGeonim7Point67Degrees returns tzais (nightfall) based on the opinion of the Geonim calculated as 45 minutes
after sunset during the summer solstice in New York, when the neshef (twilight) is the longest. The sun's
position at this time computes to {@link #ZENITH_7_POINT_67 7.75 deg} below the western horizon.
See [Igros Moshe Even Haezer 4, Ch. 4]: https://hebrewbooks.org/pdfpager.aspx?req=921&pgnum=149 (regarding
tzais for krias Shema). It is also mentioned in Rabbi Heber's
[Shaarei Zmanim]: https://hebrewbooks.org/53000 on in
[chapter 10 (page 87)]: https://hebrewbooks.org/pdfpager.aspx?req=53055&pgnum=101 and
[chapter 12 (page 108)]: https://hebrewbooks.org/pdfpager.aspx?req=53055&pgnum=122.
Also see the time of 45 minutes in
[Rabbi Simcha Bunim Cohen's]: https://en.wikipedia.org/wiki/Simcha_Bunim_Cohen
[The radiance of Shabbos]: https://www.worldcat.org/oclc/179728985 as the earliest zman for New York.
This zman is also listed in the
[Divrei Shalom Vol. III, chapter 75]: https://hebrewbooks.org/pdfpager.aspx?req=1927&pgnum=90,
and
[Bais Av Vol. III, chapter 117]: https://hebrewbooks.org/pdfpager.aspx?req=892&pgnum=431.
This zman is also listed in the Divrei Shalom etc. chapter 177
Since this zman depends on the level of light, Rabbi Yaakov Shakow presented this degree-based
calculation to Rabbi [Rabbi Shmuel Kamenetsky]: https://en.wikipedia.org/wiki/Shmuel_Kamenetsky who agreed to it.
TO-DO add hyperlinks to source of Divrei Shalom.

The method return the time.Time representing the time when the sun is 7.67 deg below sea level.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and
north of the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned. See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim7Point67Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith7Point67)
}

/*
TzaisGeonim8Point5Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated at the
sun's position at zenith8Point5 8.5 deg below the western horizon.

The method return the time.Time representing the time when the sun is 8.5 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim8Point5Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith8Point5)
}

/*
TzaisGeonim9Point3Degrees returns the tzais (nightfall) based on the calculations used in the
[Luach Itim Lebinah]: https://www.worldcat.org/oclc/243303103 as the stringent time for tzais.
It is calculated at the sun's position at zenith9Point3 9.3 deg below the western horizon.

The method return the time.Time representing the time when the sun is 9.3 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim9Point3Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith9Point3)
}

/*
TzaisGeonim9Point75Degrees returns the tzais (nightfall) based on the opinion of the Geonim calculated as 60
minutes after sunset
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/, the
day that a solar hour is 60 minutes in New York. The sun's position at this time computes to
zenith9Point75 9.75 deg below the western horizon. This is the opinion of
[Rabbi Eliyahu Henkin]: https://en.wikipedia.org/wiki/Yosef_Eliyahu_Henkin.
This also follows the opinion of
[Rabbi Shmuel Kamenetsky]: https://en.wikipedia.org/wiki/Shmuel_Kamenetsky. Rabbi Yaakov Shakow presented
these degree-based times to Rabbi Shmuel Kamenetsky who agreed to them.

TO-DO recalculate based on equinox / equilux.

The method return the time.Time representing the time when the sun is 9.75 deg below sea level.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of
the Antarctic Circle, where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisGeonim9Point75Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith9Point75)
}

/*
Tzais60 returns the tzais (nightfall) based on the opinion of the
[Chavas Yair]: https://en.wikipedia.org/wiki/Yair_Bacharach and
[Divrei Malkiel]: https://he.wikipedia.org/wiki/%D7%9E%D7%9C%D7%9B%D7%99%D7%90%D7%9C_%D7%A6%D7%91%D7%99_%D7%98%D7%A0%D7%A0%D7%91%D7%95%D7%99%D7%9D,
that the time to walk the distance of a Mil is 15 minutes for a total of 60 minutes for 4 Mil after AstronomicalCalendar.SeaLevelSunset.
See detailed documentation explaining the 60 minutes concept at Alos60.

The method return the time.Time representing 60 minutes after sea level sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos60
  - PlagHamincha60Minutes
  - ShaahZmanis60Minutes
*/
func (t *complexZmanimCalendar) Tzais60() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, gdt.GMinute(60).ToMilliseconds()), true
}

/*
TzaisAteretTorah returns tzais usually calculated as 40 minutes (configurable to any offset via
ateretTorahSunsetOffset) after sunset. Please note that Chacham Yosef Harari-Raful
of Yeshivat Ateret Torah who uses this time, does so only for calculating various other zmanai hayom
such as Sof Zman Krias Shema and Plag Hamincha. His calendars do not publish a zman
for Tzais. It should also be noted that Chacham Harari-Raful provided a 25 minute zman
for Israel. This API uses 40 minutes year round in any place on the globe by default.

The method return the time.Time representing 40 minutes (configurable via ateretTorahSunsetOffset)
after sea level sunset.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, a nil will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - AteretTorahSunsetOffset
*/
func (t *complexZmanimCalendar) TzaisAteretTorah() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, t.AteretTorahSunsetOffset().ToMilliseconds()), true
}

/*
Tzais72Zmanis return tzais (dusk) calculated as 72 minutes zmaniyos, or 1/10th of the day after
AstronomicalCalendar.SeaLevelSunset. This is the way that the
[Minchas Cohen]: https://en.wikipedia.org/wiki/Abraham_Cohen_Pimentel in Ma'amar 2:4 calculates Rebbeinu Tam's
time of tzeis. It should be noted that this calculation results in the shortest time from sunset to
tzais being during the winter solstice, the longest at the summer solstice and 72 clock minutes at the
equinox. This does not match reality, since there is no direct relationship between the length of the day and
twilight. The shortest twilight is during the equinox, the longest is during the summer solstice, and in the
winter with the shortest daylight, the twilight period is longer than during the equinoxes.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
ses
  - Alos72Zmanis
*/
func (t *complexZmanimCalendar) Tzais72Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(1.2)
}

/*
zmanisBasedOffset is a utility method to return alos (dawn) or tzais (dusk) based on a fractional day offset.
hours the number of shaaos zmaniyos (temporal hours) before sunrise or after sunset that defines dawn or dusk.
If a negative number is passed in, it will return the time of alos (dawn) (subtracting the time from sunrise)
and if a positive number is passed in, it will return the time of tzais (dusk) (adding the time to sunset).
If 0 is passed in, ok is false will be returned (since we can't tell if it is sunrise or sunset based).

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
helper.Panic will also be invoked if 0 is passed in, since we can't tell if it is sunrise or sunset based.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) zmanisBasedOffset(hours gdt.GHourH64) (tm time.Time, ok bool) {
	shaahZmanis, ok := t.ShaahZmanisGRA()
	if !ok {
		return time.Time{}, false
	}
	if hours == 0 {
		helper.Panic("hours == 0")
	}

	if hours > 0 {
		elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
		if !ok {
			return time.Time{}, false
		}
		return timeOffset(elevationAdjustedSunset, gdt.GMillisecond(float64(shaahZmanis)*float64(hours))), true
	} else {
		elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
		if !ok {
			return time.Time{}, false
		}
		return timeOffset(elevationAdjustedSunrise, gdt.GMillisecond(float64(shaahZmanis)*float64(hours))), true
	}
}

/*
Tzais90Zmanis return tzais (dusk) calculated using 90 minutes zmaniyos or 1/8th of the day after
AstronomicalCalendar.SeaLevelSunset.
This time is known in Yiddish as the achtel (an eighth) zman.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar	documentation.

see
  - Alos90Zmanis
*/
func (t *complexZmanimCalendar) Tzais90Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(1.5)
}

/*
Tzais96Zmanis return tzais (dusk) calculated using 96 minutes zmaniyos or 1/7.5 of the day after
AstronomicalCalendar.SeaLevelSunset.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Alos96Zmanis
*/
func (t *complexZmanimCalendar) Tzais96Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(1.6)
}

/*
Tzais90 return tzais (dusk) calculated as 90 minutes after sea level sunset.
This method returns tzais (nightfall) based on the opinion of the Magen Avraham that the time to walk the distance of a
Mil according to the [Rambam]: https://en.wikipedia.org/wiki/Maimonides opinion is 18 minutes for a total of 90 minutes
based on the opinion of Ula who calculated tzais as 5 Mil after sea level shkiah (sunset).
A similar calculation Tzais19Point8Degrees uses solar position calculations based on this time.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Tzais19Point8Degrees
  - Alos90
*/
func (t *complexZmanimCalendar) Tzais90() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	return timeOffset(elevationAdjustedSunset, gdt.GMinute(90).ToMilliseconds()), true
}

/*
Tzais16Point1Degrees calculates the time of tzais at the point when the sun is 16.1 deg below the horizon.
This is the sun's dip below the horizon 72 minutes after sunset according Rabbeinu Tam's calculation of tzais
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/ in Jerusalem.
The question of equinox VS equilux is complex, with Rabbi Meir Posen in the
[Ohr Meir]: https://www.worldcat.org/oclc/956316270 of the opinion that the equilux should be used.
See Yisrael Vehazmanim vol I, 34:1:4. Rabbi Yedidya Manet in his
[Zmanei Halacha Lema'aseh]: https://www.nli.org.il/en/books/NNL_ALEPH002542826/NLI (4th edition part 2, pages
and 22 and 24) and Rabbi Yonah Metzbuch (in a letter published by Rabbi Manet) are of the opinion that the
astronomical equinox should be used. The difference adds up to about 9 seconds, too trivial to make much of a
difference. For information on how this is calculated see the comments on Alos16Point1Degrees.

The method return the time.Time representing the time.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle
and north of the Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ZmanimCalendar.Tzais72
  - Alos16Point1Degrees for more information on this calculation.
*/
func (t *complexZmanimCalendar) Tzais16Point1Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith16Point1)
}

/*
Tzais18Degrees return the time.Time representing the time. If the calculation can't be computed such as northern and
southern locations even south of the Arctic Circle and north of the Antarctic Circle where the sun may
not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
For information on how this is calculated see the comments on {@link Alos18Degrees()}

see
  - Alos18Degrees
*/
func (t *complexZmanimCalendar) Tzais18Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(calculator.AstronomicalZenith)
}

/*
Tzais19Point8Degrees return the time.Time representing the time.
For information on how this is calculated see the comments on Alos19Point8Degrees

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Tzais90
  - Alos19Point8Degrees
*/
func (t *complexZmanimCalendar) Tzais19Point8Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith19Point8)
}

/*
Tzais96 return tzais (dusk) calculated as 96 minutes after sea level sunset.
For information on how this is calculated see the comments on Alos96.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Alos96
*/
func (t *complexZmanimCalendar) Tzais96() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, gdt.GMinute(96).ToMilliseconds()), true
}

/*
FixedLocalChatzos that returns the local time for fixed chatzos.
This time is noon and midnight adjusted from standard time to account for the local latitude.
The 360 deg of the globe divided by 24 calculates to 15 deg per hour with 4 minutes per degree,
so at a longitude of 0 , 15, 30 etc... ZmanimCalendar.Chatzos is at exactly 12:00 noon.
This is the time of chatzos according to the
[Aruch Hashulchan]: https://en.wikipedia.org/wiki/Aruch_HaShulchan in
[Orach Chaim 233:14]: https://hebrewbooks.org/pdfpager.aspx?req=7705&pgnum=426 and
[Rabbi Moshe Feinstein]: https://en.wikipedia.org/wiki/Moshe_Feinstein in Igros Moshe
[Orach Chaim 1:24]: https://hebrewbooks.org/pdfpager.aspx?req=916&st=&pgnum=67 and
[2:20]: https://hebrewbooks.org/pdfpager.aspx?req=14675&pgnum=191.
Lakewood, N.J., with a longitude of -74.2094, is 0.7906 away from the closest multiple of 15 at -75 deg. This
is multiplied by 4 to yield 3 minutes and 10 seconds for a chatzos of 11:56:50.
This method is not tied to the theoretical 15 deg timezones, but will adjust to the actual timezone and
[Daylight saving time]: https://en.wikipedia.org/wiki/Daylight_saving_time.

The method return the time.Time representing the local chatzos
*/
func (t *complexZmanimCalendar) FixedLocalChatzos() time.Time {
	return timeOffset(t.dateTimeFromTimeOfDay(12.0-float64(t.GeoLocation().StandardTimeOffset())/float64(timeutil.HourMillis), true), -t.GeoLocation().LocalMeanTimeOffset())
}

func isJewishDayOfMonthBetween(jewishCalendar hebrewcalendar.JewishCalendar, from jdt.JDay, to jdt.JDay) {
	if jewishCalendar.JewishDate().JDay() < from {
		helper.Panic(fmt.Sprintf("jewishCalendar.JDay() < %d)", from))
	}
	if jewishCalendar.JewishDate().JDay() > to {
		helper.Panic(fmt.Sprintf("jewishCalendar.JDay() > %d)", to))
	}

}

/*
isJewishMonthBetween11And16
Do not calculate for impossible dates, but account for extreme cases. In the extreme case of Rapa Iti in French
Polynesia on Dec 2027 when kiddush Levana 3 days can be said on Rosh Chodesh, the sof zman Kiddush Levana
will be on the 12th of the Teves. In the case of Anadyr, Russia on Jan, 2071, sof zman Kiddush Levana between the
moldos will occur is on the night of 17th of Shevat. See Rabbi Dovid Heber's Shaarei Zmanim chapter 4 (pages 28 and 32).
*/
func isJewishMonthBetween11And16(jewishCalendar hebrewcalendar.JewishCalendar) {
	isJewishDayOfMonthBetween(jewishCalendar, 11, 16)
}

/*
isJewishMonthBetween5And30
Do not calculate for impossible dates, but account for extreme cases. Tchilas zman kiddush Levana 3 days for
the extreme case of Rapa Iti in French Polynesia on Dec 2027 when kiddush Levana 3 days can be said on the evening
of the 30th, the second night of Rosh Chodesh. The 3rd day after the molad will be on the 4th of the month.
In the case of Anadyr, Russia on Jan, 2071, when sof zman kiddush levana is on the 17th of the month, the 3rd day
from the molad will be on the 5th day of Shevat. See Rabbi Dovid Heber's Shaarei Zmanim chapter 4 (pages 28 and 32).
*/
func isJewishMonthBetween5And30(jewishCalendar hebrewcalendar.JewishCalendar) {
	isJewishDayOfMonthBetween(jewishCalendar, 5, 30)
}

/*
isJewishMonthBetween2And27
Optimize to not calculate for impossible dates, but account for extreme cases. The molad in the extreme case of Rapa
Iti in French Polynesia on Dec 2027 occurs on the night of the 27th of Kislev. In the case of Anadyr, Russia on
Jan 2071, the molad will be on the 2nd day of Shevat. See Rabbi Dovid Heber's Shaarei Zmanim chapter 4 (pages 28 and 32).
*/
func isJewishMonthBetween2And27(jewishCalendar hebrewcalendar.JewishCalendar) {
	isJewishDayOfMonthBetween(jewishCalendar, 2, 27)
}

/*
isJewishMonthBetween4And9
Optimize to not calculate for impossible dates, but account for extreme cases. Tchilas zman kiddush Levana 7 days for
the extreme case of Rapa Iti in French Polynesia on Jan 2028 (when kiddush Levana 3 days can be said on the evening
of the 30th, the second night of Rosh Chodesh), the 7th day after the molad will be on the 4th of the month.
In the case of Anadyr, Russia on Jan, 2071, when sof zman kiddush levana is on the 17th of the month, the 7th day
from the molad will be on the 9th day of Shevat. See Rabbi Dovid Heber's Shaarei Zmanim chapter 4 (pages 28 and 32).
*/
func isJewishMonthBetween4And9(jewishCalendar hebrewcalendar.JewishCalendar) {
	isJewishDayOfMonthBetween(jewishCalendar, 4, 9)
}

/*
SofZmanKidushLevanaBetweenMoldos1 returns the latest time of Kidush Levana according to the
[Maharil's]: https://en.wikipedia.org/wiki/Yaakov_ben_Moshe_Levi_Moelin opinion that it is calculated as
halfway between molad and molad. This adds half the 29 days, 12 hours and 793 chalakim time
between molad and molad (14 days, 18 hours, 22 minutes and 666 milliseconds) to the month's molad.
If the time of sof zman Kiddush Levana occurs during the day (between the alos and tzais passed in
as parameters), it returns the alos passed in. If a nil alos or tzais are passed to this method,
the non-daytime adjusted time will be returned.

alos the beginning of the Jewish day.
If Kidush Levana occurs during the day (starting at alos and ending at tzais), the time returned will be alos.
If either the alos or tzais parameters are nil, no daytime adjustment will be made.

tzais the end of the Jewish day.
If Kidush Levana occurs during the day (starting at alos and ending at tzais), the time returned will be alos.
If either the alos or tzais parameter are nil, no daytime adjustment will be made.

The method return the time.Time representing the moment halfway between molad and molad.

# If the time occurs between alos and tzais, alos will be returned

see
  - JewishCalendar.SofZmanKidushLevanaBetweenMoldos
  - JewishCalendar.SofZmanKidushLevana15Days
  - JewishCalendar.SofZmanKidushLevanaBetweenMoldos
*/
func (t *complexZmanimCalendar) SofZmanKidushLevanaBetweenMoldos1(alos *time.Time, tzais *time.Time) (tm time.Time, ok bool) {
	jewishCalendar := hebrewcalendar.NewJewishCalendar(hebrewcalendar.NewJewishDate2(t.gDateTime.D))

	isJewishMonthBetween11And16(jewishCalendar)
	return t.moladBasedTime(jewishCalendar.SofZmanKidushLevanaBetweenMoldos(), alos, tzais, false)
}

/*
moladBasedTime returns the time.Time of the molad based time if it occurs on the current date.
Since Kiddush Levana can only be said during the day, there are parameters to limit it to between alos and tzais.
If the time occurs between alos and tzais, tzais will be returned.
moladBasedTime the molad based time such as molad, tchilas and sof zman Kiddush Levana
alos optional start of day to limit molad times to the end of the night before or beginning of the next night.
Ignored if either alos or tzais are nil.
tzais optional end of day to limit molad times to the end of the night before or beginning of the next night.
Ignored if either tzais or alos are nil
techila is it the start of Kiddush Levana time or the end? If it is start roll it to the next tzais,
and if it is the end, return the end of the previous night (alos passed in).
Ignored if either alos or tzais are nil.

The method return the molad based time.

If the zman does not occur during the current date, ok is false will be returned.
*/
func (t *complexZmanimCalendar) moladBasedTime(moladBasedTime time.Time, alos *time.Time, tzais *time.Time, techila bool) (tm time.Time, ok bool) {
	lastMidnight := t.midnightLastNight()
	midnightTonigh := t.midnightTonight()

	if moladBasedTime.Before(lastMidnight) {
		return time.Time{}, false
	}
	if moladBasedTime.After(midnightTonigh) {
		return time.Time{}, false
	}
	if alos != nil || tzais != nil {
		if techila && !(moladBasedTime.Before(*tzais) || moladBasedTime.After(*alos)) {
			return *tzais, true
		} else {
			return *alos, true
		}
	}

	return moladBasedTime, true
}

/*
SofZmanKidushLevanaBetweenMoldos2 returns the latest time of Kiddush Levana according to the
[Maharil's]: https://en.wikipedia.org/wiki/Yaakov_ben_Moshe_Levi_Moelin opinion that it is calculated as
halfway between molad and molad. This adds half the 29 days, 12 hours and 793 chalakim time between
molad and molad (14 days, 18 hours, 22 minutes and 666 milliseconds) to the month's molad.
The sof zman Kiddush Levana will be returned even if it occurs during the day. To limit the time to between
tzais and alos, see SofZmanKidushLevanaBetweenMoldos1.
The method return the Date representing the moment halfway between molad and molad.
If the time occurs between alos and tzais, alos will be returned
see
  - SofZmanKidushLevanaBetweenMoldos1
  - SofZmanKidushLevana15Days1
  - JewishCalendar.SofZmanKidushLevanaBetweenMoldos
*/
func (t *complexZmanimCalendar) SofZmanKidushLevanaBetweenMoldos2() (tm time.Time, ok bool) {
	return t.SofZmanKidushLevanaBetweenMoldos1(nil, nil)
}

/*
SofZmanKidushLevana15Days1 returns the latest time of Kiddush Levana calculated as 15 days after the molad.
This is the opinion brought down in the Shulchan Aruch (Orach Chaim 426).
It should be noted that some opinions hold that the
[Rema]: https://en.wikipedia.org/wiki/Moses_Isserles who brings down the opinion of the
[Maharil's]: https://en.wikipedia.org/wiki/Yaakov_ben_Moshe_Levi_Moelin of calculating
SofZmanKidushLevanaBetweenMoldos1 halfway between molad and molad is of
the opinion that the Mechaber agrees to his opinion.
Also see the Aruch Hashulchan. For additional details on the subject,
see Rabbi Dovid Heber's very detailed write-up in Siman Daled (chapter 4) of
[Shaarei Zmanim]: https://hebrewbooks.org/53000. If the time of sof zman Kiddush Levana occurs during
the day (between the alos and tzais passed in as parameters), it returns the alos passed in.
If a nil alos or tzais are passed to this method, the non-daytime adjusted time will be returned.
alos the beginning of the Jewish day. If Kidush Levana occurs during the day (starting at alos and ending at tzais),
the time returned will be alos. If either the alos or tzais parameters are nil, no daytime adjustment will be made.
tzais the end of the Jewish day. If Kidush Levana occurs during the day (starting at alos and ending at tzais),
the time returned will be alos. If either the alos or tzais parameters are nil, no daytime adjustment will be made.

The method return the Date representing the moment 15 days after the molad.

# If the time occurs between alos and tzais, alos will be returned

see
  - SofZmanKidushLevanaBetweenMoldos1
  - JewishCalendar.SofZmanKidushLevana15Days
*/
func (t *complexZmanimCalendar) SofZmanKidushLevana15Days1(alos *time.Time, tzais *time.Time) (tm time.Time, ok bool) {
	jewishCalendar := hebrewcalendar.NewJewishCalendar(hebrewcalendar.NewJewishDate2(t.gDateTime.D))

	isJewishMonthBetween11And16(jewishCalendar)
	return t.moladBasedTime(jewishCalendar.SofZmanKidushLevana15Days(), alos, tzais, false)
}

/*
SofZmanKidushLevana15Days2 returns the latest time of Kiddush Levana calculated as 15 days after the molad.
This is the opinion of the Shulchan Aruch (Orach Chaim 426). It should be noted that some opinions hold that the
[Rema]: https://en.wikipedia.org/wiki/Moses_Isserles who brings down the opinion of the
[Maharil's]: https://en.wikipedia.org/wiki/Yaakov_ben_Moshe_Levi_Moelin of calculating
SofZmanKidushLevanaBetweenMoldos1 halfway between molad and molad is of
the opinion that the Mechaber agrees to his opinion. Also see the Aruch Hashulchan.
For additional details on the subject, See Rabbi Dovid Heber's very detailed write-up in Siman Daled (chapter 4) of
[ShaareiZmanim]: https://hebrewbooks.org/53000.
The sof zman Kiddush Levana will be returned even if it occurs during the day. To limit the time to
between tzais and alos, see {@link #getSofZmanKidushLevana15Days}.

The method return the Date representing the moment 15 days after the molad. If the time occurs between alos and tzais, alos will be returned

see
  - SofZmanKidushLevana15Days1
  - SofZmanKidushLevanaBetweenMoldos1
  - JewishCalendar.SofZmanKidushLevana15Days()
*/
func (t *complexZmanimCalendar) SofZmanKidushLevana15Days2() (tm time.Time, ok bool) {
	return t.SofZmanKidushLevana15Days1(nil, nil)
}

/*
TchilasZmanKidushLevana3Days2 returns the earliest time of Kiddush Levana according to
[Rabbeinu Yonah]: https://en.wikipedia.org/wiki/Yonah_Gerondi opinion that it can be said 3 days after the
molad.
The time will be returned even if it occurs during the day when Kiddush Levana can't be said.
Use TchilasZmanKidushLevana3Days1 if you want to limit the time to night hours.

The method return the time.Time representing the moment 3 days after the molad.
see
  - TchilasZmanKidushLevana3Days1
  - TchilasZmanKidushLevana7Days1
  - JewishCalendar.TchilasZmanKidushLevana3Days
*/
func (t *complexZmanimCalendar) TchilasZmanKidushLevana3Days2() (tm time.Time, ok bool) {
	return t.TchilasZmanKidushLevana3Days1(nil, nil)
}

/*
TchilasZmanKidushLevana3Days1 returns the earliest time of Kiddush Levana according to
[Rabbeinu Yonah]: https://en.wikipedia.org/wiki/Yonah_Gerondi opinion that it can be said 3 days after the molad.
If the time of tchilas zman Kiddush Levana occurs during the day (between alos and tzais passed to
this method) it will return the following tzais. If nil is passed for either alos or tzais, the actual
tchilas zman Kiddush Levana will be returned, regardless of if it is during the day or not.
alos the beginning of the Jewish day. If Kidush Levana occurs during the day (starting at alos and ending at tzais),
the time returned will be tzais. If either the alos or tzais parameters re nil, no daytime adjustment will be made.
tzais the end of the Jewish day. If Kidush Levana occurs during the day (starting at alos and ending at	tzais),
the time returned will be tzais. If either the alos or tzais parameters are nil, no daytime adjustment will be made.

The method return the time.Time representing the moment 3 days after the molad.
If the time occurs between alos and tzais, tzais will be returned

see
  - TchilasZmanKidushLevana3Days2
  - TchilasZmanKidushLevana7Days2
  - JewishCalendar.TchilasZmanKidushLevana3Days
*/
func (t *complexZmanimCalendar) TchilasZmanKidushLevana3Days1(alos *time.Time, tzais *time.Time) (tm time.Time, ok bool) {
	jewishCalendar := hebrewcalendar.NewJewishCalendar(hebrewcalendar.NewJewishDate2(t.gDateTime.D))

	isJewishMonthBetween5And30(jewishCalendar)

	zman, ok := t.moladBasedTime(jewishCalendar.TchilasZmanKidushLevana3Days(), alos, tzais, true)
	if !ok {
		return time.Time{}, false
	}

	//Get the following month's zman kiddush Levana for the extreme case of Rapa Iti in French Polynesia on Dec 2027 when
	// kiddush Levana can be said on Rosh Chodesh (the evening of the 30th). See Rabbi Dovid Heber's Shaarei Zmanim chapter 4 (page 32)
	if jewishCalendar.JewishDate().JDay() == 30 {
		jewishCalendar.JewishDate().ForwardJMonth(1)
		zman, ok = t.moladBasedTime(jewishCalendar.TchilasZmanKidushLevana3Days(), nil, nil, true)
		if !ok {
			return time.Time{}, false
		}
	}

	return zman, true
}

/*
ZmanMolad returns the point in time of Molad as a time.Time Object.
The method return the time.Time representing the moment of the molad.
If the molad does not occur on this day, ok is false will be returned.

see
  - TchilasZmanKidushLevana3Days1, TchilasZmanKidushLevana3Days2
  - TchilasZmanKidushLevana7Days1, TchilasZmanKidushLevana7Days2
  - JewishCalendar.MoladAsDate
*/
func (t *complexZmanimCalendar) ZmanMolad() (tm time.Time, ok bool) {
	jewishCalendar := hebrewcalendar.NewJewishCalendar(hebrewcalendar.NewJewishDate2(t.gDateTime.D))

	isJewishMonthBetween2And27(jewishCalendar)
	molad, ok := t.moladBasedTime(jewishCalendar.MoladAsDate(), nil, nil, true)
	if !ok {
		return time.Time{}, false
	}

	// deal with molad that happens on the end of the previous month
	if jewishCalendar.JewishDate().JDay() > 26 {
		jewishCalendar.JewishDate().ForwardJMonth(1)
		molad, ok = t.moladBasedTime(jewishCalendar.MoladAsDate(), nil, nil, true)
		if !ok {
			return time.Time{}, false
		}
	}
	return molad, true
}

/*
midnightLastNight is used by Molad based zmanim to determine if zmanim occur during the current day.
see
  - moladBasedTime

The method return previous midnight
*/
func (t *complexZmanimCalendar) midnightLastNight() time.Time {
	return time.Date(int(t.gDateTime.D.Year), t.gDateTime.D.Month, int(t.gDateTime.D.Day), 0, 0, 0, 0, t.GeoLocation().TimeZone())
}

/*
midnightTonight is used by Molad based zmanim to determine if zmanim occur during the current day.
see
  - moladBasedTime

The method return following midnight
*/
func (t *complexZmanimCalendar) midnightTonight() time.Time {
	midnight := time.Date(int(t.gDateTime.D.Year), t.gDateTime.D.Month, int(t.gDateTime.D.Day), 0, 0, 0, 0, t.GeoLocation().TimeZone())
	// midnight.add(Calendar.DAY_OF_YEAR, 1)//roll to tonight
	return midnight.AddDate(0, 0, 1) // roll to tonight ???
}

/*
TchilasZmanKidushLevana7Days1 returns the earliest time of Kiddush Levana according to the opinions that it should not be said until 7
days after the molad.
If the time of tchilas zman Kiddush Levana occurs during the day (between ZmanimCalendar.Alos72 alos} and ZmanimCalendar.Tzais72 tzais) it
return the next tzais.
param alos the beginning of the Jewish day. If Kidush Levana occurs during the day (starting at alos and ending at tzais), the time returned will be tzais.
If either the alos or tzais parameters are nil, no daytime adjustment will be made.
tzais the end of the Jewish day. If Kidush Levana occurs during the day (starting at alos and ending at tzais),
the time returned will be tzais. If either the alos or tzais parameters are nil, no daytime adjustment will be made.

The method return the Date representing the moment 7 days after the molad.

# If the time occurs between alos and tzais, tzais will be returned

see
  - TchilasZmanKidushLevana3Days1, TchilasZmanKidushLevana3Days2
  - TchilasZmanKidushLevana7Days1, TchilasZmanKidushLevana7Days2
  - JewishCalendar.TchilasZmanKidushLevana7Days1
*/
func (t *complexZmanimCalendar) TchilasZmanKidushLevana7Days1(alos *time.Time, tzais *time.Time) (tm time.Time, ok bool) {
	jewishCalendar := hebrewcalendar.NewJewishCalendar(hebrewcalendar.NewJewishDate2(t.gDateTime.D))

	isJewishMonthBetween4And9(jewishCalendar)

	return t.moladBasedTime(jewishCalendar.TchilasZmanKidushLevana7Days(), alos, tzais, true)
}

/*
TchilasZmanKidushLevana7Days2 returns the earliest time of Kiddush Levana according to the opinions that it should not be said until 7
days after the molad. The time will be returned even if it occurs during the day when Kiddush Levana
can't be recited. Use TchilasZmanKidushLevana7Days2 if you want to limit the time to night hours.

The method return the Date representing the moment 7 days after the molad regardless of it is day or night.
- TchilasZmanKidushLevana7Days1
- hebrewcalendar.JewishCalendar.TchilasZmanKidushLevana7Days()
- TchilasZmanKidushLevana3Days1
*/
func (t *complexZmanimCalendar) TchilasZmanKidushLevana7Days2() (tm time.Time, ok bool) {
	return t.TchilasZmanKidushLevana7Days1(nil, nil)
}

/*
SofZmanAchilasChametzGRA returns the latest time one is allowed eating chametz on Erev Pesach according to
the opinion of the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
This time is identical to the ZmanimCalendar.SofZmanTfilaGRA and is provided as a convenience method for those who are
unaware how this zman is calculated. This time is 4 hours into the day based on the opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon that the day is calculated from sunrise to sunset.
This returns the time 4 * ZmanimCalendar.ShaahZmanisGRA after AstronomicalCalendar.SeaLevelSunrise sea level.

- ZmanimCalendar.ShaahZmanisGRA
- ZmanimCalendar.SofZmanTfilaGRA

The method return the time.Time one is allowed eating chametz on Erev Pesach.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) SofZmanAchilasChametzGRA() (tm time.Time, ok bool) {
	return t.SofZmanTfilaGRA()
}

/*
SofZmanAchilasChametzMGA72Minutes returns the latest time one is allowed eating chametz on Erev Pesach according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos
being ZmanimCalendar.Alos72 minutes before {AstronomicalCalendar.Sunrise.
This time is identical to the SofZmanTfilaMGA72Minutes. This time is 4 ZmanimCalendar.ShaahZmanisMGA
shaos zmaniyos} (temporal hours) after  ZmanimCalendar.Alos72 based on the opinion of the MGA that the day is
calculated from a ZmanimCalendar.Alos72 of 72 minutes before sunrise to Tzais72 nightfall of 72 minutes
after sunset. This returns the time of 4 * ZmanimCalendar.ShaahZmanisMGA after ZmanimCalendar.Alos72.

The method return the time.Time of the latest time of eating chametz.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ZmanimCalendar.ShaahZmanisMGA
  - ZmanimCalendar.Alos72
  - SofZmanTfilaMGA72Minutes
*/
func (t *complexZmanimCalendar) SofZmanAchilasChametzMGA72Minutes() (tm time.Time, ok bool) {
	return t.SofZmanTfilaMGA72Minutes()
}

/*
SofZmanAchilasChametzMGA16Point1Degrees returns the latest time one is allowed eating chametz on Erev Pesach according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos
being Alos16Point1Degrees 16.1 deg before AstronomicalCalendar.Sunrise.
This time is 4 ShaahZmanis16Point1Degrees shaos zmaniyos (solar hours) after Alos16Point1Degrees dawn
based on the opinion of the MGA that the day is calculated from dawn to nightfall with both being 16.1 deg
below sunrise or sunset. This returns the time of 4 ShaahZmanis16Point1Degrees after Alos16Point1Degrees dawn.

The method return the time.Time of the latest time of eating chametz.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanis16Point1Degrees()
  - Alos16Point1Degrees()
  - SofZmanTfilaMGA16Point1Degrees
*/
func (t *complexZmanimCalendar) SofZmanAchilasChametzMGA16Point1Degrees() (tm time.Time, ok bool) {
	return t.SofZmanTfilaMGA16Point1Degrees()
}

/*
SofZmanBiurChametzGRA returns the latest time for burning chametz on Erev Pesach according to the opinion
of the [GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon.
This time is 5 hours into the day based on the opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon that the day is calculated from
sunrise to sunset.
This returns the time 5 * ZmanimCalendar.ShaahZmanisGRA after AstronomicalCalendar.SeaLevelSunrise.

The method return the time.Time of the latest time for burning chametz on Erev Pesach.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) SofZmanBiurChametzGRA() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanisGRA, ok := t.ShaahZmanisGRA()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunrise, shaahZmanisGRA*50), true
}

/*
SofZmanBiurChametzMGA72Minutes returns the latest time for burning chametz on Erev Pesach according to the opinion of
the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos
being ZmanimCalendar.Alos72 minutes before AstronomicalCalendar.Sunrise.
This time is 5 {@link ZmanimCalendar.ShaahZmanisMGA shaos zmaniyos} (temporal hours) after ZmanimCalendar.Alos72 based on the opinion of
the MGA that the day is calculated from a  ZmanimCalendar.Alos72 of 72 minutes before sunrise to ZmanimCalendar.Tzais72 nightfall of 72 minutes after sunset.
This returns the time of 5 * ZmanimCalendar.ShaahZmanisMGA after

The method return the time.Time of the latest time for burning chametz on Erev Pesach.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - ZmanimCalendar.ShaahZmanisMGA
  - ZmanimCalendar.Alos72
*/
func (t *complexZmanimCalendar) SofZmanBiurChametzMGA72Minutes() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanisMGA, ok := t.ShaahZmanisMGA()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(alos72, shaahZmanisMGA*5), true
}

/*
SofZmanBiurChametzMGA16Point1Degrees returns the latest time for burning chametz on Erev Pesach according to the opinion
of the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern based on alos
being Alos16Point1Degrees 16.1 deg before AstronomicalCalendar.Sunrise.
This time is 5 ShaahZmanis16Point1Degrees shaos zmaniyos (solar hours) after Alos16Point1Degrees dawn
based on the opinion of the MGA that the day is calculated from dawn to nightfall with both being 16.1 deg
below sunrise or sunset. This returns the time of 5 ShaahZmanis16Point1Degrees after Alos16Point1Degrees dawn.

The method return the time.Time of the latest time for burning chametz on Erev Pesach.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the
Antarctic Circle where the sun may not reach low enough below the horizon for this calculation,
ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanis16Point1Degrees
  - Alos16Point1Degrees
*/
func (t *complexZmanimCalendar) SofZmanBiurChametzMGA16Point1Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanis16Point1Degrees, ok := t.ShaahZmanis16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(alos16Point1Degrees, shaahZmanis16Point1Degrees*5), true
}

/*
SolarMidnight that returns "solar" midnight, or the time when the sun is at its
[nadir]: https://en.wikipedia.org/wiki/Nadir.
Note: this method is experimental and might be removed.

The method return the time.Time of Solar Midnight (chatzos layla).
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) SolarMidnight() (tm time.Time, ok bool) {
	clonedCal := *t
	clonedCal.gDateTime.ToTime(nil).AddDate(0, 0, 1)
	sunset, ok := t.SeaLevelSunset()
	if !ok {
		return time.Time{}, false
	}
	sunrise, ok := clonedCal.SeaLevelSunrise()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(sunset, temporalHour(sunset, sunrise)*6), true
}

/*
sunriseBaalHatanya that returns the
[Baal Hatanya]: https://en.wikipedia.org/wiki/Shneur_Zalman_of_Liadi netz amiti (sunrise) without
AstronomicalCalculator.elevationAdjustment elevation.
This forms the base for the Baal Hatanya's dawn-based calculations that are calculated as a dip below the horizon before sunrise.

According to the Baal Hatanya, netz amiti, or true (halachic) sunrise, is when the top of the sun's
disk is visible at an elevation similar to the mountains of Eretz Yisrael. The time is calculated as the point at which
the center of the sun's disk is 1.583 deg below the horizon. This degree-based calculation can be found in Rabbi Shalom
DovBer Levine's commentary on The
[Baal Hatanya's Seder Hachnasas Shabbos]: https://www.chabadlibrary.org/books/pdf/Seder-Hachnosas-Shabbos.pdf.
From an elevation of 546 meters, the top of
[Har Hacarmel]: https://en.wikipedia.org/wiki/Mount_Carmel, the sun disappears when it is 1 deg 35' or 1.583 deg
below the sea level horizon. This in turn is based on the Gemara
[Shabbos 35a]: https://hebrewbooks.org/shas.aspx?mesechta=2&daf=35. There are other opinions brought down by
Rabbi Levine, including Rabbi Yosef Yitzchok Feigelstock who calculates it as the degrees below the horizon 4 minutes after
sunset in Yerushalayim (on the equinox). That is brought down as 1.583 deg. This is identical to the 1 deg 35' zman
and is probably a typo and should be 1.683 deg. These calculations are used by most
[Chabad]: https://en.wikipedia.org/wiki/Chabad calendars that use the Baal Hatanya's zmanim.
See [About Our Zmanim Calculations @ Chabad.org]: https://www.chabad.org/library/article_cdo/aid/3209349/jewish/About-Our-Zmanim-Calculations.htm.

Note: netz amiti is used only for calculating certain zmanim, and is intentionally unpublished. For
practical purposes, daytime mitzvos like shofar and lulav should not be done until after the
published time for netz / sunrise.

The method return the time.Time representing the exact sea-level netz amiti (sunrise) time.
If the calculation can't be	computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

- AstronomicalCalendar.Sunrise
- AstronomicalCalendar.SeaLevelSunrise
- sunsetBaalHatanya
- zenith1Point583
*/
func (t *complexZmanimCalendar) sunriseBaalHatanya() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith1Point583)
}

/*
sunsetBaalHatanya that returns the [Baal Hatanya]: https://en.wikipedia.org/wiki/Shneur_Zalman_of_Liadi
shkiah amiti (sunset) without AstronomicalCalculator.elevationAdjustment.
This forms the base for the Baal Hatanya's dusk-based calculations that are calculated as a dip below the horizon after sunset.

According to the Baal Hatanya, shkiah amiti, true (halachic) sunset, is when the top of the
sun's disk disappears from view at an elevation similar to the mountains of Eretz Yisrael.
This time is calculated as the point at which the center of the sun's disk is 1.583 degrees below the horizon.

Note: shkiah amiti is used only for calculating certain zmanim, and is intentionally unpublished. For
practical purposes, all daytime mitzvos should be completed before the published time for shkiah / sunset.

For further explanation of the calculations used for the Baal Hatanya's zmanim in this library,
see [About Our Zmanim Calculations @ Chabad.org]: https://www.chabad.org/library/article_cdo/aid/3209349/jewish/About-Our-Zmanim-Calculations.htm.

The method return the time.Time representing the exact sea-level shkiah amiti (sunset) time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - AstronomicalCalendar.Sunset
  - AstronomicalCalendar.SeaLevelSunset
  - sunriseBaalHatanya
  - zenith1Point583
*/
func (t *complexZmanimCalendar) sunsetBaalHatanya() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith1Point583)
}

/*
ShaahZmanisBaalHatanya that returns the [Baal Hatanya]: https://en.wikipedia.org/wiki/Shneur_Zalman_of_Liadi
a shaah zmanis (temporalHour). This forms the base for the Baal Hatanya's day based calculations,
that are calculated as a 1.583 deg dip below the horizon after sunset.
According to the Baal Hatanya, shkiah amiti, true (halachic) sunset, is when the top of the
sun's disk disappears from view at an elevation similar to the mountains of Eretz Yisrael.
This time is calculated as the point at which the center of the sun's disk is 1.583 degrees below the horizon.
A method that returns a shaah zmanis (temporalHour temporal hour) calculated based on the
[Baal Hatanya]: https://en.wikipedia.org/wiki/Shneur_Zalman_of_Liadi
netz amiti and shkiah amiti using a dip of 1.583 deg below the sea level horizon.
This calculation divides the day based on the opinion of the Baal Hatanya that the day runs from
sunriseBaalHatanya netz amiti to sunsetBaalHatanya shkiah amiti.
The calculations are based on a day from sunriseBaalHatanya sea level netz amiti to sunsetBaalHatanya sea level shkiah amiti.
The day is split into 12 equal parts with each one being a shaah zmanis.
This method is similar to temporalHour, but all calculations are based on a sea level sunrise and sunset.
The method return the gdt.GMillisecond length of a shaah zmanis calculated from sunriseBaalHatanya netz amiti (sunrise) to
sunsetBaalHatanya shkiah amiti
("real" sunset)}. If the calculation can't be computed such as in the Arctic Circle where there is at least one day a
year, where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
- temporalHour
- sunriseBaalHatanya
- sunsetBaalHatanya
- zenith1Point583
*/
func (t *complexZmanimCalendar) ShaahZmanisBaalHatanya() (i gdt.GMillisecond, ok bool) {
	sunriseBaalHatanya, ok := t.sunriseBaalHatanya()
	if !ok {
		return 0, false
	}
	sunsetBaalHatanya, ok := t.sunsetBaalHatanya()
	if !ok {
		return 0, false
	}
	return temporalHour(sunriseBaalHatanya, sunsetBaalHatanya), true
}

/*
AlosBaalHatanya returns the [Baal Hatanya]: https://en.wikipedia.org/wiki/Shneur_Zalman_of_Liadi alos
(dawn) calculated as the time when the sun is 16.9 deg below the eastern calculator.GeometricZenith
before AstronomicalCalendar.Sunrise.
For more information the source of 16.9 deg see calculator.zenith16Point9.

The method return The time.Time of dawn.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) AlosBaalHatanya() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith16Point9)
}

/*
SofZmanShmaBaalHatanya returns the latest zman krias shema (time to recite Shema in the morning).
This time is 3 ShaahZmanisBaalHatanya shaos zmaniyos (solar hours) after sunriseBaalHatanya
netz amiti (sunrise) based on the opinion of the Baal Hatanya that the day is calculated from
sunrise to sunset.
This returns the time 3 * ShaahZmanisBaalHatanya after sunriseBaalHatanya netz amiti (sunrise).

see
  - ZmanimCalendar.SofZmanShma
  - ShaahZmanisBaalHatanya

The method return the time.Time of the latest zman shema according to the Baal Hatanya.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) SofZmanShmaBaalHatanya() (tm time.Time, ok bool) {
	sunriseBaalHatanya, ok := t.sunriseBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	sunsetBaalHatanya, ok := t.sunsetBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanShma(sunriseBaalHatanya, sunsetBaalHatanya), true
}

/*
SofZmanTfilaBaalHatanya returns the latest zman tfilah (time to recite the morning prayers).
This time is 4 hours into the day based on the opinion of the Baal Hatanya that the day is
calculated from sunrise to sunset.
This returns the time 4 * ShaahZmanisBaalHatanya after sunriseBaalHatanya netz amiti (sunrise).

The method return the time.Time of the latest zman tfilah.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ZmanimCalendar.SofZmanTfila
  - ShaahZmanisBaalHatanya
*/
func (t *complexZmanimCalendar) SofZmanTfilaBaalHatanya() (tm time.Time, ok bool) {
	sunriseBaalHatanya, ok := t.sunriseBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	sunsetBaalHatanya, ok := t.sunsetBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	return t.SofZmanTfila(sunriseBaalHatanya, sunsetBaalHatanya), true
}

/*
SofZmanAchilasChametzBaalHatanya returns the latest time one is allowed eating chametz on Erev Pesach according to the
opinion of the Baal Hatanya.
This time is identical to the SofZmanTfilaBaalHatanya.
This time is 4 hours into the day based on the opinion of the Baal Hatanya that the day is calculated
from sunrise to sunset. This returns the time 4 ShaahZmanisBaalHatanya after sunriseBaalHatanya netz amiti (sunrise).

see
  - ShaahZmanisBaalHatanya
  - SofZmanTfilaBaalHatanya

The method return the time.Time one is allowed eating chametz on Erev Pesach.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) SofZmanAchilasChametzBaalHatanya() (tm time.Time, ok bool) {
	return t.SofZmanTfilaBaalHatanya()
}

/*
SofZmanBiurChametzBaalHatanya returns the latest time for burning chametz on Erev Pesach according to the opinion of
the Baal Hatanya. This time is 5 hours into the day based on the opinion of the Baal Hatanya that the day is calculated
from sunrise to sunset. This returns the time 5 * ShaahZmanisBaalHatanya after sunriseBaalHatanya netz amiti (sunrise).

The method return the time.Time of the latest time for burning chametz on Erev Pesach.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) SofZmanBiurChametzBaalHatanya() (tm time.Time, ok bool) {
	sunriseBaalHatanya, ok := t.sunriseBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	shaahZmanisBaalHatanya, ok := t.ShaahZmanisBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(sunriseBaalHatanya, shaahZmanisBaalHatanya*5), true
}

/*
MinchaGedolaBaalHatanya returns the time of mincha gedola. Mincha gedola is the earliest time one can pray
mincha.
The [Rambam]: https://en.wikipedia.org/wiki/Maimonides is of the opinion that it is better to delay mincha until
MinchaKetanaBaalHatanya while the
[Ra"sh]: https://en.wikipedia.org/wiki/Asher_ben_Jehiel,
[Tur]: https://en.wikipedia.org/wiki/Jacob_ben_Asher,
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon and others are of the opinion that mincha can be prayed
lechatchila starting at mincha gedola. This is calculated as 6.5 ShaahZmanisBaalHatanya
sea level solar hours after sunriseBaalHatanya netz amiti (sunrise).
This calculation is based on the opinion of the Baal Hatanya that the day is calculated from sunrise to sunset.
This returns the time 6.5 ShaahZmanisBaalHatanya after sunriseBaalHatanya netz amiti ("real" sunrise).

The method return the time.Time of the time of mincha gedola.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
- ZmanimCalendar.MinchaGedola
- ShaahZmanisBaalHatanya
- MinchaKetanaBaalHatanya
*/
func (t *complexZmanimCalendar) MinchaGedolaBaalHatanya() (tm time.Time, ok bool) {
	sunriseBaalHatanya, ok := t.sunriseBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	sunsetBaalHatanya, ok := t.sunsetBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaGedola2(sunriseBaalHatanya, sunsetBaalHatanya), true
}

/*
MinchaGedolaBaalHatanyaGreaterThan30 is a convenience method that returns the latter of MinchaGedolaBaalHatanya and
MinchaGedola30Minutes. In the winter when 1/2 of a ShaahZmanisBaalHatanya shaah zmanis is less than 30 minutes MinchaGedola30Minutes will be returned,
otherwise MinchaGedolaBaalHatanya will be returned.

The method return the time.Time of the latter of MinchaGedolaBaalHatanya and MinchaGedola30Minutes.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaGedolaBaalHatanyaGreaterThan30() (tm time.Time, ok bool) {
	minchaGedola30Minutes, ok := t.MinchaGedola30Minutes()
	if !ok {
		return time.Time{}, false
	}
	minchaGedolaBaalHatanya, ok := t.MinchaGedolaBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	if minchaGedola30Minutes.UnixMilli()-minchaGedolaBaalHatanya.UnixMilli() > 0 {
		return minchaGedola30Minutes, true
	} else {
		return minchaGedolaBaalHatanya, true
	}
}

/*
MinchaKetanaBaalHatanya returns the time of mincha ketana.
This is the preferred the earliest time to pray mincha in the opinion of the
[Rambam]: https://en.wikipedia.org/wiki/Maimonides and others.
For more information on this see the documentation on MinchaGedolaBaalHatanya.
This is calculated as 9.5 ShaahZmanisBaalHatanya sea level solar hours after sunriseBaalHatanya
netz amiti (sunrise). This calculation is calculated based on the opinion of the Baal Hatanya that the
day is calculated from sunrise to sunset. This returns the time 9.5 * ShaahZmanisBaalHatanya after sunriseBaalHatanya netz amiti (sunrise).

- ZmanimCalendar.MinchaKetana
- ShaahZmanisBaalHatanya
- MinchaGedolaBaalHatanya
The method return the time.Time of the time of mincha ketana.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) MinchaKetanaBaalHatanya() (tm time.Time, ok bool) {
	sunriseBaalHatanya, ok := t.sunriseBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	sunsetBaalHatanya, ok := t.sunsetBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	return t.MinchaKetana2(sunriseBaalHatanya, sunsetBaalHatanya), true
}

/*
PlagHaminchaBaalHatanya returns the time of plag hamincha.
This is calculated as 10.75 hours after sunrise.
This calculation is based on the opinion of the Baal Hatanya that the day is calculated from sunrise to sunset.
This returns the time 10.75 * ShaahZmanisBaalHatanya after sunriseBaalHatanya netz amiti (sunrise).

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - PlagHaminchaBaalHatanya
*/
func (t *complexZmanimCalendar) PlagHaminchaBaalHatanya() (tm time.Time, ok bool) {
	sunriseBaalHatanya, ok := t.sunriseBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	sunsetBaalHatanya, ok := t.sunsetBaalHatanya()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(sunriseBaalHatanya, sunsetBaalHatanya), true
}

/*
TzaisBaalHatanya is a method, that returns tzais (nightfall),
when the sun is 6 deg below the western geometric horizon (90 deg) after AstronomicalCalendar.Sunset.
For information on the source of this calculation see zenith6Degrees.

The method return The time.Time of nightfall.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) TzaisBaalHatanya() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith6Degrees)
}

/*
FixedLocalChatzosBasedZmanim is a utility methos to calculate zmanim based on
[Rav Moshe Feinstein]: https://en.wikipedia.org/wiki/Moshe_Feinstein as calculated in
[MTJ]: https://en.wikipedia.org/wiki/Mesivtha_Tifereth_Jerusalem,
[Yeshiva of Staten Island]: https://en.wikipedia.org/wiki/Mesivtha_Tifereth_Jerusalem, and Camp Yeshiva
of Staten Island. The day is split in two, from alos / sunrise to fixed local chatzos, and the
second half of the day, from fixed local chatzos to sunset / tzais. Morning based times are calculated
based on the first 6 hours, and afternoon times based on the second half of the day.
startOfHalfDay is the start of the half day. This would be alos or sunrise for morning based times and fixed local chatzos for the second half of the day.
endOfHalfDay is the end of the half day. This would be fixed local chatzos for morning based times and sunset or tzais for afternoon based times.
hours is the number of hours to offset the beginning of the first or second half of the day

The method return the time.Time of the latter of MinchaGedolaBaalHatanya and MinchaGedola30Minutes.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - FixedLocalChatzos
*/
func (t *complexZmanimCalendar) FixedLocalChatzosBasedZmanim(startOfHalfDay time.Time, endOfHalfDay time.Time, hours float64) time.Time {
	shaahZmanis := (endOfHalfDay.UnixMilli() - startOfHalfDay.UnixMilli()) / 6
	return time.UnixMilli(startOfHalfDay.UnixMilli() + int64(float64(shaahZmanis)*hours))
}

/*
SofZmanShmaMGA18DegreesToFixedLocalChatzos method returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion of the
calculation of sof zman krias shema (the latest time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern that the day is calculated from dawn to nightfall,
but calculated using the first half of the day only. The half a day starts at alos defined as link Alos18Degrees 18 deg
and ends at FixedLocalChatzos. Sof Zman Shema is 3 shaos zmaniyos (solar hours) after alos or half of this half-day.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos18Degrees()
  - FixedLocalChatzos
  - FixedLocalChatzosBasedZmanim
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA18DegreesToFixedLocalChatzos() (tm time.Time, ok bool) {
	alos18Degrees, ok := t.Alos18Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.FixedLocalChatzosBasedZmanim(alos18Degrees, t.FixedLocalChatzos(), 3), true
}

/*
SofZmanShmaMGA16Point1DegreesToFixedLocalChatzos method returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion of the
calculation of sof zman krias shema (the latest time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern that the
day is calculated from dawn to nightfall, but calculated using the first half of the day only. The half a day starts
at alos defined as Alos16Point1Degrees 16.1 deg and ends at FixedLocalChatzos fixed local
chatzos. Sof Zman Shema is 3 shaos zmaniyos (solar hours) after this alos or half of this half-day.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos16Point1Degrees
  - FixedLocalChatzos
  - FixedLocalChatzosBasedZmanim
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA16Point1DegreesToFixedLocalChatzos() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	fixedLocalChatzos := t.FixedLocalChatzos()
	return t.FixedLocalChatzosBasedZmanim(alos16Point1Degrees, fixedLocalChatzos, 3), true
}

/*
SofZmanShmaMGA90MinutesToFixedLocalChatzos returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion of the
calculation of sof zman krias shema (the latest time to recite Shema in the morning) according to the
opinion of the
[Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern that the
day is calculated from dawn to nightfall, but calculated using the first half of the day only.
The half a day starts at alos defined as Alos90 90 minutes before and ends at FixedLocalChatzos.
Sof Zman Shema is 3 shaos zmaniyos (solar hours) after this alos or half of this half-day.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos90
  - FixedLocalChatzos
  - FixedLocalChatzosBasedZmanim
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA90MinutesToFixedLocalChatzos() (tm time.Time, ok bool) {
	alos90, ok := t.Alos90()
	if !ok {
		return time.Time{}, false
	}
	fixedLocalChatzos := t.FixedLocalChatzos()
	return t.FixedLocalChatzosBasedZmanim(alos90, fixedLocalChatzos, 3), true
}

/*
SofZmanShmaMGA72MinutesToFixedLocalChatzos returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion of the
calculation of sof zman krias shema (the latest time to recite Shema in the morning) according to the
opinion of the [Magen Avraham (MGA)]: https://en.wikipedia.org/wiki/Avraham_Gombinern that the
day is calculated from dawn to nightfall, but calculated using the first half of the day only. The half a day starts
at alos defined as ZmanimCalendar.Alos72 72 minutes before and ends at FixedLocalChatzos
fixed local chatzos.
Sof Zman Shema is 3 shaos zmaniyos (solar hours) after this alos or half of this half-day.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ZmanimCalendar.Alos72
  - FixedLocalChatzos
  - FixedLocalChatzosBasedZmanim
*/
func (t *complexZmanimCalendar) SofZmanShmaMGA72MinutesToFixedLocalChatzos() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	return t.FixedLocalChatzosBasedZmanim(alos72, t.FixedLocalChatzos(), 3), true
}

/*
SofZmanShmaGRASunriseToFixedLocalChatzos returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion of the
calculation of sof zman krias shema the latest time to recite Shema in the morning according to the
opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon that the day is calculated from
sunrise to sunset, but calculated using the first half of the day only.
The half a day starts at AstronomicalCalendar.Sunrise and ends at FixedLocalChatzos.
Sof zman Shema is 3 shaos zmaniyos (solar hours) after sunrise or half of this half-day.

The method return the time.Time of the latest zman krias shema.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - AstronomicalCalendar.Sunrise
  - FixedLocalChatzos
  - FixedLocalChatzosBasedZmanim
*/
func (t *complexZmanimCalendar) SofZmanShmaGRASunriseToFixedLocalChatzos() (tm time.Time, ok bool) {
	if sunrise, ok := t.Sunrise(); ok {
		return t.FixedLocalChatzosBasedZmanim(sunrise, t.FixedLocalChatzos(), 3), ok
	} else {
		return time.Time{}, false
	}
}

/*
SofZmanTfilaGRASunriseToFixedLocalChatzos returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion of the
calculation of sof zman tfila (zman tfilah (the latest time to recite the morning prayers))
according to the opinion of the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon that the day is
calculated from sunrise to sunset, but calculated using the first half of the day only.
The half a day starts at AstronomicalCalendar.Sunrise and ends at FixedLocalChatzos.
Sof zman tefila is 4 shaos zmaniyos (solar hours) after sunrise or 2/3 of this half-day.

The method return the time.Time of the latest zman krias shema.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - AstronomicalCalendar.Sunrise
  - FixedLocalChatzos
  - FixedLocalChatzosBasedZmanim
*/
func (t *complexZmanimCalendar) SofZmanTfilaGRASunriseToFixedLocalChatzos() (tm time.Time, ok bool) {
	if sunrise, ok := t.Sunrise(); ok {
		return t.FixedLocalChatzosBasedZmanim(sunrise, t.FixedLocalChatzos(), 4), true
	} else {
		return time.Time{}, false
	}
}

/*
MinchaGedolaGRAFixedLocalChatzos30Minutes returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion of
the calculation of mincha gedola, the earliest time one can pray mincha
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon that is 30 minutes after FixedLocalChatzos.

The method return the time.Time of the time of mincha gedola.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ZmanimCalendar.MinchaGedola
  - FixedLocalChatzos
  - MinchaKetanaGRAFixedLocalChatzosToSunset
*/
func (t *complexZmanimCalendar) MinchaGedolaGRAFixedLocalChatzos30Minutes() time.Time {
	return timeOffset(t.FixedLocalChatzos(), gdt.GMinute(30).ToMilliseconds())
}

/*
MinchaKetanaGRAFixedLocalChatzosToSunset returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion
of the calculation of mincha ketana (the preferred time to recite the mincha prayers according to
the opinion of the
[Rambam]: https://en.wikipedia.org/wiki/Maimonides and others) calculated according
to the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon that is 3.5 shaos zmaniyos (solar hours) after FixedLocalChatzos.

The method return the time.Time of the time of mincha gedola.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ZmanimCalendar.MinchaGedola
  - FixedLocalChatzos
  - MinchaGedolaGRAFixedLocalChatzos30Minutes
*/
func (t *complexZmanimCalendar) MinchaKetanaGRAFixedLocalChatzosToSunset() (tm time.Time, ok bool) {
	sunset, ok := t.Sunset()
	if !ok {
		return time.Time{}, false
	}
	return t.FixedLocalChatzosBasedZmanim(t.FixedLocalChatzos(), sunset, 3.5), true
}

/*
PlagHaminchaGRAFixedLocalChatzosToSunset returns
[Rav Moshe Feinstein's]: https://en.wikipedia.org/wiki/Moshe_Feinstein opinion
of the calculation of plag hamincha. This method returns plag hamincha calculated according to the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon that the day ends at sunset and is 4.75 shaos
zmaniyos (solar hours) after FixedLocalChatzos.

The method return the time.Time of the time of mincha gedola.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ZmanimCalendar.PlagHamincha
  - FixedLocalChatzos
  - MinchaKetanaGRAFixedLocalChatzosToSunset
  - MinchaGedolaGRAFixedLocalChatzos30Minutes
*/
func (t *complexZmanimCalendar) PlagHaminchaGRAFixedLocalChatzosToSunset() (tm time.Time, ok bool) {
	sunset, ok := t.Sunset()
	if !ok {
		return time.Time{}, false
	}
	return t.FixedLocalChatzosBasedZmanim(t.FixedLocalChatzos(), sunset, 4.75), true
}

/*
Tzais50 return tzais (dusk) calculated as 50 minutes after sea level sunset.

This method returns tzais (nightfall) based on the opinion of Rabbi Moshe Feinstein for the New York area. This time should
not be used for latitudes different from NY area.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) Tzais50() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, gdt.GMinute(50).ToMilliseconds()), true
}

/*
SamuchLeMinchaKetanaGRA for calculating samuch lemincha ketana, / near mincha ketana time that is half an hour before
ZmanimCalendar.MinchaKetana or is 9 * ZmanimCalendar.ShaahZmanisGRA shaos zmaniyos (solar hours) after
AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise (depending on the ZmanimCalendar.IsUseElevation setting),
calculated according to the
[GRA]: https://en.wikipedia.org/wiki/Vilna_Gaon using a day starting at
sunrise and ending at sunset. This is the time that eating or other activity can't begin prior to praying mincha.
The calculation used is 9 * ShaahZmanis16Point1Degrees after Alos16Point1Degrees alos 16.1 deg.
See the
[Mechaber and Mishna Berurah 232]: https://hebrewbooks.org/pdfpager.aspx?req=60387&st=&pgnum=294 for details.

The method return the time.Time of the time of samuch lemincha ketana.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
- ZmanimCalendar.ShaahZmanisGRA
- SamuchLeMinchaKetana16Point1Degrees
*/
func (t *complexZmanimCalendar) SamuchLeMinchaKetanaGRA() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.SamuchLeMinchaKetana(elevationAdjustedSunrise, elevationAdjustedSunset), true
}

/*
SamuchLeMinchaKetana16Point1Degrees is a method for calculating samuch lemincha ketana, / near mincha ketana time that is half an hour before
MinchaGedola16Point1Degrees  or 9 * shaos zmaniyos (temporal hours) after the start of the day,
calculated using a day starting and ending 16.1 deg below the horizon. This is the time that eating or other activity
can't begin prior to praying mincha. The calculation used is 9 * ShaahZmanis16Point1Degrees after
Alos16Point1Degrees alos 16.1 deg.
See the [Mechaber and Mishna Berurah 232]: https://hebrewbooks.org/pdfpager.aspx?req=60387&st=&pgnum=294.

The method return the time.Time of the time of samuch lemincha ketana.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanis16Point1Degrees
*/
func (t *complexZmanimCalendar) SamuchLeMinchaKetana16Point1Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais16Point1Degrees, ok := t.Tzais16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.SamuchLeMinchaKetana(alos16Point1Degrees, tzais16Point1Degrees), true
}

/*
SamuchLeMinchaKetana72Minutes is a method for calculating samuch lemincha ketana, / near mincha ketana time that is half an hour before
MinchaKetana72Minutes or 9 * shaos zmaniyos (temporal hours) after the start of the day,
calculated using a day starting 72 minutes before sunrise and ending 72 minutes after sunset.
This is the time that eating or other activity can't begin prior to praying mincha. The calculation used is
9 * ShaahZmanis16Point1Degrees after Alos16Point1Degrees alos 16.1 deg.
See the [Mechaber and Mishna Berurah 232]: https://hebrewbooks.org/pdfpager.aspx?req=60387&st=&pgnum=294.

The method return the time.Time of the time of samuch lemincha ketana.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanis16Point1Degrees
*/
func (t *complexZmanimCalendar) SamuchLeMinchaKetana72Minutes() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	tzais72, ok := t.Tzais72()
	if !ok {
		return time.Time{}, false
	}
	return t.SamuchLeMinchaKetana(alos72, tzais72), true
}

/*
PlagHamincha120MinutesZmanis should be used lechumra only and returns the time of plag hamincha based on sunrise
being 120 minutes zmaniyos or 1/6th of the day before sunrise. This is calculated as 10.75 hours after
Alos120Zmanis. The formula used is 10.75 * ShaahZmanis120MinutesZmanis after Alos120Zmanis dawn.
Since the zman based on an extremely early alos and a very late tzais, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanis120MinutesZmanis
  - Alos120
  - Tzais120
  - PlagHamincha26Degrees
  - PlagHamincha120Minutes
*/
func (t *complexZmanimCalendar) PlagHamincha120MinutesZmanis() (tm time.Time, ok bool) {
	alos120Zmanis, ok := t.Alos120Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais120Zmanis, ok := t.Tzais120Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos120Zmanis, tzais120Zmanis), true
}

/*
PlagHamincha120Minutes should be used lechumra only and returns the time of plag hamincha according to the
Magen Avraham with the day starting 120 minutes before sunrise and ending 120 minutes after sunset.
This is calculated as 10.75 hours after Alos120 dawn 120 minutes.
The formula used is 10.75 ShaahZmanis120Minutes after Alos120.
Since the zman based on an extremely early alos and a very late tzais, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
	- ShaahZmanis120Minutes
	- PlagHamincha26Degrees
*/
// @Deprecated // (forRemoval=false) // add back once Java 9 is the minimum supported version
func (t *complexZmanimCalendar) PlagHamincha120Minutes() (tm time.Time, ok bool) {
	alos120, ok := t.Alos120()
	if !ok {
		return time.Time{}, false
	}
	tzais120, ok := t.Tzais120()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos120, tzais120), true
}

/*
Alos120 should be used lechumra only and returns alos (dawn) calculated using 120 minutes
before AstronomicalCalendar.SeaLevelSunrise sea level (no adjustment for elevation is made) based on the time
to walk the distance of 5 Mil(Ula) at 24 minutes a Mil. Time based offset calculations
for alos are based on the* opinion of the
[Rishonim]: https://en.wikipedia.org/wiki/Rishonim
who stated that the time of the neshef (time between dawn and sunrise) does not vary by the time of
year or location but purely depends on the time it takes to walk the distance of 5 Mil(Ula). Since
this time is extremely early, it should only be used lechumra, such as not eating after this time on a fast
day, and not as the start time for mitzvos that can only be performed during the day.

@deprecated This method should be used lechumra only (such as stopping to eat at this time on a fast day),
since it returns a very early time, and if used lekula can result in doing mitzvos hayom
too early according to most opinions. There is no current plan to remove this method from the API, and this
deprecation is intended to alert developers of the danger of using it.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Tzais120
  - Alos26Degrees
*/
func (t *complexZmanimCalendar) Alos120() (tm time.Time, ok bool) {
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunrise, -gdt.GMinute(120).ToMilliseconds()), true
}

/*
Alos120Zmanis should be used lechumra only and  method returns alos (dawn) calculated using
120 minutes zmaniyos or 1/6th of the day before AstronomicalCalendar.Sunrise or AstronomicalCalendar.SeaLevelSunrise (depending on the ZmanimCalendar.IsUseElevation setting).
This is based
on a 24-minute Mil so the time for 5 Mil is 120 minutes which is 1/6th of a day (12 * 60 /
6 = 120). The day is calculated from
AstronomicalCalendar.SeaLevelSunrise to AstronomicalCalendar.SeaLevelSunrise
or
AstronomicalCalendar.Sunrise to {@link AstronomicalCalendar.Sunset
depending on the ZmanimCalendar.IsUseElevation.
The actual calculation used is AstronomicalCalendar.Sunrise - ( ZmanimCalendar.ShaahZmanisGRA * 2 ).
Since this time is extremely early, it should only be used lechumra, such as not eating after this time on a fast day,
and not as the start time for mitzvos that can only be performed during the day.

@deprecated This method should be used lechumra only (such as stopping to eat at this time on a fast day),
since it returns a very early time, and if used lekula can result in doing mitzvos hayom
too early according to most opinions. There is no current plan to remove this method from the API, and this
deprecation is intended to alert developers of the danger of using it.

The method return the time.Time representing the time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - Alos120
  - Alos26Degrees
*/
func (t *complexZmanimCalendar) Alos120Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(-2.0)
}

/*
Alos26Degrees should be used lechumra only and returns alos (dawn) calculated when the sun is
zenith26Degrees 26 deg below the eastern geometric horizon before sunrise.
This calculation is based on the same calculation of Alos120 120 minutes, but uses a degree-based calculation instead of 120 exact minutes.
This calculation is based on the position of the sun 120 minutes before sunrise in Jerusalem
[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
which calculates to 26 deg below calculator.GeometricZenith. Since this time is extremely early, it should
only be used lechumra only, such as not eating after this time on a fast day, and not as the start time for
mitzvos that can only be performed during the day.

@deprecated This method should be used lechumra only (such as stopping to eat at this time on a fast day),
since it returns a very early time, and if used lekula can result in doing mitzvos hayom
too early according to most opinions. There is no current plan to remove this  method from the API, and this
deprecation is intended to alert developers of the danger of using it.

The method return the time.Time representing alos.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
see
  - Alos120
  - Tzais120
  - Tzais26Degrees
*/
func (t *complexZmanimCalendar) Alos26Degrees() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(zenith26Degrees)
}

/*
SofZmanShmaKolEliyahu From the GRA in Kol Eliyahu on Berachos #173 that states that zman krias shema is calculated as half the
time from AstronomicalCalendar.SeaLevelSunrise to FixedLocalChatzos.
The GRA himself seems to contradict this when he stated that zman krias shema is 1/4 of the day from
sunrise to sunset. See Sarah Lamoed #25 in Yisroel Vehazmanim Vol. III page 1016.

The method return the time.Time of the latest zman krias shema based on this calculation.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

@deprecated As per a conversation Rabbi Yisroel Twerski had with Rabbi Harfenes, this zman published in
the Yisrael Vehazmanim was based on a misunderstanding and should not be used. This deprecated method
will be removed (likely in v3.0) pending confirmation from Rabbi Harfenes.
*/
func (t *complexZmanimCalendar) SofZmanShmaKolEliyahu() (tm time.Time, ok bool) {
	chatzos := t.FixedLocalChatzos()
	elevationAdjustedSunrise, ok := t.elevationAdjustedSunrise()
	if !ok {
		return time.Time{}, false
	}
	diff := gdt.GMillisecond(chatzos.UnixMilli()-elevationAdjustedSunrise.UnixMilli()) / 2
	return timeOffset(chatzos, -diff), true
}

/*
PlagHamincha72Minutes should be used lechumra only and returns the time of plag hamincha according to
the Magen Avraham with the day starting 72 minutes before sunrise and ending 72 minutes after sunset.
This is calculated as 10.75 hours after ZmanimCalendar.Alos72.
The formula used is 10.75 ShaahZmanis72Minutes after ZmanimCalendar.Alos72.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
// @Deprecated // (forRemoval=false) // add back once Java 9 is the minimum supported version
func (t *complexZmanimCalendar) PlagHamincha72Minutes() (tm time.Time, ok bool) {
	alos72, ok := t.Alos72()
	if !ok {
		return time.Time{}, false
	}
	tzais72, ok := t.Tzais72()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos72, tzais72), true
}

/*
PlagHamincha90Minutes should be used lechumra only and returns the time of plag hamincha according to the
Magen Avraham with the day starting 90 minutes before sunrise and ending 90 minutes after sunset.
This is calculated as 10.75 hours after Alos90 dawn.
The formula used is 10.75 ShaahZmanis90Minutes after Alos90.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha90Minutes() (tm time.Time, ok bool) {
	alos90, ok := t.Alos90()
	if !ok {
		return time.Time{}, false
	}
	tzais90, ok := t.Tzais90()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos90, tzais90), true
}

/*
PlagHamincha96Minutes should be used lechumra only and returns the time of plag hamincha according to the Magen
Avraham with the day starting 96 minutes before sunrise and ending 96 minutes after sunset.
This is calculated as 10.75 hours after Alos96 dawn.
The formula used is 10.75 ShaahZmanis96Minutes after Alos96.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha96Minutes() (tm time.Time, ok bool) {
	alos96, ok := t.Alos96()
	if !ok {
		return time.Time{}, false
	}
	Tzais96, ok := t.Tzais96()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos96, Tzais96), true
}

/*
PlagHamincha96MinutesZmanis should be used lechumra only and returns the time of plag hamincha.
This is calculated as 10.75 hours after Alos96Zmanis dawn.
The formula used is 10.75 * ShaahZmanis96MinutesZmanis after Alos96Zmanis dawn.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha96MinutesZmanis() (tm time.Time, ok bool) {
	alos96Zmanis, ok := t.Alos96Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais96Zmanis, ok := t.Tzais96Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos96Zmanis, tzais96Zmanis), true
}

/*
PlagHamincha90MinutesZmanis should be used lechumra only and returns the time of plag hamincha.
This is calculated as 10.75 hours after Alos90Zmanis dawn.
The formula used is 10.75 * ShaahZmanis90MinutesZmanis after Alos90Zmanis dawn.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha90MinutesZmanis() (tm time.Time, ok bool) {
	alos90Zmanis, ok := t.Alos90Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais90Zmanis, ok := t.Tzais90Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos90Zmanis, tzais90Zmanis), true
}

/*
PlagHamincha72MinutesZmanis should be used lechumra only and returns the time of plag hamincha.
This is calculated as 10.75 hours after Alos72Zmanis.
The formula used is 10.75 * ShaahZmanis72MinutesZmanis after Alos72Zmanis dawn.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha72MinutesZmanis() (tm time.Time, ok bool) {
	alos72Zmanis, ok := t.Alos72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	tzais72Zmanis, ok := t.Tzais72Zmanis()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos72Zmanis, tzais72Zmanis), true
}

/*
PlagHamincha16Point1Degrees should be used lechumra only and returns the time of plag hamincha based on the
opinion that the day starts at Alos16Point1Degrees alos 16.1 deg and ends at Tzais16Point1Degrees tzais 16.1 deg.
This is calculated as 10.75 hours zmaniyos after Alos16Point1Degrees dawn.
The formula used is 10.75 * ShaahZmanis16Point1Degrees after Alos16Point1Degrees.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha16Point1Degrees() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais16Point1Degrees, ok := t.Tzais16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos16Point1Degrees, tzais16Point1Degrees), true
}

/*
PlagHamincha19Point8Degrees should be used lechumra only and returns the time of plag hamincha based on the
opinion that the day starts at {@link Alos19Point8Degrees() alos 19.8 deg} and ends at Tzais19Point8Degrees tzais 19.8 deg.
This is calculated as 10.75 hours zmaniyos after Alos19Point8Degrees dawn.
The formula used is 10.75 * ShaahZmanis19Point8Degrees after Alos19Point8Degrees.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha19Point8Degrees() (tm time.Time, ok bool) {
	alos19Point8Degrees, ok := t.Alos19Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais19Point8Degrees, ok := t.Tzais19Point8Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos19Point8Degrees, tzais19Point8Degrees), true
}

/*
PlagHamincha26Degrees should be used lechumra only and returns the time of plag hamincha based on the
opinion that the day starts at Alos26Degrees alos 26 deg and ends at Tzais26Degrees tzais 26 deg.
This is calculated as 10.75 hours zmaniyos after Alos26Degrees dawn.
The formula used is 10.75 * ShaahZmanis26Degrees after Alos26Degrees.
Since the zman based on an extremely early alos and a very late tzais, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.

see
  - ShaahZmanis26Degrees
  - PlagHamincha120Minutes
*/
func (t *complexZmanimCalendar) PlagHamincha26Degrees() (tm time.Time, ok bool) {
	alos26Degrees, ok := t.Alos26Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais26Degrees, ok := t.Tzais26Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos26Degrees, tzais26Degrees), true
}

/*
PlagHamincha18Degrees should be used lechumra only and returns the time of plag hamincha based on the
opinion that the day starts at Alos18Degrees alos 18 deg and ends at
Tzais18Degrees tzais 18 deg.
This is calculated as 10.75 hours zmaniyos after Alos18Degrees dawn.
The formula used is 10.75 * ShaahZmanis18Degrees after Alos18Degrees.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the time of plag hamincha.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagHamincha18Degrees() (tm time.Time, ok bool) {
	alos18Degrees, ok := t.Alos18Degrees()
	if !ok {
		return time.Time{}, false
	}
	tzais18Degrees, ok := t.Tzais18Degrees()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos18Degrees, tzais18Degrees), true
}

/*
PlagAlosToSunset should be used lechumra only and returns the time of plag hamincha based on the opinion
that the day starts at Alos16Point1Degrees alos 16.1 deg and ends at AstronomicalCalendar.Sunset.
10.75 shaos zmaniyos are calculated based on this day and added to Alos16Point1Degrees alos to reach this time.
This time is 10.75 shaos zmaniyos (temporal hours) after Alos16Point1Degrees dawn based on the opinion,
that the day is calculated from a Alos16Point1Degrees dawn of 16.1 degrees before sunrise to AstronomicalCalendar.SeaLevelSunset.
This returns the time of 10.75 * the calculated shaah zmanis after Alos16Point1Degrees dawn.
Since plag by this calculation can occur after sunset, it should only be used lechumra.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time of the plag.

If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) PlagAlosToSunset() (tm time.Time, ok bool) {
	alos16Point1Degrees, ok := t.Alos16Point1Degrees()
	if !ok {
		return time.Time{}, false
	}
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.PlagHamincha2(alos16Point1Degrees, elevationAdjustedSunset), true
}

/*
Tzais120 should be used lechumra only and returns tzais (nightfall) based on the calculations
of
[Rav Chaim Naeh]: https://en.wikipedia.org/wiki/Avraham_Chaim_Naeh
that the time to walk the distance of a Mil according to the
[Rambam]: https://en.wikipedia.org/wiki/Maimonides opinion
is 2/5 of an hour (24 minutes) for a total of 120 minutes based on the opinion of Ula who calculated
tzais as 5 Mil after sea level shkiah (sunset).
A similar calculation Tzais26Degrees uses degree-based calculations based on this 120 minute calculation.
Since the zman is extremely late and at a point that is long past the 18 deg point where the darkest point is
reached, it should only be used lechumra, such as delaying the start of nighttime mitzvos.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time representing the time.

If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) Tzais120() (tm time.Time, ok bool) {
	elevationAdjustedSunset, ok := t.elevationAdjustedSunset()
	if !ok {
		return time.Time{}, false
	}
	return timeOffset(elevationAdjustedSunset, gdt.GMinute(120).ToMilliseconds()), true
}

/*
Tzais120Zmanis should be used lechumra only and returns tzais (dusk) calculated using 120 minutes
zmaniyos after AstronomicalCalendar.SeaLevelSunset.
Since the zman is extremely late and at a point when it is long past the 18 deg point where the darkest point is reached,
it should only be used lechumra, such as delaying the start of nighttime mitzvos.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time representing the time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year, where the sun does not rise,
and one where it does not set, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) Tzais120Zmanis() (tm time.Time, ok bool) {
	return t.zmanisBasedOffset(2.0)
}

/*
Tzais26Degrees should be used lechumra only and returns tzais based on when the sun is 26 deg
below the horizon. For information on how this is calculated see the comments on Alos26Degrees.
Since the zman is extremely late and at a point when it is long past the 18 deg point where the
darkest point is reached, it should only be used lechumra such as delaying the start of nighttime
mitzvos.

@deprecated This method should be used lechumra only since it returns a very late time, and if used
lekula can result in chillul Shabbos etc. There is no current plan to remove this
method from the API, and this deprecation is intended to alert developers of the danger of using it.

The method return the time.Time representing the time.
If the calculation can't be computed such as northern and southern locations even south of the Arctic Circle and north of the Antarctic Circle,
where the sun may not reach low enough below the horizon for this calculation, ok is false will be returned.
See detailed explanation on top of the AstronomicalCalendar documentation.
*/
func (t *complexZmanimCalendar) Tzais26Degrees() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(zenith26Degrees)
}

/*
SofZmanShmaFixedLocal that returns the latest zman krias shema (time to recite Shema in the morning) calculated as 3
clock hours before FixedLocalChatzos. Note that there are opinions brought down in Yisrael Vehazmanim
[page 57]: https://hebrewbooks.org/pdfpager.aspx?req=9765&st=&pgnum=85 and Rav Yitzchak Silber's
[Sha'aos Shavos Balalacha]: https://www.worldcat.org/oclc/811253716 that this calculation is a mistake and regular
chatzos shoud be used for clock-hour calculations as opposed to fixed local chatzos. According to
these opinions it should be 3 clock hours before regular chatzos as calculated in SofZmanShma3HoursBeforeChatzos.

The method return the time.Time of the latest zman krias shema calculated as 3 clock hours before FixedLocalChatzos.

@deprecated This method of calculating sof zman Shma is considered a mistaken understanding of the proper
calculation of this zman in the opinion of Rav Yitzchak Silber's
[Sha'aos Shavos Balalacha]:https://www.worldcat.org/oclc/811253716.
On pages 316-318 he discusses Rav Yisrael
Harfenes's calculations and points to his seeming agreement that using fixed local chatzos as the focal
point is problematic. See Yisrael Vehazmanim
[page 57]: https://hebrewbooks.org/pdfpager.aspx?req=9765&st=&pgnum=85.
While the Yisrael Vehazmanim mentions this issue in vol. 1, it was not corrected in the calculations in vol. 3 and other parts of the sefer.
A competent rabbinical authority should be consulted before using this zman.
Instead, the use of SofZmanShma3HoursBeforeChatzos should be used to calculate sof zman Tfila using 3 fixed clock hours.
This will likely be removed in v3.0.
*/
func (t *complexZmanimCalendar) SofZmanShmaFixedLocal() time.Time {
	return timeOffset(t.FixedLocalChatzos(), -gdt.GMinute(180).ToMilliseconds())
}

/*
SofZmanTfilaFixedLocal returns the latest zman tfila (time to recite the morning prayers) calculated as 2 hours
before FixedLocalChatzos. See the documentation on SofZmanShmaFixedLocal showing differing opinions on how the zman is calculated.
According to many opinions SofZmanTfila2HoursBeforeChatzos should be used as opposed to this zman.

The method return the time.Time of the latest zman tfila.

@deprecated This method of calculating sof zman Tfila is considered a mistaken understanding of the proper
calculation of this zman in the opinion of Rav Yitzchak Silber's
[Sha'aos Shavos Balalacha]: https://www.worldcat.org/oclc/811253716. On pages 316-318 he discusses Rav Yisrael
Harfenes's calculations and points to his seeming agreement that using fixed local chatzos as the focal
point is problematic. See Yisrael Vehazmanim
[page 57]: https://hebrewbooks.org/pdfpager.aspx?req=9765&st=&pgnum=85.
While the Yisrael Vehazmanim mentions this issue in vol. 1, it was not corrected in the calculations in vol. 3 and other parts of the sefer.
A competent rabbinical authority should be consulted before using this zman. Instead, the use of SofZmanTfila2HoursBeforeChatzos
should be used to calculate sof zman Tfila using 2 fixed clock hours.
This will likely be removed in v3.0.
*/
func (t *complexZmanimCalendar) SofZmanTfilaFixedLocal() time.Time {
	return timeOffset(t.FixedLocalChatzos(), -gdt.GMinute(120).ToMilliseconds())
}
