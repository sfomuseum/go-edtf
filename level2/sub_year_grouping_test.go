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

	for _, str := range valid {

		_, err := ParseSubYearGroupings(str)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}
