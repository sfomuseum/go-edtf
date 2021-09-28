package cmp

import (
	"fmt"
	"github.com/sfomuseum/go-edtf/parser"
)

// IsBefore reports whether the EDTF string `this_d` is before `that_d`.
func IsBefore(this_d string, that_d string) (bool, error) {

	this_dt, err := parser.ParseString(this_d)

	if err != nil {
		return false, fmt.Errorf("Failed to parse '%s', %v", this_d, err)
	}

	that_dt, err := parser.ParseString(that_d)

	if err != nil {
		return false, fmt.Errorf("Failed to parse '%s', %v", that_d, err)
	}

	return this_dt.Before(that_dt)
}
