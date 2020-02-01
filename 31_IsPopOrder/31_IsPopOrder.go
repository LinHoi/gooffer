package main

import "fmt"

func IsPopOrder(push []int, pop []int) bool {
	if len(push) == 0 && len(pop) == 0 {
		return true
	}
	if len(push) != len(pop) {
		return false
	}

	stack := NewStack()
	pushIndex := 0
	popIndex := 0
	for popIndex < len(pop) {
		topValue, err := stack.Top()
		if err != nil {
			if pushIndex < len(push) {
				stack.Push(push[pushIndex])
				pushIndex++
			} else {
				return false
			}
		} else if topValue == pop[popIndex] {
			stack.Pop()
			popIndex++
		} else {
			for topValue != pop[popIndex] {
				if pushIndex >= len(push) {
					return false
				}
				stack.Push(push[pushIndex])
				topValue, _ = stack.Top()
				pushIndex++
			}
		}
	}
	return true
}

func main() {
	arrays := [][2][]int{
		{{1,2,3,4,5},{4,5,3,2,1}},
		{{1,2,3,4,5},{4,3,5,1,2}},
		{{},{}},
		{{1},{}},
		{{1,2,3,4,5,6,7,8,9,0},{0,9,8,7,6,5,4,3,2,1}},
	}
	for _, array := range arrays {
		result := IsPopOrder(array[0],array[1])
		if result == true {
			fmt.Printf("%v is a push seq of stack %v \n",array[1],array[0])
		}else {
			fmt.Printf("%v is not a push seq of stack %v \n",array[1],array[0])
		}
	}
}

/*
[4 5 3 2 1] is a push seq of stack [1 2 3 4 5]
[4 3 5 1 2] is not a push seq of stack [1 2 3 4 5]
[] is a push seq of stack []
[] is not a push seq of stack [1]
[0 9 8 7 6 5 4 3 2 1] is a push seq of stack [1 2 3 4 5 6 7 8 9 0]
*/