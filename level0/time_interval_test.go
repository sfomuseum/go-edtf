package level0

import (
	"testing"
)

func TestParseTimeInterval(t *testing.T) {

	valid, ok := Tests["time_interval"]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for _, str := range valid {

		d, err := ParseTimeInterval(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
		}
	}
}
