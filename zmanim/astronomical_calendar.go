package zmanim

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/zmanim/calculator"
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
	"math"
	"time"
)

/*
AstronomicalCalendar that calculates astronomical times such as Sunrise, Sunset, and twilight times.

Note: There are times when the algorithms can't calculate proper values for sunrise, sunset and twilight.
This is usually caused by trying to calculate times for areas either very far North or South, where sunrise / sunset never
happen on that date. This is common when calculating twilight with a deep dip below the horizon for locations as far
south of the North Pole as London, in the northern hemisphere. The sun never reaches this dip at certain times of the
year.
When the calculations encounter this condition, ok = false will be returned.
The reason that panic is not invoked in these cases, is because the lack of a rise/set or twilight is
not an exception, but an expected condition in many parts of the world.
*/
type AstronomicalCalendar interface {
	Sunrise() (tm time.Time, ok bool)
	Sunset() (tm time.Time, ok bool)
	SeaLevelSunrise() (tm time.Time, ok bool)
	SeaLevelSunset() (tm time.Time, ok bool)
	UTCSunrise(zenith dimension.Degrees) float64
	UTCSunset(zenith dimension.Degrees) float64
	SunriseOffsetByDegrees(offsetZenith dimension.Degrees) (tm time.Time, ok bool)
	SunsetOffsetByDegrees(offsetZenith dimension.Degrees) (tm time.Time, ok bool)
	UTCSeaLevelSunrise(zenith dimension.Degrees) float64
	UTCSeaLevelSunset(zenith dimension.Degrees) float64
	TemporalHour() (i gdt.GMillisecond, ok bool)
	SunTransit() (tm time.Time, ok bool)
	GeoLocation() calculator.GeoLocation
}

type astronomicalCalendar struct {

	// The gdt.GDateTime encapsulated by this class to track the date/time used by the class
	gDateTime gdt.GDateTime

	// geoLocation calculator.GeoLocation used for calculations.
	geoLocation calculator.GeoLocation

	/*
		astronomicalCalculator the internal calculator.AstronomicalCalculator used for calculating solar based times.
		The calculation engine used to calculate the astronomical times can be changed to a different implementation by implementing
		the abstract calculator.AstronomicalCalculator.
		A number of different calculation engine implementations are included in the calculator package.
	*/
	astronomicalCalculator calculator.AstronomicalCalculator
}

func newAstronomicalCalendar() *astronomicalCalendar {
	return &astronomicalCalendar{}
}

func (t *astronomicalCalendar) initAstronomicalCalendar(gDateTime gdt.GDateTime, geoLocation calculator.GeoLocation, astronomicalCalculator calculator.AstronomicalCalculator) {
	t.gDateTime = gDateTime
	t.geoLocation = geoLocation
	t.astronomicalCalculator = astronomicalCalculator
}

func NewAstronomicalCalendar(dateTime gdt.GDateTime, geoLocation calculator.GeoLocation, astronomicalCalculator calculator.AstronomicalCalculator) AstronomicalCalendar {
	t := newAstronomicalCalendar()

	t.initAstronomicalCalendar(dateTime, geoLocation, astronomicalCalculator)

	return t
}

/*
The Sunrise method returns a time.Time representing the calculator.AstronomicalCalculator elevationAdjustment sunrise time.
The zenith used
  - for the calculation uses calculator.GeometricZenith of 90 deg plus
  - calculator.AstronomicalCalculator elevationAdjustment. This is adjusted by the
  - calculator.AstronomicalCalculator to add approximately 50/60 of a degree to account for 34 dimension.ArcMinutes of refraction
  - and 16 dimension.ArcMinutes for the sun's radius for a total of calculator.AstronomicalCalculator adjustZenith 90.83333 deg.
  - See documentation for the specific implementation of the calculator.AstronomicalCalculator that you are using.

return the time.Time representing the exact sunrise time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day,
a year, where the sun does not rise, and one where it does not set, an ok is false will be returned.
see
- calculator.AstronomicalCalculator adjustZenith
- SeaLevelSunrise
- UTCSunrise
*/
func (t *astronomicalCalendar) Sunrise() (tm time.Time, ok bool) {
	utcSunrise := t.UTCSunrise(calculator.GeometricZenith)
	if math.IsNaN(utcSunrise) {
		return time.Time{}, false
	}
	return t.dateTimeFromTimeOfDay(utcSunrise, true), true
}

