package binary_search

import "testing"

func TestSearch(t *testing.T) {
	a := []int{1,2,3,4,5}
	println(RecursionSearch(a,5))
}
