package level0

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {

	valid, ok := Tests[DATE]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for input, tr := range valid {

		d, err := ParseDate(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}

		if tr != nil {

			err := tr.TestDate(d)

			if err != nil {
				t.Fatalf("Results failed tests '%s', %v", input, err)
			}
		}
	}

	dt, err := ParseDate("2020-02")

	if err != nil {
		t.Fatal("Failed to parse '2020-02'")
	}

	tm := dt.End.Upper.Time

	tm_str := tm.Format(time.RFC3339)
	expected_str := "2020-02-29T23:59:59Z"

	if tm_str != expected_str {
		t.Fatalf("Invalid time inputing. Expected '%s' but got '%s'", expected_str, tm_str)
	}

}
