package _3_链表

import (
	"fmt"
)

/*
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。

提示：

链表至少包含两个节点。
链表中所有节点的值都是唯一的。
给定的节点为非末尾节点并且一定是链表中的一个有效节点。
不要从你的函数中返回任何结果。
*/

// O(1)
func (node *ListNode) DeleteNode() {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 添加最后一个节点
func (node *ListNode) AddLastNode(val int) {
	if node == nil {
		node.Val = val
		node.Next = nil
		return
	}

	for node.Next != nil {
		node = node.Next
	}

	node.Next = new(ListNode)	// 下一个节点需要重新定义，不然会报空指针错误
	node.Next.Val = val
	node.Next.Next = nil
}

func (node *ListNode) PrintListNode() {
	for node.Next == nil {
		fmt.Println(node.Val, " ")
	}
}
