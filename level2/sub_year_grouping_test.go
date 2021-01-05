package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestSubYearGroupings(t *testing.T) {

	valid, ok := Tests[SUB_YEAR_GROUPINGS]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseSubYearGroupings(input)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", input, err)
				continue
			} else {
				t.Fatalf("Failed to parse '%s', %v", input, err)
			}
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Fatalf("Results failed tests '%s', %v", input, err)
			}
		}

	}
}
