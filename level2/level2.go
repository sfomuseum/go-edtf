package level2

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"regexp"
)

var re_exponential_year *regexp.Regexp
var re_significant_digits *regexp.Regexp
var re_sub_year *regexp.Regexp

func init() {

	re_exponential_year = regexp.MustCompile(`^(?i)Y(\-?\d+)E(\d+)$`)

	re_significant_digits = regexp.MustCompile(`(?:(\d{4})S(\d+)|Y(\d+)S(\d+)|Y(\d+)E(\d+)S(\d+))$`)

	re_sub_year = regexp.MustCompile(`^(\d{4})\-(2[1-9]|3[0-9]|4[0-1])$`)
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

/*

Level 2 extends the season feature of Level 1 to include the following sub-year groupings.

21     Spring (independent of location)
22     Summer (independent of location)
23     Autumn (independent of location)
24     Winter (independent of location)
25     Spring - Northern Hemisphere
26     Summer - Northern Hemisphere
27     Autumn - Northern Hemisphere
28     Winter - Northern Hemisphere
29     Spring - Southern Hemisphere
30     Summer - Southern Hemisphere
31     Autumn - Southern Hemisphere
32     Winter - Southern Hemisphere
33     Quarter 1 (3 months in duration)
34     Quarter 2 (3 months in duration)
35     Quarter 3 (3 months in duration)
36     Quarter 4 (3 months in duration)
37     Quadrimester 1 (4 months in duration)
38     Quadrimester 2 (4 months in duration)
39     Quadrimester 3 (4 months in duration)
40     Semestral 1 (6 months in duration)
41     Semestral 2 (6 months in duration)

    Example        ‘2001-34’   
    second quarter of 2001

*/

func ParseSubYearGroupings(edtf_str string) (*edtf.EDTFDate, error) {

	if !re_sub_year.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 sub year groupings string")
	}

	return nil, nil
}
