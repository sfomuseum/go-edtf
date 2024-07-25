package cmp

import (
	"testing"
)

func TestIsBetween(t *testing.T) {

	tests_between := [][3]string{
		[3]string{"2024-03-21", "2022-12", "2024-06-17"},
		[3]string{"2024-03-21", "2024~", "2024-06-17"},
		[3]string{"2024-03-21", "2024~", "2024~"},
	}

	tests_before := [][3]string{
		[3]string{"2021-03-21", "2022-12", "2024-06-17"},
		[3]string{"2024-03-21", "2024~", "2024-02~"},
	}

	tests_after := [][3]string{
		[3]string{"2021-03-21", "1984-12", "1996-12-08"},
	}

	for _, dates := range tests_between {

		d := dates[0]
		inception := dates[1]
		cessation := dates[2]

		is_between, err := IsBetween(d, inception, cessation)

		if err != nil {
			t.Fatalf("Failed to determine if '%s' is between '%s' and '%s', %v", d, inception, cessation, err)
		}

		if !is_between {
			t.Fatalf("Expected '%s' to be between '%s' and '%s'", d, inception, cessation)
		}
	}

	for _, dates := range tests_before {

		d := dates[0]
		inception := dates[1]
		cessation := dates[2]

		is_between, err := IsBetween(d, inception, cessation)

		if err != nil {
			t.Fatalf("Failed to determine if '%s' is between '%s' and '%s', %v", d, inception, cessation, err)
		}

		if is_between {
			t.Fatalf("Expected '%s' to be before '%s' (and '%s')", d, inception, cessation)
		}
	}

	for _, dates := range tests_after {

		d := dates[0]
		inception := dates[1]
		cessation := dates[2]

		is_between, err := IsBetween(d, inception, cessation)

		if err != nil {
			t.Fatalf("Failed to determine if '%s' is between '%s' and '%s', %v", d, inception, cessation, err)
		}

		if is_between {
			t.Fatalf("Expected '%s' to be after '%s' (and '%s')", d, cessation, inception)
		}
	}
}
