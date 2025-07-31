// Package hamming calculates the Hamming Distance between two DNA strands
package hamming

import "errors"

// Distance returns the Hamming Distance of two DNA strands a, and b.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA strands aren't the same length")
	}

	dist := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
