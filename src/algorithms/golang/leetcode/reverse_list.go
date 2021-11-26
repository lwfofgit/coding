package leetcode

type ListNode struct {
	Var int
	Next *ListNode
}

// 1>2>3>nil
// 3>2>1>nil
// 1.第一步先断链，从头结点开始，先把当前结点的next执行上一个节点，所以先声明一个prev表示上一个节点
// 2.如果直接current.Next = prev了，那么就不知道current一开始的next是谁了，所以断链前需要记录一下current.next
func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
