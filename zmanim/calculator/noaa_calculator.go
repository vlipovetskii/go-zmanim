package calculator

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
	"math"
)

/*
noaaCalculator implementation of sunrise and sunset methods to calculate astronomical times based on the
[NOAA]: http://noaa.gov algorithm. This calculator uses the Java algorithm based on the implementation by
[NOAA - National Oceanic and Atmospheric Administration]: http://noaa.gov
[Surface Radiation Research Branch</a>. NOAA's]: http://www.srrb.noaa.gov/highlights/sunrise/sunrise.html
[implementation]: http://www.srrb.noaa.gov/highlights/sunrise/solareqns.PDF is based on equations from
[Astronomical Algorithms]: http://www.willbell.com/math/mc1.htm by
[Jean Meeus]: http://en.wikipedia.org/wiki/Jean_Meeus. Added to the algorithm is an adjustment of the zenith
to account for elevation. The algorithm can be found in the
[Wikipedia Sunrise Equation]: http://en.wikipedia.org/wiki/Sunrise_equation article.
*/
type noaaCalculator struct {
	astronomicalCalculator
}

const (
	/*
		JulianDayJan12000 The [Julian Day]: http://en.wikipedia.org/wiki/Julian_day of January 1, 2000
	*/
	JulianDayJan12000 float64 = 2451545.0

	/*
		JulianDaysPerCentury Julian days per century
	*/
	JulianDaysPerCentury float64 = 36525.0
)

func newNOAACalculator() *noaaCalculator {
	return &noaaCalculator{}
}

func NewNOAACalculator() AstronomicalCalculator {
	t := newNOAACalculator()

	t.initAstronomicalCalculator()

	return t
}

func (t *noaaCalculator) CalculatorName() string {
	return "US National Oceanic and Atmospheric Administration Algorithm"
}

func (t *noaaCalculator) UTCSunrise(targetDateTime gdt.GDateTime, geoLocation GeoLocation, zenith dimension.Degrees, adjustForElevation bool) float64 {

	var elevation dimension.Meters = 0
	if adjustForElevation {
		elevation = geoLocation.Elevation()
	}

	adjustedZenith := t.adjustZenith(zenith, elevation)

	var sunrise = sunriseUTC(julianDay(targetDateTime), dimension.Degrees(geoLocation.Latitude()), -dimension.Degrees(geoLocation.Longitude()), adjustedZenith)
	sunrise = sunrise / 60

	// ensure that the time is >= 0 and < 24
	for sunrise < 0.0 {
		sunrise += 24.0
	}

	for sunrise >= 24.0 {
		sunrise -= 24.0
	}

	return sunrise
}

func (t *noaaCalculator) UTCSunset(targetDateTime gdt.GDateTime, geoLocation GeoLocation, zenith dimension.Degrees, adjustForElevation bool) float64 {
	var elevation dimension.Meters = 0
	if adjustForElevation {
		elevation = geoLocation.Elevation()
	}

	adjustedZenith := t.adjustZenith(zenith, elevation)

	var sunset = sunsetUTC(julianDay(targetDateTime), dimension.Degrees(geoLocation.Latitude()), -dimension.Degrees(geoLocation.Longitude()), adjustedZenith)
	sunset = sunset / 60

	// ensure that the time is >= 0 and < 24
	for sunset < 0.0 {
		sunset += 24.0
	}

	for sunset >= 24.0 {
		sunset -= 24.0
	}

	return sunset
}

/*
julianDay return the [Julian Day]: http://en.wikipedia.org/wiki/Julian_day
the Julian Day corresponding to the date Note: Number is returned for start of Day. Fractional days
should be added later.
tm the time.Time
*/
func julianDay(gDateTime gdt.GDateTime) float64 {

	a := math.Floor(float64(14-gDateTime.D.Month) / 12)
	y := float64(gDateTime.D.Year+4800) - a
	m := float64(gDateTime.D.Month) + 12*a - 3

	jdn := float64(gDateTime.D.Day) + math.Floor((153*m+2)/5) + 365*y + math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400) - 32045

	jd := jdn + float64(gDateTime.T.Hour-12)/24 + float64(gDateTime.T.Minute)/1440 + float64(gDateTime.T.Second)/86400 + float64(gDateTime.T.Nanosecond*1000)/86400000000

	return jd

	/*
		py-zmanim
					"""Month (1-12)"""
				   a = math.floor((14-dt.Month)/12)
				   y = dt.Year + 4800 - a
				   m = dt.Month + 12*a - 3

				   jdn = dt.Day + math.floor((153*m + 2)/5) + 365*y + math.floor(y/4) - math.floor(y/100) + math.floor(y/400) - 32045

				   jd = jdn + (dt.hour - 12) / 24 + dt.minute / 1440 + dt.second / 86400 + dt.microsecond / 86400000000

				   return __to_format(jd, fmt)
				...
			def __to_format(jd: float, fmt: str) -> float:
				    if fmt.lower() == 'jd':
				        return jd
	*/

	/*	var Year = int32(tm.Year())
		var Month = tm.Month()
		var Day = int32(tm.Day())

		// if Month <= 2 {
		if Month <= time.February {
			Year--
			Month += 12
		}

		var a int32 = Year / 100
		var b int32 = 2 - a + a/4

		return math.Floor(365.25*float64(Year+4716)) + math.Floor(30.6001*float64(Month+1)) + float64(Day+b) - 1524.5
	*/
}