/*
SeaLevelSunrise is a method that returns the sunrise without calculator.AstronomicalCalculator elevationAdjustment.
Non-sunrise and sunset calculations such as dawn and dusk, depend on the amount of visible light,
something that is not affected by elevation. This method returns sunrise calculated at sea level. This forms the
base for dawn calculations that are calculated as a dip below the horizon before sunrise.
return the time.Time representing the exact sea-level sunrise time. If the calculation can't be computed
such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the page.
see
- Sunrise
- UTCSeaLevelSunrise
- SeaLevelSunset
*/
func (t *astronomicalCalendar) SeaLevelSunrise() (tm time.Time, ok bool) {
	sunrise := t.UTCSeaLevelSunrise(calculator.GeometricZenith)
	if math.IsNaN(sunrise) {
		return time.Time{}, false
	}
	return t.dateTimeFromTimeOfDay(sunrise, true), true
}

/*
beginCivilTwilight is a method that returns the beginning of
[civil twilight]: https://en.wikipedia.org/wiki/Twilight#Civil_twilight
(dawn) using a zenith of calculator.CivilZenith 96 deg.
return the time.Time of the beginning of civil twilight using a zenith of 96 deg. If the calculation
can't be computed, null will be returned. See detailed explanation on top of the page.
see
- calculator.CivilZenith
*/
func (t *astronomicalCalendar) beginCivilTwilight() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(calculator.CivilZenith)
}

/*
beginNauticalTwilight is a method that returns the beginning of
[nautical twilight]: ]https://en.wikipedia.org/wiki/Twilight#Nautical_twilight using a zenith of calculator.NauticalZenith 102 deg.
return the time.Time of the beginning of nautical twilight using a zenith of 102 deg.
If the calculation can't be computed null will be returned. See detailed explanation on top of the page.
see
- calculator.NauticalZenith
*/
func (t *astronomicalCalendar) beginNauticalTwilight() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(calculator.NauticalZenith)
}

/*
beginAstronomicalTwilight is a method that returns the beginning of
[astronomical twilight]: https://en.wikipedia.org/wiki/Twilight#Astronomical_twilight using a zenith of
calculator.AstronomicalZenith 108 deg.
The method return the time.Time of the beginning of astronomical twilight using a zenith of 108 deg. If the
calculation can't be computed, null will be returned. See detailed explanation on top of the page.
see
- calculator.AstronomicalZenith
*/
func (t *astronomicalCalendar) beginAstronomicalTwilight() (tm time.Time, ok bool) {
	return t.SunriseOffsetByDegrees(calculator.AstronomicalZenith)
}

/*
Sunset method returns a time.Time representing the calculator.AstronomicalCalculator elevationAdjustment sunset time.
The zenith used for the calculation uses calculator.GeometricZenith of 90 deg plus
calculator.AstronomicalCalculator elevationAdjustment. This is adjusted by the
calculator.AstronomicalCalculator to add approximately 50/60 of a degree to account for 34 dimension.ArcMinutes of refraction
and 16 dimension.ArcMinutes for the sun's radius for a total of calculator.AstronomicalCalculator adjustZenith 90.83333 deg.
See documentation for the specific implementation of the calculator.AstronomicalCalculator that you are using.
Note: In certain cases the calculates sunset will occur before sunrise. This will typically happen when a timezone
other than the local timezone is used (calculating Los Angeles sunset using a GMT timezone for example). In this
case the sunset date will be incremented to the following date.
The method return the time.Time representing the exact sunset time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the page.
see
- calculator.AstronomicalCalculator adjustZenith
- SeaLevelSunset()
- UTCSunset
*/
func (t *astronomicalCalendar) Sunset() (tm time.Time, ok bool) {
	sunset := t.UTCSunset(calculator.GeometricZenith)
	if math.IsNaN(sunset) {
		return time.Time{}, false
	}
	return t.dateTimeFromTimeOfDay(sunset, false), true
}

/*
SeaLevelSunset is a method that returns the sunset without calculator.AstronomicalCalculator eElevationAdjustment.
Non-sunrise and sunset calculations such as dawn and dusk, depend on the amount of visible light,
something that is not affected by elevation. This method returns sunset calculated at sea level. This forms the
base for dusk calculations that are calculated as a dip below the horizon after sunset.
The method return the time.Time representing the exact sea-level sunset time.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the page.
see
- Sunset
- UTCSeaLevelSunset
- Sunset
*/
func (t *astronomicalCalendar) SeaLevelSunset() (tm time.Time, ok bool) {
	sunset := t.UTCSeaLevelSunset(calculator.GeometricZenith)
	if math.IsNaN(sunset) {
		return time.Time{}, false
	}
	return t.dateTimeFromTimeOfDay(sunset, false), true
}

