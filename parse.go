package edtf

import (
	"errors"
	"github.com/whosonfirst/go-edtf/level0"
	_ "github.com/whosonfirst/go-edtf/level1"
	_ "github.com/whosonfirst/go-edtf/level2"	
)

func ParseString(edtf_str string) (*EDTFDate, error) {

	if level0.IsLevel0(edtf_str){
		return level0.ParseLevel0(edtf_str)
	}

	return nil, errors.New("Invalid or unsupported EDTF string")
}
