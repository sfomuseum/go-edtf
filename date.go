package edtf

import (
	"fmt"
	"time"
)

type Date struct {
	DateTime    string     `json:"datetime,omitempty"`
	Timestamp   *Timestamp `json:"timestamp,omitempty"`
	YMD         *YMD       `json:"ymd"`
	Uncertain   Precision  `json:"uncertain,omitempty"`
	Approximate Precision  `json:"approximate,omitempty"`
	Unspecified Precision  `json:"unspecified,omitempty"`
	Precision   Precision  `json:"precision,omitempty"`
	Open        bool       `json:"open,omitempty"`
	Unknown     bool       `json:"unknown,omitempty"`
	Inclusivity Precision  `json:"inclusivity,omitempty"`
}

func (d *Date) SetTime(t *time.Time) {
	d.DateTime = t.Format(time.RFC3339)
	d.Timestamp = NewTimestampWithTime(t)
}

func (d *Date) String() string {
	return fmt.Sprintf("[[%T] Time: '%v' YMD: '%v']", d, d.Timestamp, d.YMD)
}
