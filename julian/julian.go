// Package julian provides types and functions for interacting with the Julian
// calendar.
package julian

import "time"

type (
	// Day models a day in the Julian calendar.
	Day float64
	// Century models a century in the Julian calendar since J2000.
	Century float64
)

// ConvertTimeToJulianDay converts a time.Time object to a julian date
func ConvertTimeToJulianDay(t time.Time) Day {
	y, m, d, hh, mm, ss, ms := t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1e6

	// Calc integer part (days)
	jday := (1461*(y+4800+(m-14)/12))/4 + (367*(m-2-12*((m-14)/12)))/12 - (3*((y+4900+(m-14)/12)/100))/4 + d - 32075

	// Calc floating point part (fraction of a day)
	jdatetime := float64(jday) + (float64(hh)-12.0)/24.0 + (float64(mm) / 1440.0) + (float64(ss) / 86400.0) + (float64(ms) / 86400000.0)

	// Adjust to UT
	_, zoneOffset := t.Zone()

	return Day(jdatetime + float64(zoneOffset)/86400)
}

// Add adds f to d.
func (d Day) Add(f float64) Day {
	return Day(d.Days() + f)
}

// Days returns the raw float64 value corresponding to d.
func (d Day) Days() float64 {
	return float64(d)
}

// ToCentury converts a Julian day to centuries since J2000.0
func (d Day) ToCentury() Century {
	return Century((d.Days() - 2451545.0) / 36525.0)
}

// Centuries returns the raw float64 value corresponding to c.
func (c Century) Centuries() float64 {
	return float64(c)
}

// ToDay converts centuries since J2000.0 to a Julian day.
func (c Century) ToDay() Day {
	return Day(c.Centuries()*36525.0 + 2451545.0)
}
