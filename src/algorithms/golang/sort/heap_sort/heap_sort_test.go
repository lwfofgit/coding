package heap_sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{0,2,4,6,1,7,8}
	Sort(a, len(a)-1)
	fmt.Println("end")
}
