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
func PrintLoopList(root *ListNode) {
	isPrinted := make(map[*ListNode]bool)
	for root != nil && !isPrinted[root] {
		fmt.Printf("%d --> ", root.Value)
		isPrinted[root] = true
		root = root.Next
	}

	fmt.Println(root.Value)
}
func PrintList(root *ListNode) {
	for root != nil {
		fmt.Printf("%d --> ", root.Value)
		root = root.Next
	}
	fmt.Println(" nil")
}

func HasLoop(root *ListNode)(meetNode *ListNode, hasLoop bool) {
	if root == nil || root.Next == nil {
		return nil, false
	}
	slowPointer , fastPointer := root.Next, root.Next.Next
	for slowPointer != fastPointer {
		if fastPointer == nil || fastPointer.Next == nil || fastPointer.Next.Next == nil {
			return nil, false
		}
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}
	return slowPointer , true
}

func LengthOfLoop(meetNode *ListNode) int {
	count := 0
	startNode := meetNode.Next
	count = 1
	for meetNode != startNode {
		count ++
		startNode = startNode.Next
	}
	return count
}

func EntryNodeOfLoop(root *ListNode) (*ListNode, error) {
	meetNode, hasLoop := HasLoop(root)
	if !hasLoop {
		return nil, errors.New("not loop in this list")
	}
	length := LengthOfLoop(meetNode)

	slowPointer := root
	fastPointer := root
	for i := 1; i <= length; i ++ {
		fastPointer = fastPointer.Next
	}

	for slowPointer != fastPointer {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}

	return slowPointer, nil
}

func mainFunc(root *ListNode){
	entry, err := EntryNodeOfLoop(root)
	fmt.Print("List: ")

	if err != nil {
		PrintList(root)
		fmt.Println(err)
	}else {
		PrintLoopList(root)
		fmt.Printf("entry pointer of list is %v with value %d \n",entry,entry.Value)
	}
}

func main() {
	node8 := NewNode(8, nil)
	node7 := NewNode(7, node8)
	node6 := NewNode(6, node7)
	node5 := NewNode(5, node6)
	node4 := NewNode(4, node5)
	node3 := NewNode(3, node4)
	node8.Next = node3
	node2 := NewNode(2, node3)
	node1 := NewNode(1, node2)
	root := NewNode(0, node1)
	mainFunc(root)
	node8.Next = nil
	mainFunc(root)
	node8.Next = root
	mainFunc(root)
	node8.Next = node8
	mainFunc(root)
}

/* OutPut
List: 0 --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 --> 3
entry pointer of list is &{3 0x246c128} with value 3
List: 0 --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 -->  nil
not loop in this list
List: 0 --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 --> 0
entry pointer of list is &{0 0x246c140} with value 0
List: 0 --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> 8 --> 8
entry pointer of list is &{8 0x246c108} with value 8
*/