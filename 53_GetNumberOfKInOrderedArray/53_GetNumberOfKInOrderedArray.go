package main

import "fmt"

func GetFirstK(array []int, k int, low, high int) int {
	if low > high {
		return -1
	}
	if low == high {
		if array[low] == k{
			return low
		} else {
			return -1
		}
	}

	middle := (low + high) / 2
	if array[middle] == k && array[middle-1] != k {
		return middle
	} else if array[middle] >= k && array[middle-1] >= k {
		return GetFirstK(array,k,low,middle-1)
	} else {
		return GetFirstK(array, k,middle+1, high)
	}
}
func GetLastK(array []int, k int , low, high int) int {
	if low > high {
		return -1
	}
	if low == high {
		if array[low] == k{
			return low
		} else {
			return -1
		}
	}
	middle := (low + high) / 2
	if array[middle] == k && array[middle+1] != k {
		return middle
	} else if array[middle] <= k && array[middle+1] <= k {
		return GetLastK(array, k,middle+1,high)
	} else {
		return GetLastK(array, k,low,middle-1)
	}
}

func GetNumberOfKInOrderedArray(array []int, k int) int {
	last := GetLastK(array, k,0,len(array)-1)
	if last == -1 {
		return 0
	} else {
		first := GetFirstK(array, k,0,len(array)-1)
		return last - first + 1
	}
}

func main() {
	arrays := [][]int{
		{1, 2, 3, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7},
		{1, 2, 3, 4, 5},
		{},
		{1},
		{5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5},
	}
	targetK := 5
	for _, array := range arrays {
		res := GetNumberOfKInOrderedArray(array, targetK)
		if res == -1 {
			fmt.Printf("%d doesn't exist in %v\n ", targetK, array)
		} else {
			fmt.Printf("the number of %d in %v is %d\n", targetK, array, res)
		}
	}
}
/* output
the number of 5 in [1 2 3 4 4 5 5 5 5 6 6 6 6 6 7 7 7 7] is 4
the number of 5 in [1 2 3 4 5] is 1
the number of 5 in [] is 0
the number of 5 in [1] is 0
the number of 5 in [5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5] is 16
 */
