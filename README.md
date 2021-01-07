# go-edtf

A Go package for parsing Extended DateTime Format (EDTF) date strings.

## Important

This is a work in progress and documentation is incomplete.

Code to parse Level 0, Level 1 and Level 2 strings has been implemented including tests.

This is almost ready to bless with an initial release but things might still change before that happens.

## Background

The following is taken from the [EDTF website](https://www.loc.gov/standards/datetime/background.html):

> EDTF defines features to be supported in a date/time string, features considered useful for a wide variety of applications

> Date and time formats are specified in ISO 8601, the International Standard for the representation of dates and times. ISO 8601-2004 provided basic date and time formats; these were not sufficiently expressive to support various semantic qualifiers and concepts than many applications find useful. For example, although it could express the concept "the year 1984", it could not express "approximately the year 1984", or "we think the year is 1984 but we're not certain". These as well as various other concepts had therefore often been represented using ad hoc conventions; EDTF provides a standard syntax for their representation.

> Further, 8601 is a complex specification describing a large number of date/time formats, in many cases providing multiple options for a given format. Thus a second aim of EDTF is to restrict the supported formats to a smaller set.

> EDTF functionality has now been integrated into ISO 8601-2019, the latest revision of ISO 8601, published in March 2019.

> EDTF was developed over the course of several years by a community of interested parties, and a draft specification was published in 2012. The draft specification is no longer publicly, readily available, because its availability has caused confusion with the official version.

## Nomenclature

### Date span

```
The full extent of something from end to end; the amount of space that something covers:
```

An `edtf.EDTFDate` instance encompasses a date span in the form of `Start` and `End` properties which are themselves `edtf.DateRange` instances.

### Date range

```
The area of variation between upper and lower limits on a particular scale
```

An `edtf.DateRange` instance encompasses upper and lower dates (for an EDTF string) in the form of `Lower` and `Upper` properties which are themselves `edtf.Date` instances.

### Date

_To be written_

## Tests

Tests are defined and handled in (3) places:

* In every `level(N)` package there are individual `_test.go` files for each feature.
* In every `level(N)` package there is a `tests.go` file that defines input values and expected response values defined as `tests.TestResult` instances.
* The `tests.TestResult` instance, its options and its methods are defined in the `tests` package. It implements a `TestDate` method that most of the individual `_test.go` files invoke.

## See also

* http://www.loc.gov/standards/datetime/
* https://www.iso.org/standard/70907.html (ISO 8601-1:2019)
* https://www.iso.org/standard/70908.html (ISO 8601-2:2019)

### Related

* https://github.com/sjansen/edtf
* https://github.com/unt-libraries/edtf-validate