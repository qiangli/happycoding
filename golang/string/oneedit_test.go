package string

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	if !oneAway("pale", "ple") {
		t.FailNow()
	}
	if !oneAway("pales", "pale") {
		t.FailNow()
	}
	if !oneAway("pale", "bale") {
		t.FailNow()
	}
	if oneAway("pale", "bae") {
		t.FailNow()
	}
}
