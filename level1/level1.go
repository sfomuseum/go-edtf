package level1

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
)

const LEVEL int = 1

var Tests map[string][]string = map[string][]string{
	"prefixed_calendar_year": []string{
		// Go is incapable of parsing these dates
		// "Y170000002",
		// "Y-17000002",
		"Y1700",
		"Y-1200",
	},
	"season": []string{
		"2001-01",
		"2001-24",
		"Spring, 2002",
		"winter, 2002",
	},
	"qualified_date": []string{
		"1984?",
		"2004-06~",
		"2004-06-11%",
	},
	"unspecified_digits": []string{
		"201X",
		"20XX",
		"2004-XX",
		"1985-04-XX",
		"1985-XX-XX",
	},
	"extended_interval_start": []string{
		"../1985-04-12",
		"../1985-04",
		"../1985",
		"/1985-04-12",
		"/1985-04",
		"/1985",
	},
	"extended_interval_end": []string{
		"1985-04-12/..",
		"1985-04/..",
		"1985/..",
		"1985-04-12/",
		"1985-04/",
		"1985/",
	},
	"negative_calendar_year": []string{
		"-1985",
	},
}

func IsLevel1(edtf_str string) bool {
	return re.Level1.MatchString(edtf_str)
}

func ParseString(edtf_str string) (*edtf.EDTFDate, error) {

	if IsLetterPrefixedCalendarYear(edtf_str) {
		return ParseLetterPrefixedCalendarYear(edtf_str)
	}

	if IsSeason(edtf_str) {
		return ParseSeason(edtf_str)
	}

	if IsQualifiedDate(edtf_str) {
		return ParseQualifiedDate(edtf_str)
	}

	if IsUnspecifiedDigits(edtf_str) {
		return ParseUnspecifiedDigits(edtf_str)
	}

	if IsNegativeCalendarYear(edtf_str) {
		return ParseNegativeCalendarYear(edtf_str)
	}

	if IsExtendedInterval(edtf_str) {
		return ParseExtendedInterval(edtf_str)
	}

	return nil, errors.New("Invalid or unsupported Level 1 EDTF string")
}
