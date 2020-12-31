package common

import (
	"testing"
	"time"
)

func TestFlipYear(t *testing.T) {

	valid := map[int]int{
		-1900: 1900,
		900:   -900,
	}

	for input, expected := range valid {

		output := FlipYear(input)

		if output != expected {
			t.Fatalf("Failed to flip year '%d'. Expected '%d' but got '%d'", input, expected, output)
		}
	}
}

func TestTimeToBCE(t *testing.T) {

	tm, err := time.Parse("2006-01-02", "1970-01-01")

	if err != nil {
		t.Fatal(err)
	}

	tm = TimeToBCE(tm)

	if tm.Year() != -1970 {
		t.Fatalf("Failed to convert time to BCE. Expected -1970 but got '%d'", tm.Year())
	}
}
