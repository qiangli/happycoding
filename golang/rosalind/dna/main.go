//http://rosalind.info/problems/dna/
package main

import (
	"fmt"
)

func cnt(s string) {
	var c uint64
	for _, b := range s {
		// map AGCT to 0123 and increment the count at bit position: 0, 16, 32, 48
		c += 1 << uint64(((b-rune('A'))%5-(b-rune('A'))%5/3)*16)
	}

	fmt.Printf("%v %v %v %v", c&0x000000000000ffff, c>>32&0x000000000000ffff, c>>16&0x000000000000ffff, c>>48&0x000000000000ffff)
}

func main() {
	sample := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	cnt(sample)
}
