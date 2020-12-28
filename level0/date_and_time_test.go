package level0

import (
	"testing"
)

func TestParseDateAndTime(t *testing.T) {

	valid, ok := Tests["date_and_time"]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for _, str := range valid {

		d, err := ParseDateAndTime(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
		}
	}
}
