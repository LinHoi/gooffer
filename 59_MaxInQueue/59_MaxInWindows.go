package main

import (
	"../datastrature/dqueue"
	"fmt"
)

func MaxInWindows(nums []int, windowSize int) []int {
	if windowSize == 0 || windowSize > len(nums) {
		return nil
	}
	dq := dqueue.New()
	resArray := make([]int, len(nums)+1-windowSize)
	index := 0
	for i := 0; i < len(nums); i++ {
		if i >= windowSize-1 {
			resArray[index] = nums[dq.Front()]
			index ++
		}
		if dq.IsEmpty() || dq.Back() >= nums[i] {
			dq.PushBack(i)
		} else if dq.Back() < nums[i] {
			for !dq.IsEmpty() && dq.Back() < nums[i] {
				dq.PopBack()
			}
			dq.PushBack(i)
		}
		for !dq.IsEmpty() && (i-dq.Front()) >= windowSize {
			dq.PopFront()
		}

	}

	return resArray
}

func MaxInWindows2(nums []int, windowSize int) []int {
	if windowSize == 0 || windowSize > len(nums) {
		return nil
	}
	dq := dqueue.New()
	resArray := make([]int, len(nums)+1-windowSize)
	index := 0
	for i := 0; i < windowSize; i++ {
		for !dq.IsEmpty() && nums[i] >= nums[dq.Back()] {
			dq.PopBack()
		}
		dq.PushBack(i)
	}

	for i := windowSize; i < len(nums); i++ {
		resArray[index] = nums[dq.Front()]
		index ++

		for !dq.IsEmpty() && nums[i] >= nums[dq.Back()] {
			dq.PopBack()
		}
		if !dq.IsEmpty() && dq.Front() <= (i-windowSize) {
			dq.PopFront()
		}
		dq.PushBack(i)
	}
	resArray[index] = nums[dq.Front()]
	return resArray
}

func main() {
	numss := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 3, 2, 4, 3, 7, 4},
		{1, 5},
		{4, 5, 7, 9, 6, 4, 3, 2, 4, 6, 8, 6, 4, 3, 6, 89, 5, 3, 5, 3, 57},
		{},
		{1},
	}
	windowSize := 4
	for _, nums := range numss {
		resArray := MaxInWindows2(nums, windowSize)
		fmt.Printf("max in %v with window size %d is %v\n", nums, windowSize, resArray)
	}
}
/*
max in [1 2 3 4 5 6 7 8] with window size 4 is [4 5 6 7 8]
max in [1 3 2 4 3 7 4] with window size 4 is [4 4 7 7]
max in [1 5] with window size 4 is []
max in [4 5 7 9 6 4 3 2 4 6 8 6 4 3 6 89 5 3 5 3 57] with window size 4 is [9 9 9 9 6 4 6 8 8 8 8 6 89 89 89 89 5 57]
max in [] with window size 4 is []
max in [1] with window size 4 is []
 */