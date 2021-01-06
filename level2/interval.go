package level2

import (
	// "fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
	//"strings"
)

/*

For Level 2 portions of a date within an interval may be designated as approximate, uncertain, or unspecified.

    Example 1                 ‘2004-06-~01/2004-06-~20’
    An interval in June 2004 beginning approximately the first and ending approximately the 20th
    Example 2                 ‘2004-06-XX/2004-07-03’
    An interval beginning on an unspecified day in June 2004 and ending July 3.


*/

func IsInterval(edtf_str string) bool {
	return re.Interval.MatchString(edtf_str)
}

func ParseInterval(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		INTERVAL 2004-06-~01/2004-06-~20 13 2004-06-~01/2004-06-~20,,2004,,06,~,01,,2004,,06,~,20
		INTERVAL 2004-06-XX/2004-07-03 13 2004-06-XX/2004-07-03,,2004,,06,,XX,,2004,,07,,03

	*/

	if !re.Interval.MatchString(edtf_str) {
		return nil, edtf.Invalid(INTERVAL, edtf_str)
	}

	sp, err := common.DateSpanFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start: sp.Start,
		End:   sp.End,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil

	/*
		// m := re.Interval.FindStringSubmatch(edtf_str)

		if !re.Interval.MatchString(edtf_str) {
			return nil, edtf.Invalid(INTERVAL, edtf_str)
		}

		parts := strings.Split(edtf_str, "/")

		start, err := common.DateRangeWithString(parts[0])

		if err != nil {
			return nil, err
		}

		end, err := common.DateRangeWithString(parts[1])

		if err != nil {
			return nil, err
		}

		d := &edtf.EDTFDate{
			Start: start,
			End:   end,
			Level: LEVEL,
			Label: INTERVAL,
			EDTF:  edtf_str,
		}

		return d, nil
	*/
}
