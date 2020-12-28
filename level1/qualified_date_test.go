package level1

import (
	"testing"
)

func TestQualifiedDate(t *testing.T) {

	valid, ok := Tests["qualified_date"]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseQualifiedDate(str)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", str, err)
		}
	}

}
