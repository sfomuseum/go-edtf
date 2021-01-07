package level2

import (
	"fmt"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/re"
	"strconv"
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
	return re.SubYearGrouping.MatchString(edtf_str)
}

func ParseSubYearGroupings(edtf_str string) (*edtf.EDTFDate, error) {

	/*

		SUB 3 2001-34,2001,34

	*/

	m := re.SubYearGrouping.FindStringSubmatch(edtf_str)

	if len(m) != 3 {
		return nil, edtf.Invalid(SUB_YEAR_GROUPINGS, edtf_str)
	}

	year := m[1]
	grouping := m[2]

	start_yyyy := year
	start_mm := ""
	start_dd := ""

	end_yyyy := year
	end_mm := ""
	end_dd := ""

	switch grouping {
	case "21", "25": // Spring (independent of location, Northern Hemisphere)
		start_mm = "03"
		start_dd = "01"
		end_mm = "05"
		end_dd = "31"
	case "22", "26": // Summer (independent of location, Northern Hemisphere)
		start_mm = "06"
		start_dd = "01"
		end_mm = "08"
		end_dd = "31"
	case "23", "27": // Autumn (independent of location, Northern Hemisphere)
		start_mm = "09"
		start_dd = "01"
		end_mm = "11"
		end_dd = "30"
	case "24", "28": // Winter (independent of location, Northern Hemisphere)
		start_mm = "12"
		start_dd = "01"
		end_mm = "02"
		end_dd = "" // leave blank to make the code look up daysforyear(year)

		y, err := strconv.Atoi(end_yyyy)

		if err != nil {
			return nil, err
		}

		end_yyyy = strconv.Itoa(y + 1)

	case "29": // Spring - Southern Hemisphere
		start_mm = "09"
		start_dd = "01"
		end_mm = "11"
		end_dd = "30"
	case "30": // Summer - Southern Hemisphere
		start_mm = "12"
		start_dd = "01"

		end_mm = "02"
		end_dd = "" // leave blank to make the code look up daysforyear(year)

		y, err := strconv.Atoi(end_yyyy)

		if err != nil {
			return nil, err
		}

		end_yyyy = strconv.Itoa(y + 1)

	case "31": // Autumn - Southern Hemisphere
		start_mm = "03"
		start_dd = "01"
		end_mm = "05"
		end_dd = "31"
	case "32": // Winter - Southern Hemisphere
		start_mm = "06"
		start_dd = "01"
		end_mm = "08"
		end_dd = "31"
	case "33": // Quarter 1 (3 months in duration)
		start_mm = "01"
		start_dd = "01"
		end_mm = "03"
		end_dd = "31"
	case "34": // Quarter 2 (3 months in duration)
		start_mm = "04"
		start_dd = "01"
		end_mm = "06"
		end_dd = "30"
	case "35": // Quarter 3 (3 months in duration)
		start_mm = "07"
		start_dd = "01"
		end_mm = "09"
		end_dd = "30"
	case "36": // Quarter 4 (3 months in duration)
		start_mm = "10"
		start_dd = "01"
		end_mm = "12"
		end_dd = "31"
	case "37": // Quadrimester 1 (4 months in duration)
		start_mm = "01"
		start_dd = "01"
		end_mm = "04"
		end_dd = "30"
	case "38": // Quadrimester 2 (4 months in duration)
		start_mm = "05"
		start_dd = "01"
		end_mm = "08"
		end_dd = "31"
	case "39": // Quadrimester 3 (4 months in duration)
		start_mm = "09"
		start_dd = "01"
		end_mm = "12"
		end_dd = "31"
	case "40": // Semestral 1 (6 months in duration)
		start_mm = "01"
		start_dd = "01"
		end_mm = "06"
		end_dd = "30"
	case "41": // Semestral 2 (6 months in duration)
		start_mm = "07"
		start_dd = "01"
		end_mm = "12"
		end_dd = "31"
	default:
		return nil, edtf.Invalid(SUB_YEAR_GROUPINGS, edtf_str)
	}

	start := fmt.Sprintf("%s-%s-%s", start_yyyy, start_mm, start_dd)
	end := fmt.Sprintf("%s-%s", end_yyyy, end_mm)

	if end_dd != "" {
		end = fmt.Sprintf("%s-%s", end, end_dd)
	}

	_str := fmt.Sprintf("%s/%s", start, end)

	sp, err := common.DateSpanFromEDTF(_str)

	if err != nil {
		return nil, err
	}

	d := &edtf.EDTFDate{
		Start:   sp.Start,
		End:     sp.End,
		EDTF:    edtf_str,
		Level:   LEVEL,
		Feature: SUB_YEAR_GROUPINGS,
	}

	return d, nil
}
