package dimension

import "math"

type Degrees float64
type Radians float64

func (d Degrees) ToRadians() Radians {
	return Radians(d * math.Pi / 180)
}

func (r Radians) ToDegrees() Degrees {
	return Degrees(r * 180 / math.Pi)
}

/*
Sin return Sin of the angle in Degrees
*/
func (d Degrees) Sin() float64 {
	return math.Sin(float64(d.ToRadians()))
}

/*
Cos Calculate cosine of the angle in Degrees,
return cosine of the angle in Degrees
*/
func (d Degrees) Cos() float64 {
	return math.Cos(float64(d.ToRadians()))
}

/*
Tan of the angle in Degrees
*/
func (d Degrees) Tan() float64 {
	return math.Tan(float64(d.ToRadians()))
}

/*
ACos return ACos of the angle in Degrees
x angle
*/
func ACos(x float64) Degrees {
	return Radians(math.Acos(x)).ToDegrees()
}

/*
ASin return ASin of the angle in Degrees
x angle
*/
func ASin(x float64) Degrees {
	return Radians(math.Asin(x)).ToDegrees()
}

/*
ATan return ATan of the angle in Degrees
x angle
*/
func ATan(x float64) Degrees {
	return Radians(math.Atan(x)).ToDegrees()
}

type Meters float64
type KM float64

func (m Meters) ToKM() KM {
	return KM(m / 1000)
}

func (km KM) ToMeters() Meters {
	return Meters(km * 1000)
}

/*
ArcMinutes minutes of a degree
minutes [minutes of arc]: https://en.wikipedia.org/wiki/Minute_of_arc#Cartography
*/
type ArcMinutes float64

func (am ArcMinutes) ToDegrees() Degrees {
	return Degrees(am)
}

/*
ArcSeconds minutes of a degree
seconds [seconds of arc]: https://en.wikipedia.org/wiki/Minute_of_arc#Cartography
*/
type ArcSeconds float64

func (am ArcSeconds) ToDegrees() Degrees {
	return Degrees(am)
}
