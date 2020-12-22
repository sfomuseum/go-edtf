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