/*
julianCenturiesFromJulianDay Convert [Julian Day]: http://en.wikipedia.org/wiki/Julian_day to centuries since J2000.0.
the centuries since 2000 Julian corresponding to the Julian Day
julianDay the Julian Day to convert
*/
func julianCenturiesFromJulianDay(julianDay float64) float64 {
	return (julianDay - JulianDayJan12000) / JulianDaysPerCentury
}

/*
julianDayFromJulianCenturies Convert centuries since J2000.0 to [Julian Day]: http://en.wikipedia.org/wiki/Julian_day.
the Julian Day corresponding to the Julian centuries passed in
julianCenturies the number of Julian centuries since J2000.0
*/
func julianDayFromJulianCenturies(julianCenturies float64) float64 {
	return julianCenturies*JulianDaysPerCentury + JulianDayJan12000
}

/*
sunGeometricMeanLongitude returns the Geometric [Mean Longitude]: http://en.wikipedia.org/wiki/Mean_longitude of the Sun.
the Geometric Mean Longitude of the Sun in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func sunGeometricMeanLongitude(julianCenturies float64) float64 {
	longitude := 280.46646 + julianCenturies*(36000.76983+0.0003032*julianCenturies)

	for longitude > 360.0 {
		longitude -= 360.0
	}

	for longitude < 0.0 {
		longitude += 360.0
	}

	return longitude // in Degrees
}

/*
sunGeometricMeanAnomaly returns the Geometric [Mean Anomaly]: http://en.wikipedia.org/wiki/Mean_anomaly of the Sun.
the Geometric Mean Anomaly of the Sun in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func sunGeometricMeanAnomaly(julianCenturies float64) float64 {
	return 357.52911 + julianCenturies*(35999.05029-0.0001537*julianCenturies) // in Degrees
}

/*
earthOrbitEccentricity return the [eccentricity of earth's orbit]: http://en.wikipedia.org/wiki/Eccentricity_%28orbit%29.
the unitless eccentricity
julianCenturies the number of Julian centuries since J2000.0
*/
func earthOrbitEccentricity(julianCenturies float64) float64 {
	return 0.016708634 - julianCenturies*(0.000042037+0.0000001267*julianCenturies) // unitless
}

