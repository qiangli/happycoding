package queue

import (
//"fmt"
)

type MyCircularQueue struct {
	head, tail int
	item       []int
}

// Constructor initializes your data structure here. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head: -1,
		tail: -1,
		item: make([]int, k),
	}
}

// EnQueue inserts an element into the circular queue. Return true if the operation is successful.
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % len(this.item)
	this.item[this.tail] = value

	return true
}

// DeQueue deletes an element from the circular queue. Return true if the operation is successful.
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	//v := this.item[this.head]
	if this.head == this.tail {
		this.head, this.tail = -1, -1
		return true
	}

	this.head = (this.head + 1) % len(this.item)

	return true
}

// Front gets the front item from the queue.
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.item[this.head]
}

// Rear gets the last item from the queue.
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.item[this.tail]
}

// IsEmpty checks whether the circular queue is empty or not.
func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == -1 {
		return true
	}
	return false
}

// IsFull checks whether the circular queue is full or not.
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%len(this.item) == this.head
}
