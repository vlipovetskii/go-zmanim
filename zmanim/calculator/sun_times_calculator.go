package calculator

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
	"math"
)

/*
sunTimesCalculator
Implementation of sunrise and sunset methods to calculate astronomical times. This calculator uses the Java algorithm
written by [Kevin Boone]: http://web.archive.org/web/20090531215353/http://www.kevinboone.com/suntimes.html
that is based on the [US Naval Observatory's]: "http://aa.usno.navy.mil/
[Almanac]: http://aa.usno.navy.mil/publications/docs/asa.php for Computer algorithm (
[Amazon]: http://www.amazon.com/exec/obidos/tg/detail/-/0160515106/,
[Barnes &amp; Noble]: http://search.barnesandnoble.com/booksearch/isbnInquiry.asp?isbn=0160515106) and is
used with his permission. Added to Kevin's code is adjustment of the zenith to account for elevation.
*/
type sunTimesCalculator struct {
	astronomicalCalculator
}

func newSunTimesCalculator() *sunTimesCalculator {
	return &sunTimesCalculator{}
}

func NewSunTimesCalculator() AstronomicalCalculator {
	t := newSunTimesCalculator()

	t.initAstronomicalCalculator()

	return t
}

func (t *sunTimesCalculator) CalculatorName() string {
	return "US Naval Almanac Algorithm"
}

func (t *sunTimesCalculator) UTCSunrise(targetDateTime gdt.GDateTime, geoLocation GeoLocation, zenith dimension.Degrees, adjustForElevation bool) float64 {
	doubleTime := math.NaN()
	elevation := dimension.Meters(0)
	if adjustForElevation {
		elevation = geoLocation.Elevation()
	}
	adjustedZenith := t.adjustZenith(zenith, elevation)

	doubleTime = timeUTC(targetDateTime, geoLocation, adjustedZenith, true)

	return doubleTime
}

func (t *sunTimesCalculator) UTCSunset(targetDateTime gdt.GDateTime, geoLocation GeoLocation, zenith dimension.Degrees, adjustForElevation bool) float64 {
	doubleTime := math.NaN()
	elevation := dimension.Meters(0)
	if adjustForElevation {
		elevation = geoLocation.Elevation()
	}
	adjustedZenith := t.adjustZenith(zenith, elevation)

	doubleTime = timeUTC(targetDateTime, geoLocation, adjustedZenith, false)

	return doubleTime
}

const (
	/*
		degPerHour The number of Degrees of longitude that corresponds to one-hour time difference.
	*/
	degPerHour = dimension.Degrees(360.0 / 24.0)
)

/*
hoursFromMeridian get time difference between location's longitude and the Meridian, in hours,
return time difference between the location's longitude and the Meridian, in hours. West of Meridian has a negative time difference
longitude the longitude
*/
func hoursFromMeridian(longitude dimension.Degrees) float64 {
	return float64(longitude / degPerHour)
}

/*
approxTimeDays calculate the approximate time of sunset or sunrise in days since midnight Jan 1st, assuming 6am and 6pm events,
return the approximate time of sunset or sunrise in days since midnight Jan 1st, assuming 6am and 6pm events.
We need this figure to derive the Sun's mean anomaly.
dayOfYear the Day of Year
hoursFromMeridian hours from the meridian
isSunrise true for sunrise and false for sunset
need this figure to derive the Sun's mean anomaly.
*/
func approxTimeDays(dayOfYear int32, hoursFromMeridian float64, isSunrise bool) float64 {
	if isSunrise {
		return float64(dayOfYear) + ((6.0 - hoursFromMeridian) / 24)
	} else { // sunset
		return float64(dayOfYear) + ((18.0 - hoursFromMeridian) / 24)
	}
}

/*
sunMeanAnomaly calculate the Sun's mean anomaly in Degrees, at sunrise or sunset, given the longitude in Degrees,
return the Sun's mean anomaly in Degrees.
dayOfYear the Day of the Year
longitude
isSunrise true for sunrise and false for sunset
*/
func sunMeanAnomaly(dayOfYear int32, longitude dimension.Degrees, isSunrise bool) dimension.Degrees {
	return dimension.Degrees(0.9856*approxTimeDays(dayOfYear, hoursFromMeridian(longitude), isSunrise)) - 3.289
}

