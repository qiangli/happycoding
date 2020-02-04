// http://rosalind.info/problems/rna/
package main

import (
	"fmt"
)

var alphabet = []rune("AGCU")

func transcribe(b rune) rune {
	idx := ((b-rune('A'))%5 - (b-rune('A'))%5/3)
	return alphabet[idx]
}

func main() {
	sample := []rune("GATGGAACTTGACTACGTAAATT")
	expected := "GAUGGAACUUGACUACGUAAAUU"

	for i := 0; i < len(sample); i++ {
		sample[i] = transcribe(sample[i])
	}
	fmt.Println(string(sample))
	if string(sample) != expected {
		fmt.Println("Not quite right")
	}
}
