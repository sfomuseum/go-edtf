package level2

import (
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/common"
	"github.com/whosonfirst/go-edtf/re"
	"sort"
	"strings"
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

	/*

		SET [1667,1668,1670..1672] 9 [1667,1668,1670..1672],[,,1672,,,..,,]
		SET [..1760-12-03] 9 [..1760-12-03],[,..,1760,12,03,,,]
		SET [1760-12..] 9 [1760-12..],[,,1760,12,,..,,]
		SET [1760-01,1760-02,1760-12..] 9 [1760-01,1760-02,1760-12..],[,,1760,12,,..,,]
		SET [1667,1760-12] 9 [1667,1760-12],[,,1760,12,,,,,]
		SET [..1984] 9 [..1984],[,..,1984,,,,,]
		SET {1667,1668,1670..1672} 9 {1667,1668,1670..1672},{,,1672,,,..,,}
		SET {1960,1961-12} 9 {1960,1961-12},{,,1961,12,,,,,}
		SET {..1984} 9 {..1984},{,..,1984,,,,,}

	*/

	m := re.SetRepresentations.FindStringSubmatch(edtf_str)

	if len(m) != 6 {
		return nil, edtf.Invalid(SET_REPRESENTATIONS, edtf_str)
	}

	class := m[1]
	candidates := m[2]

	start_ymd := ""
	end_ymd := ""

	start_open := false
	end_open := false

	inclusivity := edtf.NONE

	switch class {
	case "[":
		inclusivity = edtf.ANY
	case "{":
		inclusivity = edtf.ALL
	default:
		return nil, edtf.Invalid(SET_REPRESENTATIONS, edtf_str)
	}

	// this should be moved in to a separate method for getting
	// the list of all possible dates - we only care about the
	// bookends right now (20201231/thisisaaronland)

	possible := make([]string, 0)

	for _, date := range strings.Split(candidates, ",") {

		parts := strings.Split(date, "..")
		count := len(parts)

		switch count {
		case 1:
			possible = append(possible, date)
			continue
		case 2:

			if parts[0] != "" && parts[1] != "" { // YYYY..YYYY

				// get everything in between parts[0] and parts[1]
				// need to determine what to get (days, months, years)

				possible = append(possible, parts[0])
				possible = append(possible, parts[1])

			} else if parts[0] == "" { // ..YYYY

				// parts[1] is end (max) date
				// start (min) date is "open" or "unknown"

				possible = append(possible, parts[1])
				start_open = true

			} else { // YYYY..

				// parts[0] is start (min) date
				// end (max) date is "open" or "unknown"

				possible = append(possible, parts[0])
				end_open = true
			}

		default:
			return nil, edtf.Invalid(SET_REPRESENTATIONS, edtf_str)
		}
	}

	sort.Strings(possible)

	count := len(possible)

	switch count {
	case 0:
		return nil, edtf.Invalid(SET_REPRESENTATIONS, edtf_str)
	case 1:
		start_ymd = possible[0]
		end_ymd = start_ymd
	default:
		start_ymd = possible[0]
		end_ymd = possible[count-1]
	}

	var start *edtf.DateRange
	var end *edtf.DateRange

	if start_open {

		start = common.EmptyDateRange()
		start.Lower.Open = true
		start.Upper.Open = true

		start.Lower.Inclusivity = inclusivity
		start.Upper.Inclusivity = inclusivity

	} else {

		dr, err := common.DateRangeWithYMDStringCombined(start_ymd)

		if err != nil {
			return nil, err
		}

		start = dr
	}

	if end_open {

		end = common.EmptyDateRange()
		end.Lower.Open = true
		end.Upper.Open = true

		end.Lower.Inclusivity = inclusivity
		end.Upper.Inclusivity = inclusivity

	} else {

		dr, err := common.DateRangeWithYMDStringCombined(end_ymd)

		if err != nil {
			return nil, err
		}

		end = dr
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}
