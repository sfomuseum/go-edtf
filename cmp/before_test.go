package cmp

import (
	"testing"
)

func TestIsBefore(t *testing.T) {

	dates := map[string]string{
		"2017-09-10": "2020-10-16",
		"201X":       "2020-10-16",
		"2020-09-XX": "2020-10-16",
	}

	for this_d, that_d := range dates {

		is_before, err := IsBefore(this_d, that_d)

		if err != nil {
			t.Fatalf("Failed to determine whether this (%s) is before that (%s), %v", this_d, that_d, err)
		}

		if !is_before {
			t.Fatalf("This date (%s) is expected to be before that date (%s)", this_d, that_d)
		}
	}
}
