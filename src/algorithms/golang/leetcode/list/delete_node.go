package list

/*
题目：
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

在本题中，如果我们要删除节点 yy，我们需要知道节点 yy 的前驱节点 xx，并将 xx 的指针指向 yy 的后继节点。但由于头节点不存在前驱节点，因此我们需要在删除头节点时进行特殊判断。但如果我们添加了哑节点，那么头节点的前驱节点就是哑节点本身，此时我们就只需要考虑通用的情况即可。
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := length(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func length(head *ListNode) int {
	i := 0
	for ; head != nil; head = head.Next {
		i++
	}
	return i
}
