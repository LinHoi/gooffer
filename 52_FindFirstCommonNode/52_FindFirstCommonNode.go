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

func FindFirstCommonNode(root1, root2 *ListNode) *ListNode {
	if root1 == nil || root2 == nil {
		return nil
	}
	nodeCounts := make(map[*ListNode]int)
	for root1.Next != nil {
		nodeCounts[root1] ++
		root1 = root1.Next
	}
	for root2.Next != nil {
		nodeCounts[root2] ++
		if nodeCounts[root2] > 1 {
			return root2
		}
		root2 = root2.Next
	}
	return nil
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

	res := FindFirstCommonNode(root1, root2)
	if res != nil {
		fmt.Printf("first common node of root1 and root2 is %v with value %d\n", res,res.Value)
	} else {
		fmt.Println("there is not common node ")
	}

	node4.Next = node7
	res = FindFirstCommonNode(root1, root2)
	if res != nil {
		fmt.Printf("first common node of root1 and root2 is %v with value %d\n", res,res.Value)
	} else {
		fmt.Println("there is not common node ")
	}
}
/* output
there is not common node
first common node of root1 and root2 is &{6 0xc00005a1c0} with value 6
 */
