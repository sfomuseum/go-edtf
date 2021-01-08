# go-edtf

A Go package for parsing Extended DateTime Format (EDTF) date strings. It is compliant with Level 0, Level 1 and Level 2 of the EDTF specification (2019).

* [Background](#background)
* [Features](#features)
* [Nomenclature and Type Definitions](#nomenclature-and-type-definitions)
* [Documentation](#documentation)
* [Example](#example)
* [Tools](#tools)
* [Tests](#tests)
* [Reporting Bugs and Issues](#reporting-bugs-and-issues)

## Background

The following is taken from the [EDTF website](https://www.loc.gov/standards/datetime/background.html):

> EDTF defines features to be supported in a date/time string, features considered useful for a wide variety of applications

> Date and time formats are specified in ISO 8601, the International Standard for the representation of dates and times. ISO 8601-2004 provided basic date and time formats; these were not sufficiently expressive to support various semantic qualifiers and concepts than many applications find useful. For example, although it could express the concept "the year 1984", it could not express "approximately the year 1984", or "we think the year is 1984 but we're not certain". These as well as various other concepts had therefore often been represented using ad hoc conventions; EDTF provides a standard syntax for their representation.

> Further, 8601 is a complex specification describing a large number of date/time formats, in many cases providing multiple options for a given format. Thus a second aim of EDTF is to restrict the supported formats to a smaller set.

> EDTF functionality has now been integrated into ISO 8601-2019, the latest revision of ISO 8601, published in March 2019.

> EDTF was developed over the course of several years by a community of interested parties, and a draft specification was published in 2012. The draft specification is no longer publicly, readily available, because its availability has caused confusion with the official version.

## Features

### Level 0

| Name | Implementation | Tests | Notes |
| --- | --- | --- | --- |
| [Date](https://www.loc.gov/standards/datetime/) | [yes](level0/date.go) | [yes](level0/date_test.go) | |
| [Date and Time](https://www.loc.gov/standards/datetime/) | [yes](level0/date_and_time.go) | [yes](level0/date_and_time_test.go) | |
| [Time Interval](https://www.loc.gov/standards/datetime/) | [yes](level0/time_interval.go) | [yes](level0/time_interval_test.go) | |

### Level 1

| Name | Implementation | Tests | Notes |
| --- | --- | --- | --- |
| [Letter-prefixed calendar year](https://www.loc.gov/standards/datetime/) | [yes](level1/letter_prefixed_calendar_year.go) | [yes](level1/letter_prefixed_calendar_year_test.go) | Calendar years greater (or less) than 9999 are not supported yet. |
| [Season](https://www.loc.gov/standards/datetime/) | [yes](level1/season.go) | [yes](level1/season_test.go) | |
| [Qualification of a date (complete)](https://www.loc.gov/standards/datetime/) | [yes](level1/qualified_date.go) | [yes](level1/qualified_date_test.go) | |
| [Unspecified digit(s) from the right](https://www.loc.gov/standards/datetime/) | [yes](level1/unspecified_digits.go) | [yes](level1/unspecified_digits_test.go) | |
| [Extended Interval (L1)](https://www.loc.gov/standards/datetime/) | [yes](level1/extended_interval.go) | [yes](level1/extended_interval_test.go) | |
| [Negative calendar year](https://www.loc.gov/standards/datetime/) | [yes](level1/negative_calendar_year.go) | [yes](level1/negative_calendar_year_test.go) | |

### Level 2

| Name | Implementation | Tests | Notes |
| --- | --- | --- | --- |
| [Exponential year](https://www.loc.gov/standards/datetime/) | [yes](level2/exponential_year.go) | [yes](level2/exponential_year_test.go) | Calendar years greater (or less) than 9999 are not supported yet. |
| [Significant digits](https://www.loc.gov/standards/datetime/) | [yes](level2/significant_digits.go) | [yes](level2/significant_digits_test.go) | |
| [Sub-year groupings](https://www.loc.gov/standards/datetime/) | [yes](level2/sub_year_grouping.go) | [yes](level2/sub_year_grouping_test.go) | Compound phrases, like "second quarter of 2001" are not supported yet. |
| [Set representation](https://www.loc.gov/standards/datetime/) | [yes](level2/set_representation.go) | [yes](level2/set_representation_test.go) | |
| [Qualification](https://www.loc.gov/standards/datetime/) | [yes](level2/qualification.go) | [yes](level2/qualification_test.go) | |
| [Unspecified Digit](https://www.loc.gov/standards/datetime/) | [yes](level2/unspecified_digit.go) | [yes](level2/unspecified_digit_test.go) | Years with a leading unspecified digit, for example "X999", are not supported yet |
| [Interval](https://www.loc.gov/standards/datetime/) | [yes](level2/interval.go) | [yes](level2/interval.go) | |

## Nomenclature and Type Definitions

### Date Spans (or `edtf.EDTFDate`)

The word `span` is defined as:

```
The full extent of something from end to end; the amount of space that something covers:
```

An `edtf.EDTFDate` instance is a struct that represents a date span in the form of `Start` and `End` properties which are themselves `edtf.DateRange` instances. It also contains properties denoting the EDTF feature and level associated with EDTF string used to create the instance.

```
type EDTFDate struct {
	Start   *DateRange `json:"start"`
	End     *DateRange `json:"end"`
	EDTF    string     `json:"edtf"`
	Level   int        `json:"level"`
	Feature string     `json:"feature"`
}
```

### Date ranges (or `edtf.DateRange`)

The word `range` is defined as:

```
The area of variation between upper and lower limits on a particular scale
```

A `edtf.DateRange` instance encompasses upper and lower dates (for an EDTF string). It is a struct with `Lower` and `Upper` properties which are themselves `edtf.Date` instances.

```   		    	     
type DateRange struct {
	EDTF  string `json:"edtf"`
	Lower *Date  `json:"lower"`
	Upper *Date  `json:"upper"`
}
```

### Date (or `edtf.Date`)

A `edtf.Date` instance is the upper or lower end of a date range. It is a struct that contains atomic date and time information as well as a number of flags denoting precision and other granularities defined in the EDTF specification.

```
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
```

#### Notes

* `DateTime` strings are encoded as `RFC3339` strings, in UTC.

### Timestamp (or `edtf.Timestamp`)

A `edtf.Timestamp` instance represents the Unix timestamp for the date. When serialized (or "marshal-ed") as JSON a `edtf.Timestamp` instance is encoded as a signed 64-bit integer. When a JSON-encoded `edtf.EDTFDate` is unserialized (or "unmarshal"-ed) the integer is converted to a `edtf.Timestamp` instance.

The `Timestamp` is defined as a custom struct rather than an integer so that when it is serialized there is no confusion about whether a value of "0" means "January 1, 1970" or "this property is missing". It means the former. If the property is `nil` then it should be excluded from the JSON representation entirely.

Because the Go language imposes limits on the minimum and maximum date it can represent (-9999 and 9999 respectively) this element _may_ be `nil`.

### YMD (or `edtf.YMD`)

A `edtf.YMD instance a struct containing numeric year, month and day properties. It is designed to supplement "time" elements or, in cases where a "time" element is not possible to replace it. 

```
type YMD struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}
```

### Precision (or `edtf.Precision`)

"Precision" a 32-bit integer (as well as a Go language `Precision` instance with its own method) that uses bitsets to represent granularity.

The following named granularities are defined as constants:

| Name | Value | Notes |
| --- | --- | --- |
| NONE | 0 | |
| ALL | 2 | |
| ANY | 4 | |
| DAY | 8 | |
| WEEK | 16 |
| MONTH | 32 | |
| YEAR | 64 | | 
| DECADE | 128 | |
| CENTURY | 256 | | 
| MILLENIUM | 512 | |

## Documentation

Proper Go language documentation [is incomplete](https://github.com/sfomuseum/go-edtf/issues/12) but can be viewed in its current state at: https://godoc.org/github.com/sfomuseum/go-edtf

## Example

```
package main

import (
	"flag"
	"github.com/sfomuseum/go-edtf/parser"
	"log"
)

func main() {

	flag.Parse()

	for _, raw := range flag.Args() {

		if !parser.IsValid(raw){
			continue
		}

		level, feature, _ := parser.Matches(raw)

		log.Printf("%s is a Level %d (%s) string\n", raw, level, feature)
		
		d, _ := parser.ParseString(raw)

		log.Printf("%s span %v to %v\n", d.Lower(), d.Upper())
	}
}
```

_Error handling removed for the sake of brevity._

## Tools

To build binary versions of these tools run the `cli` Makefile target. For example:

```
$> make cli
go build -mod vendor -o bin/parse cmd/parse/main.go
```

### parse

Parse one or more EDTF strings and return a list of JSON-encoded `edtf.EDTFDate` objects.

```
$> ./bin/parse -h
Parse one or more EDTF strings and return a list of JSON-encoded edtf.EDTFDate objects.
Usage:
	 ./bin/parse edtf_string(N) edtf_string(N)
```

For example:

```
$> ./bin/parse 2004-06-XX/2004-07-03 '{1667,1668,1670..1672}' | jq

[
  {
    "start": {
      "edtf": "2004-06-XX",
      "lower": {
        "datetime": "2004-06-01T00:00:00Z",
        "timestamp": 1086048000,
        "ymd": {
          "year": 2004,
          "month": 6,
          "day": 1
        },
        "precision": 32
      },
      "upper": {
        "datetime": "2004-06-30T23:59:59Z",
        "timestamp": 1088639999,
        "ymd": {
          "year": 2004,
          "month": 6,
          "day": 30
        },
        "precision": 32
      }
    },
    "end": {
      "edtf": "2004-07-03",
      "lower": {
        "datetime": "2004-07-03T00:00:00Z",
        "timestamp": 1088812800,
        "ymd": {
          "year": 2004,
          "month": 7,
          "day": 3
        },
        "precision": 64
      },
      "upper": {
        "datetime": "2004-07-03T23:59:59Z",
        "timestamp": 1088899199,
        "ymd": {
          "year": 2004,
          "month": 7,
          "day": 3
        },
        "precision": 64
      }
    },
    "edtf": "2004-06-XX/2004-07-03",
    "level": 2,
    "feature": "Interval"
  },
  {
    "start": {
      "edtf": "1667",
      "lower": {
        "datetime": "1667-01-01T00:00:00Z",
        "timestamp": -9561715200,
        "ymd": {
          "year": 1667,
          "month": 1,
          "day": 1
        },
        "precision": 64,
        "inclusivity": 2
      },
      "upper": {
        "datetime": "1667-12-31T23:59:59Z",
        "timestamp": -9530179201,
        "ymd": {
          "year": 1667,
          "month": 12,
          "day": 31
        },
        "precision": 64,
        "inclusivity": 2
      }
    },
    "end": {
      "edtf": "1672",
      "lower": {
        "datetime": "1672-01-01T00:00:00Z",
        "timestamp": -9403948800,
        "ymd": {
          "year": 1672,
          "month": 1,
          "day": 1
        },
        "precision": 64,
        "inclusivity": 2
      },
      "upper": {
        "datetime": "1672-12-31T23:59:59Z",
        "timestamp": -9372326401,
        "ymd": {
          "year": 1672,
          "month": 12,
          "day": 31
        },
        "precision": 64,
        "inclusivity": 2
      }
    },
    "edtf": "{1667,1668,1670..1672}",
    "level": 2,
    "feature": "Set representation"
  }
]
```

## Tests

Tests are defined and handled in (3) places:

* In every `level(N)` package there are individual `_test.go` files for each feature.
* In every `level(N)` package there is a `tests.go` file that defines input values and expected response values defined as `tests.TestResult` instances.
* The `tests.TestResult` instance, its options and its methods are defined in the `tests` package. It implements a `TestDate` method that most of the individual `_test.go` files invoke.

## Reporting Bugs and Issues

There might still be bugs, implementation gotchas or other issues. If you encounter any of these [please report them here](https://github.com/sfomuseum/go-edtf/issues).

## See also

* http://www.loc.gov/standards/datetime/
* https://www.iso.org/standard/70907.html (ISO 8601-1:2019)
* https://www.iso.org/standard/70908.html (ISO 8601-2:2019)

### Related

* https://github.com/sjansen/edtf
* https://github.com/unt-libraries/edtf-validate