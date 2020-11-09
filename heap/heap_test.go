package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	h := BuildMaxHeap([]int{})
	assert.Equal(t, 0, h.Len())
	assert.Equal(t, nil, h.Max())
}
func TestBuildMaxHeap(t *testing.T) {
	h := BuildMaxHeap([]int{1, 2, 5, 6, 3, 8})
	assert.Equal(t, 6, h.Len())
	assert.Equal(t, 8, h.Max())
}
func TestExtractMax(t *testing.T) {
	h := BuildMaxHeap([]int{1, 2, 5, 6, 3, 8})

	max, _ := h.ExtractMax()
	assert.Equal(t, 8, max)

	max, _ = h.ExtractMax()
	assert.Equal(t, 6, max)
}
func TestHeapSort(t *testing.T) {
	a := []int{1, 2, 5, 6, 3, 8}
	r := Sort(a)
	assert.Equal(t, []int{1, 2, 3, 5, 6, 8}, r)
}
