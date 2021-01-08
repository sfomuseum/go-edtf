package parser

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/level0"
	"github.com/sfomuseum/go-edtf/level1"
	"github.com/sfomuseum/go-edtf/level2"
)

func IsValid(edtf_str string) bool {

	if level0.IsLevel0(edtf_str) {
		return true
	}

	if level1.IsLevel1(edtf_str) {
		return true
	}

	if level2.IsLevel2(edtf_str) {
		return true
	}

	return false
}

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

	return nil, edtf.Unrecognized("Invalid or unsupported EDTF string", edtf_str)
}

func Matches(edtf_str string) (int, string, error) {

	if level0.IsLevel0(edtf_str) {

		feature, err := level0.Matches(edtf_str)

		if err != nil {
			return -1, "", err
		}

		return level0.LEVEL, feature, nil
	}

	if level1.IsLevel1(edtf_str) {

		feature, err := level1.Matches(edtf_str)

		if err != nil {
			return -1, "", err
		}

		return level1.LEVEL, feature, nil
	}

	if level2.IsLevel2(edtf_str) {

		feature, err := level2.Matches(edtf_str)

		if err != nil {
			return -1, "", err
		}

		return level2.LEVEL, feature, nil
	}

	return -1, "", edtf.Unrecognized("Invalid or unsupported EDTF string", edtf_str)
}
