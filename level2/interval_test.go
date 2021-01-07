package level2

import (
	"github.com/sfomuseum/go-edtf"
	"testing"
)

func TestInterval(t *testing.T) {

	valid, ok := Tests[INTERVAL]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseInterval '%s'", input)

		d, err := ParseInterval(input)

		if err != nil {

			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", input, err)
				continue
			} else {
				t.Logf("Failed to parse '%s', %v", input, err)
				t.Fail()
				continue
			}
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Logf("Results failed tests '%s', %v", input, err)
				t.Fail()
				continue
			}
		}

	}
}
