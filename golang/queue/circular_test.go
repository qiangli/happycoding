package queue

import (
	"fmt"
	"testing"
)

func TestCircular(t *testing.T) {
	circularQueue := Constructor(3) // set the size to be 3
	circularQueue.EnQueue(1)        // return true
	circularQueue.EnQueue(2)        // return true
	circularQueue.EnQueue(3)        // return true
	circularQueue.EnQueue(4)        // return false, the queue is full
	circularQueue.Rear()            // return 3
	circularQueue.IsFull()          // return true
	circularQueue.DeQueue()         // return true
	circularQueue.EnQueue(4)        // return true
	circularQueue.Rear()            // return 4

	fmt.Println(circularQueue)
}
