package edtf

import (
	"fmt"
	"time"
)

const UNCERTAIN string = "?"
const APPROXIMATE string = "~"
const UNCERTAIN_AND_APPROXIMATE string = "%"
const OPEN string = ".."
const UNKNOWN string = ""
const NEGATIVE string = "-"

const HMS_LOWER string = "00:00:00"
const HMS_UPPER string = "23:59:59"

const (
	NONE      Precision = 0
	ALL       Precision = 1 << iota // 2
	ANY                             // 4
	DAY                             // 8
	WEEK                            // 16
	MONTH                           // 32
	YEAR                            // 64
	DECADE                          // 128
	CENTURY                         // 256
	MILLENIUM                       // 512
)

const MAX_YEARS int = 9999 // This is a Golang thing

type EDTFDate struct {
	// Span *DateSpan `json:"span"`
	Start *DateRange `json:"start"`
	End   *DateRange `json:"end"`
	EDTF  string     `json:"edtf"`
	Level int        `json:"level"`
	Label string     `json:"label"`
}

func (d *EDTFDate) Lower() *time.Time {
	return d.Start.Lower.Time
}

func (d *EDTFDate) Upper() *time.Time {
	return d.End.Upper.Time
}

type DateSpan struct {
	Start *DateRange `json:"start"`
	End   *DateRange `json:"end"`
}

func (s *DateSpan) String() string {
	return fmt.Sprintf("[[%T] Start: '%v' End: '%v']", s, s.Start, s.End)
}

type DateRange struct {
	EDTF  string `json:"edtf"`
	Lower *Date  `json:"lower"`
	Upper *Date  `json:"upper"`
}

func (r *DateRange) String() string {
	return fmt.Sprintf("[[%T] Lower: '%v' Upper: '%v'[", r, r.Lower, r.Upper)
}

type Date struct {
	Time *time.Time `json:"time,omitempty"`
	YMD  *YMD       `json:"ymd"`
	// HMS string `json:"hms,omitempty"`
	Uncertain   Precision `json:"uncertain,omitempty"`
	Approximate Precision `json:"approximate,omitempty"`
	Unspecified Precision `json:"unspecified,omitempty"`
	Precision   Precision `json:"precision,omitempty"`
	Open        bool      `json:"open,omitempty"`
	Unknown     bool      `json:"unknown,omitempty"`
	Inclusivity Precision `json:"inclusivity,omitempty"`
}

func (d *Date) String() string {
	return fmt.Sprintf("[[%T] Time: '%v' YMD: '%v']", d, d.Time, d.YMD)
}

type YMD struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (ymd *YMD) String() string {
	return fmt.Sprintf("[%T] Y: '%d' M: '%d' D: '%d'", ymd, ymd.Year, ymd.Month, ymd.Day)
}

func (ymd *YMD) Equals(other_ymd *YMD) bool {

	if ymd.Year != other_ymd.Year {
		return false
	}

	if ymd.Month != other_ymd.Month {
		return false
	}

	if ymd.Day != other_ymd.Day {
		return false
	}

	return true
}

/*

Eventually this should be generated from the components pieces
collected during parsing and compared against Raw but this will
do for now (20201223/thisisaaronland)

*/

func (d *EDTFDate) String() string {
	return d.EDTF
}

// https://stackoverflow.com/questions/48050522/using-bitsets-in-golang-to-represent-capabilities

type Precision uint32

func (f Precision) HasFlag(flag Precision) bool { return f&flag != 0 }
func (f *Precision) AddFlag(flag Precision)     { *f |= flag }
func (f *Precision) ClearFlag(flag Precision)   { *f &= ^flag }
func (f *Precision) ToggleFlag(flag Precision)  { *f ^= flag }

func (f *Precision) IsAnnual() bool {
	return f.HasFlag(YEAR)
}

func (f *Precision) IsMonthly() bool {
	return f.HasFlag(MONTH)
}

func (f *Precision) IsDaily() bool {
	return f.HasFlag(DAY)
}
