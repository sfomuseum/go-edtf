package edtf

import (
	"fmt"
)

var deprecated map[string]string

func init() {

	deprecated = map[string]string{
		OPEN_2012:        OPEN,
		UNSPECIFIED_2012: UNSPECIFIED,
	}

}

// IsDeprecated returns a boolean flag indicating whether 'str' is considered a deprecated EDTF value.
func IsDeprecated(str string) bool {

	for test, _ := range deprecated {

		if str == test {
			return true
		}
	}

	return false
}

// ReplaceDeprecated returns the current value for 'old'.
func ReplaceDeprecated(old string) (string, error) {

	new, ok := deprecated[old]

	if !ok {
		err := fmt.Errorf("Unknown or unsupported EDTF string '%s' : %v", old, deprecated)
		return "", err
	}

	return new, nil
}
