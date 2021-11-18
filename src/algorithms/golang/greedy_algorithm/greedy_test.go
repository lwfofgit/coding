package greedy_algorithm

import (
	"fmt"
	"testing"
)

func TestGetNums(t *testing.T) {
	v := 6
	fmt.Println(GetNums(v,0))
}

func TestTestStr(t *testing.T) {
	str := "abcdef"
	fmt.Println(TestStr(str, 3))
	fmt.Println(str[1:3])
	fmt.Println(str[:len(str)-1])
}

func TestMinValue(t *testing.T) {
	i := 36785
	k := 3
	fmt.Println(MinValue(i, k))
}
