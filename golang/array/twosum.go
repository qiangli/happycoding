package array

func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		x := target - nums[i]
		memo[x] = i
	}

	for i := 0; i < len(nums); i++ {
		j, ok := memo[nums[i]]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
