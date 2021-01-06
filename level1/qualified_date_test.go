package level1

import (
	"testing"
)

func TestQualifiedDate(t *testing.T) {

	valid, ok := Tests[QUALIFIED_DATE]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, tr := range valid {

		t.Logf("ParseQualifiedDate '%s'", input)

		d, err := ParseQualifiedDate(input)

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
