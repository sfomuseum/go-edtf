package edtf

import (
	"fmt"
	"strconv"
	"strings"
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
	Start   *DateRange `json:"start"`
	End     *DateRange `json:"end"`
	EDTF    string     `json:"edtf"`
	Level   int        `json:"level"`
	Feature string     `json:"feature"`
}

func (d *EDTFDate) Lower() *time.Time {

	ts := d.Start.Lower.Timestamp

	if ts == nil {

	}

	return ts.Time()
}

func (d *EDTFDate) Upper() *time.Time {

	ts := d.End.Upper.Timestamp

	if ts == nil {
		return nil
	}

	return ts.Time()
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
	DateTime    string     `json:"datetime,omitempty"`
	Timestamp   *Timestamp `json:"timestamp,omitempty"`
	YMD         *YMD       `json:"ymd"`
	Uncertain   Precision  `json:"uncertain,omitempty"`
	Approximate Precision  `json:"approximate,omitempty"`
	Unspecified Precision  `json:"unspecified,omitempty"`
	Precision   Precision  `json:"precision,omitempty"`
	Open        bool       `json:"open,omitempty"`
	Unknown     bool       `json:"unknown,omitempty"`
	Inclusivity Precision  `json:"inclusivity,omitempty"`
}

func (d *Date) SetTime(t *time.Time) {
	d.DateTime = t.Format(time.RFC3339)
	d.Timestamp = NewTimestampWithTime(t)
}

func (d *Date) String() string {
	return fmt.Sprintf("[[%T] Time: '%v' YMD: '%v']", d, d.Timestamp, d.YMD)
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

type Timestamp struct {
	timestamp int64
}

func NewTimestampWithTime(t *time.Time) *Timestamp {
	return &Timestamp{t.Unix()}
}

func (ts *Timestamp) Time() *time.Time {

	t := time.Unix(ts.Unix(), 0)
	return &t
}

func (ts *Timestamp) Unix() int64 {
	return ts.timestamp
}

func (ts *Timestamp) UnmarshalJSON(b []byte) error {

	s := strings.Trim(string(b), `"`)
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return err
	}

	*ts = Timestamp{i}
	return nil
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	str_ts := strconv.FormatInt(ts.timestamp, 10)
	return []byte(str_ts), nil
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
