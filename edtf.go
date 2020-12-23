package edtf

import (
	"time"
)

type EDTFDate struct {
	Upper *DateRange
	Lower *DateRange
	Raw   string
	Level int
}

type DateRange struct {
	Upper *Date
	Lower *Date
}

type Date struct {
	// String      string
	Time        *time.Time
	Uncertain   bool
	Approximate bool
	Unknown     bool
	Open        bool
	Precision   string
}

/*

Eventually this should be generated from the components pieces
collected during parsing and compared against Raw but this will
do for now (20201223/thisisaaronland)

*/

func (d *EDTFDate) String() string {
	return d.Raw
}
