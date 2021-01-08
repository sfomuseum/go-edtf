package common

import (
	"github.com/sfomuseum/go-edtf"
	"strings"
	"time"
)

func DateSpanFromEDTF(edtf_str string) (*edtf.DateSpan, error) {

	parts := strings.Split(edtf_str, "/")
	count := len(parts)

	is_multi := false

	var left_edtf string
	var right_edtf string

	switch count {
	case 2:
		left_edtf = parts[0]
		right_edtf = parts[1]
		is_multi = true
	case 1:
		left_edtf = parts[0]
	default:
		return nil, edtf.Invalid("date span", edtf_str)
	}

	if !is_multi {
		return dateSpanFromYMD(left_edtf)
	}

	left_span, err := dateSpanFromEDTF(left_edtf)

	if err != nil {
		return nil, err
	}

	right_span, err := dateSpanFromEDTF(right_edtf)

	if err != nil {
		return nil, err
	}

	left_span.Start.Upper = left_span.End.Upper

	right_span.End.Lower = right_span.Start.Lower

	left_span.End = right_span.End

	return left_span, nil
}

// specifically from one half of a FOO/BAR string

func dateSpanFromEDTF(edtf_str string) (*edtf.DateSpan, error) {

	var span *edtf.DateSpan

	switch edtf_str {
	case edtf.UNKNOWN:

		span = UnknownDateSpan()

		span.Start.EDTF = edtf_str
		span.End.EDTF = edtf_str

	case edtf.OPEN:

		span = OpenDateSpan()

		span.Start.EDTF = edtf_str
		span.End.EDTF = edtf_str

	default:

		ds, err := dateSpanFromYMD(edtf_str)

		if err != nil {
			return nil, err
		}

		span = ds
	}

	return span, nil
}

func dateSpanFromYMD(edtf_str string) (*edtf.DateSpan, error) {

	str_range, err := StringRangeFromYMD(edtf_str)

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
		YMD:         start_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
		Precision:   str_range.Precision,
	}

	start_upper := &edtf.Date{
		YMD:         start_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
		Precision:   str_range.Precision,
	}

	end_lower := &edtf.Date{
		YMD:         end_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
		Precision:   str_range.Precision,
	}

	end_upper := &edtf.Date{
		YMD:         end_ymd,
		Uncertain:   str_range.Uncertain,
		Approximate: str_range.Approximate,
		Precision:   str_range.Precision,
	}

	if start_lower_t != nil {
		start_lower.SetTime(start_lower_t)
	}

	if start_upper_t != nil {
		start_upper.SetTime(start_upper_t)
	}

	if end_lower_t != nil {
		end_lower.SetTime(end_lower_t)
	}

	if end_upper_t != nil {
		end_upper.SetTime(end_upper_t)
	}

	start_range := &edtf.DateRange{
		EDTF:  edtf_str,
		Lower: start_lower,
		Upper: start_upper,
	}

	end_range := &edtf.DateRange{
		EDTF:  edtf_str,
		Lower: end_lower,
		Upper: end_upper,
	}

	sp := &edtf.DateSpan{
		Start: start_range,
		End:   end_range,
	}

	return sp, nil
}

func EmptyDateSpan() *edtf.DateSpan {

	start := EmptyDateRange()
	end := EmptyDateRange()

	sp := &edtf.DateSpan{
		Start: start,
		End:   end,
	}

	return sp
}

func UnknownDateSpan() *edtf.DateSpan {

	start := UnknownDateRange()
	end := UnknownDateRange()

	sp := &edtf.DateSpan{
		Start: start,
		End:   end,
	}

	return sp
}

func OpenDateSpan() *edtf.DateSpan {

	start := OpenDateRange()
	end := OpenDateRange()

	sp := &edtf.DateSpan{
		Start: start,
		End:   end,
	}

	return sp
}
