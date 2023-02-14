package calculator

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
	"math"
)

type AstronomicalCalculator interface {
	/*
	 GetCalculatorName returns the name of the algorithm.
	*/
	CalculatorName() string
	/*
		GetUTCSunrise a method that calculates UTC sunrise as well as any time based on an angle above or below sunrise.
		The UTC time of sunrise in 24 hours format. 5:45:00 AM will return 5.75.0. If an error was encountered in
		the calculation expected behavior for some locations such as near the poles, NaN will be returned.
		This abstract method is implemented by the classes that extend this class.
		calendar Used to calculate Day of Year.
		geoLocation The location information used for astronomical calculating sun times.
		zenith the azimuth below the vertical zenith of 90 Degrees. for sunrise typically the adjustZenith zenith used for the calculation uses geometric zenith of 90&deg; and {@link #adjustZenith adjusts}
		this slightly to account for solar refraction and the sun's radius. Another example would be BeginNauticalTwilight that passes
		NauticalZenith to this method.
		adjustForElevation should the time be adjusted for elevation
	*/
	UTCSunrise(targetDateTime gdt.GDateTime, geoLocation GeoLocation, zenith dimension.Degrees, adjustForElevation bool) float64
	/*
	 GetUTCSunset a method that calculates UTC sunset as well as any time based on an angle above or below sunset.
	 The UTC time of sunset in 24 hours format. 5:45:00 AM will return 5.75.0. If an error was encountered in
	 the calculation expected behavior for some locations such as near the poles,
	 NaN will be returned.
	 This abstract method is implemented by the classes that extend this class.
	 calendar Used to calculate Day of Year.
	 geoLocation The location information used for astronomical calculating sun times.
	 zenith the azimuth below the vertical zenith of 90 deg;. For sunset typically the adjustZenith
	 used for the calculation uses geometric zenith of 90 deg; and adjustZenith adjusts
	 this slightly to account for solar refraction and the sun's radius. Another example would be
	 EndNauticalTwilight that passes
	 NAUTICAL_ZENITH to this method.
	 adjustForElevation should the time be adjusted for elevation
	*/
	UTCSunset(targetDateTime gdt.GDateTime, geoLocation GeoLocation, zenith dimension.Degrees, adjustForElevation bool) float64
}

/*
astronomicalCalculator An abstract class that all sun time calculating classes extend. This allows the algorithm used to be changed at
runtime, easily allowing comparison the results of using different algorithms.
*/
type astronomicalCalculator struct {
	/*
		The commonly used average solar refraction.
			Calendar calculations lists a more accurate global average of 34.478885263888294
			 refraction value to be used when calculating sunrise and sunset.
			The default value is 34 dimension.ArcMinutes. The [Errata and Notes for Calendrical Calculations: The Millennium Edition]: https://web.archive.org/web/20150915094635/http://emr.cs.iit.edu/home/reingold/calendar-book/second-edition/errata.pdf
			 by Edward M. Reingold and Nachum Dershowitz
			 lists the actual average refraction value as 34.478885263888294 or approximately 34' 29". The refraction value as well
			 as the solarRadius and elevation adjustment are added to the zenith used to calculate sunrise and sunset.
	*/
	refraction dimension.ArcMinutes

	/*
	 The commonly used average solarRadius in minutes of a degree.
	 The default value is 16 dimension.ArcMinutes. The sun's radius as it appears from earth is
	 almost universally given as 16 arc minutes but in fact it differs by the time of the Year. At the
	 [perihelion]: https://en.wikipedia.org/wiki/Perihelion it has an apparent radius of 16.293, while at the
	 [aphelion]: https://en.wikipedia.org/wiki/Aphelion it has an apparent radius of 15.755. There is little
	 affect for most location, but at high and low latitudes the difference becomes more apparent. My Calculations for
	 the difference at the location of the [Royal Observatory, Greenwich]: https://www.rmg.co.uk/royal-observatory
	 shows only a 4.494 seconds difference between the perihelion and aphelion radii, but moving into the arctic circle the
	 difference becomes more noticeable. Tests for Tromso, Norway (latitude 69.672312, longitude 19.049787) show that
	 on May 17, the rise of the midnight sun, a 2 minute 23 seconds difference is observed between the perihelion and
	 aphelion radii using the USNO algorithm, but only 1 minute and 6 seconds difference using the NOAA algorithm.
	 Areas farther north show an even greater difference. Note that these test are not real valid test cases because
	 they show the extreme difference on days that are not the perihelion or aphelion, but are shown for illustrative
	 purposes only.
	*/
	solarRadius dimension.ArcMinutes

	/*
	 The commonly used average earthRadius in KM. At this time, this only affects elevation adjustment and not the
	 sunrise and sunset calculations. The value currently defaults to 6356.9 dimension.KM.
	*/
	earthRadius dimension.KM
}

