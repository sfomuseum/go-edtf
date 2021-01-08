package edtf

import (
	"fmt"
)

type DateRange struct {
	EDTF  string `json:"edtf"`
	Lower *Date  `json:"lower"`
	Upper *Date  `json:"upper"`
}

func (r *DateRange) String() string {
	return fmt.Sprintf("[[%T] Lower: '%v' Upper: '%v'[", r, r.Lower, r.Upper)
}
