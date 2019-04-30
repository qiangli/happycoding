package maps

import (
	"fmt"
)

func keys(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// items in a but not in b
func difference(a, b []string) []string {
	m := map[string]bool{}
	for _, v := range b {
		m[v] = true
	}
	c := []string{}
	for _, v := range a {
		_, ok := m[v]
		if !ok {
			c = append(c, v)
		}
	}
	return c
}

// items in both a and b
func intersection(a, b []string) []string {
	ma := map[string]bool{}
	for _, v := range a {
		ma[v] = true
	}
	c := []string{}
	for _, v := range b {
		_, ok := ma[v]
		if ok {
			c = append(c, v)
		}
	}
	return c
}

func diff(ma, mb map[string]string) {
	a := keys(ma)
	b := keys(mb)

	da := difference(a, b)
	db := difference(b, a)

	var dab []string
	for _, v := range intersection(a, b) {
		va := ma[v]
		vb := mb[v]
		if va != vb {
			dab = append(dab, v)
		}
	}

	dab = append(dab, da...)
	dab = append(dab, db...)
	fmt.Printf("diff keys: %v\n", dab)
}