/*
endCivilTwilight is a method that returns the end of
[civil twilight]: https://en.wikipedia.org/wiki/Twilight#Civil_twilight
using a zenith of calculator.CivilZenith 96 deg.
The method return the time.Time of the end of civil twilight using a zenith of calculator.CivilZenith 96 deg.
If the calculation can't be computed, ok is false will be returned.
See detailed explanation on top of the page.
see
- calculator.CivilZenith
*/
func (t *astronomicalCalendar) endCivilTwilight() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(calculator.CivilZenith)
}

/*
endNauticalTwilight is a method that returns the end of nautical twilight using a zenith of calculator.NauticalZenith 102 deg.
The method return the time.Time of the end of nautical twilight using a zenith of calculator.NauticalZenith 102 deg
If the calculation can't be computed, ok is false will be returned.
See detailed explanation on top of the page.
see
- calculator.NauticalZenith
*/
func (t *astronomicalCalendar) endNauticalTwilight() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(calculator.NauticalZenith)
}

/*
endAstronomicalTwilight is a method that returns the end of astronomical twilight using a zenith of calculator.AstronomicalZenith 108 deg.
The method return the time.Time of the end of astronomical twilight using a zenith of calculator.AstronomicalZenith 108 deg.
If the calculation can't be computed, ok is false will be returned.
See detailed explanation on top of the page.
see
- calculator.AstronomicalZenith
*/
func (t *astronomicalCalendar) endAstronomicalTwilight() (tm time.Time, ok bool) {
	return t.SunsetOffsetByDegrees(calculator.AstronomicalZenith)
}

/*
timeOffset is a utility method that returns a date offset by the offset time passed in.
Please note that the level of light during twilight is not affected by elevation,
so if this is being used to calculate an offset before sunrise or
after sunset with the intent of getting a rough "level of light" calculation, the sunrise or sunset time passed
to this method should be sea level sunrise and sunset.
tm is the start time
offset is the offset in milliseconds to add to the time.
return the time.Time with the offset in gdt.GMillisecond added to it
*/
func timeOffset(tm time.Time, offset gdt.GMillisecond) time.Time {
	if offset == math.MaxInt64 {
		helper.Panic("fmt.Sprintf(\"%v is not UTC\", tm)")
	}
	tm.Add(time.Duration(offset) * time.Millisecond)
	return tm.Add(time.Duration(offset) * time.Millisecond)
}

/*
SunriseOffsetByDegrees is a utility method that returns the time of an offset by degrees below or above the horizon of
Sunrise.
Note that the degree offset is from the vertical, so for a calculation of 14 deg
before sunrise, an offset of 14 + calculator.GeometricZenith = 104 would have to be passed as a parameter.
offsetZenith the dimension.Degrees before Sunrise to use in the calculation. For time after sunrise use
negative numbers.
Note that the degree offset is from the vertical, so for a calculation of 14 deg;
before sunrise, an offset of 14 + calculator.GeometricZenith = 104 would have to be passed as a parameter.
The method return the time.Time of the offset after (or before) Sunrise. If the calculation
can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the page.
*/
func (t *astronomicalCalendar) SunriseOffsetByDegrees(offsetZenith dimension.Degrees) (tm time.Time, ok bool) {
	dawn := t.UTCSunrise(offsetZenith)
	if math.IsNaN(dawn) {
		return time.Time{}, false
	}
	return t.dateTimeFromTimeOfDay(dawn, true), true
}

/*
SunsetOffsetByDegrees is a utility method that returns the time of an offset by degrees below or above the horizon of Sunset
. Note that the degree offset is from the vertical, so for a calculation of 14 deg; after sunset, an
offset of 14 + calculator.GeometricZenith = 104 would have to be passed as a parameter.
offsetZenith the dimension.Degrees after Sunset to use in the calculation.
For time before sunset use negative numbers.
Note that the degree offset is from the vertical, so for a calculation of 14 deg; after
sunset, an offset of 14 + calculator.GeometricZenith = 104 would have to be passed as a parameter.
The method return the time.Time of the offset after (or before) Sunset.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the page.
*/
func (t *astronomicalCalendar) SunsetOffsetByDegrees(offsetZenith dimension.Degrees) (tm time.Time, ok bool) {
	sunset := t.UTCSunset(offsetZenith)
	if math.IsNaN(sunset) {
		return time.Time{}, false
	}
	return t.dateTimeFromTimeOfDay(sunset, false), true
}

