// Package strand provides function for working with a DNA strand
package strand

// ToRNA returns the RNA complement to the DNA string passed
//  The mapping looks like:
//  G -> C
//  C -> G
//  A -> U
//  T -> A
func ToRNA(dna string) string {
	rna := ""
	for i := 0; i < len(dna); i++ {
		switch string(dna[i]) {
		case "G":
			rna = rna + "C"
		case "C":
			rna = rna + "G"
		case "A":
			rna = rna + "U"
		case "T":
			rna = rna + "A"
		}
	}
	return rna
}
