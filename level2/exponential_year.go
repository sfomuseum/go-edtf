package level2

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"github.com/whosonfirst/go-edtf/common"
	"math/big"
)

/*

Exponential year

'Y' at the beginning of the string (which indicates "year", as in level 1) may be followed by an integer, followed by 'E' followed by a positive integer. This signifies "times 10 to the power of". Thus 17E8 means "17 times 10 to the eighth power".

    Example        ‘Y-17E7’
    the calendar year -17*10 to the seventh power= -170000000

*/

func IsExponentialYear(edtf_str string) bool {
	return re_exponential_year.MatchString(edtf_str)
}

func ParseExponentialYear(edtf_str string) (*edtf.EDTFDate, error) {

	/*
		EXP 5 Y-17E7,-17E7,-,17,7
		EXP 5 Y10E7,10E7,,10,7
	*/

	m := re_exponential_year.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, errors.New("Invalid Level 2 exponential year string")
	}

	notation := m[1]
	prefix := m[2]

	flt, _, err := big.ParseFloat(notation, 10, 0, big.ToNearestEven)

	if err != nil {
		return nil, err
	}

	var i = new(big.Int)
	yyyy, _ := flt.Int(i)

	if yyyy.Int64() > int64(9999) || yyyy.Int64() < 0 {
		return nil, errors.New("Unsupported level 2 exponential year string")
	}

	yyyy_i := int(yyyy.Int64())

	start, err := common.DateRangeWithYMD(yyyy_i, 0, 0)

	if err != nil {
		return nil, err
	}

	if prefix == edtf.NEGATIVE {
		start.Lower.Time = calendar.ToBCE(start.Lower.Time)
		start.Upper.Time = calendar.ToBCE(start.Upper.Time)
		start.Lower.BCE = true
		start.Upper.BCE = true
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
