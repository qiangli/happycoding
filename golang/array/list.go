package array

type ArrayList struct {
	array []int
	size  int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		array: make([]int, 4),
		size:  0,
	}
}

func (r *ArrayList) full() bool {
	return r.size == len(r.array)
}

func (r *ArrayList) resize() {
	na := make([]int, len(r.array)*2)
	//copy(na, r.array)
	for i, v := range r.array {
		na[i] = v
	}
	r.array = na
}

func (r *ArrayList) Append(v int) {
	if r.full() {
		r.resize()
	}
	r.array[r.size] = v
	r.size++
}

func (r *ArrayList) Insert(index int, v int) {
	if r.full() {
		r.resize()
	}
	//shift to make room
	for i := r.size; i >= index; i-- {
		r.array[i+1] = r.array[i]
	}
	r.array[index] = v
	r.size++
}

// Remove deletes the element at specified location and returns the removed value
func (r *ArrayList) Remove(index int) int {
	v := r.array[index]
	for i := index; i < r.size; i++ {
		r.array[i] = r.array[i+1]
	}
	r.size--
	return v
}

func (r ArrayList) Get(index int) int {
	return r.array[index]
}

func (r ArrayList) Size() int {
	return r.size
}
