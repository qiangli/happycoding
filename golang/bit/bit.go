package bit

func set(n int, i byte) int {
	return n | (1 << i)
}

func get(n int, i byte) bool {
	return n&(1<<i) > 0
}

func clear(n int, i byte) int {
	return n & ^(1 << i)
}
