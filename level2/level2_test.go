package level2

import (
	"testing"
)

func TestIsLevel2(t *testing.T) {

	valid := make([]string, 0)

	for _, candidates := range Tests {

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

	for _, candidates := range Tests {

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

func TestSetRepresentations(t *testing.T) {

	valid, ok := Tests[SET_REPRESENTATIONS]

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

	valid, ok := Tests[GROUP_QUALIFICATION]

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

	valid, ok := Tests[INDIVIDUAL_QUALIFICATION]

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

	valid, ok := Tests[UNSPECIFIED_DIGIT]

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

	valid, ok := Tests[INTERVAL]

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
