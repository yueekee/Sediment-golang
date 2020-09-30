package main

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
说明：
给定的 n 保证是有效的。

进阶：
你能尝试使用一趟扫描实现吗？
*/

func main() {

}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// sentinel -> head -> 2 -> 3 -> 4
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{}
	sentinel.Next = head

	var pre *ListNode
	cur := sentinel
	index := 1

	for head != nil {
		if index >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		index ++
	}

	pre.Next = pre.Next.Next
	return sentinel.Next
}

