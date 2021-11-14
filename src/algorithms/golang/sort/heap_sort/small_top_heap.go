package heap_sort

import "github.com/lwfofgit/coding/src/algorithms/golang/common"

//小顶堆
type SmallTopHeap struct {
	data []int
	count int
	capacity int
}

func CreateHeap(capacity int) *SmallTopHeap {
	return &SmallTopHeap{
		count: 0,
		capacity: capacity,
		data: make([]int, capacity+1),
	}
}

//[0,2,1,3]
func BuildSmallTopHeap(data []int) *SmallTopHeap {
	heap := CreateHeap(len(data))
	for i := range data {
		heap.data[i+1] = data[i]
	}

	for i := len(data)/2; i>=1; i-- {
		heap.Heapify(i)
	}
	return heap
}

func (h *SmallTopHeap)Heapify(i int) {
	if h.count >= h.capacity {
		return
	}

	/*
	    int maxPos = i;
	    if (i*2 <= n && a[i] < a[i*2]) maxPos = i*2;
	    if (i*2+1 <= n && a[maxPos] < a[i*2+1]) maxPos = i*2+1;
	    if (maxPos == i) break;
	    swap(a, i, maxPos);
	    i = maxPos;
	*/
	for {
		minPos := i
		if 2*i <= h.capacity && h.data[i] > h.data[2*i] {
			minPos = 2*i
		}
		if 2*i+1 <= h.capacity && h.data[minPos] > h.data[2*i +1]{
			minPos = 2*i + 1
		}
		if minPos == i {
			return
		}
		common.Swap(h.data, i, minPos)
		i = minPos
		return
	}
}
