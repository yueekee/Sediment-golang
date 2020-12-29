package main

import (
	"strings"
)

func main() {

}

// 参考文章：https://juejin.cn/post/6844903798507307015#heading-0

type Node struct {
	tag      string
	id       string
	classes  string
	children []*Node
}

// 广度优先搜索（BFS）
func findById(root *Node, id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.id == id {
			return nextUp
		}
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return nil
}

// 深度优先搜索（DFS）-使用递归
func findByIdDFS(node *Node, id string) *Node {
	if node.id == id {
		return node
	}

	if len(node.children) > 0 {
		for _, child := range node.children {
			findByIdDFS(child, id)
		}
	}
	return nil
}

func (n *Node) hasClass(className string) bool {
	classes := strings.Fields(n.classes)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}

func findAllByClassName(root *Node, className string) []*Node {
	result := make([]*Node, 0)
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.hasClass(className) {
			result = append(result, nextUp)
		}
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return result
}

//func (node *Node) remove() {
//	// Remove the node from it's parents children collection
//	for idx, sibling := range n.parent.children {
//		if sibling == node {
//			node.parent.children = append(
//				node.parent.children[:idx],
//				node.parent.children[idx+1:]...,
//			)
//		}
//	}
//
//	// If the node has any children, set their parent to nil and set the node's children collection to nil
//	if len(node.children) != 0 {
//		for _, child := range node.children {
//			child.parent = nil
//		}
//		node.children = nil
//	}
//}
