package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestSetRepresentations(t *testing.T) {

	valid, ok := Tests[SET_REPRESENTATIONS]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for input, _ := range valid {

		_, err := ParseSetRepresentations(input)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Skipping '%s', %v", input, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", input, err)
			}
		}
	}
}
