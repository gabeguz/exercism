package pangram

import "strings"

func IsPangram(input string) bool {
	// example pangram : "the quick brown fox jumps over the lazy dog"
	// iterate through the alphabet (a...z) and continue until you find 
	// a letter that is not in the string... if you make it through all
	// letters, then you have a pangram otherwise, you don't

	// convert input to lowercase
	lowercase := strings.ToLower(input)

	alphabet := "abcdefghijklmnopqrstuvwxzy"

	// assume every string is a pangram until shown otherwise
	pangram := true

	for x := 0; x < len(alphabet); x++ {
		if (strings.Contains(lowercase, string(alphabet[x]))) {
			continue
		} else {
		 pangram = false
	 	}
	}
	return pangram
}
