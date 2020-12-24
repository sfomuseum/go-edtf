package level0

import (
	"testing"
	"time"
)

var test_strings map[string][]string

func init() {

	test_strings = map[string][]string{
		"date": []string{
			"1985-04-12",
			"1985-04",
			"1985",
		},
		"date_time": []string{
			"1985-04-12T23:20:30",
			"1985-04-12T23:20:30Z",
			"1985-04-12T23:20:30-04",
			"1985-04-12T23:20:30+04:30",
		},
		"time_interval": []string{
			"1964/2008",
			"2004-06/2006-08",
			"2004-02-01/2005-02-08",
			"2004-02-01/2005-02",
			"2004-02-01/2005",
			"2005/2006-02",
		},
	}
}

func TestIsLevel0(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range test_strings {

		for _, str := range candidates {
			valid = append(valid, str)
		}
	}

	invalid := []string{
		"Dec. 1970",
		"Last week",
		"Y170000002",
		"c 1960",
	}

	for _, str := range valid {

		ok := IsLevel0(str)

		if !ok {
			t.Fatalf("Expected '%s' to parse as Level 0 string", str)
		}

		ParseString(str)
	}

	for _, str := range invalid {

		ok := IsLevel0(str)

		if ok {
			t.Fatalf("Expected '%s' to not parse as Level 0 string", str)
		}
	}

}

func TestParseString(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range test_strings {

		for _, str := range candidates {
			valid = append(valid, str)
		}
	}

	for _, str := range valid {

		d, err := ParseString(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
		}

	}
}

func TestParseDate(t *testing.T) {

	valid, ok := test_strings["date"]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for _, str := range valid {

		d, err := ParseDate(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
		}
	}

	dt, err := ParseDate("2020-02")

	if err != nil {
		t.Fatal("Failed to parse '2020-02'")
	}

	tm := dt.Upper.Upper.Time

	tm_str := tm.Format(time.RFC3339)
	expected_str := "2020-02-29T23:59:59Z"

	if tm_str != expected_str {
		t.Fatalf("Invalid time string. Expected '%s' but got '%s'", expected_str, tm_str)
	}

}

func TestParseDateTime(t *testing.T) {

	valid, ok := test_strings["date_time"]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for _, str := range valid {

		d, err := ParseDateTime(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
		}
	}
}

func TestParseTimeInterval(t *testing.T) {

	valid, ok := test_strings["time_interval"]

	if !ok {
		t.Fatal("Failed to load test strings")
	}

	for _, str := range valid {

		d, err := ParseTimeInterval(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		if d.String() != str {
			t.Fatalf("Failed to stringify '%s', %v", str, err)
		}
	}
}
