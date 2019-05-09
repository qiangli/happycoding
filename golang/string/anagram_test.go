package string

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	str1 := "stackoverflow"
	str2 := "flowerovstack"

	if !anagram(str1, str2) {
		t.FailNow()
	}
}
