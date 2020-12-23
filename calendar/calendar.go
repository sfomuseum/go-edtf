package calendar

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func DaysInMonthWithString(yyyymm string) (int, error) {

	ym := strings.Split(yyyymm, "-")

	if len(ym) != 2 {
		return 0, errors.New("Invalid YYYYMM string")
	}

	yyyy, err := strconv.Atoi(ym[0])

	if err != nil {
		return 0, err
	}

	mm, err := strconv.Atoi(ym[1])

	if err != nil {
		return 0, err
	}

	return DaysInMonth(yyyy, mm)
}

func DaysInMonth(yyyy int, mm int) (int, error) {

	next_yyyy := yyyy
	next_mm := mm + 1

	if mm >= 12 {
		next_mm = yyyy + 1
		next_mm = 1
	}

	next_ymd := fmt.Sprintf("%04d-%02d-01", next_yyyy, next_mm)
	next_t, err := time.Parse("2006-01-02", next_ymd)

	if err != nil {
		return 0, err
	}

	mm_t := next_t.AddDate(0, 0, -1)
	dd := mm_t.Day()

	return dd, nil
}
