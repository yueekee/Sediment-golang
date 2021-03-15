package _3_链表

/*
反转一个单链表。
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归的方式
// 1->2->3->4->5
func reverseList2(head *ListNode) *ListNode {
	// 递归的条件是，当前为空，或者下一个节点为空
	if head == nil || head.Next == nil {
		return head
	}

	// cur 就是最后一个节点
	cur := reverseList2(head.Next)
	// 对链表1->2->3->4->5，cur=5
	// head=4，4.Next.Next = 5.Next，这里即5->4
	head.Next.Next = head
	// 防止列表循环，将head.Next设置为空
	head.Next = nil
	// 返回cur，即最后一个节点
	return cur
}
