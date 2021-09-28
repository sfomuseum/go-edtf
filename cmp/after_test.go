package cmp

import (
	"testing"
)

func TestIsAfter(t *testing.T) {

	dates := map[string]string{
		"2021-09-10": "2020-10-16",
		"203X":       "2020-10-16",
	}

	for this_d, that_d := range dates {

		is_after, err := IsAfter(this_d, that_d)

		if err != nil {
			t.Fatalf("Failed to determine whether this (%s) is after that (%s), %v", this_d, that_d, err)
		}

		if !is_after {
			t.Fatalf("This date (%s) is expected to be after that date (%s)", this_d, that_d)
		}
	}
}
