package main

import "fmt"

func main(){
	stack := NewStack()
	if min , err := stack.Min(); err == nil {
		fmt.Printf("Min in stack is %d\n",min)
	}else {
		fmt.Println(err)
	}

	stack.Push(99999)
	if min , err := stack.Min(); err == nil {
		fmt.Printf("Min in stack is %d\n",min)
	}else {
		fmt.Println(err)
	}

	stack.Push(999)
	if min , err := stack.Min(); err == nil {
		fmt.Printf("Min in stack is %d\n",min)
	}else {
		fmt.Println(err)
	}

	stack.Push(3)
	if min , err := stack.Min(); err == nil {
		fmt.Printf("Min in stack is %d\n",min)
	}else {
		fmt.Println(err)
	}

	stack.Push(1)
	if min , err := stack.Min(); err == nil {
		fmt.Printf("Min in stack is %d\n",min)
	}else {
		fmt.Println(err)
	}

	stack.Pop()
	if min , err := stack.Min(); err == nil {
		fmt.Printf("Min in stack is %d\n",min)
	}else {
		fmt.Println(err)
	}
}

/*OutPut
stack is empty
Min in stack is 99999
Min in stack is 999
Min in stack is 3
Min in stack is 1
Min in stack is 3
*/
