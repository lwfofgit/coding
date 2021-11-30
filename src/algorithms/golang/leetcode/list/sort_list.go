package list

/*
题目：
	给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

解题思路：自顶向下归并排序
	找中点（快慢双指针）
	链表排序
	合并排序链表
时间复杂度：O（N logN），N为链表长度
空间复杂度：O（logN），N为链表长度

class Solution {
    public ListNode sortList(ListNode head) {
        return sortList(head, null);
    }
    // 链表排序(二分法)
    public ListNode sortList(ListNode head, ListNode tail) {
        // 递归出口(节点个数 == 0 、1)
        if(head == null || head.next == null) return head;

        ListNode mid = mid(head);
        ListNode node = mid.next; // 松手前先保存
        mid.next = null;

        ListNode L1 = sortList(head, mid);
        ListNode L2 = sortList(node, tail);
        return merge(L1, L2);
    }
    // 找中点
    ListNode mid(ListNode head) {
        ListNode p = head;
        ListNode q = head;
        while(q.next != null && q.next.next != null) {
            p = p.next;
            q = q.next.next;
        }
        return p;
    }
    // 链表合并(递归法)
    ListNode merge(ListNode A, ListNode B) {

        if(A == null) return B;
        if(B == null) return A;

        if(A.val < B.val) {
            A.next = merge(A.next, B);
            return A;
        }else {
            B.next = merge(A, B.next);
            return B;
        }
    }
}
*/

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := mid(head)
	node := mid.Next

	l1 := sort(head, mid)
	l2 := sort(node, tail)
	return merge(l1, l2)
}


// 快慢指针找中点
func mid(head *ListNode) *ListNode {
	quick := head
	slow := head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}

// 递归合并链表
func merge(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		a.Next = merge(a.Next,b)
		return a
	}else {
		b.Next = merge(a, b.Next)
		return b
	}
}
/*
   // 链表合并(递归法)
   ListNode merge(ListNode A, ListNode B) {

       if(A == null) return B;
       if(B == null) return A;

       if(A.val < B.val) {
           A.next = merge(A.next, B);
           return A;
       }else {
           B.next = merge(A, B.next);
           return B;
       }
   }
*/

