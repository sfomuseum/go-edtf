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

	for _, str := range valid {

		_, err := ParseSetRepresentations(str)

		if err != nil {
			if edtf.IsNotImplemented(err) {
				t.Logf("Failed to parse '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}
