package level1

import (
	"testing"
)

func TestExtendedIntervalEnd(t *testing.T) {

	valid, ok := Tests[EXTENDED_INTERVAL_END]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseExtendedIntervalEnd '%s'", input)

		d, err := ParseExtendedIntervalEnd(input)

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

func TestExtendedIntervalStart(t *testing.T) {

	valid, ok := Tests[EXTENDED_INTERVAL_START]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseExtendedIntervalStart '%s'", input)

		d, err := ParseExtendedIntervalStart(input)

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
