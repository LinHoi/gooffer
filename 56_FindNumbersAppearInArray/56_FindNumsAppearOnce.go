package main

import "fmt"

func XorAll(array []int) int {
	res := array[0]
	for i := 1; i < len(array); i++ {
		res ^= array[i]
	}
	return res
}
func FindFirstOneInBinaryNumber(num int) int {
	index := 0
	for (num & 1) == 0 {
		num = num >> 1
		index ++
	}
	return index
}
func HasOneInIndexBit(num int, index uint) bool {
	num = num >> index
	if num & 1 == 0 {
		return false
	} else {
		return true
	}
}

func FindNumsAppearOnce(array []int) (int,int) {
	xorAll := XorAll(array)
	firstOneIndex := uint(FindFirstOneInBinaryNumber(xorAll))
	var appearOnceA,appearOnceB int
	for i := range array {
		if HasOneInIndexBit(array[i], firstOneIndex) {
			appearOnceA ^= array[i]
		} else {
			appearOnceB ^= array[i]
		}
	}
	return appearOnceA, appearOnceB
}
func main() {
	arrays := [][]int{
		{1,2,3,4,1,2},
		{2,3,2,3,4,5,6,6,7,7},
		{1,2,3,4,5,6,7,8,9,0,10,11,12,12,11,10,9,8,7,6,5,4,3,2},
	}
	for _, array := range arrays {
		a, b := FindNumsAppearOnce(array)
		fmt.Printf("numbers appear once in %v are %d and %d\n",array,a,b)
	}
}
/* output
numbers appear once in [1 2 3 4 1 2] are 3 and 4
numbers appear once in [2 3 2 3 4 5 6 6 7 7] are 5 and 4
numbers appear once in [1 2 3 4 5 6 7 8 9 0 10 11 12 12 11 10 9 8 7 6 5 4 3 2] are 1 and 0
 */