package common

import (
	"github.com/sfomuseum/go-edtf"
	"math/big"
)

// Parse a string in exponential notation in to a year value in numeric form.
func ParseExponentialNotation(notation string) (int, error) {

	flt, _, err := big.ParseFloat(notation, 10, 0, big.ToNearestEven)

	if err != nil {
		return 0, err
	}

	var i = new(big.Int)
	yyyy, _ := flt.Int(i)

	if yyyy.Int64() > int64(edtf.MAX_YEARS) || yyyy.Int64() < int64(0-edtf.MAX_YEARS) {
		return 0, edtf.Unsupported("exponential notation", notation)
	}

	yyyy_i := int(yyyy.Int64())
	return yyyy_i, nil
}
