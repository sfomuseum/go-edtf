// package parser provides methods for parsing and validating EDTF strings.
package parser

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/common"
	"github.com/sfomuseum/go-edtf/level0"
	"github.com/sfomuseum/go-edtf/level1"
	"github.com/sfomuseum/go-edtf/level2"
	_ "log"
)

// Return a boolean value indicating whether a string is a valid EDTF date.
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

	switch edtf_str {
	case edtf.OPEN, edtf.UNKNOWN:
		return true
	default:
		return false
	}
}

// Parse a string in to an edtf.EDTFDate instance.
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

	if edtf_str == edtf.OPEN {
		sp := common.OpenDateSpan()

		d := &edtf.EDTFDate{
			Start:   sp.Start,
			End:     sp.End,
			EDTF:    edtf_str,
			Level:   -1,
			Feature: "Open",
		}

		return d, nil
	}

	if edtf_str == edtf.UNKNOWN {

		sp := common.UnknownDateSpan()

		d := &edtf.EDTFDate{
			Start:   sp.Start,
			End:     sp.End,
			EDTF:    edtf_str,
			Level:   -1,
			Feature: "Unknown",
		}

		return d, nil
	}

	return nil, edtf.Unrecognized("Invalid or unsupported EDTF string", edtf_str)
}

// Determine which EDTF level and corresponding EDTF feature a string matches.
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
