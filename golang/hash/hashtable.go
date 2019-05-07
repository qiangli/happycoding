package hash

import (
	"fmt"
)

type entry struct {
	key  interface{}
	val  interface{}
	next *entry
}

type Hashtable struct {
	bucket []*entry
}

// NewHashtable returns a new hashtable
func NewHashtable() *Hashtable {
	n := 10
	return &Hashtable{
		bucket: make([]*entry, n),
	}
}

func (r *Hashtable) hash(i interface{}) int {
	hash := func(s string) int {
		n := len(s)
		v := int(s[n-1])
		m := 31
		for j := n - 2; j >= 0; j-- {
			v += int(s[j]) * m
			m *= m
		}
		return v
	}
	v := 0
	switch i.(type) {
	case int:
		v = i.(int)
	case string:
		s := i.(string)
		v = hash(s)
	default:
		s := fmt.Sprintf("%v", i)
		v = hash(s)
	}
	mod := v % len(r.bucket)
	return mod
}

func (r *Hashtable) Put(key interface{}, val interface{}) {
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
	c := head
	for {
		if c.next == nil {
			c.next = e
			return
		}
		c = c.next
	}
}

func (r *Hashtable) Get(key interface{}) (interface{}, bool) {
	idx := r.hash(key)
	head := r.bucket[idx]
	c := head
	for {
		if c == nil {
			return nil, false
		}
		if c.key == key {
			return c.val, true
		}
		c = c.next
	}
}

func (r *Hashtable) Delete(key interface{}) bool {
	idx := r.hash(key)
	head := r.bucket[idx]
	if head == nil {
		return false
	}
	if head.key == key {
		r.bucket[idx] = head.next
		return true
	}
	p := head
	c := head.next
	for {
		if c == nil {
			return false
		}
		if c.key == key {
			p.next = c.next
			return true
		}
		p = c
		c = c.next
	}
}
