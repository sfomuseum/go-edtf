package common

import (
	"errors"
	"strconv"
)

func YMDFromStrings(str_yyyy string, str_mm string, str_dd string) (int, int, int, error) {


	if str_yyyy == "" {
		return 0, 0, 0, errors.New("Missing year")
	}

	if str_mm == "" && str_dd != "" {
		return 0, 0, 0, errors.New("Missing month")
	}

	yyyy, err := strconv.Atoi(str_yyyy)

	if err != nil {
		return 0, 0, 0, err
	}

	mm := 1
	dd := 1

	if str_mm != "" {

		m, err := strconv.Atoi(str_mm)

		if err != nil {
			return 0, 0, 0, err
		}

		mm = m
	}

	if str_dd != "" {

		d, err := strconv.Atoi(str_dd)

		if err != nil {
			return 0, 0, 0, err
		}

		dd = d
	}

	return yyyy, mm, dd, nil
}
