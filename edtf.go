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

const NONE int = 0
const ALL int = 1
const ANY int = 2

const ANNUAL int = 1
const MONTHLY int = 2
const DAILY int = 3

const MAX_YEARS int = 9999 // This is a Golang thing

type EDTFDate struct {
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

type DateRange struct {
	Lower *Date `json:"lower"`
	Upper *Date `json:"upper"`
}

type Date struct {
	EDTF        string     `json:"edtf"`
	Time        *time.Time `json:"time,omitempty"`
	Uncertain   int        `json:"uncertain,omitempty"`
	Approximate int        `json:"approximate,omitempty"`
	Unspecified int        `json:"unspecified,omitempty"`
	Open        bool       `json:"open,omitempty"`
	Unknown     bool       `json:"unknown,omitempty"`
	Inclusivity int        `json:"inclusivity,omitempty"`
}

/*

Eventually this should be generated from the components pieces
collected during parsing and compared against Raw but this will
do for now (20201223/thisisaaronland)

*/

func (d *EDTFDate) String() string {
	return d.EDTF
}
