package main

import "fmt"

func MoreThenHalfNum(num []int) int {
	result := num[0]
	count := 1
	for i := 1; i < len(num); i++ {
		if num[i] == result {
			count++
		} else {
			count--
		}
		if count == 0 {
			result = num[i]
			count = 1
		}

	}
	return result
}

func main() {
	nums := [][]int{
		{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5},
		{1},
		{1,2,2},
	}
	for _, num := range nums {
		result := MoreThenHalfNum(num)
		fmt.Println(result)
	}
}
/*
5
1
2
 */
