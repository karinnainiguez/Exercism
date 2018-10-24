package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(num int) (string, error) {
	var newNum strings.Builder

	if num <= 0 || num > 3000 {
		return "", errors.New("Number must be between 1 and 3000")
	}

	// M == 1000
	// D == 500
	// C == 100
	// L == 50
	// X == 10
	// V == 5
	// I == 1
	if num > 1000 {
		count := num / 1000
		for i := 0; i < count; i++ {
			newNum.WriteString("M")
		}
		num %= 1000
	}
	if num > 500 {
		count := num / 500
		for i := 0; i < count; i++ {
			newNum.WriteString("D")
		}
		num %= 500

	}
	if num > 100 {
		count := num / 500
		for i := 0; i < count; i++ {
			newNum.WriteString("D")
		}
		num %= 500
	}
	if num > 50 {
		count := num / 50
		for i := 0; i < count; i++ {
			newNum.WriteString("L")
		}
		num %= 50
	}
	if num > 10 {
		count := num / 10
		for i := 0; i < count; i++ {
			newNum.WriteString("X")
		}
		num %= 10
	}
	if num > 5 {
		count := num / 5
		for i := 0; i < count; i++ {
			newNum.WriteString("V")
		}
		num %= 5
	}

	for num > 0 {
		newNum.WriteString("I")
		num -= 1
	}

	return newNum.String(), nil
}
