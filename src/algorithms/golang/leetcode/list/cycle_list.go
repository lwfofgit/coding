package list

/*
题目：
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
如果链表中存在环，则返回 true 。 否则，返回 false 。

题解：
方法一：哈希表
使用哈希表来存储所有已经访问过的节点。
每次我们到达一个节点，如果该节点已经存在于哈希表中，则说明该链表是环形链表，否则就将该节点加入哈希表中。
重复这一过程，直到我们遍历完整个链表即可。

*/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[*ListNode]bool)
	for ; head != nil; head = head.Next {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
	}
	return false
}

//func hasCycleV2(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return false
//	}
//
//	fast, slow := head.Next, head
//	return false
//}
