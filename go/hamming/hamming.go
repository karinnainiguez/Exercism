package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	// if the two strings are of varying lengths, return error
	if len(a) != len(b) {
		return 0, errors.New("not the same length")
	}
	// start a var at 0 for distance
	distance := 0
	// iterate thorugh strings one letter at a time
	for i := 0; i < len(a); i++ {
		// if the strings are not equal at index i, increment distance
		if !(strings.EqualFold(string(a[i]), string(b[i]))) {
			distance++
		}
	}
	return distance, nil
}
