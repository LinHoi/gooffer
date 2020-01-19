package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}
func NewNode(value int, next *ListNode) *ListNode {
	return &ListNode{value, next}
}
func PrintList(root *ListNode) {
	for root != nil {
		fmt.Printf("%d --> ", root.Value)
		root = root.Next
	}
	fmt.Println(" nil")
}

func MergeList(root1, root2 *ListNode) *ListNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	var preNode = &ListNode{0,nil}
	res := preNode
	nextNode1 , nextNode2 := root1 , root2

	for nextNode1 != nil && nextNode2 != nil{
		if nextNode1.Value <= nextNode2.Value {
			preNode.Next = nextNode1
			preNode = nextNode1
			nextNode1  = nextNode1.Next

		}else {
			preNode.Next = nextNode2
			preNode = nextNode2
			nextNode2 = nextNode2.Next
		}
	}
	if nextNode1 != nil {
		preNode.Next = nextNode1
	}
	if nextNode2 != nil {
		preNode.Next = nextNode2
	}
	return res.Next
}

func main() {
	node8 := NewNode(8, nil)
	node7 := NewNode(6, node8)
	node6 := NewNode(4, node7)
	root2 := NewNode(2, node6)


	node4 := NewNode(9, nil)
	node3 := NewNode(7, node4)
	node2 := NewNode(5, node3)
	node1 := NewNode(3, node2)
	root1 := NewNode(1, node1)
	fmt.Print("List 1: ")
	PrintList(root1)
	fmt.Print("List 2: ")
	PrintList(root2)
	newRoot := MergeList(root1, root2)
	fmt.Print("Merged List: ")
	PrintList(newRoot)

	var root3 *ListNode = nil
	fmt.Print("List 3: ")
	PrintList(root3)
	newRoot = MergeList(newRoot,root3)
	fmt.Print("Merged List: ")
	PrintList(newRoot)
}

/* OutPut
List 1: 1 --> 3 --> 5 --> 7 --> 9 -->  nil
List 2: 2 --> 4 --> 6 --> 8 -->  nil
Merged List: 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 --> 9 -->  nil
List 3:  nil
Merged List: 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 --> 9 -->  nil
*/