package unix

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
	_ "log"
)

// DateSpan is a struct containing Unix timestamps for a range of (two) dates. Dates before 1970-01-01 are represented as negative values.
type DateSpan struct {
	// Start is the Unix timestamp for the starting date.
	Start int64
	// End is the Unix timestamp for the ending date.
	End int64
}

// DateRange is a struct containing inner and outer `DateSpan` instances for an EDTF date string.
type DateRange struct {
	// Outer is a `DateSpan` instance which matches the lower value of a starting date range and the upper value of an ending date range.
	Outer *DateSpan
	// Outer is a `DateSpan` instance which matches the upper value of a starting date range and the lower value of an ending date range.
	Inner *DateSpan
}

// DeriveRanges will parse 'edtf_str' and return a boolean flag signaling that it was possible to derive date ranges and, when possible, a
// `DateRange` instance containing Unix timestamps. For example some EDTF date strings like ".." (indicating an "open" or "ongoing" date)
// are valid EDTF but not suitable for deriving a date range.
func DeriveRanges(edtf_str string) (bool, *DateRange, error) {

	if !isValid(edtf_str) {
		return false, nil, nil
	}

	edtf_dt, err := parser.ParseString(edtf_str)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to parse '%s', %w", edtf_str, err)
	}

	start := edtf_dt.Start
	end := edtf_dt.End

	start_lower := start.Lower
	start_upper := start.Upper

	end_lower := end.Lower
	end_upper := end.Upper

	if start_lower == nil {
		return false, nil, nil
	}

	if start_upper == nil {
		return false, nil, nil
	}

	if end_lower == nil {
		return false, nil, nil
	}

	if end_upper == nil {
		return false, nil, nil
	}

	start_lower_ts := start_lower.Timestamp
	start_upper_ts := start_upper.Timestamp

	end_lower_ts := end_lower.Timestamp
	end_upper_ts := end_upper.Timestamp

	if start_lower_ts == nil {
		return false, nil, nil
	}

	if start_upper_ts == nil {
		return false, nil, nil
	}

	if end_lower_ts == nil {
		return false, nil, nil
	}

	if end_upper_ts == nil {
		return false, nil, nil
	}

	outer_start := start_lower_ts.Unix()
	outer_end := end_upper_ts.Unix()

	inner_start := start_upper_ts.Unix()
	inner_end := end_lower_ts.Unix()

	outer := &DateSpan{
		Start: outer_start,
		End:   outer_end,
	}

	inner := &DateSpan{
		Start: inner_start,
		End:   inner_end,
	}

	r := &DateRange{
		Outer: outer,
		Inner: inner,
	}

	return true, r, nil
}

func isValid(edtf_str string) bool {

	if edtf.IsOpen(edtf_str) {
		return false
	}

	if edtf.IsUnknown(edtf_str) {
		return false
	}

	if edtf.IsUnspecified(edtf_str) {
		return false
	}

	return true
}
