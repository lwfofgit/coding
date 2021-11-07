package heap_sort

import "github.com/lwfofgit/data_structure/src/algorithms/golang/common"

//小顶堆
type BigTopHeap struct {
	data     []int
	count    int
	capacity int
}

func CreateBigHeap(capacity int) *BigTopHeap {
	return &BigTopHeap{
		count:    0,
		capacity: capacity,
		data:     make([]int, capacity+1),
	}
}

//[0,2,1,3]
func BuildBigTopHeap(data []int) *BigTopHeap {
	heap := CreateBigHeap(len(data))
	for i := range data {
		heap.data[i+1] = data[i]
	}

	for i := len(data) / 2; i >= 1; i-- {
		heap.Heapify(i)
	}
	return heap
}

func (h *BigTopHeap) Heapify(i int) {
	if h.count >= h.capacity {
		return
	}
	for {
		maxPos := i
		if 2*i <= h.capacity && h.data[i] < h.data[2*i] {
			maxPos = 2 * i
		}
		if 2*i+1 <= h.capacity && h.data[maxPos] < h.data[2*i+1] {
			maxPos = 2*i + 1
		}
		if maxPos == i {
			return
		}
		common.Swap(h.data, i, maxPos)
		i = maxPos
		return
	}
}
