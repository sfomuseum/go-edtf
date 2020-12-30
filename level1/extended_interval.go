package level1

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
)

/*

Extended Interval (L1)

    A null string may be used for the start or end date when it is unknown.
    Double-dot (“..”) may be used when either the start or end date is not specified, either because there is none or for any other reason.
    A modifier may appear at the end of the date to indicate "uncertain" and/or "approximate"

Open end time interval

    Example 1          ‘1985-04-12/..’
    interval starting at 1985 April 12th with day precision; end open
    Example 2          ‘1985-04/..’
    interval starting at 1985 April with month precision; end open
    Example 3          ‘1985/..’
    interval starting at year 1985 with year precision; end open

Open start time interval

    Example 4          ‘../1985-04-12’
    interval with open start; ending 1985 April 12th with day precision
    Example 5          ‘../1985-04’
    interval with open start; ending 1985 April with month precision
    Example 6          ‘../1985’
    interval with open start; ending at year 1985 with year precision

Time interval with unknown end

    Example 7          ‘1985-04-12/’
    interval starting 1985 April 12th with day precision; end unknown
    Example 8          ‘1985-04/’
    interval starting 1985 April with month precision; end unknown
    Example 9          ‘1985/’
    interval starting year 1985 with year precision; end unknown

Time interval with unknown start

    Example 10       ‘/1985-04-12’
    interval with unknown start; ending 1985 April 12th with day precision
    Example 11       ‘/1985-04’
    interval with unknown start; ending 1985 April with month precision
    Example 12       ‘/1985’
    interval with unknown start; ending year 1985 with year precision

*/

func IsExtendedInterval(edtf_str string) bool {

	if re.IntervalEnd.MatchString(edtf_str) {
		return true
	}

	if re.IntervalStart.MatchString(edtf_str) {
		return true
	}

	return true
}

func ParseExtendedInterval(edtf_str string) (*edtf.EDTFDate, error) {

	if re.IntervalStart.MatchString(edtf_str) {
		return ParseExtendedIntervalStart(edtf_str)
	}

	if re.IntervalEnd.MatchString(edtf_str) {
		return ParseExtendedIntervalEnd(edtf_str)
	}

	return nil, edtf.Invalid(EXTENDED_INTERVAL, edtf_str)
}

func ParseExtendedIntervalStart(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		START 5 ../1985-04-12,..,1985,04,12
		START 5 ../1985-04,..,1985,04,
		START 5 ../1985,..,1985,,
		START 5 /1985-04-12,,1985,04,12
		START 5 /1985-04,,1985,04,
		START 5 /1985,,1985,,

	*/

	m := re.IntervalStart.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, edtf.Invalid(EXTENDED_INTERVAL_START, edtf_str)
	}

	start_dt := m[1]
	end_yyyy := m[2]
	end_mm := m[3]
	end_dd := m[4]

	end, err := common.DateRangeWithYMDString(end_yyyy, end_mm, end_dd)

	if err != nil {
		return nil, err
	}

	start := &edtf.DateRange{
		Lower: &edtf.Date{},
		Upper: &edtf.Date{},
	}

	switch start_dt {
	case edtf.OPEN:
		start.Upper.Open = true
		start.Lower.Open = true
	case edtf.UNKNOWN:
		start.Upper.Unknown = true
		start.Lower.Unknown = true
	default:
		return nil, edtf.Invalid(EXTENDED_INTERVAL_START, edtf_str)
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}

func ParseExtendedIntervalEnd(edtf_str string) (*edtf.EDTFDate, error) {

	/*
		END 5 1985/..,1985,,,..
		END 5 1985/,1985,,,
	*/

	m := re.IntervalEnd.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, edtf.Invalid(EXTENDED_INTERVAL_END, edtf_str)
	}

	return nil, nil

	start_yyyy := m[1]
	start_mm := m[2]
	start_dd := m[3]
	end_dt := m[4]

	start, err := common.DateRangeWithYMDString(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	end := &edtf.DateRange{
		Lower: &edtf.Date{},
		Upper: &edtf.Date{},
	}

	switch end_dt {
	case edtf.OPEN:
		end.Upper.Open = true
		end.Lower.Open = true
	case edtf.UNKNOWN:
		end.Upper.Unknown = true
		end.Lower.Unknown = true
	default:
		return nil, edtf.Invalid(EXTENDED_INTERVAL_END, edtf_str)
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