/*
sunTrueLongitudeFromSunMeanAnomaly return the Sun's true longitude in Degrees. The result is an angle &gt;= 0 and &lt;= 360.
sunMeanAnomaly the Sun's mean anomaly in Degrees
*/
func sunTrueLongitudeFromSunMeanAnomaly(sunMeanAnomaly dimension.Degrees) dimension.Degrees {
	sunMeanAnomalySin := sunMeanAnomaly.Sin()
	l := sunMeanAnomaly + dimension.Degrees(1.916*sunMeanAnomalySin) + dimension.Degrees(0.020*sunMeanAnomalySin) + 282.634

	// get longitude into 0-360 degree range
	if l >= 360.0 {
		l = l - 360.0
	}

	if l < 0 {
		l = l + 360.0
	}

	return l
}

/*
sunRightAscensionHours calculates the Sun's right ascension in hours,
return the Sun's right ascension in hours in angles; 0 and; 360.
sunTrueLongitude the Sun's true longitude in Degrees; 0 and; 360.
*/
func sunRightAscensionHours(sunTrueLongitude dimension.Degrees) float64 {
	a := 0.91764 * sunTrueLongitude.Tan()
	ra := dimension.ATan(a)

	lQuadrant := math.Floor(float64(sunTrueLongitude)/90.0) * 90.0
	raQuadrant := math.Floor(float64(ra)/90.0) * 90.0
	ra = ra + dimension.Degrees(lQuadrant-raQuadrant)

	return float64(ra / degPerHour) // convert to hours
}

/*
cosLocalHourAngle calculate the cosine of the Sun's local hour angle,
return the cosine of the Sun's local hour angle
sunTrueLongitude the sun's true longitude
latitude the latitude
zenith the zenith
*/
func cosLocalHourAngle(sunTrueLongitude dimension.Degrees, latitude dimension.Degrees, zenith dimension.Degrees) float64 {
	sinDec := 0.39782 * sunTrueLongitude.Sin()
	cosDec := dimension.ASin(sinDec).Cos()
	return (zenith.Cos() - (sinDec * latitude.Sin())) / (cosDec * latitude.Cos())
}

/*
localMeanTime calculate local mean time of rising or setting. By 'local' is meant the exact time at the location, assuming that
there were no time zone. That is, the time difference between the location and the Meridian depended entirely on
the longitude. We can't do anything with this time directly; we must convert it to UTC and then to a local time,
return the fractional number of hours since midnight as a double.
localHour the local hour
sunRightAscensionHours the sun's right ascension in hours
approxTimeDays approximate time days
*/
func localMeanTime(localHour float64, sunRightAscensionHours float64, approxTimeDays float64) float64 {
	return localHour + sunRightAscensionHours - (0.06571 * approxTimeDays) - 6.622
}

/*
timeUTC get sunrise or sunset time in UTC, according to flag. This time is returned as
a double and is not adjusted for time-zone,
return the time as a double. If an error was encountered in the calculation
expected behavior for some locations such as near the poles,
NaN will be returned.
calendar the Calendar object to extract the Day of Year for calculation
geoLocation the GeoLocation object that contains the latitude and longitude
zenith Sun's zenith, in Degrees
isSunrise True for sunrise and false for sunset.
*/
func timeUTC(targetDateTime gdt.GDateTime, geoLocation GeoLocation, zenith dimension.Degrees, isSunrise bool) float64 {

	dayOfYear := int32(targetDateTime.ToTime(nil).YearDay())

	sunMeanAnomaly := sunMeanAnomaly(dayOfYear, dimension.Degrees(geoLocation.Longitude()), isSunrise)
	sunTrueLong := sunTrueLongitudeFromSunMeanAnomaly(sunMeanAnomaly)
	sunRightAscensionHours := sunRightAscensionHours(sunTrueLong)
	cosLocalHourAngle := cosLocalHourAngle(sunTrueLong, dimension.Degrees(geoLocation.Latitude()), zenith)

	localHourAngle := dimension.Degrees(0)

	if isSunrise {
		localHourAngle = 360.0 - dimension.ACos(cosLocalHourAngle)
	} else { // sunset
		localHourAngle = dimension.ACos(cosLocalHourAngle)
	}

	localHour := float64(localHourAngle / degPerHour)

	localMeanTime := localMeanTime(localHour, sunRightAscensionHours,
		approxTimeDays(dayOfYear, hoursFromMeridian(dimension.Degrees(geoLocation.Longitude())), isSunrise))
	processedTime := localMeanTime - hoursFromMeridian(dimension.Degrees(geoLocation.Longitude()))
	for processedTime < 0.0 {
		processedTime += 24.0
	}
	for processedTime >= 24.0 {
		processedTime -= 24.0
	}
	return processedTime

}
