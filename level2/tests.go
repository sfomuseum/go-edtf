package level2

var Tests map[string][]string = map[string][]string{
	EXPONENTIAL_YEAR: []string{
		"Y-17E7", // TO DO - https://github.com/whosonfirst/go-edtf/issues/5
		"Y10E7",  // TO DO
		"Y20E2",
	},
	SIGNIFICANT_DIGITS: []string{
		"1950S2",
		"Y171010000S3",
		"Y-20E2S3",
		"Y3388E2S3",
		"Y-20E2S3",
	},
	SUB_YEAR_GROUPINGS: []string{
		"2001-34",
		// "second quarter of 2001",	// TO DO
	},
	SET_REPRESENTATIONS: []string{
		"[1667,1668,1670..1672]",
		"[..1760-12-03]",
		"[1760-12..]",
		"[1760-01,1760-02,1760-12..]",
		"[1667,1760-12]",
		"[..1984]",
		"{1667,1668,1670..1672}",
		"{1960,1961-12}",
		"{..1984}",
	},
	GROUP_QUALIFICATION: []string{
		"2004-06-11%",
		"2004-06~-11",
		"2004?-06-11",
	},
	INDIVIDUAL_QUALIFICATION: []string{
		"?2004-06-~11",
		"2004-%06-11",
	},
	UNSPECIFIED_DIGIT: []string{
		"156X-12-25",
		"15XX-12-25",
		// "XXXX-12-XX",	// TO DO
		"1XXX-XX",
		"1XXX-12",
		"1984-1X",
	},
	INTERVAL: []string{
		"2004-06-~01/2004-06-~20",
		"2004-06-XX/2004-07-03",
	},
}
