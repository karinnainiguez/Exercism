package scrabble

import "strings"

// Score returns the total score of a word provided
func Score(word string) int {
	totalScore := 0

	for i := 0; i < len(word); i++ {
		switch strings.ToUpper(string(word[i])) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			totalScore++
		case "D", "G":
			totalScore += 2
		case "B", "C", "M", "P":
			totalScore += 3
		case "F", "H", "V", "W", "Y":
			totalScore += 4
		case "K":
			totalScore += 5
		case "J", "X":
			totalScore += 8
		case "Q", "Z":
			totalScore += 10
		}
	}

	return totalScore
}
