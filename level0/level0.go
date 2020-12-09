package level0

import (
	"errors"
	"regexp"
	"github.com/whosonfirst/go-edtf"
)

var re_date *regexp.Regexp
var re_date_time *regexp.Regexp
var re_time_interval *regexp.Regexp

func init() {

	re_date = regexp.MustCompile(`^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`)

	re_date_time = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})(Z|(\+|-)(\d{2})(\:(\d{2}))?)?$`)

	re_time_interval = regexp.MustCompile(`^(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?\/(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?$`)
}

func ParseDate(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_date.MatchString(edtf_str){
		return nil, errors.New("Invalid Level 0 date string")
	}

	return nil, nil
}

func ParseDateTime(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_date_time.MatchString(edtf_str){
		return nil, errors.New("Invalid Level 0 date and time string")
	}

	return nil, nil	
}

func ParseTimeInterval(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_time_interval.MatchString(edtf_str){
		return nil, errors.New("Invalid Level 0 time interval string")
	}

	return nil, nil	
}
