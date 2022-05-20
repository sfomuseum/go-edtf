package unix

import (
	"github.com/sfomuseum/go-edtf"
	"testing"
)

func TestDeriveRanges(t *testing.T) {

	tests := map[string]*DateRange{
		"2021-05-25": &DateRange{Outer: &DateSpan{Start: 1621900800, End: 1621987199}, Inner: &DateSpan{Start: 1621900800, End: 1621987199}},
		"2021-11-09": &DateRange{Outer: &DateSpan{Start: 1636416000, End: 1636502399}, Inner: &DateSpan{Start: 1636416000, End: 1636502399}},
		edtf.UNKNOWN: nil,
		edtf.OPEN:    nil,
	}

	for edtf_str, expected_ranges := range tests {

		date_ranges, err := DeriveRanges(edtf_str)

		if err != nil {
			t.Fatalf("Failed to derive ranges for '%s', %v", edtf_str, err)
		}

		if date_ranges == nil {
			continue
		}

		if date_ranges.Inner.Start != expected_ranges.Inner.Start {
			t.Fatalf("Unexpected for inner.start timestamp for '%s'. Expected '%d' but got '%d'", edtf_str, expected_ranges.Inner.Start, date_ranges.Inner.Start)
		}

		if date_ranges.Inner.End != expected_ranges.Inner.End {
			t.Fatalf("Unexpected for inner.end timestamp for '%s'. Expected '%d' but got '%d'", edtf_str, expected_ranges.Inner.End, date_ranges.Inner.End)
		}

		if date_ranges.Outer.Start != expected_ranges.Outer.Start {
			t.Fatalf("Unexpected for outer.start timestamp for '%s'. Expected '%d' but got '%d'", edtf_str, expected_ranges.Outer.Start, date_ranges.Outer.Start)
		}

		if date_ranges.Outer.End != expected_ranges.Outer.End {
			t.Fatalf("Unexpected for outer.end timestamp for '%s'. Expected '%d' but got '%d'", edtf_str, expected_ranges.Outer.End, date_ranges.Outer.End)
		}

	}

}
