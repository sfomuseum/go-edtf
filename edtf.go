package edtf

import (
	"time"
)

type EDTFDate struct {
	Upper DateRange
	Lower DateRange
	Raw   string
	Level int
}

type DateRange struct {
	Upper Date
	Lower Date
}

type Date struct {
	String      string
	Time        time.Time
	Uncertain   bool
	Approximate bool
	Unknown     bool
	Open        bool
	Precision   string
}
