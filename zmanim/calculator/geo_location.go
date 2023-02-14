package calculator

import (
	"fmt"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
	"math"
	"time"
)

type GeoLocation interface {
	// Elevation and other getters
	//
	Elevation() dimension.Meters
	Latitude() float64
	Longitude() float64
	LocationName() string
	TimeZone() *time.Location
	StandardTimeOffset() gdt.GMillisecond
	LocalMeanTimeOffset() gdt.GMillisecond
	AntimeridianAdjustment() int32
	GeodesicInitialBearing(location GeoLocation) float64
	GeodesicFinalBearing(location GeoLocation) float64
	GeodesicDistance(location GeoLocation) float64
	RhumbLineDistance(location GeoLocation) float64
	// SetElevation and other setters
	//
	SetElevation(elevation dimension.Meters)
	SetLatitude1(latitude float64)
	SetLatitude2(degrees dimension.Degrees, minutes dimension.ArcMinutes, seconds dimension.ArcSeconds, direction string)
	SetLongitude1(longitude float64)
	SetLongitude2(degrees dimension.Degrees, minutes dimension.ArcMinutes, seconds dimension.ArcSeconds, direction string)
	SetLocationName(name string)
	SetTimeZone(timeZone *time.Location)
	// VincentyFormula ...
	VincentyFormula(location GeoLocation, formula int32) float64
}

/*
geoLocation
--- Time zones https://yourbasic.org/golang/time-change-convert-location-timezone/
fmt.Println(t.Location(), t.Format("15:04"))
UTC 19:32
Local 20:32
Asia/Shanghai 03:32
America/Metropolis <time unknown>
type Location https://pkg.go.dev/time#Location
*/
type geoLocation struct {
	/*
		latitude the latitude in a double format such as 40.095965 for Lakewood, NJ.
		Note: For latitudes south of the equator, a negative value should be used.
	*/
	latitude float64
	/*
	   longitude double the longitude in a double format such as -74.222130 for Lakewood, NJ.
	   Note: For longitudes east of the [Prime]: https://en.wikipedia.org/wiki/Prime_Meridian
	   Meridian (Greenwich), a negative value should be used.
	*/
	longitude float64
	// locationName the location name for display use such as "Lakewood, NJ"
	locationName string
	/*
	   timeZone the TimeZone for the location.
	*/
	timeZone *time.Location
	// elevation the elevation above sea level in dimension.Meters. Elevation is not used in most algorithms used for calculating
	elevation dimension.Meters
}

const (

	/*
		Distance constant for a distance type calculation.
	*/
	Distance = 0

	/*
		InitialBearing constant for an initial bearing type calculation.
	*/
	InitialBearing = 1

	/*
		FinalBearing constant for a final bearing type calculation.
	*/
	FinalBearing = 2
)

func newGeoLocation() *geoLocation {
	return &geoLocation{}
}

func NewGeoLocation() GeoLocation {
	t := newGeoLocation()

	t.SetLocationName("Greenwich, England")
	t.SetLongitude1(0) // added for clarity
	t.SetLatitude1(51.4772)
	t.SetTimeZone(timeutil.GmtTimezoneOrPanic())

	return t
}

func NewGeoLocation1(name string, latitude float64, longitude float64, timeZone *time.Location) GeoLocation {
	t := newGeoLocation()

	t.SetLocationName(name)
	t.SetLongitude1(longitude)
	t.SetLatitude1(latitude)
	t.SetTimeZone(timeZone)

	return t
}

func NewGeoLocation2(name string, latitude float64, longitude float64, elevation dimension.Meters, timeZone *time.Location) GeoLocation {
	t := newGeoLocation()

	t.SetLocationName(name)
	t.SetLongitude1(longitude)
	t.SetLatitude1(latitude)
	t.SetTimeZone(timeZone)

	t.SetElevation(elevation)

	return t
}

func (t *geoLocation) Elevation() dimension.Meters {
	return t.elevation
}

