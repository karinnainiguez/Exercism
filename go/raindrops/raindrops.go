package raindrops

import (
	"bytes"
	"strconv"
)

// Convert takes an integer and returns a string
func Convert(num int) string {
	var response bytes.Buffer

	// If the number has 3 as a factor, add 'Pling' to response
	if num%3 == 0 {
		response.WriteString("Pling")
	}

	// If the number has 5 as a factor, add 'Plang' to response
	if num%5 == 0 {
		response.WriteString("Plang")
	}

	// If the number has a 7 as a factor, add 'Plong' to response
	if num%7 == 0 {
		response.WriteString("Plong")
	}

	// If the number does NOT have 3, 5, or 7 as factor, add num to response
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		response.WriteString(strconv.Itoa(num))
	}

	return response.String()
}
