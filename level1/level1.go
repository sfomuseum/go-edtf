package level1

import (
	"errors"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/re"
)

const LEVEL int = 1

const LETTER_PREFIXED_CALENDAR_YEAR string = "letter prefixed calendar year"
const SEASON string = "season"
const QUALIFIED_DATE string = "qualified date"
const UNSPECIFIED_DIGITS string = "unspecified digits"
const EXTENDED_INTERVAL string = "extended interval"
const EXTENDED_INTERVAL_START string = "extended interval start"
const EXTENDED_INTERVAL_END string = "extended interval end"
const NEGATIVE_CALENDAR_YEAR string = "negative calendar year"

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

	return "", errors.New("Invalid Level 1 string")
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