func (t *geoLocation) SetElevation(elevation dimension.Meters) {
	if elevation < 0 {
		// throw new IllegalArgumentException("Elevation cannot be negative");
		helper.Panic("Elevation cannot be negative")
	}
	/*
				How to check for NaN in golang https://stackoverflow.com/questions/31246192/how-to-check-for-nan-in-golang
				math.IsNaN() Function in Golang with Examples https://www.geeksforgeeks.org/math-isnan-function-in-golang-with-examples/
				https://go.dev/src/runtime/float.go
		func isFinite(f float64) bool {
			return !isNaN(f - f)
		}
			// isInf reports whether f is an infinity.
			func isInf(f float64) bool {
				return !isNaN(f) && !isFinite(f)
			}
	*/
	if math.IsNaN(float64(elevation)) || math.IsNaN(float64(elevation-elevation)) {
		// throw new IllegalArgumentException("Elevation must not be NaN or infinite");
		panic("Elevation must not be NaN or infinite")
	}
	t.elevation = elevation
}

func (t *geoLocation) SetLatitude1(latitude float64) {
	if latitude > 90 || latitude < -90 {
		panic("Latitude must be between -90 and  90")
	}
	t.latitude = latitude
}

/*
SetLatitude2 method to set the latitude in Degrees, minutes and seconds.
Degrees The Degrees of latitude to set between 0&deg; and 90&deg;. For example 40 would be used for Lakewood, NJ.
An IllegalArgumentException will be thrown if the value exceeds the limit.
minutes [minutes of arc]: https://en.wikipedia.org/wiki/Minute_of_arc#Cartography
seconds [seconds of arc]: https://en.wikipedia.org/wiki/Minute_of_arc#Cartography
direction N for north and S for south. A panic will be if the value is not S or N.
*/
func (t *geoLocation) SetLatitude2(degrees dimension.Degrees, minutes dimension.ArcMinutes, seconds dimension.ArcSeconds, direction string) {

	if degrees < 0 {
		panic(fmt.Sprintf("degrees %v is negative.", degrees))
	}

	if minutes < 0 {
		panic(fmt.Sprintf("minutes %v is negative.", minutes))
	}

	if seconds < 0 {
		panic(fmt.Sprintf("seconds %v is negative.", seconds))
	}

	tempLat := float64(degrees) + ((float64(minutes) + (float64(seconds) / 60.0)) / 60.0)
	if tempLat > 90 || tempLat < 0 {
		panic("Latitude must be between 0 and  90. Use direction of S instead of negative.")
	}

	if direction == "S" {
		tempLat *= -1
	} else if direction != "N" {
		panic("Latitude direction must be N or S")
	}
	t.latitude = tempLat

}

func (t *geoLocation) Latitude() float64 {
	return t.latitude
}

/*
SetLongitude1 method to set the longitude in a double format.
longitude The Degrees of longitude to set in a double format between -180 deg; and 180 deg;.
A panic will be if the value exceeds the limit. For example -74.2094 would be
used for Lakewood, NJ. Note: for longitudes east of the <a
[Prime Meridian]: https://en.wikipedia.org/wiki/Prime_Meridian (Greenwich) a negative value
should be used.
*/
func (t *geoLocation) SetLongitude1(longitude float64) {
	if longitude > 180 || longitude < -180 {
		panic("Longitude must be between -180 and  180")
	}
	t.longitude = longitude
}

/*
SetLongitude2 method to set the longitude in Degrees, minutes and seconds.
Degrees the Degrees of longitude to set between 0 deg; and 180 deg;. As an example 74 would be set for Lakewood, NJ.
A panic will be if the value exceeds the limits.
minutes [minutes of arc]: https://en.wikipedia.org/wiki/Minute_of_arc#Cartography
seconds [seconds of arc]: https://en.wikipedia.org/wiki/Minute_of_arc#Cartography
direction E for east of the [Prime Meridian]: https://en.wikipedia.org/wiki/Prime_Meridian or W for west of it.
A panic will be if the value is not E or W.
*/
func (t *geoLocation) SetLongitude2(degrees dimension.Degrees, minutes dimension.ArcMinutes, seconds dimension.ArcSeconds, direction string) {

	if degrees < 0 {
		panic(fmt.Sprintf("degrees %v is negative.", degrees))
	}

	if minutes < 0 {
		panic(fmt.Sprintf("minutes %v is negative.", minutes))
	}

	if seconds < 0 {
		panic(fmt.Sprintf("seconds %v is negative.", seconds))
	}

	longTemp := float64(degrees) + ((float64(minutes) + (float64(seconds) / 60.0)) / 60.0)
	if longTemp > 180 || t.longitude < 0 {
		panic("Longitude must be between 0 and  180.  Use a direction of W instead of negative.")
	}

	if direction == "W" {
		longTemp *= -1
	} else if direction != "E" {
		panic("Longitude direction must be E or W")
	}
	t.longitude = longTemp
}

