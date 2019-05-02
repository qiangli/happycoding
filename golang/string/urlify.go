package string

import (
	"fmt"
)

func URLify(a []byte, length int) {
	if length <= 0 {
		return
	}
	//space count
	cnt := 0
	for i := 0; i < length; i++ {
		if a[i] == ' ' {
			cnt++
		}
	}

	spacer := func(a []byte, k int) int {
		sp := []byte("%20")
		for i, c := range sp {
			a[k-i] = c
		}
		return k - 3
	}

	k := length - 1 + 2*cnt
	j := length - 1
	for i := length - 1; i > 0; {
		if a[i] == ' ' {
			fmt.Printf("sp: %v %v %v %v\n", string(a), i, j, k)

			for j > i {
				fmt.Printf("%v %v %v %v\n", string(a), i, j, k)

				a[k] = a[j]
				k--
				j--
			}
			k = spacer(a, k)
			j--
		}
		i--
	}
}
