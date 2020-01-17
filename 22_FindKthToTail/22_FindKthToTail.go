package main

import (
	"errors"
	"fmt"
)

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

func FindKthToTail(root *ListNode, k int) (int ,error){
	if root == nil ||  k <= 0 {
		return 0 ,errors.New("fail to find " +fmt.Sprint(k)+"th from tail")
	}
	slowPointer, fastPointer := root, root
	for i := 1; i <= k; i++ {
		if fastPointer == nil {
			return 0, errors.New(fmt.Sprint(k)+" is longer then the length of list")
		}
		fastPointer = fastPointer.Next
	}
	for fastPointer != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}
	return slowPointer.Value, nil
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
	ks := []int{
		-22,
		-1,
		0,
		1,
		99,
		8,
		4,
		9,
		10,
	}
	fmt.Print("List: ")
	PrintList(root)
	for _, k := range ks {
		res , err := FindKthToTail(root, k)
		if err != nil {
			fmt.Println(err)
		}else {
			fmt.Printf("%dth of list from Tail is %d \n", k, res )
		}

	}
}
/*Output
List: 0 --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 -->  nil
fail to find -22th to tail
fail to find -1th to tail
fail to find 0th to tail
1th of list from Tail is 8
99 is longer then the length of list
8th of list from Tail is 1
4th of list from Tail is 5
9th of list from Tail is 0
10 is longer then the length of list
*/