func (t *geoLocation) Longitude() float64 {
	return t.longitude
}

func (t *geoLocation) LocationName() string {
	return t.locationName
}

func (t *geoLocation) SetLocationName(name string) {
	t.locationName = name
}

func (t *geoLocation) TimeZone() *time.Location {
	return t.timeZone
}

func (t *geoLocation) SetTimeZone(timeZone *time.Location) {
	t.timeZone = timeZone
}

/*
StandardTimeOffset returns the amount of time in milliseconds to add to UTC to get standard time in this time zone
See TimeZone.java public abstract int getRawOffset()
*/
func (t *geoLocation) StandardTimeOffset() gdt.GMillisecond {
	/*
		From py-zmanim
		def standard_time_offset(self) -> int:
		   now = datetime.now(tz=self.time_zone)
		   return int((now.utcoffset() - now.dst()).total_seconds()) * 1000

	*/
	now := time.Now() // Now returns the current local time
	_, destOffset := now.In(t.TimeZone()).Zone()
	_, utcOffset := now.UTC().Zone()

	return gdt.GSecond(destOffset - utcOffset).ToMilliseconds()
}

/*
LocalMeanTimeOffset a method that will return the location's local mean time offset in gdt.GMillisecond from local
[standard time]: https://en.wikipedia.org/wiki/Standard_time. The globe is split into 360&deg;, with
15 deg; per hour of the Day. For a local that is at a longitude that is evenly divisible by 15 (longitude % 15 == 0),
at solar SunTransit noon (with adjustment for the [equation of time]: https://en.wikipedia.org/wiki/Equation_of_time) the sun should be directly overhead,
so a user who is 1&deg; west of this will have noon at 4 minutes after standard time noon, and conversely, a user
who is 1 deg; east of the 15 deg; longitude will have noon at 11:56 AM. Lakewood, N.J., whose longitude is
-74.2094, is 0.7906 away from the closest multiple of 15 at -75 deg;. This is multiplied by 4 to yield 3 minutes
and 10 seconds earlier than standard time. The offset returned does not account for the [Daylight saving time]: https://en.wikipedia.org/wiki/Daylight_saving_time offset since this class is
unaware of dates.
*/
func (t *geoLocation) LocalMeanTimeOffset() gdt.GMillisecond {
	/**
	TimeZone GetRawOffset() Method in Java with Examples https://www.geeksforgeeks.org/timezone-getrawoffset-method-in-java-with-examples/
	The GetRawOffset() method of TimeZone class in Java is used to know the amount of time in milliseconds needed to be added to the UTC to get the standard time in this TimeZone.
	*/
	return gdt.GMillisecond(t.Longitude()*4*timeutil.MinuteMillis - float64(t.StandardTimeOffset()))
}

/*
AntimeridianAdjustment adjust the date for [antimeridian]: https://en.wikipedia.org/wiki/180th_meridian crossover. This is
needed to deal with edge cases such as Samoa that use a different calendar date than expected based on their geographic location.
The actual Time Zone offset may deviate from the expected offset based on the longitude. Since the 'absolute time'
calculations are always based on longitudinal offset from UTC for a given date, the date is presumed to only
increase East of the Prime Meridian, and to only decrease West of it. For Time Zones that cross the antimeridian,
the date will be artificially adjusted before calculation to conform with this presumption.
For example, Apia, Samoa with a longitude of -171.75 uses a local offset of +14:00.  When calculating sunrise for
2018-02-03, the calculator should operate using 2018-02-02 since the expected zone is -11.  After determining the
UTC time, the local DST offset of [UTC+14:00]: https://en.wikipedia.org/wiki/UTC%2B14:00 should be applied
to bring the date back to 2018-02-03.
*/
func (t *geoLocation) AntimeridianAdjustment() int32 {
	localHoursOffset := t.LocalMeanTimeOffset() / timeutil.HourMillis

	if localHoursOffset >= 20 { // if the offset is 20 hours or more in the future (never expected anywhere other
		// than a location using a timezone across the anti meridian to the east such as Samoa)
		return 1 // roll the date forward a Day
	} else if localHoursOffset <= -20 { // if the offset is 20 hours or more in the past (no current location is known
		//that crosses the antimeridian to the west, but better safe than sorry)
		return -1 // roll the date back a Day
	}
	return 0 //99.999% of the world will have no adjustment
}

