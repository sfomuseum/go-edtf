package level0

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
)

/*

Time Interval

EDTF Level 0 adopts representations of a time interval where both the start and end are dates: start and end date only; that is, both start and duration, and duration and end, are excluded. Time of day is excluded.

    Example 1          ‘1964/2008’ is a time interval with calendar year precision, beginning sometime in 1964 and ending sometime in 2008.
    Example 2          ‘2004-06/2006-08’ is a time interval with calendar month precision, beginning sometime in June 2004 and ending sometime in August of 2006.
    Example 3          ‘2004-02-01/2005-02-08’ is a time interval with calendar day precision, beginning sometime on February 1, 2004 and ending sometime on February 8, 2005.
    Example 4          ‘2004-02-01/2005-02’ is a time interval beginning sometime on February 1, 2004 and ending sometime in February 2005. Since the start endpoint precision (day) is different than that of the end endpoint (month) the precision of the time interval at large is undefined.
    Example 5          ‘2004-02-01/2005’ is a time interval beginning sometime on February 1, 2004 and ending sometime in 2005. The start endpoint has calendar day precision and the end endpoint has calendar year precision. Similar to the previous example, the precision of the time interval at large is undefined.
    Example 6          ‘2005/2006-02’ is a time interval beginning sometime in 2005 and ending sometime in February 2006.

*/

func IsTimeInterval(edtf_str string) bool {
	return re.TimeInterval.MatchString(edtf_str)
}

func ParseTimeInterval(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.TimeInterval.MatchString(edtf_str) {
		return nil, edtf.Invalid(TIME_INTERVAL, edtf_str)
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
		Feature: TIME_INTERVAL,
	}

	return d, nil
}
