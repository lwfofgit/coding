package heap_sort

import (
	"fmt"
	"testing"
)

func TestSmallTopHeap_Heapify(t *testing.T) {
	 a := []int{2,1,3,6,5,8} // [0,1,2,3,6,5,8]
	 h := BuildSmallTopHeap(a)
	 fmt.Println(&h.data)
}