/*
sunEquationOfCenter returns the [equation of center]: http://en.wikipedia.org/wiki/Equation_of_the_center for the sun.
the equation of center for the sun in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func sunEquationOfCenter(julianCenturies float64) float64 {
	m := sunGeometricMeanAnomaly(julianCenturies)

	mrad := dimension.Degrees(m).ToRadians() // math.ToRadians(m)
	sinm := math.Sin(float64(mrad))
	sin2m := math.Sin(float64(mrad + mrad))
	sin3m := math.Sin(float64(mrad + mrad + mrad))

	return sinm*(1.914602-julianCenturies*(0.004817+0.000014*julianCenturies)) + sin2m*(0.019993-0.000101*julianCenturies) + sin3m*0.000289 // in Degrees
}

/*
sunTrueLongitudeFromJulianCenturies return the true longitude of the Sun.
the sun's true longitude in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func sunTrueLongitudeFromJulianCenturies(julianCenturies float64) float64 {
	sunLongitude := sunGeometricMeanLongitude(julianCenturies)
	center := sunEquationOfCenter(julianCenturies)

	return sunLongitude + center // in Degrees
}

/*
sunApparentLongitude return the apparent longitude of the Sun.
Sun's apparent longitude in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func sunApparentLongitude(julianCenturies float64) float64 {
	sunTrueLongitude := sunTrueLongitudeFromJulianCenturies(julianCenturies)

	omega := 125.04 - 1934.136*julianCenturies
	lambda := sunTrueLongitude - 0.00569 - 0.00478*math.Sin(float64(dimension.Degrees(omega).ToRadians()))
	return lambda // in Degrees
}

/*
meanObliquityOfEcliptic returns the mean [obliquity of the ecliptic]: http://en.wikipedia.org/wiki/Axial_tilt (Axial tilt).
the mean obliquity in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func meanObliquityOfEcliptic(julianCenturies float64) float64 {
	seconds := 21.448 - julianCenturies*(46.8150+julianCenturies*(0.00059-julianCenturies*(0.001813)))
	return 23.0 + (26.0+(seconds/60.0))/60.0 // in Degrees
}

/*
obliquityCorrection returns the corrected [obliquity of the ecliptic]: http://en.wikipedia.org/wiki/Axial_tilt (Axial tilt).
the corrected obliquity in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func obliquityCorrection(julianCenturies float64) float64 {
	obliquityOfEcliptic := meanObliquityOfEcliptic(julianCenturies)

	omega := 125.04 - 1934.136*julianCenturies
	return obliquityOfEcliptic + 0.00256*math.Cos(float64(dimension.Degrees(omega).ToRadians())) // in Degrees
}

/*
sunDeclination return the [declination]: http://en.wikipedia.org/wiki/Declination of the sun.
the sun's declination in Degrees
julianCenturies the number of Julian centuries since J2000.0
*/
func sunDeclination(julianCenturies float64) dimension.Degrees {
	correction := dimension.Degrees(obliquityCorrection(julianCenturies)).ToRadians()
	apparentLongitude := dimension.Degrees(sunApparentLongitude(julianCenturies)).ToRadians()
	sint := math.Sin(float64(correction)) * math.Sin(float64(apparentLongitude))
	return dimension.Radians(math.Asin(sint)).ToDegrees()
}

/*
equationOfTime return the [Equation of Time]: http://en.wikipedia.org/wiki/Equation_of_time - the difference between
true solar time and mean solar time.
equation of time in minutes of time
julianCenturies the number of Julian centuries since J2000.0
*/
func equationOfTime(julianCenturies float64) float64 {
	epsilon := obliquityCorrection(julianCenturies)
	geomMeanLongSun := sunGeometricMeanLongitude(julianCenturies)
	eccentricityEarthOrbit := earthOrbitEccentricity(julianCenturies)
	geomMeanAnomalySun := sunGeometricMeanAnomaly(julianCenturies)

	y := math.Tan(float64(dimension.Degrees(epsilon).ToRadians() / 2.0))
	y *= y

	sin2l0 := math.Sin(float64(2.0 * dimension.Degrees(geomMeanLongSun).ToRadians()))
	sinm := math.Sin(float64(dimension.Degrees(geomMeanAnomalySun).ToRadians()))
	cos2l0 := math.Cos(float64(2.0 * dimension.Degrees(geomMeanLongSun).ToRadians()))
	sin4l0 := math.Sin(float64(4.0 * dimension.Degrees(geomMeanLongSun).ToRadians()))
	sin2m := math.Sin(float64(2.0 * dimension.Degrees(geomMeanAnomalySun).ToRadians()))

	equationOfTime := y*sin2l0 - 2.0*eccentricityEarthOrbit*sinm + 4.0*eccentricityEarthOrbit*y*sinm*cos2l0 - 0.5*y*y*sin4l0 - 1.25*eccentricityEarthOrbit*eccentricityEarthOrbit*sin2m
	return float64(dimension.Radians(equationOfTime).ToDegrees() * 4.0) // in minutes of time
}

/*
sunHourAngleAtSunrise return the [hour angle]: http://en.wikipedia.org/wiki/Hour_angle of the sun at sunrise for the
latitude.
hour angle of sunrise in Radians
lat the latitude of observer in Degrees
solarDec the declination angle of sun in Degrees
zenith the zenith
*/
func sunHourAngleAtSunrise(lat dimension.Degrees, solarDec dimension.Degrees, zenith dimension.Degrees) float64 {
	latRad := lat.ToRadians()
	sdRad := solarDec.ToRadians()
	zenithRad := zenith.ToRadians()

	return math.Acos(math.Cos(float64(zenithRad))/(math.Cos(float64(latRad))*math.Cos(float64(sdRad))) - math.Tan(float64(latRad))*math.Tan(float64(sdRad))) // in Radians
}

