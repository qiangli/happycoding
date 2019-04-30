package maps

import (
	"testing"
)

func TestDiff(t *testing.T) {
	m1 := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}

	m2 := map[string]string{
		"k4": "v4",
		"k5": "v5",
		"k3": "v3",
	}

	diff(m1, m2)
}
