package re

// Common

const PATTERN_YEAR string = `(\-?\d{4})`

// Level 0

const PATTERN_DATE string = `(\-?\d{4})(?:-([0][1-9]|1[0-2])(?:-([012][1-9]|3[01]))?)?`

const PATTERN_DATE_AND_TIME string = PATTERN_DATE + `T(\d{2}):(\d{2}):(\d{2})(Z|(\+|-)(\d{2})(\:(\d{2}))?)?`

const PATTERN_TIME_INTERVAL string = PATTERN_DATE + `/` + PATTERN_DATE

// Level 1

const PATTERN_LETTER_PREFIXED_CALENDAR_YEAR string = `Y(\-?\d+)`

const PATTERN_SEASON string = PATTERN_YEAR + `\-(0[1-9]|1[0-2]|2[1-4])|(?i)(spring|summer|fall|winter)\s*,\s*(\d{4})`

const PATTERN_QUALIFIED_DATE string = PATTERN_DATE + `(\?|~|%)`

const PATTERN_UNSPECIFIED_DIGITS string = `(?:([0-9X]{4})(?:-([0X][1-9X]|[1X][0-2X])(?:-([012X][1-9X]|[3X][01X]))?)?)`

const PATTERN_INTERVAL_START = `(\.\.)?\/` + PATTERN_DATE

const PATTERN_INTERVAL_END = PATTERN_DATE + `\/(\.\.)?`

const PATTERN_NEGATIVE_YEAR = `\-` + PATTERN_YEAR

// Level 2

const PATTERN_EXPONENTIAL_YEAR string = `(?i)Y((\-)?(\d+)E(\d+))`

const PATTERN_SIGNIFICANT_DIGITS string = `(?:` + PATTERN_YEAR + `|` + PATTERN_LETTER_PREFIXED_CALENDAR_YEAR + `|` + PATTERN_EXPONENTIAL_YEAR + `)S(\d+)`

const PATTERN_SUB_YEAR string = `(\d{4})\-(1[0-2]|2[1-9]|3[0-9]|4[0-1])`

const PATTERN_SET_REPRESENTATIONS string = `(\[|\{)(\.\.)?(?:(?:(\d{4})(?:-(\d{2})(?:-(\d{2}))?)?)(,|\.\.)?)+(\.\.)?(\}|\])`

const PATTERN_GROUP_QUALIFICATION string = `(?:(\d{4})(%|~|\?)?(?:-(\d{2})(%|~|\?)?(?:-(\d{2})(%|~|\?)?)?)?)`

const PATTERN_INDIVIDUAL_QUALIFICATION string = `(?:(%|~|\?)?(\d{4})(?:-(%|~|\?)?(\d{2})(?:-(%|~|\?)?(\d{2}))?)?)`

const PATTERN_UNSPECIFIED_DIGIT string = `([0-9X]{4})(?:-([0-9X]{2})(?:-([0-9X]{2}))?)?`

const PATTERN_INTERVAL string = `(%|~|\?)?([0-9X]{4})(?:-(%|~|\?)?([0-9X]{2})(?:-(%|~|\?)?([0-9X]{2}))?)?\/(%|~|\?)?([0-9X]{4})(?:-(%|~|\?)?([0-9X]{2})(?:-(%|~|\?)?([0-9X]{2}))?)?`
