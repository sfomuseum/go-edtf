package edtf

import (
	"fmt"
)

type DateSpan struct {
	Start *DateRange `json:"start"`
	End   *DateRange `json:"end"`
}

func (s *DateSpan) String() string {
	return fmt.Sprintf("[[%T] Start: '%v' End: '%v']", s, s.Start, s.End)
}
