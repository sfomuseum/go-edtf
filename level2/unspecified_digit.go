package level2

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
	// "strconv"
	// "strings"
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

	if !re.UnspecifiedDigit.MatchString(edtf_str) {
		return nil, edtf.Invalid(UNSPECIFIED_DIGIT, edtf_str)
	}

	sp, err := common.DateSpanFromEDTF(edtf_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: UNSPECIFIED_DIGIT,
	}

	return d, nil
}
