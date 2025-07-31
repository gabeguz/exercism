// Package isogram determines if a given string is an isogram or not
package isogram

import "strings"
import "fmt"

// IsIsogram returns true if the string is an isogram and false if not
func IsIsogram(word string) bool {
    if len(word) == 0 {
        return true
    }

    // string pre-processsing lowercase the word, remove spaces and hyphens
	word = strings.ToLower(word)
    word = strings.ReplaceAll(word, " ", "")
    word = strings.ReplaceAll(word, "-", "")

    // count how many times a letter appears in the word
    for i := 0 ; i < len(word); i++ {
        fmt.Println("Checking " + string(word[i]) + " in " + word)
        if strings.Count(word, string(word[i])) > 1 {
            return false
        }
    }
    return true
}