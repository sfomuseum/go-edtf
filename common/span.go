package common

import (
	_ "fmt"
	"github.com/whosonfirst/go-edtf"
	"strings"
	"time"
)

func DateSpanFromEDTF(edtf_str string) (*edtf.DateSpan, error) {

	parts := strings.Split(edtf_str, "/")

	left_edtf := parts[0]

	sp, err := dateSpanFromEDTF(left_edtf)

	if err != nil {
		return nil, err
	}

	if len(parts) == 2 {

		right_edtf := parts[1]

		right_sp, err := dateSpanFromEDTF(right_edtf)

		if err != nil {
			return nil, err
		}

		sp.Start.Upper = sp.End.Upper

		right_sp.End.Lower = right_sp.Start.Lower

		sp.End = right_sp.End
	}

	return sp, nil
}

func dateSpanFromEDTF(edtf_str string) (*edtf.DateSpan, error) {

	str_range, err := StringRangeFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	start := str_range.Start
	end := str_range.End

	start_ymd, err := YMDFromStringDate(start)

	if err != nil {
		return nil, err
	}

	end_ymd, err := YMDFromStringDate(end)

	if err != nil {
		return nil, err
	}

	var start_lower_t *time.Time
	var start_upper_t *time.Time

	var end_lower_t *time.Time
	var end_upper_t *time.Time

	// fmt.Println("START", start)
	// fmt.Println("END", end)

	if end.Equals(start) {

		st, err := TimeWithYMD(start_ymd, edtf.HMS_LOWER)

		if err != nil {
			return nil, err
		}

		et, err := TimeWithYMD(end_ymd, edtf.HMS_UPPER)

		if err != nil {
			return nil, err
		}

		start_lower_t = st
		start_upper_t = st

		end_lower_t = et
		end_upper_t = et

	} else {

		sl, err := TimeWithYMD(start_ymd, edtf.HMS_LOWER)

		if err != nil {
			return nil, err
		}

		su, err := TimeWithYMD(start_ymd, edtf.HMS_UPPER)

		if err != nil {
			return nil, err
		}

		el, err := TimeWithYMD(end_ymd, edtf.HMS_LOWER)

		if err != nil {
			return nil, err
		}

		eu, err := TimeWithYMD(end_ymd, edtf.HMS_UPPER)

		if err != nil {
			return nil, err
		}

		start_lower_t = sl
		start_upper_t = su
		end_lower_t = el
		end_upper_t = eu

		/*
			fmt.Printf("START LOWER %v\n", sl)
			fmt.Printf("START UPPER %v\n", su)
			fmt.Printf("END LOWER %v\n", el)
			fmt.Printf("END UPPER %v\n", eu)
		*/
	}

	//

	start_lower := &edtf.Date{
		Time:        start_lower_t,
		YMD:         start_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
	}

	start_upper := &edtf.Date{
		Time:        start_upper_t,
		YMD:         start_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
	}

	end_lower := &edtf.Date{
		Time:        end_lower_t,
		YMD:         end_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
	}

	end_upper := &edtf.Date{
		Time:        end_upper_t,
		YMD:         end_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
	}

	start_range := &edtf.DateRange{
		Lower: start_lower,
		Upper: start_upper,
	}

	end_range := &edtf.DateRange{
		Lower: end_lower,
		Upper: end_upper,
	}

	sp := &edtf.DateSpan{
		Start: start_range,
		End:   end_range,
	}

	return sp, nil
}
