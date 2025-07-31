// Package scrabble provides tools for scoring words in the game of Scrabble.
package scrabble

import "strings"

var letterValues = map[string]int{
	"A": 1,
	"E": 1,
	"I": 1,
	"O": 1,
	"U": 1,
	"L": 1,
	"N": 1,
	"R": 1,
	"S": 1,
	"T": 1,
	"D": 2,
	"G": 2,
	"B": 3,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

// Score accepts an ASCII string, representing a word, composed of any of the
// 26 valid scrabble letters and returs that word's Scrabble score.
func Score(input string) int {

	score := 0

	for i := 0; i < len(input); i++ {
		score += letterValues[strings.ToUpper(string(input[i]))]
	}

	return score
}
