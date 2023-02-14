package test

import (
	"github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/jdt"
)

const (
	StandardMondayChaseirim  jdt.JYear = 5777
	StandardTuesdayKesidran  jdt.JYear = 5762
	StandardShabbosShelaimim jdt.JYear = 5770
	LeapMondayShelaimim      jdt.JYear = 5776
	LeapTuesdayKesidran      jdt.JYear = 5755
	LeapThursdayChaseirim    jdt.JYear = 5765
)

/*
var AllYearTypes = []jdt.JYear{
	StandardMondayChaseirim,
	StandardMondayShelaimim,
	StandardTuesdayKesidran,
	StandardThursdayKesidran,
	StandardThursdayShelaimim,
	StandardShabbosChaseirim,
	StandardShabbosShelaimim,
	LeapMondayChaseirim,
	LeapMondayShelaimim,
	LeapTuesdayKesidran,
	LeapThursdayChaseirim,
	LeapThursdayShelaimim,
	LeapShabbosChaseirim,
	LeapShabbosShelaimim,
}
*/

/*
def all_days_matching(year, matcher, in_israel=False, use_modern_holidays=False):
    calendar = JewishCalendar(year, 7, 1)
    calendar.in_israel = in_israel
    calendar.use_modern_holidays = use_modern_holidays
    collection = {}
    while calendar.jewish_year == year:
        sd = matcher(calendar)
        if sd:
            if sd not in collection:
                collection[sd] = []
            collection[sd] += [f'{calendar.jewish_month}-{calendar.jewish_day}']
        calendar.forward()
    return collection
*/
