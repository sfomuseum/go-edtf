package level1

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/re"
)

const LEVEL int = 1

const LETTER_PREFIXED_CALENDAR_YEAR string = "Letter-prefixed calendar year"
const SEASON string = "Seasons"
const QUALIFIED_DATE string = "Qualification of a date (complete)"
const UNSPECIFIED_DIGITS string = "Unspecified digit(s) from the right"
const EXTENDED_INTERVAL string = "Extended Interval"
const EXTENDED_INTERVAL_START string = "Extended Interval (Start)"
const EXTENDED_INTERVAL_END string = "Extended Interval (End)"
const NEGATIVE_CALENDAR_YEAR string = "Negative calendar year"

func IsLevel1(edtf_str string) bool {
	return re.Level1.MatchString(edtf_str)
}

func Matches(edtf_str string) (string, error) {

	if IsLetterPrefixedCalendarYear(edtf_str) {
		return LETTER_PREFIXED_CALENDAR_YEAR, nil
	}

	if IsSeason(edtf_str) {
		return SEASON, nil
	}

	if IsQualifiedDate(edtf_str) {
		return QUALIFIED_DATE, nil
	}

	if IsUnspecifiedDigits(edtf_str) {
		return UNSPECIFIED_DIGITS, nil
	}

	if IsNegativeCalendarYear(edtf_str) {
		return NEGATIVE_CALENDAR_YEAR, nil
	}

	if IsExtendedInterval(edtf_str) {

		if re.IntervalStart.MatchString(edtf_str) {
			return EXTENDED_INTERVAL_START, nil
		}

		if re.IntervalEnd.MatchString(edtf_str) {
			return EXTENDED_INTERVAL_END, nil
		}
	}

	return "", edtf.Invalid("Invalid Level 1 string", edtf_str)
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

	return nil, edtf.Invalid("Invalid or unsupported Level 1 EDTF string", edtf_str)
}
