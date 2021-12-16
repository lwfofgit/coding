package tree

import . "github.com/lwfofgit/coding/src/algorithms/golang/leetcode"

func midOrder(root *TreeNode) []int {
	var v []int
	if root == nil {
		return v
	}
	v = append(v, midOrder(root.Left)...)
	v = append(v, root.Val)
	v = append(v, midOrder(root.Right)...)
	return v
}

// 顺序：左根右
// 1.先完成左根遍历，则入栈顺序为：根左，不停循环把左边子树压入栈；
// 2.然后出栈，左边第一个先出栈，出栈后值存储到列表，然后把该节点右节点复制给当前节点，如果当前节点没有
// 右节点，就不会循环入栈动作，就直接从栈中回去下一个节点，即根，然后把根的值加入列表；
// 3.把根的右节点压入栈
func midOrderV2(root *TreeNode) []int {
	v := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			v = append(v, root.Val)
			root = root.Right
		}
	}
	return v
}

// 顺序：根左右, 前序和中序类似，只是先入栈过程中先把根的值记录到列表，出栈时如果是根节点就把它的右节点入栈
func preOrder(root *TreeNode) []int {
	v := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		if root != nil {
			v = append(v, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return v
}

/*
实现：
	这里要增加一个辅助节点来节点上一次放入结果数组的值，当一个节点左右都是空的时候，就可以放入结果集，
当上一个放入结果集的节点是他的孩子节点的时候，说明他的孩子已经访问完成了，就可以再次放入了(根节点)，
还有一点要注意的是，当一个节点的左右不为空时，要先加入右孩子，再加入左孩子，这样才能先访问左孩子。

*/

func postOrder(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	if root == nil {
		return res
	}
	pre := &TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		// 左右子孩子都为空说明是叶子节点，可以放入值，因为右孩子先入栈，左孩子后入
		// 需要用pre来记录一下上一个放入值的节点是那个，如果当前节点的上一个节点是它的
		// 左或是右，说明它的孩子访问完毕了，那么它也就可以放入值了
		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			res = append(res, cur.Val)
			pre = cur
			stack = stack[:len(stack)-1]
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return res
}
