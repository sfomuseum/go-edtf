package level2

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
)

/*

Exponential year

'Y' at the beginning of the string (which indicates "year", as in level 1) may be followed by an integer, followed by 'E' followed by a positive integer. This signifies "times 10 to the power of". Thus 17E8 means "17 times 10 to the eighth power".

    Example        ‘Y-17E7’
    the calendar year -17*10 to the seventh power= -170000000

*/

func IsExponentialYear(edtf_str string) bool {
	return re.ExponentialYear.MatchString(edtf_str)
}

func ParseExponentialYear(edtf_str string) (*edtf.EDTFDate, error) {

	/*
		EXP 5 Y-17E7,-17E7,-,17,7
		EXP 5 Y10E7,10E7,,10,7
	*/

	m := re.ExponentialYear.FindStringSubmatch(edtf_str)

	if len(m) != 4 {
		return nil, edtf.Invalid(EXPONENTIAL_YEAR, edtf_str)
	}

	notation := m[1]

	yyyy_i, err := common.ParseExponentialNotation(notation)

	if err != nil {
		return nil, err
	}

	start, err := common.DateRangeWithYMD(yyyy_i, 0, 0)

	if err != nil {
		return nil, err
	}

	end := start

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
