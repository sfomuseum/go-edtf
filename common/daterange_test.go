package common

import (
	"testing"
)

func TestDateRangeWithYMDString(t *testing.T) {

	valid := [][]string{
		[]string{"2020", "12", "28"},
		[]string{"2020", "02", ""}, // ensure 29 days
		[]string{"2019", "02", ""}, // ensure 28 days
		[]string{"1900", "", ""},
		[]string{"1450", "03", ""},
		[]string{"-1900", "05", ""},
		[]string{"-1972", "03", "19"},
	}

	for _, ymd := range valid {

		_, err := DateRangeWithYMDString(ymd[0], ymd[1], ymd[2])

		if err != nil {
			t.Fatalf("Failed to parse %s, %v", ymd, err)
		}
	}
}
