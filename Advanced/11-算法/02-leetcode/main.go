package main

import listnode "github.com/liankui/Sediment-golang/Advanced/11-算法/02-leetcode/03-链表"

func main() {
	node := listnode.NewListNode()
	//node := new(listnode.ListNode)
	node.AddLastNode(1)
	node.AddLastNode(2)
	node.AddLastNode(3)
	node.ShowListNode()
}
