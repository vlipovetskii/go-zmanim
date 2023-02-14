package calculator

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil"
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"
	"github.com/vlipovetskii/go-zmanim/helper"
	"github.com/vlipovetskii/go-zmanim/helper/assert"
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
	"testing"
)

func TestGMT(t *testing.T) {
	geoLocation := NewGeoLocation()

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, geoLocation.LocationName(), "Greenwich, England")
	assert.Equal(t, tag, geoLocation.Longitude(), float64(0))
	assert.Equal(t, tag, geoLocation.Latitude(), 51.4772)
	// self.assertTrue(gmt.time_zone._filename.endswith('/GMT'))
	assert.Equal(t, tag, geoLocation.Elevation(), dimension.Meters(0))

}

func TestLatitudeNumeric(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetLatitude1(33.3)

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, 33.3, geoLocation.Latitude())
}

func TestLatitudeCartographyNorth(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetLatitude2(41, 7, 5.17296, "N")

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, 41.1181036, geoLocation.Latitude())
}

func TestLatitudeCartographySouth(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetLatitude2(41, 7, 5.17296, "S")

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, -41.1181036, geoLocation.Latitude())
}

func TestLongitudeNumeric(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetLatitude1(23.4)

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, 23.4, geoLocation.Latitude())
}

func TestLongitudeCartographyEast(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetLongitude2(41, 7, 5.17296, "E")

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, 41.1181036, geoLocation.Longitude())
}

func TestLongitudeCartographyWest(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetLongitude2(41, 7, 5.17296, "W")

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, -41.1181036, geoLocation.Longitude())
}

func TestTimeZoneWithString(t *testing.T) {
	/*
	   geo = GeoLocation.GMT()
	   geo.time_zone = 'America/New_York'
	   self.assertTrue(geo.time_zone._filename.endswith('/America/New_York'))
	*/
}

func TestTimeZoneWithObject(t *testing.T) {
	/*
	   geo = GeoLocation.GMT()
	   geo.time_zone = tz.gettz('America/New_York')
	   self.assertTrue(geo.time_zone._filename.endswith('/America/New_York'))
	*/
}

func TestAntiMeridianAdjustmentForGMT(t *testing.T) {
	geoLocation := NewGeoLocation()

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, int32(0), geoLocation.AntimeridianAdjustment())
}

func TestAntiMeridianAdjustmentForStandardZone(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetTimeZone(timeutil.LoadLocationOrPanic("America/New_York"))

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, int32(0), geoLocation.AntimeridianAdjustment())
}

func TestAntiMeridianAdjustmentForEastwardCrossover(t *testing.T) {
	geoLocation := SamoaGeoLocation()

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, int32(-1), geoLocation.AntimeridianAdjustment())
}

func TestAntiMeridianAdjustmentForWestwardCrossover(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetLongitude1(179)
	geoLocation.SetTimeZone(timeutil.LoadLocationOrPanic("Etc/GMT+12"))

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, int32(1), geoLocation.AntimeridianAdjustment())
}

func TestLocalMeanTimeOffsetForGMT(t *testing.T) {
	geoLocation := NewGeoLocation()

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, gdt.GMillisecond(0), geoLocation.LocalMeanTimeOffset())
}

func TestLocalMeanTimeOffsetOnCenterMeridian(t *testing.T) {
	geoLocation := NewGeoLocation1("Sample", 40, -75, timeutil.LoadLocationOrPanic("America/New_York"))

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, gdt.GMillisecond(0), geoLocation.LocalMeanTimeOffset())
}

func TestLocalMeanTimeOffsetEastOfCenterMeridian(t *testing.T) {
	geoLocation := NewGeoLocation1("Sample", 40, -74, timeutil.LoadLocationOrPanic("America/New_York"))

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, gdt.GMillisecond(1*4*timeutil.MinuteMillis), geoLocation.LocalMeanTimeOffset())
}

func TestLocalMeanTimeOffsetWestOfCenterMeridian(t *testing.T) {
	geoLocation := NewGeoLocation1("Sample", 40, -76.25, timeutil.LoadLocationOrPanic("America/New_York"))

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, gdt.GMillisecond(-1.25*4*timeutil.MinuteMillis), geoLocation.LocalMeanTimeOffset())
}

func TestStandardTimeOffsetForGMT(t *testing.T) {
	geoLocation := NewGeoLocation()

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, gdt.GMillisecond(0), geoLocation.StandardTimeOffset())
}

func TestStandardTimeOffsetForStandardTimezone(t *testing.T) {
	geoLocation := NewGeoLocation()
	geoLocation.SetTimeZone(timeutil.LoadLocationOrPanic("America/New_York"))

	tag := helper.CurrentFuncName()

	assert.Equal(t, tag, -gdt.GHour(5).ToMilliseconds(), geoLocation.StandardTimeOffset())
}

func TestStandardTimeOffsetAt(t *testing.T) {
	/*
	   def test_time_zone_offset_at(self):
	       expected = [('2017-03-12T06:30:00Z', 'US/Eastern', -5),
	                   ('2017-03-12T07:00:00Z', 'US/Eastern', -4),
	                   ('2017-03-12T09:30:00Z', 'US/Pacific', -8),
	                   ('2017-03-12T10:00:00Z', 'US/Pacific', -7),
	                   ('2017-03-23T23:30:00Z', 'Asia/Jerusalem', 2),
	                   ('2017-03-24T00:00:00Z', 'Asia/Jerusalem', 3)]

	       def test_entry(time, tz):
	           geo = GeoLocation('Sample', 0, 0, tz)
	           return time, tz, geo.time_zone_offset_at(parser.parse(time))

	       for entry in expected:
	           self.assertEqual(test_entry(entry[0], entry[1]), entry)

	*/
}
