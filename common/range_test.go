package common

import (
	"testing"
)

func TestStringRangeWithString(t *testing.T) {

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

		t.Logf("StringRangeFromYMD '%s'", ymd)

		_, err := StringRangeFromYMD(ymd)

		if err != nil {
			t.Logf("Failed to parse %s, %v", ymd, err)
			t.Fail()
			continue
		}
	}

	for _, ymd := range invalid {

		t.Logf("StringRangeFromYMD (false) '%s'", ymd)

		_, err := StringRangeFromYMD(ymd)

		if err == nil {
			t.Logf("False positive parsing %s, %v", ymd, err)
			t.Fail()
			continue
		}
	}

}
