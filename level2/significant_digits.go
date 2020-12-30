package level2

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-edtf"
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
	return re_significant_digits.MatchString(edtf_str)
}

func ParseSignificantDigits(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_significant_digits.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 significant digits string")
	}

	/*

		SIGN 8 1950S2,1950,2,,,,,
		SIGN 8 Y171010000S3,,,171010000,3,,,
		SIGN 8 Y3388E2S3,,,,,3388,2,3

	*/

	m := re_significant_digits.FindStringSubmatch(edtf_str)

	fmt.Println("SIGN", len(m), strings.Join(m, ","))
	return nil, nil
}
