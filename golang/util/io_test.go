package util

import (
	"testing"
)

func TestScan(t *testing.T) {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var a, b int
	scanf("%d %d\n", &a, &b)
	printf("%d\n", a+b)
}
