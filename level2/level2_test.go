package level2

import (
	"testing"
)

var test_strings map[string][]string

func init() {

	test_strings = map[string][]string{
		"exponential_year": []string{
			"Y-17E7",
		},
		"significant_digits": []string{
			"1950S2",
			"Y171010000S3",
			"Y3388E2S3",
		},
		"sub_year_groupings": []string{
			"2001-34",
			// "second quarter of 2001"
		},
		"set_representations": []string{
			"[1667,1668,1670..1672]",
			"[..1760-12-03]",
			"[1760-12..]",
			"[1760-01,1760-02,1760-12..]",
			"[1667,1760-12]",
			"[..1984]",
			"{1667,1668,1670..1672}",
			"{1960,1961-12}",
			"{..1984}",
		},
		"group_qualification": []string{
			"2004-06-11%",
			"2004-06~-11",
			"2004?-06-11",
		},
		"individual_qualification": []string{
			"?2004-06-~11",
			"2004-%06-11",
		},
		"unspecified_digit": []string{
			"156X-12-25",
			"15XX-12-25",
			"XXXX-12-XX",
			"1XXX-XX",
			"1XXX-12",
			"1984-1X",
		},
		"interval": []string{
			"2004-06-~01/2004-06-~20",
			"2004-06-XX/2004-07-03",
		},
	}

}

func TestIsLevel2(t *testing.T) {

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

		ok := IsLevel2(str)

		if !ok {
			t.Fatalf("Expected '%s' to parse as Level 2 string", str)
		}

		ParseString(str)
	}

	for _, str := range invalid {

		ok := IsLevel2(str)

		if ok {
			t.Fatalf("Expected '%s' to not parse as Level 2 string", str)
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

		_, err := ParseString(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}

		/*
			if d.String() != str {
				t.Fatalf("Failed to stringify '%s', %v", str, err)
			}
		*/

	}
}

func TestExponentialYear(t *testing.T) {

	valid, ok := test_strings["exponential_year"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseExponentialYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestSignificantDigits(t *testing.T) {

	valid, ok := test_strings["significant_digits"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseSignificantDigits(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestSubYearGroupings(t *testing.T) {

	valid, ok := test_strings["sub_year_groupings"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseSubYearGroupings(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestSetRepresentations(t *testing.T) {

	valid, ok := test_strings["set_representations"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseSetRepresentations(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestGroupQualification(t *testing.T) {

	valid, ok := test_strings["group_qualification"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseGroupQualification(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestIndividualQualification(t *testing.T) {

	valid, ok := test_strings["individual_qualification"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseIndividualQualification(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestUnspecifiedDigit(t *testing.T) {

	valid, ok := test_strings["unspecified_digit"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseUnspecifiedDigit(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestInterval(t *testing.T) {

	valid, ok := test_strings["interval"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseInterval(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}
