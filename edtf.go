package edtf

import (
	"time"
)

const UNCERTAIN string = "?"
const APPROXIMATE string = "~"
const UNCERTAIN_AND_APPROXIMATE string = "%"

type EDTFDate struct {
	Start *DateRange `json:"start"`
	End   *DateRange `json:"end"`
	EDTF  string     `json:"edtf"`
	Level int        `json:"level"`
	Label string     `json:"label"`
}

type DateRange struct {
	Lower *Date `json:"lower"`
	Upper *Date `json:"upper"`
}

type Date struct {
	EDTF        string     `json:"edtf"`
	Time        *time.Time `json:"time,omitempty"`
	Uncertain   bool       `json:"uncertain,omitempty"`
	Approximate bool       `json:"approximate,omitempty"`
	Unknown     bool       `json:"unknown,omitempty"`
	Unspecified bool       `json:"unspecified,omitempty"`
	Open        bool       `json:"open,omitempty"`
	Precision   string     `json:"precision,omitempty"`
}

/*

Eventually this should be generated from the components pieces
collected during parsing and compared against Raw but this will
do for now (20201223/thisisaaronland)

*/

func (d *EDTFDate) String() string {
	return d.EDTF
}
