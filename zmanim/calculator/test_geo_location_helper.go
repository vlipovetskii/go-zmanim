package calculator

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
)

func LakewoodGeoLocation() GeoLocation {
	return NewGeoLocation2("Lakewood, NJ", 40.0721087, -74.2400243, 15, timeutil.LoadLocationOrPanic("America/New_York"))
}

func SamoaGeoLocation() GeoLocation {
	return NewGeoLocation2("Apia, Samoa", -13.8599098, -171.8031745, 1858, timeutil.LoadLocationOrPanic("Pacific/Apia"))
}

func JerusalemGeoLocation() GeoLocation {
	return NewGeoLocation2("Jerusalem, Israel", 31.7781161, 35.233804, 740, timeutil.LoadLocationOrPanic("Asia/Jerusalem"))
}

func LosAngelesGeoLocation() GeoLocation {
	return NewGeoLocation2("Los Angeles, CA", 34.0201613, -118.6919095, 71, timeutil.LoadLocationOrPanic("America/Los_Angeles"))
}

func TokyoGeoLocation() GeoLocation {
	return NewGeoLocation2("Tokyo, Japan", 35.6733227, 139.6403486, 40, timeutil.LoadLocationOrPanic("Asia/Tokyo"))
}

/*
func ArcticNunavutGeoLocation() GeoLocation {
	return NewGeoLocation2("Fort Conger, NU Canada", 81.7449398, -64.7945858, 127, timeutil.LoadLocationOrPanic("America/Toronto"))
}
*/

func BasicTestGeoLocations() []GeoLocation {
	return []GeoLocation{
		LakewoodGeoLocation(),
		JerusalemGeoLocation(),
		LosAngelesGeoLocation(),
		TokyoGeoLocation(),
		// ArcticNunavutGeoLocation(),
		SamoaGeoLocation(),
	}
}

func HooperBayGeoLocation() GeoLocation {
	return NewGeoLocation2("Hooper Bay, Alaska", 61.520182, -166.1740437, 8, timeutil.LoadLocationOrPanic("America/Anchorage"))
}

func DaneborgGeoLocation() GeoLocation {
	return NewGeoLocation2("Daneborg, Greenland", 74.2999996, -20.2420877, 0, timeutil.LoadLocationOrPanic("America/Godthab"))
}
