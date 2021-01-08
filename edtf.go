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

const MAX_YEARS int = 9999 // This is a Golang thing

type EDTFDate struct {
	Start   *DateRange `json:"start"`
	End     *DateRange `json:"end"`
	EDTF    string     `json:"edtf"`
	Level   int        `json:"level"`
	Feature string     `json:"feature"`
}

func (d *EDTFDate) Lower() (*time.Time, error) {

	ts := d.Start.Lower.Timestamp

	if ts == nil {
		return nil, NotSet()
	}

	return ts.Time(), nil
}

func (d *EDTFDate) Upper() (*time.Time, error) {

	ts := d.End.Upper.Timestamp

	if ts == nil {
		return nil, NotSet()
	}

	return ts.Time(), nil
}

/*

Eventually this should be generated from the components pieces
collected during parsing and compared against Raw but this will
do for now (20201223/thisisaaronland)

*/

func (d *EDTFDate) String() string {
	return d.EDTF
}
