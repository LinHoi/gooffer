package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewNode(value int, next *ListNode) *ListNode {
	return &ListNode{value, next}
}

func DeleteNode(listHead **ListNode, toBeDeleted *ListNode) {
	if listHead == nil || toBeDeleted == nil {
		return
	}
	if toBeDeleted.Next != nil {
		next := toBeDeleted.Next
		toBeDeleted.Value = next.Value
		toBeDeleted.Next = next.Next

		next = nil
	} else if *listHead == toBeDeleted {
		toBeDeleted = nil
		*listHead = nil
	} else {
		node := *listHead
		for node.Next != toBeDeleted {
			node = node.Next
		}
		node.Next = nil
		toBeDeleted = nil
	}
}
func PrintList(root *ListNode) {
	for root != nil {
		fmt.Printf("%d --> ", root.Value)
		root = root.Next
	}
	fmt.Println(" nil")
}

func main() {
	node4 := NewNode(4, nil)
	node3 := NewNode(3, node4)
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)
	root  := NewNode(0, node1)
	fmt.Println("Origin List:")
	PrintList(root)
	fmt.Println("List after delete node 2:")
	DeleteNode(&root, node2)
	PrintList(root)
	DeleteNode(&root, root)
	fmt.Println("List after delete root:")
	PrintList(root)
	fmt.Println("List after delete tail")
	DeleteNode(&root, node4)
	PrintList(root)
}

/* output
Origin List:
0 --> 1 --> 2 --> 3 --> 4 -->  nil
List after delete node 2:
0 --> 1 --> 3 --> 4 -->  nil
List after delete root:
1 --> 3 --> 4 -->  nil
List after delete tail
1 --> 3 -->  nil
*/
