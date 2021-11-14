package binary_search

import "fmt"

type node struct{
	Data interface{}
	Left *node
	Right *node
}

func CreateNode(data interface{}) *node  {
	return &node{
		Data: data,
	}
}

/*
      1
    2    3
  4   5 6  7

pre: 1 2 4 5 3 6 7
mid: 4 2 5 1 6 3 7
post:4 5 2 6 7 3 1
*/
func CreateTree() *node {
	root := CreateNode(1)

	root.Left = CreateNode(2)
	root.Right = CreateNode(3)

	root.Left.Left = CreateNode(4)
	root.Left.Right = CreateNode(5)

	root.Right.Left = CreateNode(6)
	root.Right.Right = CreateNode(7)
	return root
}

// 先打印自己，然后在打印左子树，最后打印右子树
func PreOrder(root *node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func MidOrder(root *node) {
	if root == nil {
		return
	}
	MidOrder(root.Left)
	fmt.Print(root.Data)
	MidOrder(root.Right)
}

func PostOrder(root *node) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Print(root.Data)
}
