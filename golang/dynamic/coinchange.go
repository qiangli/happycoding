package dynamic

func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var compute func(int) int
	memo := make(map[int]int)

	compute = func(a int) int {
		if a == 0 {
			return 0
		}
		if a < 0 {
			return -1
		}

		//
		v, ok := memo[a]
		if ok {
			//fmt.Println("cached:", v)
			return v
		}
		v = 1 << 31
		for _, c := range coins {
			k := a - c
			x := compute(k)
			if x >= 0 {
				v = min(v, x+1)
				ok = true
			}
		}
		if !ok {
			v = -1
		}
		//fmt.Println("caching", v, "ok?", ok)
		memo[a] = v
		return v
	}

	return compute(amount)
}

// Greedy algorithm wont work for the following:
// coins := []int{186,419,83,408}
// amount := 6249
//
//func coinChange(coins []int, amount int) int {
//     if len(coins) == 0 || amount <= 0 {
//         return 0
//     }
//     sort.Sort(sort.IntSlice(coins))
//     fmt.Println("coins:", coins)
//     var a []int
//     for i := len(coins) - 1; i >= 0 && amount > 0; {
//         if coins[i] > amount {
//             i--
//             continue
//         }
//         a = append(a, coins[i])
//         amount -= coins[i]
//     }
//     if amount == 0 {
//         fmt.Println("coin:", a)
//         return  len(a)
//     }
//     return -1
//}
