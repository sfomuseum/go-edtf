package tests

type TestResult struct {
	Start *DateRange
	End   *DateRange
}

type DateRange struct {
	Lower *Date
	Upper *Date
}

type Date struct {
	Time string
}
