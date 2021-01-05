package common

import (
	"github.com/whosonfirst/go-edtf"
)

func DateSpanWithString(edtf_str string) (*edtf.DateSpan, error) {

	start, err := DateRangeWithString(edtf_str)

	if err != nil {
		return nil, err
	}

	end := start

	sp := &edtf.DateSpan{
		Start: start,
		End:   end,
	}

	return sp, nil
}
