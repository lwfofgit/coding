package list

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	ma := make(map[*ListNode]bool)
	ma[headA] = true
	for headA.Next != nil {
		ma[headA.Next] = true
		headA = headA.Next
	}


	for headB != nil {
		if _,ok := ma[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}