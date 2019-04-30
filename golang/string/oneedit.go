package string

import ()

func oneAway(s1, s2 string) bool {
	return oneEdit([]byte(s1), []byte(s2))
}

func oneEdit(s1, s2 []byte) bool {
	l1 := len(s1)
	l2 := len(s2)

	if l1 == l2 && oneReplace(s1, s2) {
		return true
	}

	if l1+1 == l2 && oneDelete(s1, s2) {
		return true
	}

	if l1 == l2+1 && oneDelete(s2, s1) {
		return true
	}

	return false
}

func oneReplace(s1, s2 []byte) bool {
	c := 0
	for i, v := range s1 {
		if v != s2[i] {
			c++
		}
		if c > 1 {
			return false
		}
	}
	return true
}

func oneDelete(s1, s2 []byte) bool {
	diff := false
	i := -1
	for _, v := range s1 {
		i++
		if v != s2[i] {
			if diff {
				return false
			}
			i++
			diff = true
		}
	}
	return true
}
