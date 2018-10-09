package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT = 4 // not a triangle
	Equ = 3 // equilateral
	Iso = 2 // isosceles
	Sca = 0 // scalene
)

// KindFromSides return the Kind of triangle based on sides provided
func KindFromSides(a, b, c float64) Kind {
	// if any of the sides are negative, NaN, or Inf ==> NaT
	for _, i := range []float64{a, b, c} {
		if i <= 0 || math.IsNaN(i) || math.IsInf(i, 0) {
			return NaT
		}
	}
	// sides must add up to greater or equal to third side
	if a+b < c || b+c < a || (a+c) < b {
		return NaT
	}
	// Equ
	if a == b && b == c {
		return Equ
	}

	// Iso
	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
