package util

// Iterator over a collection
type Iterator interface {
	HasNext() bool
	Next() int
}

// Collection represents a group of elements.
type Collection interface {
	//add(int) bool
	Clear()
	//contains(int) bool
	Iterator() Iterator
	//remove(int) bool
	Size() int
}
