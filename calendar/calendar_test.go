package calendar

import (
	"testing"
)

func TestDaysInMonth(t *testing.T) {

	tests := map[string]int{
		"2019-02": 28,
		"2020-02": 29,
		"1985-06": 30,
		"1950-12": 31,
	}

	for str_ym, expected_days := range tests {

		t.Logf("DaysInMonthWithString '%s'", str_ym)
			
		days, err := DaysInMonthWithString(str_ym)

		if err != nil {
			t.Logf("Failed to determine days in month for '%s', %v", str_ym, err)
			t.Fail()
			continue
		}

		if days != expected_days {
			t.Logf("Failed to determined days in month for '%s', expected '%d' but got '%d'", str_ym, expected_days, days)
			t.Fail()
			continue
		}
	}
}
