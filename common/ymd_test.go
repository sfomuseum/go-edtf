package common

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestYMDFromString(t *testing.T) {

	valid := map[string]*edtf.YMD{
		"2020-02":    &edtf.YMD{Year: 2020, Month: 2, Day: 29},
		"1972-06":    &edtf.YMD{Year: 1972, Month: 6, Day: 30},
		"1972":       &edtf.YMD{Year: 1972, Month: 1, Day: 31}, // should be 1 ?
		"-200-05-30": &edtf.YMD{Year: -200, Month: 5, Day: 30},
	}

	for input, expected := range valid {

		output, err := YMDFromString(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if !output.Equals(expected) {
			t.Logf("Unexpected results for '%s'. Expected '%s' but got '%s'", input, expected, output)
			t.Fail()
			continue
		}
	}
}
