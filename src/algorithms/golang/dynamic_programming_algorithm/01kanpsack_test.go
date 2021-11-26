package dynamic_programming_algorithm

import (
	"fmt"
	"testing"
)

func TestKnapsack(t *testing.T) {
	w := []int{2,2,2,4,6}
	fmt.Println(Knapsack(w, 5, 9))
}
