package heaptypes

/*  Integer heaps
	Low and High return values

!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

*/

// An IntHeaplow is a min-heap of ints.
type IntHeaplow []int

func (h IntHeaplow) Len() int           { return len(h) }
func (h IntHeaplow) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeaplow) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeaplow) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeaplow) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An IntHeaphigh is a min-heap of ints.
type IntHeaphigh []int

func (h IntHeaphigh) Len() int           { return len(h) }
func (h IntHeaphigh) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeaphigh) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeaphigh) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeaphigh) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
