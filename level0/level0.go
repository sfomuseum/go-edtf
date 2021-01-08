package level0

import (
	"errors"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/re"
)

const LEVEL int = 0

const DATE string = "Date"
const DATE_AND_TIME string = "Date and Time"
const TIME_INTERVAL string = "Time Interval"

func IsLevel0(edtf_str string) bool {
	return re.Level0.MatchString(edtf_str)
}

func Matches(edtf_str string) (string, error) {

	if IsDate(edtf_str) {
		return DATE, nil
	}

	if IsDateAndTime(edtf_str) {
		return DATE_AND_TIME, nil
	}

	if IsTimeInterval(edtf_str) {
		return TIME_INTERVAL, nil
	}

	return "", errors.New("Invalid Level 0 string")
}

func ParseString(edtf_str string) (*edtf.EDTFDate, error) {

	if IsDate(edtf_str) {
		return ParseDate(edtf_str)
	}

	if IsDateAndTime(edtf_str) {
		return ParseDateAndTime(edtf_str)
	}

	if IsTimeInterval(edtf_str) {
		return ParseTimeInterval(edtf_str)
	}

	return nil, errors.New("Invalid Level 0 string")
}
