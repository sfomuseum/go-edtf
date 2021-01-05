package level1

import (
	"testing"
)

func TestQualifiedDate(t *testing.T) {

	valid, ok := Tests[QUALIFIED_DATE]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, _ := range valid {

		_, err := ParseQualifiedDate(input)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", input, err)
		}
	}

}
