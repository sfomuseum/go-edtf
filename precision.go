package edtf

import ()

const (
	NONE      Precision = 0
	ALL       Precision = 1 << iota // 2
	ANY                             // 4
	DAY                             // 8
	WEEK                            // 16
	MONTH                           // 32
	YEAR                            // 64
	DECADE                          // 128
	CENTURY                         // 256
	MILLENIUM                       // 512
)

// https://stackoverflow.com/questions/48050522/using-bitsets-in-golang-to-represent-capabilities

type Precision uint32

func (f Precision) HasFlag(flag Precision) bool { return f&flag != 0 }
func (f *Precision) AddFlag(flag Precision)     { *f |= flag }
func (f *Precision) ClearFlag(flag Precision)   { *f &= ^flag }
func (f *Precision) ToggleFlag(flag Precision)  { *f ^= flag }

func (f *Precision) IsAnnual() bool {
	return f.HasFlag(YEAR)
}

func (f *Precision) IsMonthly() bool {
	return f.HasFlag(MONTH)
}

func (f *Precision) IsDaily() bool {
	return f.HasFlag(DAY)
}