/*
GeodesicInitialBearing calculate the initial [geodesic]: https://en.wikipedia.org/wiki/Great_circle bearing between this
Object and a second Object passed to this method using [Thaddeus Vincenty's inverse formula See T Vincenty]: https://en.wikipedia.org/wiki/Thaddeus_Vincenty
[Direct and Inverse Solutions of Geodesics on the Ellipsoid with application of nested equations]: https://www.ngs.noaa.gov/PUBS_LIB/inverse.pdf, Survey Review, vol XXII no 176, 1975
*/
func (t *geoLocation) GeodesicInitialBearing(location GeoLocation) float64 {
	return t.VincentyFormula(location, InitialBearing)
}

/*
GeodesicFinalBearing calculate the final [geodesic]: https://en.wikipedia.org/wiki/Great_circle bearing between this Object
and a second Object passed to this method using [Thaddeus Vincenty's inverse formula See T Vincenty]: https://en.wikipedia.org/wiki/Thaddeus_Vincenty,
[Direct and  - Inverse Solutions of Geodesics on the Ellipsoid with application of nested equations]: https://www.ngs.noaa.gov/PUBS_LIB/inverse.pdf, Survey Review, vol
XXII no 176, 1975
*/
func (t *geoLocation) GeodesicFinalBearing(location GeoLocation) float64 {
	return t.VincentyFormula(location, FinalBearing)
}

/*
GeodesicDistance Calculate [geodesic distance]: https://en.wikipedia.org/wiki/Great-circle_distance in Meters between
this Object and a second Object passed to this method using [Thaddeus Vincenty's]: https://en.wikipedia.org/wiki/Thaddeus_Vincenty inverse formula See T Vincenty,
[Direct and Inverse Solutions of Geodesics on the Ellipsoid with application of nested equations]: https://www.ngs.noaa.gov/PUBS_LIB/inverse.pdf, Survey Review, vol XXII no 176, 1975
*/
func (t *geoLocation) GeodesicDistance(location GeoLocation) float64 {
	return t.VincentyFormula(location, Distance)
}

/*
VincentyFormula calculate [geodesic distance]: https://en.wikipedia.org/wiki/Great-circle_distance in Meters between
this Object and a second Object passed to this method using [Thaddeus Vincenty's> inverse formula See T Vincenty]: https://en.wikipedia.org/wiki/Thaddeus_Vincenty,
[Direct and Inverse Solutions of Geodesics on the Ellipsoid with application of nested equations]: https://www.ngs.noaa.gov/PUBS_LIB/inverse.pdf, Survey Review, vol XXII no 176, 1975
*/
func (t *geoLocation) VincentyFormula(location GeoLocation, formula int32) float64 {
	var a = 6378137.0
	var b = 6356752.3142
	var f = 1 / 298.257223563 // WGS-84 ellipsiod
	var L = float64(dimension.Degrees(location.Longitude() - t.Longitude()).ToRadians())
	var U1 = math.Atan((1 - f) * math.Tan(float64(dimension.Degrees(t.Latitude()).ToRadians())))
	var U2 = math.Atan((1 - f) * math.Tan(float64(dimension.Degrees(location.Latitude()).ToRadians())))
	var sinU1 = math.Sin(U1)
	var cosU1 = math.Cos(U1)
	var sinU2 = math.Sin(U2)
	var cosU2 = math.Cos(U2)

	lambda := L
	lambdaP := 2 * math.Pi
	var iterLimit float64 = 20
	var sinLambda float64 = 0
	var cosLambda float64 = 0
	var sinSigma float64 = 0
	var cosSigma float64 = 0
	var sigma float64 = 0
	var sinAlpha float64 = 0
	var cosSqAlpha float64 = 0
	var cos2SigmaM float64 = 0
	var C float64

	// while (Math.abs(lambda - lambdaP) > 1e-12 && --iterLimit > 0) {
	for math.Abs(lambda-lambdaP) > 1e-12 {

		iterLimit--
		if iterLimit <= 0 {
			break
		}

		sinLambda = math.Sin(lambda)
		cosLambda = math.Cos(lambda)
		sinSigma = math.Sqrt((cosU2*sinLambda)*(cosU2*sinLambda) + (cosU1*sinU2-sinU1*cosU2*cosLambda)*(cosU1*sinU2-sinU1*cosU2*cosLambda))
		if sinSigma == 0 {
			return 0 // co-incident points
		}
		cosSigma = sinU1*sinU2 + cosU1*cosU2*cosLambda
		sigma = math.Atan2(sinSigma, cosSigma)
		sinAlpha = cosU1 * cosU2 * sinLambda / sinSigma
		cosSqAlpha = 1 - sinAlpha*sinAlpha
		cos2SigmaM = cosSigma - 2*sinU1*sinU2/cosSqAlpha
		if math.IsNaN(cos2SigmaM) {
			cos2SigmaM = 0 // equatorial line: cosSqAlpha=0 (§6)
		}
		C = f / 16 * cosSqAlpha * (4 + f*(4-3*cosSqAlpha))
		lambdaP = lambda
		lambda = L + (1-C)*f*sinAlpha*(sigma+C*sinSigma*(cos2SigmaM+C*cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)))
	}

	if iterLimit == 0 {
		return math.NaN() // formula failed to converge
	}

	uSq := cosSqAlpha * (a*a - b*b) / (b * b)
	A := 1 + uSq/16384*(4096+uSq*(-768+uSq*(320-175*uSq)))
	B := uSq / 1024 * (256 + uSq*(-128+uSq*(74-47*uSq)))
	deltaSigma := B * sinSigma * (cos2SigmaM + B/4*(cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)-B/6*cos2SigmaM*(-3+4*sinSigma*sinSigma)*(-3+4*cos2SigmaM*cos2SigmaM)))
	distance := b * A * (sigma - deltaSigma)

	// initial bearing
	fwdAz := float64(dimension.Radians(math.Atan2(cosU2*sinLambda, cosU1*sinU2-sinU1*cosU2*cosLambda)).ToDegrees())
	// final bearing
	revAz := float64(dimension.Radians(math.Atan2(cosU1*sinLambda, -sinU1*cosU2+cosU1*sinU2*cosLambda)).ToDegrees())
	if formula == Distance {
		return distance
	} else if formula == InitialBearing {
		return fwdAz
	} else if formula == FinalBearing {
		return revAz
	} else { // should never happen
		return math.NaN()
	}
}

