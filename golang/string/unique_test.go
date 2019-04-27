package string

import (
	"fmt"
	"testing"
)

func TestIsUnique(t *testing.T) {
	a := []byte("notunique")

	fmt.Printf("unique %v %v %v %v\n", a, isUnique(a), isUnique2(a), isUnique3(a))
	if isUnique(a) || isUnique2(a) || isUnique3(a) {
		t.FailNow()
	}

	a = []byte("uniq")

	fmt.Printf("unique %v %v %v %v\n", a, isUnique(a), isUnique2(a), isUnique3(a))
	if !isUnique(a) || !isUnique2(a) || !isUnique3(a) {
		t.FailNow()
	}
}
