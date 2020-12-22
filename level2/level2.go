package level2

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"regexp"
)

var re_exponential_year *regexp.Regexp
var re_significant_digits *regexp.Regexp

func init() {

	re_exponential_year = regexp.MustCompile(`^(?i)Y(\-?\d+)E(\d+)$`)

	re_significant_digits = regexp.MustCompile(`(?:(\d{4})S(\d+)|Y(\d+)S(\d+)|Y(\d+)E(\d+)S(\d+))$`)
}

/*

Exponential year

'Y' at the beginning of the string (which indicates "year", as in level 1) may be followed by an integer, followed by 'E' followed by a positive integer. This signifies "times 10 to the power of". Thus 17E8 means "17 times 10 to the eighth power".

    Example        ‘Y-17E7’
    the calendar year -17*10 to the seventh power= -170000000


*/

func ParseExponentialYear(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_exponential_year.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 exponential year string")
	}

	return nil, nil
}

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

func ParseSignificantDigits(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_significant_digits.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 significant digits string")
	}

	return nil, nil

}
