package level2

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
	"strconv"
	"strings"
)

/*

Significant digits

A year (expressed in any of the three allowable forms: four-digit, 'Y' prefix, or exponential) may be followed by 'S', followed by a positive integer indicating the number of significant digits.

    Example 1      ‘1950S2’
    some year between 1900 and 1999, estimated to be 1950
    Example 2      ‘Y171010000S3’
    some year between 171010000 and 171010999, estimated to be 171010000
    Example 3       ‘Y3388E2S3’
    some year between 338000 and 338999, estimated to be 338800.

*/

func IsSignificantDigits(edtf_str string) bool {
	return re.SignificantDigits.MatchString(edtf_str)
}

func ParseSignificantDigits(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		SIGN 5 1950S2,1950,,,2
		SIGN 5 Y171010000S3,,171010000,,3
		SIGN 5 Y-20E2S3,,,-20E2,3
		SIGN 5 Y3388E2S3,,,3388E2,3
		SIGN 5 Y-20E2S3,,,-20E2,3

	*/

	m := re.SignificantDigits.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
	}

	str_yyyy := m[1]
	str_year := m[2]
	notation := m[3]
	str_digits := m[4]

	var yyyy int

	if str_yyyy != "" {

		y, err := strconv.Atoi(str_yyyy)

		if err != nil {
			return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
		}

		yyyy = y

	} else if str_year != "" {

		if len(str_year) > 4 {
			return nil, edtf.Unsupported(SIGNIFICANT_DIGITS, edtf_str)
		}

		y, err := strconv.Atoi(str_year)

		if err != nil {
			return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
		}

		yyyy = y

	} else if notation != "" {

		y, err := common.ParseExponentialNotation(notation)

		if err != nil {
			return nil, err
		}

		yyyy = y

	} else {
		return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
	}

	if yyyy > edtf.MAX_YEARS {
		return nil, edtf.Unsupported(SIGNIFICANT_DIGITS, edtf_str)
	}

	digits, err := strconv.Atoi(str_digits)

	if err != nil {
		return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
	}

	if len(strconv.Itoa(digits)) > len(strconv.Itoa(yyyy)) {
		return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
	}

	str_yyyy = strconv.Itoa(yyyy)
	prefix_yyyy := str_yyyy[0 : len(str_yyyy)-digits]

	first := strings.Repeat("0", digits)
	last := strings.Repeat("9", digits)

	start_yyyy := prefix_yyyy + first
	end_yyyy := prefix_yyyy + last

	_str := fmt.Sprintf("%s/%s", start_yyyy, end_yyyy)

	if strings.HasPrefix(start_yyyy, "-") && strings.HasPrefix(end_yyyy, "-") {
		_str = fmt.Sprintf("%s/%s", end_yyyy, start_yyyy)
	}

	sp, err := common.DateSpanFromEDTF(_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: SIGNIFICANT_DIGITS,
	}

	return d, nil
}
