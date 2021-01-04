package level2

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
	"strconv"
	"strings"
)

/*

Unspecified Digit

For level 2 the unspecified digit, 'X', may occur anywhere within a component.

    Example 1                 ‘156X-12-25’
    December 25 sometime during the 1560s
    Example 2                 ‘15XX-12-25’
    December 25 sometime during the 1500s
    Example 3                ‘XXXX-12-XX’
    Some day in December in some year
    Example 4                 '1XXX-XX’
    Some month during the 1000s
    Example 5                  ‘1XXX-12’
    Some December during the 1000s
    Example 6                  ‘1984-1X’
    October, November, or December 1984

*/

func IsUnspecifiedDigit(edtf_str string) bool {
	return re.UnspecifiedDigit.MatchString(edtf_str)
}

func ParseUnspecifiedDigit(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		UNSPEC 156X-12-25 4 156X-12-25,156X,12,25
		UNSPEC 15XX-12-25 4 15XX-12-25,15XX,12,25
		UNSPEC 1XXX-XX 4 1XXX-XX,1XXX,XX,
		UNSPEC 1XXX-12 4 1XXX-12,1XXX,12,
		UNSPEC 1984-1X 4 1984-1X,1984,1X,

	*/

	m := re.UnspecifiedDigit.FindStringSubmatch(edtf_str)

	if len(m) != 4 {
		return nil, edtf.Invalid(UNSPECIFIED_DIGIT, edtf_str)
	}

	yyyy := m[1]
	mm := m[2]
	dd := m[3]

	start_yyyy := yyyy
	end_yyyy := yyyy

	start_mm := mm
	start_dd := dd

	end_mm := mm
	end_dd := dd

	precision := edtf.NONE

	if strings.HasSuffix(yyyy, "X") {

		start_m := int64(0)
		end_m := int64(0)

		start_c := int64(0)
		end_c := int64(900)

		start_d := int64(0)
		end_d := int64(90)

		start_y := int64(0)
		end_y := int64(9)

		if string(yyyy[0]) == "X" {
			return nil, edtf.NotImplemented(UNSPECIFIED_DIGIT, edtf_str)
		} else {

			m, err := strconv.ParseInt(string(yyyy[0]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_m = m * 1000
			end_m = start_m
		}

		if string(yyyy[1]) != "X" {

			c, err := strconv.ParseInt(string(yyyy[1]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_c = c * 100
			end_c = start_c
		}

		if string(yyyy[2]) != "X" {

			d, err := strconv.ParseInt(string(yyyy[2]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_d = d * 10
			end_d = start_d
		}

		if string(yyyy[3]) != "X" {

			y, err := strconv.ParseInt(string(yyyy[3]), 10, 32)

			if err != nil {
				return nil, err
			}

			start_y = y * 1
			end_y = start_y
		}

		start_ymd := start_m + start_c + start_d + start_y
		end_ymd := end_m + end_c + end_d + end_y

		// fmt.Printf("OMG '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, start_m, start_c, start_d, start_y, start_ymd)
		// fmt.Printf("WTF '%s' '%d' '%d' '%d' '%d' '%d'\n", yyyy, end_m, end_c, end_d, end_y, end_ymd)

		start_yyyy = strconv.FormatInt(start_ymd, 10)
		end_yyyy = strconv.FormatInt(end_ymd, 10)

		precision.AddFlag(edtf.ANNUAL)
	}

	if strings.HasSuffix(mm, "X") {

		// this does not account for 1985-24, etc.

		if strings.HasPrefix(mm, "X") {

			start_mm = "01"
			end_mm = "12"
		} else {

			start_mm = "10"
			end_mm = "12"
		}

		precision.AddFlag(edtf.MONTHLY)
	}

	if strings.HasSuffix(dd, "X") {

		switch string(dd[0]) {
		case "X":
			start_dd = "01"
			end_dd = ""
		case "1":
			start_dd = "10"
			end_dd = "19"
		case "2":
			start_dd = "20"
			end_dd = "29"
		case "3":
			start_dd = "30"
			end_dd = ""
		default:
			return nil, edtf.Invalid(UNSPECIFIED_DIGIT, edtf_str)
		}

		precision.AddFlag(edtf.DAILY)
	}

	start, err := common.DateRangeWithYMDString(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	end, err := common.DateRangeWithYMDString(end_yyyy, end_mm, end_dd)

	if err != nil {
		return nil, err
	}

	if precision != edtf.NONE {

		start.Lower.Unspecified = precision
		start.Upper.Unspecified = precision

		end.Lower.Unspecified = precision
		end.Upper.Unspecified = precision
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
