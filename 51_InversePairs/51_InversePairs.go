package main

import "fmt"

func InversePairs(array []int) int {
	copyArray := make([]int, len(array))
	for i := range array {
		copyArray[i] = array[i]
	}
	return InversePairsCore(array, copyArray)
}

func InversePairsCore(array, copyArray []int) int {
	if len(array) <= 1 {
		return 0
	}

	middle := len(array) / 2
	left := InversePairsCore(array[:middle], copyArray)
	right := InversePairsCore(array[middle:], copyArray)


	arrayLength := len(array)
	for i := 0; i< arrayLength; i++ {
		copyArray[i] = array[i]
	}
	count := 0
	leftIndex, rightIndex, newIndex := middle-1, len(array)-1, len(array)-1
	for leftIndex >= 0 && rightIndex >= middle {
		if copyArray[leftIndex] > copyArray[rightIndex] {
			count += len(array) - middle
			array[newIndex] = copyArray[leftIndex]
			newIndex --
			leftIndex --
		} else {
			array[newIndex] = copyArray[rightIndex]
			newIndex --
			rightIndex --
		}
	}
	for leftIndex >= 0 {
		array[newIndex] = copyArray[leftIndex]
		newIndex --
		leftIndex --
	}
	for rightIndex >= middle  {
		array[newIndex] = copyArray[rightIndex]
		newIndex --
		rightIndex --
	}

	return count + left + right
}

func main() {
	arrays := [][]int{
		{3,4,5,6,7,8,9},
		{3,2,1},
		{2,1},
		{9,8,7,6,5,4,3,2,1},
		{},
	}
	for _, array := range arrays {
		fmt.Print(array)
		res := InversePairs(array)
		fmt.Printf(": %d Inverse Pairs \n",res)
	}
}
/* output
[3 4 5 6 7 8 9]: 0 Inverse Pairs
[3 2 1]: 3 Inverse Pairs
[2 1]: 1 Inverse Pairs
[9 8 7 6 5 4 3 2 1]: 36 Inverse Pairs
[]: 0 Inverse Pairs
 */