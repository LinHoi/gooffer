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
func ReverseList(root *ListNode) *ListNode {
	var reversedHead *ListNode
	var preNode *ListNode
	reversedHead = nil
	preNode = nil
	node := root
	for node != nil {
		nextNode := node.Next
		if nextNode == nil {
			reversedHead = node
		}
		node.Next = preNode
		preNode = node
		node = nextNode
	}
	return reversedHead


}

func main() {
	node8 := NewNode(8, nil)
	node7 := NewNode(7, node8)
	node6 := NewNode(6, node7)
	node5 := NewNode(5, node6)
	node4 := NewNode(4, node5)
	node3 := NewNode(3, node4)
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)
	root := NewNode(0, node1)

	fmt.Print("List: ")
	PrintList(root)
	reversedRoot := ReverseList(root)
	fmt.Print("Reverse List: ")
	PrintList(reversedRoot)

	root = nil
	fmt.Print("List: ")
	PrintList(root)
	reversedRoot = ReverseList(root)
	fmt.Print("Reverse List: ")
	PrintList(reversedRoot)

	root = NewNode(0,nil)
	fmt.Print("List: ")
	PrintList(root)
	reversedRoot = ReverseList(root)
	fmt.Print("Reverse List: ")
	PrintList(reversedRoot)
}

/*OutPut
List: 0 --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 -->  nil
Reverse List: 8 --> 7 --> 6 --> 5 --> 4 --> 3 --> 2 --> 1 --> 0 -->  nil
List:  nil
Reverse List:  nil
List: 0 -->  nil
Reverse List: 0 -->  nil

*/