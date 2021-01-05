package level0

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
	"time"
	// "fmt"
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

	var start_lower_t *time.Time
	var start_upper_t *time.Time

	var end_lower_t *time.Time
	var end_upper_t *time.Time

	if r_end.Equals(r_start) {

		st, err := common.TimeWithYMD(start_ymd, edtf.HMS_LOWER)

		if err != nil {
			return nil, err
		}

		et, err := common.TimeWithYMD(end_ymd, edtf.HMS_UPPER)

		if err != nil {
			return nil, err
		}

		start_lower_t = st
		start_upper_t = st

		end_lower_t = et
		end_upper_t = et

	} else {

		sl, err := common.TimeWithYMD(start_ymd, edtf.HMS_LOWER)

		if err != nil {
			return nil, err
		}

		su, err := common.TimeWithYMD(start_ymd, edtf.HMS_UPPER)

		if err != nil {
			return nil, err
		}

		el, err := common.TimeWithYMD(end_ymd, edtf.HMS_LOWER)

		if err != nil {
			return nil, err
		}

		eu, err := common.TimeWithYMD(end_ymd, edtf.HMS_UPPER)

		if err != nil {
			return nil, err
		}

		start_lower_t = sl
		start_upper_t = su
		end_lower_t = el
		end_upper_t = eu
	}

	start_lower := &edtf.Date{
		Time: start_lower_t,
		YMD:  start_ymd,
	}

	start_upper := &edtf.Date{
		Time: start_upper_t,
		YMD:  start_ymd,
	}

	end_lower := &edtf.Date{
		Time: end_lower_t,
		YMD:  end_ymd,
	}

	end_upper := &edtf.Date{
		Time: end_upper_t,
		YMD:  end_ymd,
	}

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
