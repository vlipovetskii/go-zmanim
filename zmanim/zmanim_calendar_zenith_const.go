package zmanim

import "github.com/vlipovetskii/go-zmanim/zmanim/calculator"

const (
	/*
			zenith16Point1 is the zenith of 16.1 deg below geometric zenith 90 deg.
			This calculation is used for determining alos (dawn) and tzais (nightfall) in some opinions.
			It is based on the calculation that the time between dawn and sunrise (and sunset to nightfall) is 72 minutes,
			the time that is takes to walk 4 mil at 18 minutes
			a mil [Rambam]: https://en.wikipedia.org/wiki/Maimonides and others.
			The sun's position at 72 minutes before AstronomicalCalendar.Sunrise in Jerusalem
			[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/ is
			16.1 deg below calculator.GeometricZenith.
			see
		- ZmanimCalendar.AlosHashachar
			- ComplexZmanimCalendar.Alos16Point1Degrees
			- ComplexZmanimCalendar.Tzais16Point1Degrees
			- ComplexZmanimCalendar.SofZmanShmaMGA16Point1Degrees
			- ComplexZmanimCalendar.SofZmanTfilaMGA16Point1Degrees
			- ComplexZmanimCalendar.MinchaGedola16Point1Degrees
			- ComplexZmanimCalendar.MinchaKetana16Point1Degrees
			- ComplexZmanimCalendar.PlagHamincha16Point1Degrees
			- ComplexZmanimCalendar.PlagAlos16Point1ToTzaisGeonim7Point083Degrees
			- ComplexZmanimCalendar.SofZmanShmaAlos16Point1ToSunset
	*/
	zenith16Point1 = calculator.GeometricZenith + 16.1

	/*
		zenith8Point5 is the zenith of 8.5 deg below geometric zenith 90 deg.
		This calculation is used for calculating alos
		(dawn) and tzais (nightfall) in some opinions. This calculation is based on the position of the sun 36
		minutes after AstronomicalCalendar.Sunset in Jerusalem
		[around the equinox / equilux<]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/,
		which is 8.5 deg below calculator.GeometricZenith.
		The [Ohr Meir]: https://www.worldcat.org/oclc/29283612 considers this the time that 3 small stars are visible,
		which is later than the required 3 medium stars.
		see
		- ZmanimCalendar.Tzais
		- ComplexZmanimCalendar.TzaisGeonim8Point5Degrees
	*/
	zenith8Point5 = calculator.GeometricZenith + 8.5

	/*
		zenith3Point7 is the zenith of 3.7 deg below calculator.GeometricZenith 90 deg.
		This calculation is used for calculating tzais (nightfall) based on the opinion
		of the Geonim that tzais is the time it takes to walk 3/4 of a Mil at 18 minutes a Mil,
		or 13.5 minutes after sunset. The sun is 3.7 deg; below calculator.GeometricZenith at this time in Jerusalem
		[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/.
		TzaisGeonim3Point7Degrees
	*/
	zenith3Point7 = calculator.GeometricZenith + 3.7

	/*
		zenith3Point8 is the zenith of 3.8 deg; below calculator.GeometricZenith 90 deg.
		This calculation is used for calculating tzais (nightfall) based on the opinion of the Geonim> that tzais is the
		time it takes to walk 3/4 of a Mil at 18 minutes a Mil, or 13.5 minutes after sunset.
		The sun is 3.8 deg below calculator.GeometricZenith at this time in Jerusalem
		[around the equinox / equilux]: https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/.
		TzaisGeonim3Point8Degrees
	*/
	zenith3Point8 = calculator.GeometricZenith + 3.8

	/**
	 * The zenith of 5.95&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> (nightfall) according to some opinions. This calculation is based on the position of
	 * the sun 24 minutes after sunset in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a>,
	 * which calculates to 5.95&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getTzaisGeonim5Point95Degrees()
	 */
	zenith5Point95 = calculator.GeometricZenith + 5.95

	/**
	 * The zenith of 7.083&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This is often referred to as
	 * 7&deg;5' or 7&deg; and 5 minutes. This calculation is used for calculating <em>alos</em> (dawn) and
	 * <em>tzais</em> (nightfall) according to some opinions. This calculation is based on observation of 3 medium sized
	 * stars by Dr. Baruch Cohen in his calendar published in in 1899 in Strasbourg, France. This calculates to
	 * 7.0833333&deg; below {@link #GeometricZenith geometric zenith}. The <a href="https://hebrewbooks.org/1053">Sh"Ut
	 * Melamed Leho'il</a> in Orach Chaim 30 agreed to this <em>zman</em>, as did the Sh"Ut Bnei Tziyon, Tenuvas Sadeh and
	 * it is very close to the time of the <a href="https://hebrewbooks.org/22044">Mekor Chesed</a> of the Sefer chasidim.
	 * It is close to the position of the sun 30 minutes after sunset in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a>, but not
	 * Exactly. The actual position of the sun 30 minutes after sunset in Jerusalem at the equilux is 7.205&deg; and
	 * 7.199&deg; at the equinox. See Hazmanim Bahalacha vol 2, pages 520-521 for details.
	 *
	 * @see #getTzaisGeonim7Point083Degrees()
	 * @see #getBainHasmashosRT13Point5MinutesBefore7Point083Degrees()
	 */
	zenith7Point83 = calculator.GeometricZenith + 7 + (5.0 / 60)

	/**
	 * The zenith of 10.2&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>misheyakir</em> according to some opinions. This calculation is based on the position of the sun
	 * 45 minutes before {@link #getSunrise sunrise} in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a> which
	 * calculates to 10.2&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getMisheyakir10Point2Degrees()
	 */
	zenith10Point2 = calculator.GeometricZenith + 10.2

	/**
	 * The zenith of 11&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>misheyakir</em> according to some opinions. This calculation is based on the position of the sun
	 * 48 minutes before {@link #getSunrise sunrise} in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a>, which
	 * calculates to 11&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getMisheyakir11Degrees()
	 */
	zenith11Degrees = calculator.GeometricZenith + 11

	/**
	 * The zenith of 11.5&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>misheyakir</em> according to some opinions. This calculation is based on the position of the sun
	 * 52 minutes before {@link #getSunrise sunrise} in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a>, which
	 * calculates to 11.5&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getMisheyakir11Point5Degrees()
	 */
	zenith11Point5 = calculator.GeometricZenith + 11.5

	/**
	 * The zenith of 13.24&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating Rabbeinu Tam's <em>bain hashmashos</em> according to some opinions.
	 * NOTE: See comments on {@link #getBainHasmashosRT13Point24Degrees} for additional details about the degrees.
	 *
	 * @see #getBainHasmashosRT13Point24Degrees
	 *
	 */
	zenith13Point24 = calculator.GeometricZenith + 13.24

	/**
	 * The zenith of 19&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>alos</em> according to some opinions.
	 *
	 * @see #getAlos19Degrees()
	 * @see #ZENITH_19_POINT_8
	 */
	zenith19Degrees = calculator.GeometricZenith + 19

	/**
	 * The zenith of 19.8&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>alos</em> (dawn) and <em>tzais</em> (nightfall) according to some opinions. This calculation is
	 * based on the position of the sun 90 minutes after sunset in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a> which
	 * calculates to 19.8&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getTzais19Point8Degrees()
	 * @see #getAlos19Point8Degrees()
	 * @see #getAlos90()
	 * @see #getTzais90()
	 * @see #ZENITH_19_DEGREES
	 */
	zenith19Point8 = calculator.GeometricZenith + 19.8

	/**
	 * The zenith of 26&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>alos</em> (dawn) and <em>tzais</em> (nightfall) according to some opinions. This calculation is
	 * based on the position of the sun {@link #getAlos120() 120 minutes} after sunset in Jerusalem o<a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a> which
	 * calculates to 26&deg; below {@link #GeometricZenith geometric zenith}. Since the level of darkness when the sun is
	 * 26&deg; and at a point when the level of darkness is long past the 18&deg; point where the darkest point is reached,
	 * it should only be used <em>lechumra</em> such as delaying the start of nighttime <em>mitzvos</em> or avoiding eating
	 * this early on a fast day.
	 *
	 * @see #getAlos26Degrees()
	 * @see #getTzais26Degrees()
	 * @see #getAlos120()
	 * @see #getTzais120()
	 */
	zenith26Degrees = calculator.GeometricZenith + 26.0

	/**
	 * The zenith of 4.37&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> (nightfall) according to some opinions. This calculation is based on the position of
	 * the sun {@link #getTzaisGeonim4Point37Degrees() 16 7/8 minutes} after sunset (3/4 of a 22.5-minute <em>Mil</em>)
	 * in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a>,
	 * which calculates to 4.37&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getTzaisGeonim4Point37Degrees()
	 */
	zenith4Point37 = calculator.GeometricZenith + 4.37

	/**
	 * The zenith of 4.61&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> (nightfall) according to some opinions. This calculation is based on the position of
	 * the sun {@link #getTzaisGeonim4Point37Degrees() 18 minutes} after sunset (3/4 of a 24-minute <em>Mil</em>) in
	 * Jerusalem <a href="https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox
	 * / equilux</a>, which calculates to 4.61&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getTzaisGeonim4Point61Degrees()
	 */
	zenith4Point61 = calculator.GeometricZenith + 4.61

	/**
	 * The zenith of 4.8&deg; below {@link #GeometricZenith geometric zenith} (90&deg;).
	 * @see #getTzaisGeonim4Point8Degrees()
	 */
	zenith4Point8 = calculator.GeometricZenith + 4.8

	/**
	 * The zenith of 3.65&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> (nightfall) according to some opinions. This calculation is based on the position of
	 * the sun {@link #getTzaisGeonim3Point65Degrees() 13.5 minutes} after sunset (3/4 of an 18-minute <em>Mil</em>)
	 * in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a> which
	 * calculates to 3.65&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getTzaisGeonim3Point65Degrees()
	 */
	zenith3Point65 = calculator.GeometricZenith + 3.65

	/**
	 * The zenith of 3.676&deg; below {@link #GeometricZenith geometric zenith} (90&deg;).
	 */
	zenith3Point676 = calculator.GeometricZenith + 3.676

	/**
	 * The zenith of 5.88&deg; below {@link #GeometricZenith geometric zenith} (90&deg;).
	 */
	zenith5Point88 = calculator.GeometricZenith + 5.88

	/**
	 * The zenith of 1.583&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>netz amiti</em> (sunrise) and <em>shkiah amiti</em> (sunset) based on the opinion of the
	 * <a href="https://en.wikipedia.org/wiki/Shneur_Zalman_of_Liadi">Baal Hatanya</a>.
	 *
	 * @see #sunriseBaalHatanya()
	 * @see #sunsetBaalHatanya()
	 */
	zenith1Point583 = calculator.GeometricZenith + 1.583

	/**
	 * The zenith of 16.9&deg; below geometric zenith (90&deg;). This calculation is used for determining <em>alos</em>
	 * (dawn) based on the opinion of the Baal Hatanya. It is based on the calculation that the time between dawn
	 * and <em>netz amiti</em> (sunrise) is 72 minutes, the time that is takes to walk 4 <em>mil</em> at 18 minutes
	 * a <em>mil</em> (<a href="https://en.wikipedia.org/wiki/Maimonides">Rambam</a> and others). The sun's position at 72
	 * minutes before {@link #sunriseBaalHatanya <em>netz amiti</em> (sunrise)} in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a> is
	 * 16.9&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getAlosBaalHatanya()
	 */
	zenith16Point9 = calculator.GeometricZenith + 16.9

	/**
	 * The zenith of 6&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> / nightfall based on the opinion of the Baal Hatanya. This calculation is based on the
	 * position of the sun 24 minutes after {@link #getSunset sunset} in Jerusalem <a href=
	 * "https://kosherjava.com/2022/01/12/equinox-vs-equilux-zmanim-calculations/">around the equinox / equilux</a>, which
	 * is 6&deg; below {@link #GeometricZenith geometric zenith}.
	 *
	 * @see #getTzaisBaalHatanya()
	 */
	zenith6Degrees = calculator.GeometricZenith + 6

	/**
	 * The zenith of 6.45&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> (nightfall) according to some opinions. This is based on the calculations of <a href=
	 * "https://en.wikipedia.org/wiki/Yechiel_Michel_Tucazinsky">Rabbi Yechiel Michel Tucazinsky</a> of the position of
	 * the sun no later than {@link #getTzaisGeonim6Point45Degrees() 31 minutes} after sunset in Jerusalem, and at the
	 * height of the summer solstice, this <em>zman</em> is 28 minutes after <em>shkiah</em>. This computes to 6.45&deg;
	 * below {@link #GeometricZenith geometric zenith}. This calculation is found in the <a href=
	 * "https://hebrewbooks.org/pdfpager.aspx?req=50536&st=&pgnum=51">Birur Halacha Yoreh Deah 262</a> it the commonly
	 * used <em>zman</em> in Israel. It should be noted that this differs from the 6.1&deg;/6.2&deg; calculation for
	 * Rabbi Tucazinsky's time as calculated by the Hazmanim Bahalacha Vol II chapter 50:7 (page 515).
	 *
	 * @see #getTzaisGeonim6Point45Degrees()
	 */
	zenith6Point45 = calculator.GeometricZenith + 6.45

	/**
	 * The zenith of 7.65&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>misheyakir</em> according to some opinions.
	 *
	 * @see #getMisheyakir7Point65Degrees()
	 */
	zenith7Point65 = calculator.GeometricZenith + 7.65

	/**
	 * The zenith of 7.67&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> according to some opinions.
	 *
	 * @see #getTzaisGeonim7Point67Degrees()
	 */
	zenith7Point67 = calculator.GeometricZenith + 7.67

	/**
	 * The zenith of 9.3&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>tzais</em> (nightfall) according to some opinions.
	 *
	 * @see #getTzaisGeonim9Point3Degrees()
	 */
	zenith9Point3 = calculator.GeometricZenith + 9.3

	/**
	 * The zenith of 9.5&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>misheyakir</em> according to some opinions.
	 *
	 * @see #getMisheyakir9Point5Degrees()
	 */
	zenith9Point5 = calculator.GeometricZenith + 9.5

	/**
	 * The zenith of 9.75&deg; below {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating <em>alos</em> (dawn) and <em>tzais</em> (nightfall) according to some opinions.
	 *
	 * @see #getTzaisGeonim9Point75Degrees()
	 */
	zenith9Point75 = calculator.GeometricZenith + 9.75

	/**
	 * The zenith of 2.1&deg; above {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating the start of <em>bain hashmashos</em> (twilight) of 13.5 minutes before sunset converted to degrees
	 * according to the Yereim. As is traditional with degrees below the horizon, this is calculated without refraction
	 * and from the center of the sun. It would be 0.833&deg; less without this.
	 *
	 * @see #getBainHasmashosYereim2Point1Degrees()
	 */
	zenithMinus2Point1 = calculator.GeometricZenith - 2.1

	/**
	 * The zenith of 2.8&deg; above {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating the start of <em>bain hashmashos</em> (twilight) of 16.875 minutes before sunset converted to degrees
	 * according to the Yereim. As is traditional with degrees below the horizon, this is calculated without refraction
	 * and from the center of the sun. It would be 0.833&deg; less without this.
	 *
	 * @see #getBainHasmashosYereim2Point8Degrees()
	 */
	zenithMinus2Point8 = calculator.GeometricZenith - 2.8

	/**
	 * The zenith of 3.05&deg; above {@link #GeometricZenith geometric zenith} (90&deg;). This calculation is used for
	 * calculating the start of <em>bain hashmashos</em> (twilight) of 18 minutes before sunset converted to degrees
	 * according to the Yereim. As is traditional with degrees below the horizon, this is calculated without refraction
	 * and from the center of the sun. It would be 0.833&deg; less without this.
	 *
	 * @see #getBainHasmashosYereim3Point05Degrees()
	 */
	zenithMinus3Point05 = calculator.GeometricZenith - 3.05
)
