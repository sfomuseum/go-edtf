package level1

import (
	"testing"
)

func TestSeason(t *testing.T) {

	valid, ok := Tests[SEASON]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseSeason(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}
