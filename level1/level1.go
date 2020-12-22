package level1

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"regexp"
)

var re_calendaryear *regexp.Regexp
var re_season *regexp.Regexp
var re_qualified *regexp.Regexp
var re_unspecified *regexp.Regexp

func init() {

	re_calendaryear = regexp.MustCompile(`^Y(\-)?(\d+)$`)

	re_season = regexp.MustCompile(`^(\d{4})-(0[1-9]|2[1-4])|(?i)(spring|summer|fall|winter)\s*,\s*(\d{4})$`)

	re_qualified = regexp.MustCompile(`^(?:(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?)(\?|~|%)$`)

	re_unspecified = regexp.MustCompile(`^(?:(\d{3})(X)|(\d{2})(XX)|(\d{4})-(XX)|(\d{4})\-(\d{2})\-(XX)|(\d{4})\-(XX)\-(XX))$`)
}

func ParseLetterPrefixedCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_calendaryear.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 calendar year string")
	}

	return nil, nil
}

func ParseSeason(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_season.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 season string")
	}

	return nil, nil
}

func ParseQualifiedDate(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_qualified.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 qualified date string")
	}

	return nil, nil
}

func ParseUnspecifiedDigits(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_unspecified.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 1 unspecified digits string")
	}

	return nil, nil
}

func ParseExtendedInterval(edtf_str string) (*edtf.EDTFDate, error) {

	return nil, nil
}

func ParseNegativeCalendarYear(edtf_str string) (*edtf.EDTFDate, error) {

	return nil, nil
}