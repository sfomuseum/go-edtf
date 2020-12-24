package level1

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/calendar"
	"strconv"
	"strings"
	"time"
)

/*

Seasons

The values 21, 22, 23, 24 may be used used to signify ' Spring', 'Summer', 'Autumn', 'Winter', respectively, in place of a month value (01 through 12) for a year-and-month format string.

    Example                   2001-21     Spring, 2001

*/

func IsSeason(edtf_str string) bool {
	return re_season.MatchString(edtf_str)
}

func ParseSeason(edtf_str string) (*edtf.EDTFDate, error) {

	/*
		SEASON 5 [2001-01 2001 01  ]
		SEASON 5 [2001-24 2001 24  ]
		SEASON 5 [Spring, 2002   Spring 2002]
		SEASON 5 [winter, 2002   winter 2002]
	*/

	m := re_season.FindStringSubmatch(edtf_str)

	if len(m) != 5 {
		return nil, errors.New("Invalid Level 1 season string")
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
			start_mm = 1
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 3

		case "SPRING":

			start_yyyy = yyyy
			start_mm = 4
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 6

		case "SUMMER":

			start_yyyy = yyyy
			start_mm = 7
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 9

		case "FALL":

			start_yyyy = yyyy
			start_mm = 10
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 12

		default:
			return nil, errors.New("Invalid season")
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
		case 21:

			start_yyyy = yyyy
			start_mm = 1
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 3

		case 22:

			start_yyyy = yyyy
			start_mm = 4
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 6

		case 23:

			start_yyyy = yyyy
			start_mm = 7
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 9

		case 24:

			start_yyyy = yyyy
			start_mm = 10
			start_dd = 1

			end_yyyy = yyyy
			end_mm = 12

		default:

			start_yyyy = yyyy
			start_mm = mm
			start_dd = 1

			end_yyyy = yyyy
			end_mm = mm
		}

		dm, err := calendar.DaysInMonth(uint(end_yyyy), uint(end_mm))

		if err != nil {
			return nil, err
		}

		end_dd = int(dm)
	}

	start, err := dateRange(start_yyyy, start_mm, start_dd)

	if err != nil {
		return nil, err
	}

	end, err := dateRange(end_yyyy, end_mm, end_dd)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start: start,
		End:   end,
		EDTF:  edtf_str,
		Level: LEVEL,
	}

	return d, nil
}

func dateRange(yyyy int, mm int, dd int) (*edtf.DateRange, error) {

	lower_hms := "00:00:00"
	upper_hms := "23:59:59"

	lower_str := fmt.Sprintf("%04d-%02d-%02dT%s", yyyy, mm, dd, lower_hms)
	upper_str := fmt.Sprintf("%04d-%02d-%02dT%s", yyyy, mm, dd, upper_hms)

	lower_t, err := time.Parse("2006-01-02T15:04:05", lower_str)

	if err != nil {
		return nil, err
	}

	upper_t, err := time.Parse("2006-01-02T15:04:05", upper_str)

	if err != nil {
		return nil, err
	}

	lower_d := &edtf.Date{
		Time: &lower_t,
	}

	upper_d := &edtf.Date{
		Time: &upper_t,
	}

	dt := &edtf.DateRange{
		Lower: lower_d,
		Upper: upper_d,
	}

	return dt, nil
}
