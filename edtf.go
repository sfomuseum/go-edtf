package edtf

import (
	"fmt"
	"time"
)

const UNCERTAIN string = "?"
const APPROXIMATE string = "~"
const UNCERTAIN_AND_APPROXIMATE string = "%"
const OPEN string = ".."
const OPEN_2012 string = "open"
const UNSPECIFIED string = ""
const UNSPECIFIED_2012 string = "uuuu"
const UNKNOWN string = UNSPECIFIED // this code was incorrectly referring to "unspecified" as "unknown"
const UNKNOWN_2012 string = UNSPECIFIED_2012

const NEGATIVE string = "-"

const HMS_LOWER string = "00:00:00"
const HMS_UPPER string = "23:59:59"

const MAX_YEARS int = 9999 // This is a Golang thing

// Return a boolean value indicating whether a string is considered to be an "open" EDTF date.
func IsOpen(s string) bool {

	switch s {
	case OPEN, OPEN_2012:
		return true
	default:
		return false
	}
}

// Return a boolean value indicating whether a string is considered to be an "unspecified" EDTF date.
func IsUnspecified(s string) bool {

	switch s {
	case UNSPECIFIED, UNSPECIFIED_2012:
		return true
	default:
		return false
	}
}

// Return a boolean value indicating whether a string is considered to be an "unknown" EDTF date.
func IsUnknown(s string) bool {

	switch s {
	case UNKNOWN, UNKNOWN_2012:
		return true
	default:
		return false
	}
}

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

// After reports whether the EDTFDate instance `d` is after `u`.
func (d *EDTFDate) After(u *EDTFDate) (bool, error) {

	if IsOpen(d.EDTF) {
		return false, nil
	}

	u_t, err := u.Upper()

	if err != nil {
		return false, fmt.Errorf("Failed to derive upper time for inception date (%s), %w", u.EDTF, err)
	}

	t, err := d.Lower()

	if err != nil {
		return false, fmt.Errorf("Failed to derive lower time for cessation date (%s), %w", d.EDTF, err)
	}

	if u_t.After(*t) {
		return false, nil
	}

	return true, nil
}

// Before reports whether the EDTFDate instance `d` is after `u`.
func (d *EDTFDate) Before(u *EDTFDate) (bool, error) {

	if IsOpen(d.EDTF) {
		return false, nil
	}

	u_t, err := u.Lower()

	if err != nil {
		return false, fmt.Errorf("Failed to derive lower time for inception date (%s), %w", u.EDTF, err)
	}

	t, err := d.Upper()

	if err != nil {
		return false, fmt.Errorf("Failed to derive upper time for cessation date (%s), %w", d.EDTF, err)
	}

	if u_t.Before(*t) {
		return false, nil
	}

	return true, nil
}
