package level2

import (
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
	"strings"
)

/*

For Level 2 portions of a date within an interval may be designated as approximate, uncertain, or unspecified.

    Example 1                 ‘2004-06-~01/2004-06-~20’
    An interval in June 2004 beginning approximately the first and ending approximately the 20th
    Example 2                 ‘2004-06-XX/2004-07-03’
    An interval beginning on an unspecified day in June 2004 and ending July 3.


*/

func IsInterval(edtf_str string) bool {
	return re.Interval.MatchString(edtf_str)
}

func ParseInterval(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		INTERVAL 2004-06-~01/2004-06-~20 13 2004-06-~01/2004-06-~20,,2004,,06,~,01,,2004,,06,~,20
		INTERVAL 2004-06-XX/2004-07-03 13 2004-06-XX/2004-07-03,,2004,,06,,XX,,2004,,07,,03

	*/

	m := re.Interval.FindStringSubmatch(edtf_str)

	fmt.Println("INTERVAL", edtf_str, len(m), strings.Join(m, ","))

	if len(m) != 13 {
		return nil, edtf.Invalid(INTERVAL, edtf_str)
	}

	return nil, edtf.NotImplemented(INTERVAL, edtf_str)
}
