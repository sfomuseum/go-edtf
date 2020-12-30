package level2

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
)

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

func IsSubYearGrouping(edtf_str string) bool {
	return re.SubYear.MatchString(edtf_str)
}

func ParseSubYearGroupings(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.SubYear.MatchString(edtf_str) {
		return nil, errors.New("Invalid Level 2 sub year groupings string")
	}

	return nil, nil
}
