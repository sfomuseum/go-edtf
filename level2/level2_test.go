package level2

import (
	"testing"
)

func TestExponentialYear(t *testing.T) {

	valid := []string{
		"Y-17E7",
	}

	for _, str := range valid {

		_, err := ParseExponentialYear(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestSignificantDigits(t *testing.T) {

	valid := []string{
		"1950S2",
		"Y171010000S3",
		"Y3388E2S3",
	}

	for _, str := range valid {

		_, err := ParseSignificantDigits(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestSubYearGroupings(t *testing.T) {

	valid := []string{
		"2001-34",
		// "second quarter of 2001"
	}

	for _, str := range valid {

		_, err := ParseSubYearGroupings(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestSetRepresentations(t *testing.T) {

	valid := []string{
		"[1667,1668,1670..1672]",
		"[..1760-12-03]",
		"[1760-12..]",
		"[1760-01,1760-02,1760-12..]",
		"[1667,1760-12]",
		"[..1984]",
		"{1667,1668,1670..1672}",
		"{1960,1961-12}",
		"{..1984}",
	}

	for _, str := range valid {

		_, err := ParseSetRepresentations(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestGroupQualification(t *testing.T) {

	valid := []string{
		"2004-06-11%",
		"2004-06~-11",
		"2004?-06-11",
	}

	for _, str := range valid {

		_, err := ParseGroupQualification(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestIndividualQualification(t *testing.T) {

	valid := []string{
		"?2004-06-~11",
		"2004-%06-11",
	}

	for _, str := range valid {

		_, err := ParseIndividualQualification(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestUnspecifiedDigit(t *testing.T) {

	valid := []string{
		"156X-12-25",
		"15XX-12-25",
		"XXXX-12-XX",
		"1XXX-XX",
		"1XXX-12",
		"1984-1X",
	}

	for _, str := range valid {

		_, err := ParseUnspecifiedDigit(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}

func TestInterval(t *testing.T) {

	valid := []string{
		"2004-06-~01/2004-06-~20",
		"2004-06-XX/2004-07-03",
	}

	for _, str := range valid {

		_, err := ParseInterval(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}
}
