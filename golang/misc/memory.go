package misc

import (
	"fmt"
	"math/rand"
	"time"
)

// https://www.filipekberg.se/2013/07/01/memory-access-pattern-matters/
// https://software.intel.com/en-us/articles/optimize-data-structures-and-memory-access-patterns-to-improve-data-locality
// https://channel9.msdn.com/Events/Build/2013/4-329
func matrix(N int) int {
	// init
	A := make([][]int, N)
	B := make([][]int, N)
	C := make([][]int, N)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		B[i] = make([]int, N)
		C[i] = make([]int, N)
		for j := 0; j < N; j++ {
			A[i][j] = rand.Int()
			B[i][j] = rand.Int()
			C[i][j] = rand.Int()
		}
	}

	fmt.Printf("init done\n")

	// multiply
	start1 := time.Now()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	elapsed1 := time.Since(start1)
	fmt.Printf("time taken: %v\n", elapsed1)

	// fast version
	start2 := time.Now()

	for i := 0; i < N; i++ {
		for k := 0; k < N; k++ {
			for j := 0; j < N; j++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	elapsed2 := time.Since(start2)
	fmt.Printf("time taken: %v\n", elapsed2)

	diff := elapsed2 - elapsed1
	fmt.Printf("Difference: %v\n", diff)

	return int(diff)
}
