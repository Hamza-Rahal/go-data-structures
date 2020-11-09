package heap

// MaxHeap heap that stores integers
type MaxHeap struct {
	array []int
}

// NewMaxHeap constructor for MaxHeap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Max retuns the biggest element in the heap
// complexit: O(1)
func (h *MaxHeap) Max() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return h.array[0]
}

// maxHeapify transforms a tree whose sub trees are max heaps into a max heap
// complexity: O(log(n)) n is the length of the array
func (h *MaxHeap) maxHeapify(i int) {
	lchild := 2*i + 1
	rchild := 2*i + 2
	parent := i
	l := len(h.array)

	if lchild < l && h.array[lchild] > h.array[parent] {
		parent = lchild
	}
	if rchild < l && h.array[rchild] > h.array[parent] {
		parent = rchild
	}

	if parent != i {
		// swaps parent and max(rchild, lchild)
		v := h.array[i]
		h.array[i] = h.array[parent]
		h.array[parent] = v

		h.maxHeapify(parent)
	}
}

// Len returns the size of the heap
func (h *MaxHeap) Len() int {
	return len(h.array)
}

// ExtractMax removes the root of the heap and rebuilds it to fix potential heap property violations
// complexity: O(n)
func (h *MaxHeap) ExtractMax() (int, error) {
	if len(h.array) == 0 {
		return 0, &errorString{"The heap is empty"}
	}
	max := h.array[0]
	h.array = BuildMaxHeap(h.array[1:]).array
	return max, nil
}

// BuildMaxHeap builds a max heap out of the provided array
// complexity: O(n)
func BuildMaxHeap(array []int) *MaxHeap {
	h := MaxHeap{array}
	l := len(array)
	if l >= 2 {
		for i := (l - 1) / 2; i >= 0; i-- {
			// in-place algorithm
			h.maxHeapify(i)
		}
	}
	return &h
}

// Sort sorts the provided array
// complexity: O(nlog(n))
func Sort(array []int) []int {
	n := len(array)
	for i := 0; i < n-1; i++ {
		BuildMaxHeap(array[:n-i])
		// swap the head and the tail of the unsorted area of the array
		head := array[0]
		array[0] = array[n-1-i]
		array[n-1-i] = head
	}
	return array
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