/*
UTCSunrise is a method that returns the sunrise in UTC time without correction for time zone offset from GMT and without using
daylight savings time.
zenith the dimension.Degrees below the horizon. For time after sunrise use negative numbers.
The method return the time in the format: 18.75 for 18:45:00 UTC/GMT.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, math.NaN will be returned.
See detailed explanation on top of the page.
*/
func (t *astronomicalCalendar) UTCSunrise(zenith dimension.Degrees) float64 {
	return t.astronomicalCalculator.UTCSunrise(t.adjustedDateTime(), t.GeoLocation(), zenith, true)
}

/*
UTCSeaLevelSunrise is a method that returns the sunrise in UTC time without correction for time zone offset from GMT and without using
daylight savings time. Non-sunrise and sunset calculations such as dawn and dusk, depend on the amount of visible
light, something that is not affected by elevation. This method returns UTC sunrise calculated at sea level. This
forms the base for dawn calculations that are calculated as a dip below the horizon before sunrise.
zenith is the dimension.Degrees below the horizon. For time after sunrise use negative numbers.
The method return the time in the format: 18.75 for 18:45:00 UTC/GMT.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, math.NaN will be returned.
See detailed explanation on top of the page.
see
- UTCSunrise
- UTCSeaLevelSunset
*/
func (t *astronomicalCalendar) UTCSeaLevelSunrise(zenith dimension.Degrees) float64 {
	return t.astronomicalCalculator.UTCSunrise(t.adjustedDateTime(), t.GeoLocation(), zenith, false)
}

/*
UTCSunset is a method that returns the sunset in UTC time without correction for time zone offset from GMT and without using
daylight savings time.
zenith is the dimension.Degrees below the horizon. For time after sunset use negative numbers.
The method return the time in the format: 18.75 for 18:45:00 UTC/GMT.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, math.NaN will be returned.
See detailed explanation on top of the page.
see
- UTCSeaLevelSunset
*/
func (t *astronomicalCalendar) UTCSunset(zenith dimension.Degrees) float64 {
	return t.astronomicalCalculator.UTCSunset(t.adjustedDateTime(), t.GeoLocation(), zenith, true)
}

/*
UTCSeaLevelSunset is a method that returns the sunset in UTC time without correction for elevation, time zone offset from GMT and
without using daylight savings time. Non-sunrise and sunset calculations such as dawn and dusk, depend on the
amount of visible light, something that is not affected by elevation. This method returns UTC sunset calculated
at sea level. This forms the base for dusk calculations that are calculated as a dip below the horizon after
sunset.
zenith is the dimension.degrees below the horizon. For time before sunset use negative numbers.
The method return the time in the format: 18.75 for 18:45:00 UTC/GMT.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, math.NaN will be returned.
See detailed explanation on top of the page.
see
- UTCSunset
- UTCSeaLevelSunrise
*/
func (t *astronomicalCalendar) UTCSeaLevelSunset(zenith dimension.Degrees) float64 {
	return t.astronomicalCalculator.UTCSunset(t.adjustedDateTime(), t.GeoLocation(), zenith, false)
}

/*
TemporalHour is a method that returns calculator.AstronomicalCalculator elevationAdjustment temporal (solar) hour.
The day from Sunrise to Sunset is split into 12 equal parts with each one being a temporal hour.
see
- Sunrise
- Sunset
- temporalHour
The method return the gdt.GMillisecond length of a temporal hour.
If the calculation can't be computed, ok is false will be returned.
See detailed explanation on top of the page.
see
- temporalHour
*/
func (t *astronomicalCalendar) TemporalHour() (i gdt.GMillisecond, ok bool) {
	seaLevelSunrise, ok := t.SeaLevelSunrise()
	if !ok {
		return 0, false
	}
	seaLevelSunset, ok := t.SeaLevelSunset()
	if !ok {
		return 0, false
	}
	return temporalHour(seaLevelSunrise, seaLevelSunset), true
}

/*
temporalHour is a utility method that will allow the calculation
of a temporal (solar) hour based on the sunrise and sunset
passed as parameters to this method. An example of the use of this method would be the calculation of a
non-elevation adjusted temporal hour by passing in
SeaLevelSunrise and SeaLevelSunset as parameters.
sunrise is the start of the day.
sunset is the end of the day.
return the gdt.GMillisecond length of the temporal hour.
*/
func temporalHour(sunrise time.Time, sunset time.Time) gdt.GMillisecond {
	return gdt.GMillisecond(sunset.Sub(sunrise).Milliseconds()) / 12
}