/*
RhumbLineDistance returns the [rhumb line]: https://en.wikipedia.org/wiki/Rhumb_line distance from the current location
to the GeoLocation passed in.
*/
func (t *geoLocation) RhumbLineDistance(location GeoLocation) float64 {
	var earthRadius float64 = 6378137 // Earth's radius in meters (WGS-84)
	dLat := float64(dimension.Degrees(location.Latitude()).ToRadians() - dimension.Degrees(t.Latitude()).ToRadians())
	dLon := math.Abs(float64(dimension.Degrees(location.Longitude()).ToRadians() - dimension.Degrees(t.Longitude()).ToRadians()))
	dPhi := math.Log(math.Tan(float64(dimension.Degrees(location.Latitude()).ToRadians()/2+math.Pi/4))) / math.Tan(float64(dimension.Degrees(t.Latitude()).ToRadians()/2+math.Pi/4))
	q := dLat / dPhi

	// if (!(Math.abs(q) <= Double.MAX_VALUE)) {
	if !(math.Abs(q) <= math.MaxFloat64) {
		q = math.Cos(float64(dimension.Degrees(t.Latitude()).ToRadians()))
	}
	// if dLon over 180° take shorter rhumb across 180° meridian:
	if dLon > math.Pi {
		dLon = 2*math.Pi - dLon
	}
	d := math.Sqrt(dLat*dLat + q*q*dLon*dLon)
	return d * earthRadius
}

func (t *geoLocation) String() string {
	/*
			StringBuilder sb = new StringBuilder();
		sb.append("\nLocation Name:\t\t\t").append(getLocationName());
		sb.append("\nLatitude:\t\t\t").append(getLatitude()).append("\u00B0");
		sb.append("\nLongitude:\t\t\t").append(getLongitude()).append("\u00B0");
		sb.append("\nElevation:\t\t\t").append(getElevation()).append(" Meters");
		sb.append("\nTimezone ID:\t\t\t").append(getTimeZone().getID());
		sb.append("\nTimezone Display Name:\t\t").append(getTimeZone().getDisplayName())
				.append(" (").append(getTimeZone().getDisplayName(false, TimeZone.SHORT)).append(")");
		sb.append("\nTimezone GMT Offset:\t\t").append(getTimeZone().getRawOffset() / HOUR_MILLIS);
		sb.append("\nTimezone DST Offset:\t\t").append(getTimeZone().getDSTSavings() / HOUR_MILLIS);
		return sb.toString();

	*/
	return "" //fmt.Sprintf("%b", b)
}
