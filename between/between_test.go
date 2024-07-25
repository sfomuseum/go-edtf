package between

import (
	"testing"
)

func TestIsBetweenString(t *testing.T) {

	tests_between := [][3]string{
		[3]string{ "2024-03-21", "2022-12", "2024-06-17" },
	}

	for _, dates := range tests_between {

		d := dates[0]
		inception := dates[1]
		cessation := dates[2]

		is_between, err := IsBetweenString(d, inception, cessation)

		if err != nil {
			t.Fatalf("Failed to determine if '%s' is between '%s' and '%s', %v", d, inception, cessation, err)
		}

		if !is_between {
			t.Fatalf("Expected '%s' to be between '%s' and '%s'", d, inception, cessation)
		}
	}
		
}
