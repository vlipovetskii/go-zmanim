package calculator

import (
	"github.com/vlipovetskii/go-zmanim/zmanim/dimension"
)

const (
	/*
		GeometricZenith
		90 deg below the vertical. Used as a basis for most calculations since the location of the sun is 90 deg below
		the horizon at sunrise and sunset.
		Note: it is important to note that for sunrise and sunset the calculator.adjustZenith
		is required to account for the radius of the sun and refraction. The adjusted zenith should not
		be used for calculations above or below 90 deg since they are usually calculated as an offset to 90 deg.
	*/
	GeometricZenith dimension.Degrees = 90

	// CivilZenith Sun's zenith at civil twilight 96 deg.
	CivilZenith dimension.Degrees = 96

	// NauticalZenith Sun's zenith at nautical twilight 102 deg.
	NauticalZenith dimension.Degrees = 102

	// AstronomicalZenith Sun's zenith at astronomical twilight (108&deg;).
	AstronomicalZenith dimension.Degrees = 108
)
