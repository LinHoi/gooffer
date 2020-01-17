package main

import "fmt"

func IsEven(num int) bool {
	return num & 1 == 1
}
func isDevideBy3(num int) bool {
	return num % 3 == 0
}

func ReOrder(array []int,isFunc  func(int) bool) {
	low, high := 0, len(array)-1
	for low < high {
		for isFunc(array[low]) {
			low ++
		}
		for !isFunc(array[high]){
			high --
		}
		if low < high {
			array[low], array[high] = array[high], array[low]
		}
	}
}

func main() {
	arrays := [][]int{
		{1,23,4,5,6,7,5,6,78,46,3,5,9},
		{123,43,46,4,68,4,63,787,445,365,3},
		{12,3,43,5,56,563,3,5787,433,756,345,57,66,3445,4454,65},
	}
	for _ ,array := range arrays {
		ReOrder(array,IsEven)
		fmt.Printf("ReOrder array by even is %v \n", array)
	}
	for _ ,array := range arrays {
		ReOrder(array,isDevideBy3)
		fmt.Printf("ReOrder array by being able to be divided by 3 is %v \n", array)
	}
}
/* OutPut
ReOrder array by even is [1 23 9 5 5 7 5 3 78 46 6 6 4]
ReOrder array by even is [123 43 3 365 445 787 63 4 68 4 46]
ReOrder array by even is [65 3 43 5 3445 563 3 5787 433 57 345 756 66 56 4454 12]
ReOrder array by being able to be divided by 3 is [6 6 9 78 3 7 5 5 5 46 23 1 4]
ReOrder array by being able to be divided by 3 is [123 63 3 365 445 787 43 4 68 4 46]
ReOrder array by being able to be divided by 3 is [12 3 66 756 345 57 3 5787 433 563 3445 5 43 56 4454 65]
*/