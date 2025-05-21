package string

// assume ascii char size 128
func isUnique(a []byte) bool {
	if len(a) > 128 {
		return false
	}
	set := make([]bool, 128)
	for i := 0; i < len(a); i++ {
		v := a[i]
		if set[v] {
			return false
		}
		set[v] = true
	}
	return true
}

func isUnique2(a []byte) bool {
	sz := len(a)
	if sz > 128 {
		return false
	}
	var flag [16]uint64
	set := func(x byte) {
		i := x / 16
		r := x % 16
		flag[i] = flag[i] | (1 << r)
	}
	get := func(x byte) bool {
		i := x / 16
		r := x % 16
		return flag[i]&(1<<r) != 0
	}
	for i := 0; i < sz; i++ {
		v := a[i]
		if get(v) {
			return false
		}
		set(v)
	}
	return true
}

func isUnique3(a []byte) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}