/*
sunHourAngleAtSunset returns the [hour angle]: http://en.wikipedia.org/wiki/Hour_angle of the sun at sunset for the
latitude.
TO-DO: use - sunHourAngleAtSunrise implementation to avoid duplication of code.
the hour angle of sunset in Radians
lat the latitude of observer in Degrees
solarDec the declination angle of sun in Degrees
zenith the zenith
*/
func sunHourAngleAtSunset(lat dimension.Degrees, solarDec dimension.Degrees, zenith dimension.Degrees) float64 {
	latRad := lat.ToRadians()
	sdRad := solarDec.ToRadians()

	hourAngle := math.Acos(math.Cos(float64(zenith.ToRadians()))/(math.Cos(float64(latRad))*math.Cos(float64(sdRad))) - math.Tan(float64(latRad))*math.Tan(float64(sdRad)))
	return -hourAngle // in Radians
}

/*
unused function - remove ?
func calcLongitude(longitude float64) float64 {
	// [How do I use modulus for float/double?]: https://stackoverflow.com/questions/2947044/how-do-i-use-modulus-for-float-double
	// [Floating Point %]: https://www.mindprod.com/jgloss/modulus.html#FLOATINGPOINT
	// % on floating-point operations behaves analogously to the integer remainder operator; this may be compared with the C library function fmod.
	// [What is the equivalent of C++ fmod in Go?]: https://stackoverflow.com/questions/57214080/what-is-the-equivalent-of-c-fmod-in-go
	return math.Mod(-(longitude * 360.0 / 24.0), 360.0)
}
*/

/*
solarElevation return the [Solar Elevation]: http://en.wikipedia.org/wiki/Celestial_coordinate_system for the
horizontal coordinate system at the given location at the given time. Can be negative if the sun is below the
horizon. Not corrected for altitude.
solar elevation in Degrees - horizon is 0 Degrees, civil twilight is -6 Degrees
cal time of calculation
lat latitude of location for calculation
lon longitude of location for calculation
*/
/*
unused function - remove ?
func solarElevation(gDateTime gdt.GDateTime, lat float64, lon float64) float64 {
	julianDay := julianDay(gDateTime)
	julianCenturies := julianCenturiesFromJulianDay(julianDay)

	eot := equationOfTime(julianCenturies)

	longitude := float64(gDateTime.T.Hour+12.0) + (float64(gDateTime.T.Minute)+eot+float64(gDateTime.T.Second)/60.0)/60.0

	longitude = calcLongitude(longitude)
	hourAngleRad := dimension.Degrees(lon - longitude).ToRadians()
	declination := sunDeclination(julianCenturies)
	decRad := declination.ToRadians()
	latRad := dimension.Degrees(lat).ToRadians()

	return float64(dimension.Radians(math.Asin((math.Sin(float64(latRad)) * math.Sin(float64(decRad))) + (math.Cos(float64(latRad)) * math.Cos(float64(decRad)) * math.Cos(float64(hourAngleRad))))).ToDegrees())

}
*/

/*
solarAzimuth return the [Solar Azimuth]: http://en.wikipedia.org/wiki/Celestial_coordinate_system for the
horizontal coordinate system at the given location at the given time. Not corrected for altitude. True south is 0 Degrees.
cal time of calculation
lat latitude of location for calculation
lon longitude of location for calculation
*/
/*
unused function - remove ?
func solarAzimuth(gDateTime gdt.GDateTime, lat float64, lon float64) float64 {
	julianDay := julianDay(gDateTime)
	julianCenturies := julianCenturiesFromJulianDay(julianDay)

	eot := equationOfTime(julianCenturies)

	longitude := float64(gDateTime.T.Hour+12.0) + (float64(gDateTime.T.Minute)+eot+float64(gDateTime.T.Second)/60.0)/60.0

	longitude = calcLongitude(longitude)
	hourAngleRad := dimension.Degrees(lon - longitude).ToRadians()
	declination := sunDeclination(julianCenturies)
	decRad := declination.ToRadians()
	latRad := dimension.Degrees(lat).ToRadians()

	return float64(dimension.Radians(math.Atan(math.Sin(float64(hourAngleRad))/((math.Cos(float64(hourAngleRad))*math.Sin(float64(latRad)))-(math.Tan(float64(decRad))*math.Cos(float64(latRad)))))).ToDegrees() + 180)

}
*/

