package common

import (
	"github.com/whosonfirst/go-edtf"
	"math/big"
)

func ParseExponentialNotation(notation string) (int, error) {

	flt, _, err := big.ParseFloat(notation, 10, 0, big.ToNearestEven)

	if err != nil {
		return 0, err
	}

	var i = new(big.Int)
	yyyy, _ := flt.Int(i)

	if yyyy.Int64() > int64(9999) || yyyy.Int64() < 0 {
		return 0, edtf.Unsupported("exponential notation", notation)
	}

	yyyy_i := int(yyyy.Int64())
	return yyyy_i, nil
}
