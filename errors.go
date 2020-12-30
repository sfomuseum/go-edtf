package edtf

import (
	"fmt"
)

type NotImplementedError struct {
	edtf_str string
	label    string
}

func (e *NotImplementedError) Error() string {
	return fmt.Sprintf("Not implemented '%s' (%s)", e.edtf_str, e.label)
}

func NotImplemented(label string, edtf_str string) error {
	return &NotImplementedError{
		edtf_str: edtf_str,
		label:    label,
	}
}

func IsNotImplemented(e error) bool {

	switch e.(type) {
	case *NotImplementedError:
		return true
	default:
		return false
	}
}

type InvalidError struct {
	edtf_str string
	label    string
}

func (e *InvalidError) Error() string {
	return fmt.Sprintf("Invalid EDTF string '%s' (%s)", e.edtf_str, e.label)
}

func Invalid(label string, edtf_str string) error {
	return &InvalidError{
		edtf_str: edtf_str,
		label:    label,
	}
}

func IsInvalid(e error) bool {

	switch e.(type) {
	case *InvalidError:
		return true
	default:
		return false
	}
}

type UnsupportedError struct {
	edtf_str string
	label    string
}

func (e *UnsupportedError) Error() string {
	return fmt.Sprintf("Unsupported EDTF string '%s' (%s)", e.edtf_str, e.label)
}

func Unsupported(label string, edtf_str string) error {
	return &UnsupportedError{
		edtf_str: edtf_str,
		label:    label,
	}
}

func IsUnsupported(e error) bool {

	switch e.(type) {
	case *UnsupportedError:
		return true
	default:
		return false
	}
}
