package level1

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/calendar"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
	"strconv"
	"strings"
)

/*

Seasons

The values 21, 22, 23, 24 may be used used to signify ' Spring', 'Summer', 'Autumn', 'Winter', respectively, in place of a month value (01 through 12) for a year-and-month format string.

    Example                   2001-21     Spring, 2001

*/

func IsSeason(edtf_str string) bool {
	return re.Season.MatchString(edtf_str)
}

func ParseSeason(edtf_str string) (*edtf.EDTFDate, error) {

	/*
		SEASON 5 [2001-01 2001 01  ]
		SEASON 5 [2001-24 2001 24  ]
		SEASON 5 [Spring, 2002   Spring 2002]
		SEASON 5 [winter, 2002   winter 2002]
	*/

	m := re.Season.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, edtf.Invalid(SEASON, edtf_str)
	}

	var start_yyyy int
	var start_mm int
	var start_dd int

	var end_yyyy int
	var end_mm int
	var end_dd int

	if m[1] == "" {

		season := m[3]
		str_yyyy := m[4]

		yyyy, err := strconv.Atoi(str_yyyy)

		if err != nil {
			return nil, err
		}

		switch strings.ToUpper(season) {
		case "WINTER":

			start_yyyy = yyyy
			start_mm = 12
			start_dd = 1

			end_yyyy = yyyy + 1
			end_mm = 2

		case "SPRING":

			start_yyyy = yyyy
			start_mm = 3
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 5

		case "SUMMER":

			start_yyyy = yyyy
			start_mm = 6
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 8

		case "FALL":

			start_yyyy = yyyy
			start_mm = 9
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 11

		default:
			return nil, edtf.Invalid(SEASON, edtf_str)
		}

	} else {

		str_yyyy := m[1]
		str_mm := m[2]

		yyyy, err := strconv.Atoi(str_yyyy)

		if err != nil {
			return nil, err
		}

		mm, err := strconv.Atoi(str_mm)

		if err != nil {
			return nil, err
		}

		switch mm {
		case 21: // spring

			start_yyyy = yyyy
			start_mm = 3
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 5

		case 22: // summer

			start_yyyy = yyyy
			start_mm = 6
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 8

		case 23: // autumn

			start_yyyy = yyyy
			start_mm = 9
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 11

		case 24: // winter

			start_yyyy = yyyy
			start_mm = 12
			start_dd = 1

			end_yyyy = yyyy + 1
			end_mm = 2

		default:

			start_yyyy = yyyy
			start_mm = mm
			start_dd = 1

			end_yyyy = yyyy
			end_mm = mm
		}

	}

	dm, err := calendar.DaysInMonth(end_yyyy, end_mm)

	if err != nil {
		return nil, err
	}

	end_dd = dm

	_str := fmt.Sprintf("%04d-%02d-%02d/%04d-%02d-%02d", start_yyyy, start_mm, start_dd, end_yyyy, end_mm, end_dd)

	sp, err := common.DateSpanFromEDTF(_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: SEASON,
	}

	return d, nil
}
