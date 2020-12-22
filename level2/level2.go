package level2

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"regexp"
)

var re_exponential_year *regexp.Regexp
var re_significant_digits *regexp.Regexp

func init() {

	re_exponential_year = regexp.MustCompile(`^(?i)Y(\-?\d+)E(\d+)$`)

	re_significant_digits = regexp.MustCompile(`(?:(\d{4})S(\d+)|Y(\d+)S(\d+)|Y(\d+)E(\d+)S(\d+))$`)
}

func ParseExponentialYear(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_exponential_year.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 exponential year string")
	}

	return nil, nil
}

func ParseSignificantDigits(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_significant_digits.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 significant digits string")
	}

	return nil, nil

}
