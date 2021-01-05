package level0

import (
	"testing"
)

func TestParseDateAndTime(t *testing.T) {

	valid, ok := Tests[DATE_AND_TIME]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseDateAndTime(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
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
