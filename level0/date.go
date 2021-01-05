package level0

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
	// "time"
	"fmt"
)

/*

Date

    complete representation:            [year][“-”][month][“-”][day]
    Example 1          ‘1985-04-12’ refers to the calendar date 1985 April 12th with day precision.
    reduced precision for year and month:   [year][“-”][month]
    Example 2          ‘1985-04’ refers to the calendar month 1985 April with month precision.
    reduced precision for year:  [year]
    Example 3          ‘1985’ refers to the calendar year 1985 with year precision.

*/

func IsDate(edtf_str string) bool {
	return re.Date.MatchString(edtf_str)
}

func ParseDate(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.Date.MatchString(edtf_str) {
		return nil, edtf.Invalid(DATE, edtf_str)
	}

	r, err := common.StringRangeFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}
	
	r_start := r.Start
	r_end := r.End
	
	start_ymd, err := common.YMDFromStringDate(r_start)

	if err != nil {
		return nil, err
	}
	
	end_ymd, err := common.YMDFromStringDate(r_end)

	if err != nil {
		return nil, err
	}

	fmt.Println(edtf_str, start_ymd, end_ymd)
	
	lower_t, err := common.TimeWithYMD(start_ymd, edtf.HMS_LOWER)

	if err != nil {
		return nil, err
	}

	upper_t, err := common.TimeWithYMD(end_ymd, edtf.HMS_UPPER)

	if err != nil {
		return nil, err
	}
	
	start_lower := &edtf.Date{
		YMD: start_ymd,
	}

	start_upper := &edtf.Date{
		YMD: start_ymd,
	}

	end_lower := &edtf.Date{
		YMD: end_ymd,
	}

	end_upper := &edtf.Date{
		YMD: end_ymd,
	}

	start_lower.Time = lower_t
	start_upper.Time = lower_t		
	
	end_lower.Time = upper_t
	end_upper.Time = upper_t
		
	start := &edtf.DateRange{
		Lower: start_lower,
		Upper: start_upper,
	}

	end := &edtf.DateRange{
		Lower: end_lower,
		Upper: end_upper,
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
