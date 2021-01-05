package common

import (
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"strings"
)

func DateSpanWithString(edtf_str string) (*edtf.DateSpan, error) {

	parts := strings.Split(edtf_str, "/")

	fmt.Println("SPAN PARTS", edtf_str, parts)

	first := parts[0]

	start_r, err := StringRangeFromEDTF(first)

	if err != nil {
		return nil, err
	}

	fmt.Println("SPAN STRING RANGE", edtf_str, start_r)

	start, err := DateRangeWithStringRange(start_r)

	if err != nil {
		return nil, err
	}

	fmt.Println("SPAN DATE RANGE", edtf_str, start)

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
