package main

import (
	"fmt"
)

// a bad solution
func MinInRotatedArray(array []int) (minIndex  int, err error){
	if array[0] < array[len(array) -1] {
		return 0 , nil
	}
	low := 0
	hight := len(array)
	for low <= hight {
		middle := (low + hight) / 2
		if array[middle]<=array[middle+1] && array[middle] < array[middle-1]{
			return middle, nil
		}else {
			if array[middle] > array[middle+1] {
				low = middle + 1
			}else {
				hight = middle
			}
		}
	}
	return hight , nil
}

func MinInRotatedArray2(array []int) (minIndex int, err error){
	low, mid, high := 0, 0, len(array)-1
	for array[low] >= array[high] {
		if high - low  == 1 {
			return high, nil
		}
		mid = (low + high) / 2
		if array[low] == array[high] && array[low] == array[mid] {
			return findMinSimple(array), nil
			//return -1, errors.New("can not deal with such array ", )
		}
		if array[mid] >= array[low] {
			low = mid
		}else if array[mid] <= array[high] {
			high = mid
		}

	}
	return mid, nil
}

func findMinSimple(array []int) (index int) {
	for i := range array {
		if (i+1) <= len(array)-1 && array[i] > array[i+1]{
			return i+1
		}
	}
	return 0
}


func main() {
	arrays := [6][]int{
		{3, 4, 5, 6, 1, 2, 3},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 1, 1, 1, 1, 1, 1, 2, 1},
		{1, 1, 2, 2, 3, 3, 4, 4, 5},
		{3, 3, 3, 3, 1, 1, 1, 1},
		{1,1,1,1,1,1,1,1,1,1,1},
	}
	for _, array := range arrays {
		index, err := MinInRotatedArray2(array)
		if err != nil {
			fmt.Println(err.Error(),array)
		} else {
			fmt.Printf("min in array : %v is array[%d]= %d \n", array, index, array[index])
		}
	}
}

