package level0

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
)

const LEVEL int = 0

var Tests map[string][]string = map[string][]string{
	"date": []string{
		"1985-04-12",
		"1985-04",
		"1985",
	},
	"date_and_time": []string{
		"1985-04-12T23:20:30",
		"1985-04-12T23:20:30Z",
		"1985-04-12T23:20:30-04",
		"1985-04-12T23:20:30+04:30",
	},
	"time_interval": []string{
		"1964/2008",
		"2004-06/2006-08",
		"2004-02-01/2005-02-08",
		"2004-02-01/2005-02",
		"2004-02-01/2005",
		"2005/2006-02",
	},
}

func IsLevel0(edtf_str string) bool {
	return re.Level0.MatchString(edtf_str)
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
