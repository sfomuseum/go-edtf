package cmp

import (
	"fmt"
)

func IsBetween(d string, inception string, cessation string) (bool, error) {

	is_before_inception, err := IsBefore(d, inception)

	if err != nil {
		return false, fmt.Errorf("Failed to determine if date is before inception date, %w", err)
	}

	if is_before_inception {
		return false, nil
	}

	is_after_cessation, err := IsAfter(d, cessation)

	if err != nil {
		return false, fmt.Errorf("Failed to determine if date is after cessation date, %w", err)
	}

	if is_after_cessation {
		return false, nil
	}

	return true, nil
}
