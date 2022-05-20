package unix

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
	_ "log"
)

type DateSpan struct {
	Start int64
	End   int64
}

type DateRange struct {
	Outer *DateSpan
	Inner *DateSpan
}

// DeriveRanges will parse 'edtf_str' and return a `DateRange` instance containing Unix timestamps.
func DeriveRanges(edtf_str string) (*DateRange, error) {

	if !isValid(edtf_str) {
		return nil, nil
	}

	edtf_dt, err := parser.ParseString(edtf_str)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse '%s', %w", edtf_str, err)
	}

	start := edtf_dt.Start
	end := edtf_dt.End

	start_lower := start.Lower
	start_upper := start.Upper

	end_lower := end.Lower
	end_upper := end.Upper

	if start_lower == nil {
		return nil, nil
	}

	if start_upper == nil {
		return nil, nil
	}

	if end_lower == nil {
		return nil, nil
	}

	if end_upper == nil {
		return nil, nil
	}

	start_lower_ts := start_lower.Timestamp
	start_upper_ts := start_upper.Timestamp

	end_lower_ts := end_lower.Timestamp
	end_upper_ts := end_upper.Timestamp

	if start_lower_ts == nil {
		return nil, nil
	}

	if start_upper_ts == nil {
		return nil, nil
	}

	if end_lower_ts == nil {
		return nil, nil
	}

	if end_upper_ts == nil {
		return nil, nil
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

	return r, nil
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

