package common

import (
// "testing"
)

/*
func TestDateRangeWithString(t *testing.T) {

	valid := []string{
		"1984-05-31",
		"2020",
		"1972-11",
		"1972-XX-31",
		"1972-12-15?",
		"~1972-12-15",
		"-1967",
		"2004-06-~20",
	}

	invalid := []string{
		"~2021?",
		// daterange_test.go:39: False positive parsing 2018-XX?-12, <nil>
		// "2018-XX?-12",
	}

	for _, ymd := range valid {

		_, err := DateRangeWithString(ymd)

		if err != nil {
			t.Fatalf("Failed to parse %s, %v", ymd, err)
		}
	}

	for _, ymd := range invalid {

		_, err := DateRangeWithString(ymd)

		if err == nil {
			t.Fatalf("False positive parsing %s, %v", ymd, err)
		}
	}

}

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
*/
