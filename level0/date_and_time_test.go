package level0

import (
	"testing"
)

func TestParseDateAndTime(t *testing.T) {

	valid, ok := Tests[DATE_AND_TIME]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for input, _ := range valid {

		d, err := ParseDateAndTime(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}

		if d.String() != input {
			t.Fatalf("Failed to stringify '%s', %v", input, err)
		}
	}
}
