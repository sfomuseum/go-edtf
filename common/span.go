package common

import (
	"github.com/whosonfirst/go-edtf"
	"strings"
)

func DateSpanWithString(edtf_str string) (*edtf.DateSpan, error) {

	parts := strings.Split(edtf_str, "/")
	first := parts[0]

	start_r, err := StringRangeFromEDTF(first)

	if err != nil {
		return nil, err
	}

	start, err := DateRangeWithStringRange(start_r)

	if err != nil {
		return nil, err
	}

	var end *edtf.DateRange

	if len(parts) == 2 {

		last := parts[1]

		end_r, err := StringRangeFromEDTF(last)

		if err != nil {
			return nil, err
		}

		r, err := DateRangeWithStringRange(end_r)

		if err != nil {
			return nil, err
		}

		end = r

	} else {
		end = start
	}

	sp := &edtf.DateSpan{
		Start: start,
		End:   end,
	}

	return sp, nil
}
