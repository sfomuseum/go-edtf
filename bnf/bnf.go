package bnf

/*

The syntax used in this BNF description is:

    = used for assignment
    | used for or
    ?    used for  zero or one
    *  used for  zero or more
    +  used for one or more
    Surrounding parentheses used for grouping, thus
        ( x | y)+
        means "one or more of x and y"
        (A B) | C
        means either A followed by B, or C (parentheses used to determine how "or" is evaluated)
    newline also used for grouping to reduce use of parentheses; thus
    A = B C
         | D

    means

    A = (B C) | D
    (*...*) used for comment e.g (* this is a comment *)

https://www.loc.gov/standards/datetime/pre-submission.html#bnf

*/

/*

dateTimeString = level0Expression | level1Expression | level2Expression

(* ************************** Level 0 *************************** *)

level0Expression = date | dateAndTime | L0Interval

(* ** date ** *)

date =  year | yearMonth | yearMonthDay

(* ** dateAndTime ** *)

dateAndTime = date "T" time

time = baseTime zoneOffset?

baseTime = hour ":" minute ":" second | "24:00:00"

          (* Zone *)

  zoneOffset = "Z"
      | ("+" | "-")
                     (zoneOffsetHour (":" minute)?
                     | "14:00"
                     | "00:" oneThru59 )

zoneOffsetHour = oneThru13

(* ** level 0 interval ** *)

L0Interval = date "/" date

(* ** Definition for year ** *)

year = positiveYear | negativeYear | "0000"

positiveYear =
       positiveDigit digit digit digit
     | digit positiveDigit digit digit
     | digit digit positiveDigit digit
     | digit digit digit positiveDigit

negativeYear = "-" positiveYear


(* ** Other Auxiliary Assignments for Level 0 ** *)


year = digit digit digit digit
month = oneThru12
monthDay =
      ("01" |"03" |"05" |"07" |"08" |"10" |"12") "-" oneThru31
      | ("04" |"06" |"09" |"11") "-" oneThru30
      | "02-" oneThru29
yearMonth = year "-" month
yearMonthDay =   year "-" monthDay
hour = zeroThru23
minute = zeroThru59
second =  zeroThru59

oneThru12 = "01" | "02" | "03" | "04" | "05" | "06" | "07" | "08" | "09" | "10" | "11" | "12"
oneThru13 = oneThru12 | "13"
oneThru23 = oneThru13 | "14" | "15" | "16" | "17" | "18" | "19" | "20" | "21" | "22" | "23"
zeroThru23 =  "00" | oneThru23
oneThru29 = oneThru23 | "24" | "25" | "26" | "27" | "28" | "29"
oneThru30 = oneThru29 | "30"
oneThru31 = oneThru30 | "31"
oneThru59 = oneThru31 | "32" | "33"| "34"| "35"| "36"| "37"| "38"| "39"| "40"| "41"| "42" | "43"
                | "44"| "45"| "46" | "47"| "48"| "49" |"50"|"51"|"52"|"53"|"54"|"55"|"56"|"57"|"58"|"59"
zeroThru59 = "00" | oneThru59

digit = positiveDigit | "0"
positiveDigit = "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"

day = oneThru31

(* ************************** Level 1 *************************** *)

level1Expression =  uncertainOrApproxDate
                        | unspecified
                        | L1Interval
                        | longYearSimple
                        | season

(* *** uncertainOrApproxDate *** *)

uncertainOrApproxDate =  date UASymbol


(* *** unspecified *** *)

unspecified =    yearWithOneOrTwoUnspecifedDigits
                   | monthUnspecified
                   | dayUnspecified
                   | dayAndMonthUnspecified

          yearWithOneOrTwoUnspecifedDigits = digit digit (digit|'u') 'u'
          monthUnspecified = year "-uu'
          dayUnspecified = yearMonth "-uu'
          dayAndMonthUnspecified = year "-uu-uu'



(* *** L1Interval *** *)

L1interval = L1Start "/" L1End

      L1Start = ( dateOrSeason UASymbol?) | "unknown"
      L1End   = L1Start | "open"



(* *** Long Year - Simple Form *** *)

longYearSimple = "y" "-"? positiveDigit digit digit digit digit+



(* *** Season (unqualified) *** *)

season = year "-" seasonNumber



(* ** Auxiliary Assignments for Level 1 ** *)

UASymbol = ("?" | "~" | "?~")
seasonNumber = "21" | "22" | "23" | "24"
dateOrSeason = date | season

(* ************************** Level 2 *************************** *)

level2Expression =   internalUncertainOrApproximate
                         | internalUnspecified
                         | choiceList
                         |inclusiveList
                         |maskedPrecision
                         | L2Interval
                         |longYearScientific
                         |seasonQualified

(* ** Internal Uncertain or Approximate** *)

internalUncertainOrApproximate =  IUABase | "(" IUABase ")" UASymbol

    IUABase =    year UASymbol "-" month ("-(" day ")" UASymbol)?
                   | year UASymbol "-" monthDay UASymbol?
                   | year UASymbol? "-(" month ")" UASymbol ("-(" day ")" UASymbol)?
                   | year UASymbol? "-(" month ")" UASymbol ( "-" day )?
                   | yearMonth UASymbol "-(" day ")" UASymbol
                   | yearMonth UASymbol "-" day
                   | yearMonth "-(" day ")" UASymbol
                   | year "-(" monthDay ")" UASymbol
                   | season UASymbol



(* ** Internal Unspecified** *)

internalUnspecified =
     yearWithU
   | yearMonthWithU
   | yearMonthDayWithU

      yearWithU =
               "u" digitOrU digitOrU digitOrU
             | digitOrU  "u" digitOrU digitOrU
             | digitOrU  digitOrU "u" digitOrU
             | digitOrU  digitOrU digitOrU "u"


      yearMonthWithU =
                  (year | yearWithU) "-" monthWithU
                | yearWithU "-" month

      yearMonthDayWithU =
                      (yearWithU | year) "-" monthDayWithU
                    | yearWithU "-" monthDay

     monthDayWithU=
                     (month | monthWithU) "-" dayWithU
                   | monthWithU "-" day

     monthWithU = oneThru12 | "0u" | "1u" | ("u" digitOrU)

     dayWithU =
                  oneThru31
               | "u" dugitOrU
               | oneThru3 "u"

 digitOrU = positiveDigitOrU | "0"
 positiveDigitOrU = positiveDigit | "u"
 oneThru3 = "1" | "2" | "3"



(* ** Inclusive list and choice list** *)

   choiceList =   "[" listContent "]"
   inclusiveList = "{" listContent "}"


listContent = earlier ("," listElement)*
                   | (earlier ",")? (listElement ",")* later
                   | listElement ("," listElement)+
                   | consecutives

    listElement =    date
                      | dateWithInternalUncertainty
                      | uncertainOrApproxDate
                      | unspecified
                      | consecutives

      earlier =  ".." date
      later = date ".."
      consecutives = yearMonthDay ".." yearMonthDay
                          | yearMonth ".." yearMonth
                          | year ".." year



(* ** Masked precision ** *)

maskedPrecision =  digit digit ((digit "x") | "xx" )



(* ** L2Interval ** *)

L2Interval =    dateOrSeason "/" dateWithInternalUncertainty
                 |dateWithInternalUncertainty "/"dateOrSeason
                 |dateWithInternalUncertainty "/" dateWithInternalUncertainty


(* ** Long Year - Scientific Form ** *)

longYearScientific = "y" "-"? positiveInteger "e" positiveInteger ("p" positiveInteger)?

positiveInteger = positiveDigit digit*

(* ** SeasonQualified ** *)

seasonQualified = season "^" seasonQualifier

     seasonQualifier = qualifyingString



(* ** Auxiliary Assignments for Level 2 ** *)

 dateWithInternalUncertainty =
                             internalUncertainOrApproximate
                         | internalUnspecified

qualifyingString = Any sequence of characters that does not include whitespace.

*/
