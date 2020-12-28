package level0

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {

	valid, ok := Tests["date"]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for _, str := range valid {

		d, err := ParseDate(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
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
		t.Fatalf("Invalid time string. Expected '%s' but got '%s'", expected_str, tm_str)
	}

}