/*
sunriseUTC return the [Universal Coordinated Time]: http://en.wikipedia.org/wiki/Universal_Coordinated_Time (UTC)
of sunrise for the given Day at the given location on earth
the time in minutes from zero UTC
julianDay the Julian Day
latitude the latitude of observer in Degrees
longitude the longitude of observer in Degrees
zenith the zenith
*/
func sunriseUTC(julianDay float64, latitude dimension.Degrees, longitude dimension.Degrees, zenith dimension.Degrees) float64 {
	julianCenturies := julianCenturiesFromJulianDay(julianDay)

	// Find the time of solar noon at the location, and use that declination. This is better than start of the
	// Julian Day

	noonmin := solarNoonUTC(julianCenturies, longitude)
	tnoon := julianCenturiesFromJulianDay(julianDay + noonmin/1440.0)

	// First pass to approximate sunrise (using solar noon)

	eqTime := equationOfTime(tnoon)
	solarDec := sunDeclination(tnoon)
	hourAngle := sunHourAngleAtSunrise(latitude, solarDec, zenith)

	delta := longitude - dimension.Radians(hourAngle).ToDegrees()
	timeDiff := 4 * delta                       // in minutes of time
	timeUTC := 720 + float64(timeDiff) - eqTime // in minutes

	// Second pass includes fractional Julian Day in gamma calc

	newt := julianCenturiesFromJulianDay(julianDayFromJulianCenturies(julianCenturies) + timeUTC/1440.0)
	eqTime = equationOfTime(newt)
	solarDec = sunDeclination(newt)
	hourAngle = sunHourAngleAtSunrise(latitude, solarDec, zenith)
	delta = longitude - dimension.Radians(hourAngle).ToDegrees()
	timeDiff = 4 * delta
	timeUTC = 720 + float64(timeDiff) - eqTime // in minutes

	return timeUTC
}

/*
solarNoonUTC return the [Universal Coordinated Time]: http://en.wikipedia.org/wiki/Universal_Coordinated_Time (UTC)
[solar noon]: http://en.wikipedia.org/wiki/Noon#Solar_noon for the given Day at the given location on earth.
the time in minutes from zero UTC
julianCenturies the number of Julian centuries since J2000.0
longitude the longitude of observer in Degrees
*/
func solarNoonUTC(julianCenturies float64, longitude dimension.Degrees) float64 {
	// First pass uses approximate solar noon to calculate eqtime
	tnoon := julianCenturiesFromJulianDay(julianDayFromJulianCenturies(julianCenturies) + float64(longitude/360.0))
	eqTime := equationOfTime(tnoon)
	solNoonUTC := 720 + float64(longitude*4) - eqTime // min

	newt := julianCenturiesFromJulianDay(julianDayFromJulianCenturies(julianCenturies) - 0.5 + solNoonUTC/1440.0)

	eqTime = equationOfTime(newt)

	return 720 + float64(longitude*4) - eqTime // min
}

/*
sunsetUTC return the [Universal Coordinated Time]: http://en.wikipedia.org/wiki/Universal_Coordinated_Time (UTC)
of sunset for the given Day at the given location on earth
julianDay the Julian Day
the time in minutes from zero Universal Coordinated Time (UTC)
latitude of observer in Degrees
longitude of observer in Degrees
zenith
*/
func sunsetUTC(julianDay float64, latitude dimension.Degrees, longitude dimension.Degrees, zenith dimension.Degrees) float64 {
	julianCenturies := julianCenturiesFromJulianDay(julianDay)

	// Find the time of solar noon at the location, and use that declination. This is better than start of the
	// Julian Day

	noonmin := solarNoonUTC(julianCenturies, longitude)
	tnoon := julianCenturiesFromJulianDay(julianDay + noonmin/1440.0)

	// First calculates sunrise and approx length of Day

	eqTime := equationOfTime(tnoon)
	solarDec := sunDeclination(tnoon)
	hourAngle := sunHourAngleAtSunset(latitude, solarDec, zenith)

	delta := longitude - dimension.Radians(hourAngle).ToDegrees()
	timeDiff := 4 * delta
	timeUTC := 720 + float64(timeDiff) - eqTime

	// Second pass includes fractional Julian Day in gamma calc

	newt := julianCenturiesFromJulianDay(julianDayFromJulianCenturies(julianCenturies) + timeUTC/1440.0)
	eqTime = equationOfTime(newt)
	solarDec = sunDeclination(newt)
	hourAngle = sunHourAngleAtSunset(latitude, solarDec, zenith)

	delta = longitude - dimension.Radians(hourAngle).ToDegrees()
	timeDiff = 4 * delta
	timeUTC = 720 + float64(timeDiff) - eqTime // in minutes

	return timeUTC
}
