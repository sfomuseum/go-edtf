# go-edtf

A Go package for parsing Extended DateTime Format (EDTF) dates.

## Important

This is a work in progress and does not work yet. It is, currently, a naive parser using regular expressions. Help is welcome :D

## Background

The following is taken from the [EDTF website](https://www.loc.gov/standards/datetime/background.html):

> EDTF defines features to be supported in a date/time string, features considered useful for a wide variety of applications

> Date and time formats are specified in ISO 8601, the International Standard for the representation of dates and times. ISO 8601-2004 provided basic date and time formats; these were not sufficiently expressive to support various semantic qualifiers and concepts than many applications find useful. For example, although it could express the concept "the year 1984", it could not express "approximately the year 1984", or "we think the year is 1984 but we're not certain". These as well as various other concepts had therefore often been represented using ad hoc conventions; EDTF provides a standard syntax for their representation.

> Further, 8601 is a complex specification describing a large number of date/time formats, in many cases providing multiple options for a given format. Thus a second aim of EDTF is to restrict the supported formats to a smaller set.

> EDTF functionality has now been integrated into ISO 8601-2019, the latest revision of ISO 8601, published in March 2019.

> EDTF was developed over the course of several years by a community of interested parties, and a draft specification was published in 2012. The draft specification is no longer publicly, readily available, because its availability has caused confusion with the official version.

## See also

* http://www.loc.gov/standards/datetime/
* https://www.iso.org/standard/70907.html (ISO 8601-1:2019)
* https://www.iso.org/standard/70908.html (ISO 8601-2:2019)

### Related

* https://github.com/sjansen/edtf
* https://github.com/unt-libraries/edtf-validate