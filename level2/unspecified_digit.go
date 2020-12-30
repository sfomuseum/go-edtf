package level2

import (
	_ "fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
	_ "strings"
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

	if !re.UnspecifiedDigit.MatchString(edtf_str) {
		return nil, edtf.Invalid(UNSPECIFIED_DIGIT, edtf_str)
	}

	return nil, edtf.NotImplemented(UNSPECIFIED_DIGIT, edtf_str)
}
