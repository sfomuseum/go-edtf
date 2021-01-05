package edtf

import (
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
	NONE Precision = 0
	ALL  Precision = 1 << iota
	ANY
	ANNUAL
	MONTHLY
	DAILY
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

type DateRange struct {
	EDTF  string `json:"edtf"`
	Lower *Date  `json:"lower"`
	Upper *Date  `json:"upper"`
}

type Date struct {
	Time *time.Time `json:"time,omitempty"`
	YMD  *YMD       `json:"ymd"`
	// HMS string `json:"hms,omitempty"`
	Uncertain   Precision `json:"uncertain,omitempty"`
	Approximate Precision `json:"approximate,omitempty"`
	Unspecified Precision `json:"unspecified,omitempty"`
	Open        bool      `json:"open,omitempty"`
	Unknown     bool      `json:"unknown,omitempty"`
	Inclusivity Precision `json:"inclusivity,omitempty"`
}

type YMD struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
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
