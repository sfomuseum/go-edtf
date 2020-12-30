package level2

import (
	"github.com/whosonfirst/go-edtf"
	"testing"
)

func TestInterval(t *testing.T) {

	valid, ok := Tests[INTERVAL]

	if !ok {
		t.Fatalf("Failed to load test strings")
	}

	for _, str := range valid {

		_, err := ParseInterval(str)

		if err != nil {

			if edtf.IsNotImplemented(err) {
				t.Logf("Failed to parse '%s', %v", str, err)
			} else {
				t.Fatalf("Failed to parse '%s', %v", str, err)
			}
		}
	}
}
