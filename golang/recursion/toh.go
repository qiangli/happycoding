package recursion

import (
	"fmt"
)

// Tower of Hanoi
// https://en.wikipedia.org/wiki/Tower_of_Hanoi

// move stack of disks from peg 0 to 2
func toh(stack int) {
	moveStack(stack, 0, 2, 1)
}

func moveStack(stack, p0, p2, p1 int) {
	if stack == 0 {
		return
	}
	moveStack(stack-1, p0, p1, p2)
	moveDisk(stack, p0, p2)
	moveStack(stack-1, p1, p2, p0)
}

func moveDisk(disk, from, to int) {
	fmt.Printf("move disk: %d from %d to: %d\n", disk, from, to)
}
