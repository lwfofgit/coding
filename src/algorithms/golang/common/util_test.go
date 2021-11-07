package common

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T)  {
	//t.Name()
	a := []int{1,2,3}
	Swap(a, 1,2)
	fmt.Println(a)
}
