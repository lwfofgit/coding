package string_match

import (
	"fmt"
	"testing"
)

func TestBK(t *testing.T) {
	parent := "abcdefghm"
	child := "def"

	i := BK(parent,child)
	fmt.Println(i)
}
