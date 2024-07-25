package between

import (
	"fmt"

	"github.com/sfomuseum/go-edtf"	
	"github.com/sfomuseum/go-edtf/parser"
)

func IsBetweenString(d_str string, inception_str string, cessation_str string) (bool, error) {

	d, err := parser.ParseString(d_str)

	if err != nil {
		return false, err
	}

	inception_d, err := parser.ParseString(inception_str)

	if err != nil {
		return false, err
	}

	cessation_d, err := parser.ParseString(cessation_str)

	if err != nil {
		return false, err
	}
	
	return IsBetween(d, inception_d, cessation_d)
}

func IsBetween(d *edtf.EDTFDate, inception_d *edtf.EDTFDate, cessation_d *edtf.EDTFDate) (bool, error) {
	return false, fmt.Errorf("Not implemented")
}
