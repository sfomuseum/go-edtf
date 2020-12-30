package level2

import (
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
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

	if !re.SignificantDigits.MatchString(edtf_str) {
		return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
	}

	/*

		SIGN 9 1950S2,1950,,,,,,,2
		SIGN 9 Y171010000S3,,,171010000,,,,,3
		SIGN 9 Y-1S3,,-,1,,,,,3
		SIGN 9 Y3388E2S3,,,,3388E2,,3388,2,3
		SIGN 9 Y-20E2S3,,,,-20E2,-,20,2,3

	*/

	m := re.SignificantDigits.FindStringSubmatch(edtf_str)

	fmt.Println("SIGN", len(m), strings.Join(m, ","))

	if len(m) != 9 {
		return nil, edtf.Invalid(SIGNIFICANT_DIGITS, edtf_str)
	}

	str_yyyy := m[1]
	// prefix := m[3]
	str_year := m[3]
	notation := m[4]
	digits := m[8]

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

	fmt.Println("YYYY IS", yyyy)

	if yyyy > edtf.MAX_YEARS {
		return nil, edtf.Unsupported(SIGNIFICANT_DIGITS, edtf_str)
	}

	fmt.Println("DIGITS", edtf_str, yyyy, digits)

	return nil, edtf.NotImplemented(SIGNIFICANT_DIGITS, edtf_str)
}
