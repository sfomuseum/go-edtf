package level2

import (
	"testing"
)

func TestExponentialYear(t *testing.T) {

	valid, ok := Tests["exponential_year"]

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
