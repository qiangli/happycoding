package string

import (
	"testing"
)

func TestContain(t *testing.T) {
	if !contain("geekforgeeks", "for") {
		t.FailNow()
	}

	if !contain("geekforgeeks", "geeks") {
		t.FailNow()
	}

	if contain("geekforgeeks", "fore") {
		t.FailNow()
	}
}
