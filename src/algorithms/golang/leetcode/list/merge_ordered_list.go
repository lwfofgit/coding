package list

/*
题目：
	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
	1->2->4
	1->3->4
	合并之后：1->1->2->3->4->4

题解：
	1.迭代方式：注意初始化一个哨兵节点，prev := preHead很关键，如果直接用preHead迭代，当执行
preHead = preHead.Next后，preHead就会往后移了，到最后就找不到头了
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	preHead := &ListNode{}
	prev := preHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return preHead.Next
}
