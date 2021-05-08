package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func (node *ListNode) ShowListNode() {
	if node == nil {
		fmt.Println("node 为 nil")
		return
	}
	for node != nil {
		fmt.Printf("%v -> ", node.Val)
		node = node.Next
	}
	fmt.Println("NULL")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