/*
SunTransit is a method that returns sundial or solar noon. It occurs when the Sun is
[transiting]: https://en.wikipedia.org/wiki/Transit_%28astronomy%29 the
[celestial meridian]: https://en.wikipedia.org/wiki/Meridian_%28astronomy%29.
In this class it is calculated as halfway between sea level sunrise and sea level sunset, which can be slightly off the real transit
time due to changes in declination (the lengthening or shortening day).
The method return the time.Time representing Sun's transit.
If the calculation can't be computed such as in the Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok is false will be returned.
See detailed explanation on top of the page.
see
- sunTransit
- temporalHour
*/
func (t *astronomicalCalendar) SunTransit() (tm time.Time, ok bool) {
	seaLevelSunrise, ok := t.SeaLevelSunrise()
	if !ok {
		return time.Time{}, false
	}
	seaLevelSunset, ok := t.SeaLevelSunset()
	if !ok {
		return time.Time{}, false
	}
	return t.sunTransit(seaLevelSunrise, seaLevelSunset), true
}

/*
sunTransit is a method that returns sundial or solar noon. It occurs when the Sun is
[transiting]: https://en.wikipedia.org/wiki/Transit_%28astronomy%29 the
[celestial meridian]: https://en.wikipedia.org/wiki/Meridian_%28astronomy%29.
In this class it is calculated as halfway between the sunrise and sunset passed to this method. This time can be slightly off the
real transit time due to changes in declination (the lengthening or shortening day).
startOfDay the start of day for calculating the sun's transit. This can be sea level sunrise, visual sunrise (or
any arbitrary start of day) passed to this method.
endOfDay the end of day for calculating the sun's transit. This can be sea level sunset, visual sunset (or any
arbitrary end of day) passed to this method.
The method return the time.Time representing Sun's transit. If the calculation can't be computed such as in the
Arctic Circle where there is at least one day a year,
where the sun does not rise, and one where it does not set, ok s false will be returned.
See detailed explanation on top of the page.
*/
func (t *astronomicalCalendar) sunTransit(startOfDay time.Time, endOfDay time.Time) time.Time {
	temporalHour := temporalHour(startOfDay, endOfDay)
	return timeOffset(startOfDay, temporalHour*6)
}

/*
dateTimeFromTimeOfDay is a method that returns a time.Time from the time passed in as a parameter.
timeOfDay is the time to be set as the time for the time.Time.
The time expected is in the format: 18.75 for 6:45:00 PM.time is sunrise and false if it is sunset
isSunrise
The method return the time.Time.
*/
func (t *astronomicalCalendar) dateTimeFromTimeOfDay(timeOfDay float64, isSunrise bool) time.Time {
	if math.IsNaN(timeOfDay) {
		helper.Panic("math.IsNaN(timeOfDay)")
	}

	calculatedTime := timeOfDay
	hours := math.Floor(calculatedTime) // retain only the hours
	calculatedTime -= hours
	calculatedTime *= 60
	minutes := math.Floor(calculatedTime) // retain only the minutes
	calculatedTime -= minutes
	calculatedTime *= 60
	seconds := math.Floor(calculatedTime) // retain only the seconds
	calculatedTime -= seconds             // remaining NOT milliseconds, but microseconds
	nanoseconds := math.Floor(calculatedTime * float64(time.Second))

	adjustedDateTime := t.adjustedDateTime()
	year, month, day := adjustedDateTime.D.Year, adjustedDateTime.D.Month, adjustedDateTime.D.Day
	utcDateTime := time.Date(int(year), month, int(day), int(hours), int(minutes), int(seconds), int(nanoseconds), timeutil.GmtTimezoneOrPanic())

	// adjust date if utc time reflects a wraparound from the local offset
	// Check if a date transition has occurred, or is about to occur - this indicates the date of the event is
	// actually not the target date, but the day prior or after
	localOffsetHours := math.Floor(t.GeoLocation().Longitude() / 15)
	if isSunrise && localOffsetHours+hours > 18 {
		utcDateTime = utcDateTime.AddDate(0, 0, -1) //add(Calendar.DAY_OF_MONTH, -1)
	} else if !isSunrise && localOffsetHours+hours < 6 {
		utcDateTime = utcDateTime.AddDate(0, 0, 1) // cal.add(Calendar.DAY_OF_MONTH, 1)
	}

	return t.convertDateTimeForZone(gdt.NewGDateTime1(utcDateTime))

}

