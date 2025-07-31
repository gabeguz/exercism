// Package raindrops allows you to convert integers to their raindrop equivalent
package raindrops

import "strconv"

// Convert takes an integer and returns a string in the form PlingPlangPlong
// depending on if the integer is factorable by 3, 5, or 7.   If it's not
// factorable by any of those numbers, the number itlself is returned, as a
// string.
func Convert(input int) string {

	var out string

	nofactors := false

	if input%3 == 0 {
		out = "Pling"
		nofactors = true
	}

	if input%5 == 0 {
		out = out + "Plang"
		nofactors = true
	}

	if input%7 == 0 {
		out = out + "Plong"
		nofactors = true
	}

	if nofactors == false {
		out = strconv.Itoa(input)
	}

	return out
}
