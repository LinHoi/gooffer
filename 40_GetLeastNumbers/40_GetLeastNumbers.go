package main

import "fmt"

func Partition(nums []int, start,end int) int {
	target := nums[start]
	for end > start {
		for nums[end] >= target && end > start {
			end --
		}
		nums[start] = nums[end]
		for nums[start] < target && end > start {
			start ++
		}
		nums[end] = nums[start]
	}
	nums[start] = target
	return start
}

func GetLeastNumbers(input []int, k int) (output []int){
	if k > len(input) {
		return nil
	}
	start, end := 0, len(input) - 1
	partition := Partition(input,start,end)

	for partition != k {
		if partition > k {
			partition = Partition(input,start,partition-1)
		} else {
			partition = Partition(input,partition+1, end)
		}
	}

	return input[0:k]
}

func main() {
	inputs := [][]int {
		{1,244,2,5,677,7,88,90,5},
		{344,4,234,2345,6,346,767,545},
		{2,43,45,56,33,45,34,5,45,6,76,89,9,90,345,4545,24536,4554,252345,654,23,7676,475,686,4747,9879,3464,67865,3464},
	}
	for _, input := range inputs {
		output := GetLeastNumbers(input,3)
		fmt.Printf("Least %d num of %v is %v\n",3,input,output)
	}
	for _, input := range inputs {
		output := GetLeastNumbers(input,10)
		fmt.Printf("Least %d num of %v is %v\n",10,input,output)
	}
}

/*
Least 3 num of [1 2 5 5 90 7 88 244 677] is [1 2 5]
Least 3 num of [6 4 234 344 2345 346 767 545] is [6 4 234]
Least 3 num of [2 5 6 9 23 33 34 43 45 45 76 89 56 90 345 4545 24536 4554 252345 654 45 7676 475 686 4747 9879 3464 67865 3464] is [2 5 6]
Least 10 num of [1 2 5 5 90 7 88 244 677] is []
Least 10 num of [6 4 234 344 2345 346 767 545] is []
Least 10 num of [2 5 6 9 23 33 34 43 45 45 45 56 76 90 345 4545 24536 4554 252345 654 89 7676 475 686 4747 9879 3464 67865 3464] is [2 5 6 9 23 33 34 43 45 45]
 */