func (t *astronomicalCalendar) convertDateTimeForZone(utcDateTime gdt.GDateTime) time.Time {
	return utcDateTime.ToTime(nil).In(t.GeoLocation().TimeZone())
}

/*
sunriseSolarDipFromOffset returns the dip below the horizon before sunrise that matches the offset minutes on passed in as a parameter.
For example passing in 72 minutes for a calendar set to the equinox in Jerusalem returns a value close to 16.1 deg
Please note that this method is very slow and inefficient and should NEVER be used in a loop.
TO-DO: Improve efficiency.
offset gdt.GMinute
return the dimension.Degrees below the horizon before sunrise that match the offset in minutes passed it as a parameter.
*/
func (t *astronomicalCalendar) sunriseSolarDipFromOffset(offset gdt.GMinute) (f dimension.Degrees, ok bool) {
	offsetByDegrees, ok := t.SeaLevelSunrise()
	if !ok {
		return 0, false
	}
	seaLevelSunrise, ok := t.SeaLevelSunrise()
	if !ok {
		return 0, false
	}
	offsetByTime := timeOffset(seaLevelSunrise, -(offset.ToMilliseconds()))

	degrees := dimension.Degrees(0)
	incrementor := dimension.Degrees(0.0001)

	for (offset < 0 && offsetByDegrees.UnixMilli() < offsetByTime.UnixMilli()) || (offset > 0 && offsetByDegrees.UnixMilli() > offsetByTime.UnixMilli()) {
		if offset > 0 {
			degrees += incrementor
		} else {
			degrees -= incrementor
		}
		offsetByDegrees, ok = t.SunriseOffsetByDegrees(calculator.GeometricZenith + degrees)
		if !ok {
			return 0, false
		}
	}
	return degrees, true
}

/*
sunsetSolarDipFromOffset returns the dip below the horizon after sunset that matches the offset minutes on passed in as a parameter.
For example passing in 72 minutes for a calendar set to the equinox in Jerusalem returns a value close to 16.1 deg
Please note that this method is very slow and inefficient and should NEVER be used in a loop.
TO-DO: Improve efficiency.
offset gdt.GMinute
The method return the dimension.Degrees below the horizon after sunset that match the offset in minutes passed it as a parameter.
see
- SunriseSolarDipFromOffset
*/
func (t *astronomicalCalendar) sunsetSolarDipFromOffset(offset gdt.GMinute) (f dimension.Degrees, ok bool) {
	offsetByDegrees, ok := t.SeaLevelSunset()
	if !ok {
		return 0, false
	}
	seaLevelSunset, ok := t.SeaLevelSunset()
	if !ok {
		return 0, false
	}

	offsetByTime := timeOffset(seaLevelSunset, offset.ToMilliseconds())

	//BigDecimal degrees = new BigDecimal(0);
	//BigDecimal incrementor = new BigDecimal("0.001");
	degrees := dimension.Degrees(0)
	incrementor := dimension.Degrees(0.0001)

	for (offset > 0 && offsetByDegrees.UnixMilli() < offsetByTime.UnixMilli()) || (offset < 0 && offsetByDegrees.UnixMilli() > offsetByTime.UnixMilli()) {
		if offset > 0 {
			// degrees = degrees.add(incrementor);
			degrees += incrementor
		} else {
			// degrees = degrees.subtract(incrementor);
			degrees -= incrementor
		}

		sunsetOffsetByDegrees, ok := t.SunsetOffsetByDegrees(calculator.GeometricZenith + degrees)
		if !ok {
			return 0, false
		}
		offsetByDegrees = sunsetOffsetByDegrees
	}
	return degrees, true
}

/*
adjustedDateTime adjusts the gdt.GDateTime to deal with edge cases where the location crosses the anti-meridian.
see
- GeoLocation AntimeridianAdjustment
The method return the adjusted gdt.GDateTime
*/
func (t *astronomicalCalendar) adjustedDateTime() gdt.GDateTime {
	offset := t.GeoLocation().AntimeridianAdjustment()
	adjustedTime := t.gDateTime.ToTime(nil).AddDate(0, 0, int(offset))
	return gdt.NewGDateTime1(adjustedTime)
}

func (t *astronomicalCalendar) String() string {
	panic("implement me")
}

func (t *astronomicalCalendar) GDateTime() gdt.GDateTime {
	return t.gDateTime
}

func (t *astronomicalCalendar) GeoLocation() calculator.GeoLocation {
	return t.geoLocation
}
