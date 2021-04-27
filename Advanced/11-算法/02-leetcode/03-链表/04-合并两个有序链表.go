package _3_链表

func mergeTwoListNode(Node1, Node2 *ListNode) *ListNode {
	pre := &ListNode{}
	res := pre

	for Node1 != nil && Node2 != nil {
		if Node1.Val < Node2.Val {
			pre.Next = Node1
			Node1 = Node1.Next
		} else {
			pre.Next = Node2
			Node2 = Node2.Next
		}
		pre = pre.Next
	}

	if Node1 != nil {
		pre.Next = Node1
	}
	if Node2 != nil {
		pre.Next = Node2
	}
	return res.Next
}
