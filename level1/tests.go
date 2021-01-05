package level1

var Tests map[string][]string = map[string][]string{
	LETTER_PREFIXED_CALENDAR_YEAR: []string{
		"Y170000002", // TO DO
		"Y-17000002", // TO DO
		"Y1700",
		"Y-1200",
	},
	SEASON: []string{
		"2001-01",
		"2001-24",
		// "-0011-24",	// TO DO
		// "-0301-05",	// TO DO
		"Spring, 2002",
		"winter, 2002",
		// "Summer, -1980",	// TO DO
	},
	QUALIFIED_DATE: []string{
		"1984?",
		"2004-06~",
		"2004-06-11%",
	},
	UNSPECIFIED_DIGITS: []string{
		"201X",
		"20XX",
		"2004-XX",
		"1985-04-XX",
		"1985-XX-XX",
	},
	EXTENDED_INTERVAL_START: []string{
		"../1985-04-12",
		"../1985-04",
		"../1985",
		"/1985-04-12",
		"/1985-04",
		"/1985",
	},
	EXTENDED_INTERVAL_END: []string{
		"1985-04-12/..",
		"1985-04/..",
		"1985/..",
		"1985-04-12/",
		"1985-04/",
		"1985/",
	},
	NEGATIVE_CALENDAR_YEAR: []string{
		"-1985",
	},
}
