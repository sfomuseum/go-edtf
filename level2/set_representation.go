package level2

import (
	_ "fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/re"
	_ "strings"
)

/*

Set representation

    Square brackets wrap a single-choice list (select one member).
    Curly brackets wrap an inclusive list (all members included).
    Members of the set are separated by commas.
    No spaces are allowed, anywhere within the expression.
    Double-dots indicates all the values between the two values it separates, inclusive.
    Double-dot at the beginning or end of the list means "on or before" or "on or after" respectively.
    Elements immediately preceeding and/or following as well as the elements represented by a double-dot, all have the same precision. Otherwise, different elements may have different precisions

One of a set

    Example 1       [1667,1668,1670..1672]
    One of the years 1667, 1668, 1670, 1671, 1672
    Example 2         [..1760-12-03]
    December 3, 1760; or some earlier date
    Example 3          [1760-12..]
    December 1760, or some later month
    Example 4         [1760-01,1760-02,1760-12..]
    January or February of 1760 or December 1760 or some later month
    Example 5          [1667,1760-12]
    Either the year 1667 or the month December of 1760.
    Example 6         [..1984]
    The year 1984 or an earlier year

All Members

    Example 7          {1667,1668,1670..1672}
    All of the years 1667, 1668, 1670, 1671, 1672
    Example 8            {1960,1961-12}
    The year 1960 and the month December of 1961.
    Example 9         {..1984}
    The year 1984 and all earlier years

*/

func IsSetRepresentation(edtf_str string) bool {
	return re.SetRepresentations.MatchString(edtf_str)
}

func ParseSetRepresentations(edtf_str string) (*edtf.EDTFDate, error) {

	if !re.SetRepresentations.MatchString(edtf_str) {
		return nil, edtf.Invalid(SET_REPRESENTATIONS, edtf_str)
	}

	return nil, edtf.NotImplemented(SET_REPRESENTATIONS, edtf_str)
}
