package hash

import (
//"fmt"
)

type entry struct {
	key  int
	val  interface{}
	next *entry
}

type Hashtable struct {
	bucket []*entry
}

// NewHashtable returns a new hashtable
func NewHashtable() *Hashtable {
	n := 100
	return &Hashtable{
		bucket: make([]*entry, n),
	}
}

func (r *Hashtable) hash(v int) int {
	mod := v % len(r.bucket)
	return mod
}

func (r *Hashtable) Put(key int, val interface{}) {
	idx := r.hash(key)
	head := r.bucket[idx]
	e := &entry{
		key: key,
		val: val,
	}
	if head == nil {
		r.bucket[idx] = e
		return
	}
	for c := head; ; {
		if c.next == nil {
			c.next = e
			return
		}
		c = c.next
	}
}

func (r *Hashtable) Get(key int) (interface{}, bool) {
	idx := r.hash(key)
	head := r.bucket[idx]
	for c := head; ; {
		if c == nil {
			return 0, false
		}
		if c.key == key {
			return c.val, true
		}
		c = c.next
	}
}