func (t *astronomicalCalculator) initAstronomicalCalculator() {
	t.refraction = dimension.ArcMinutes(float64(34) / 60)
	t.solarRadius = dimension.ArcMinutes(float64(16) / 60)
	t.earthRadius = 6356.9
}

/*
elevationAdjustment method to return the adjustment to the zenith required to account for the elevation. Since a person at a higher
elevation can see farther below the horizon, the calculation for sunrise / sunset is calculated below the horizon
used at sea level. This is only used for sunrise and sunset and not times before or after it such as
BeginNauticalTwilight nautical twilight since those
calculations are based on the level of available light at the given dip below the horizon, something that is not
affected by elevation, the adjustment should only be made, if the zenith == 90 deg; adjustZenith adjusted
for refraction and solar radius. The algorithm used is
elevationAdjustment = Math.ToDegrees(Math.acos(earthRadiusInMeters / (earthRadiusInMeters + elevationMeters)));
The source of this algorithm is [Calendrical Calculations]: http://www.calendarists.com by Edward M.
Reingold and Nachum Dershowitz. An alternate algorithm that produces an almost identical (but not accurate)
result found in Ma'aglay Tzedek by Moishe Kosower and other sources is:
elevationAdjustment = 0.0347 * Math.sqrt(elevationMeters) in dimension.Meters.
The method return the UTC time of sunrise in 24 hours format. 5:45:00 AM will return 5.75.0.
If an error was encountered in the calculation (expected behavior for some locations such as near the poles),
math.NaN will be returned.
*/
func (t *astronomicalCalculator) elevationAdjustment(elevation dimension.Meters) dimension.Degrees {
	return dimension.Radians(math.Acos(float64(t.earthRadius / (t.earthRadius + elevation.ToKM())))).ToDegrees()
}

/*
adjustZenith adjusts the zenith of astronomical sunrise and sunset to account for solar refraction, solar radius and
elevation.
The zenith adjusted to include the solarRadius sun's radius, refraction
and elevationAdjustment. This will only be adjusted for sunrise and sunset, if the zenith == 90 deg;
The value for Sun's zenith and true rise/set Zenith (used in this class and subclasses) is the angle
that the center of the Sun makes to a line perpendicular to the Earth's surface. If the Sun were a point and the
Earth were without an atmosphere, true sunset and sunrise would correspond to a 90&deg; zenith. Because the Sun
is not a point, and because the atmosphere refracts light, this 90&deg; zenith does not, in fact, correspond to
true sunset or sunrise, instead the center of the Sun's disk must lie just below the horizon for the upper edge
to be obscured. This means that a zenith of just above 90 deg; must be used. The Sun subtends an angle of 16
minutes of arc this can be changed via the {@link #setSolarRadius(double)} method , and atmospheric refraction
accounts for 34 minutes or so (this can be changed via the {@link #setRefraction(double)} method), giving a total
of 50 arcminutes. The total value for ZENITH is 90+(5/6) or 90.8333333&deg; for true sunrise/sunset. Since a
person at an elevation can see below the horizon of a person at sea level, this will also adjust the zenith to
account for elevation if available. Note that this will only adjust the value if the zenith is exactly 90 Degrees.
For values below and above this no correction is done. As an example, astronomical twilight is when the sun is
18 deg; below the horizon or {@link com.kosherjava.zmanim.AstronomicalCalendar#ASTRONOMICAL_ZENITH 108&deg;
below the zenith}. This is traditionally calculated with none of the mentioned above adjustments. The same goes
for various tzais and alos times such as the
ZmanimCalendar#ZENITH_16_POINT_1 16.1 deg; dip used in
ComplexZmanimCalendar.getAlos16Point1Degrees.
zenith the azimuth below the vertical zenith of 90 deg;. For sunset typically the adjustZenith zenith
used for the calculation uses geometric zenith of 90 deg; and adjustZenith adjusts
this slightly to account for solar refraction and the sun's radius. Another example would be
AstronomicalCalendar.getEndNauticalTwilight that passes
NauticalZenith to this method.
elevation in dimension.Meters.
*/
func (t *astronomicalCalculator) adjustZenith(zenith dimension.Degrees, elevation dimension.Meters) dimension.Degrees {
	if zenith != GeometricZenith {
		return zenith
	} else {
		return zenith + (t.solarRadius.ToDegrees() + t.refraction.ToDegrees() + t.elevationAdjustment(elevation))
	}
}
