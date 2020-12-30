package level2

import (
	"testing"	
)

func TestSubYearGroupings(t *testing.T) {

	valid, ok := Tests["sub_year_groupings"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseSubYearGroupings(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

