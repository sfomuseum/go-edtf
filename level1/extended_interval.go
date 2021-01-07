package level1

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
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

	if !re.IntervalStart.MatchString(edtf_str) {
		return nil, edtf.Invalid(EXTENDED_INTERVAL_START, edtf_str)
	}

	sp, err := common.DateSpanFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: EXTENDED_INTERVAL_START,
	}

	return d, nil
}

func ParseExtendedIntervalEnd(edtf_str string) (*edtf.EDTFDate, error) {

	/*
		END 5 1985/..,1985,,,..
		END 5 1985/,1985,,,
	*/

	if !re.IntervalEnd.MatchString(edtf_str) {
		return nil, edtf.Invalid(EXTENDED_INTERVAL_END, edtf_str)
	}

	sp, err := common.DateSpanFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: EXTENDED_INTERVAL_END,
	}

	return d, nil
}
