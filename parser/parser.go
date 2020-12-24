package parser

import (
	"errors"
	"github.com/whosonfirst/go-edtf"
	"github.com/whosonfirst/go-edtf/level0"
	"github.com/whosonfirst/go-edtf/level1"
	"github.com/whosonfirst/go-edtf/level2"
)

func ParseString(edtf_str string) (*edtf.EDTFDate, error) {

	if level0.IsLevel0(edtf_str) {
		return level0.ParseString(edtf_str)
	}

	if level1.IsLevel1(edtf_str) {
		return level1.ParseString(edtf_str)
	}

	if level2.IsLevel2(edtf_str) {
		return level2.ParseString(edtf_str)
	}

	return nil, errors.New("Invalid or unsupported EDTF string")
}
