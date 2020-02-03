package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isGreater(source, target int) bool {
	if source == target {
		return false
	}
	sourcestr := strconv.Itoa(source)
	targetStr := strconv.Itoa(target)
	sourceLen := len(sourcestr)
	targetLen := len(targetStr)
	var newLen int
	if sourceLen <= targetLen {
		newLen = sourceLen
	} else {
		newLen = targetLen
	}
	index := 0
	for newLen - index > 0 {
		if sourcestr[index] > targetStr[index] {
			return true
		} else if sourcestr[index] < targetStr[index] {
			return false
		} else {
			index ++
		}
	}

	if sourceLen <= targetLen {
		return true
	} else {
		return false
	}
}

func isLess(source, target int ) bool {
	return !isGreater(source,target)
}

func NewSort(nums []int)  {
	sort.Slice(nums,func (i, j int) bool {
		return isLess(nums[i],nums[j])
	})
}

func PrintSlice(nums []int) {

	for _, num := range nums {
		fmt.Print(num)
	}
}
func main() {
	numss := [][]int{
		{1, 2, 34, 566, 87, 55, 57, 778, 457, 44, 55, 45, 456, 4567, 688},
		{12,35,57,687,344,6757,354,68,36456,6868,36456,7868,36546,786,34564,9879,4565886,6575,465756,475889,3577768,467},
		{0},
		{},
	}
	for _, nums := range numss{
		NewSort(nums)
		fmt.Printf("Min Number of %v is:\n   ",nums)
		PrintSlice(nums)
		fmt.Println()
	}
}
/* output
Min Number of [1 2 34 44 4567 456 457 45 55 55 566 57 688 778 87] is:
   12344445674564574555555665768877887
Min Number of [12 344 34564 354 3577768 35 36456 36456 36546 4565886 465756 467 475889 57 6575 6757 6868 687 68 7868 786 9879] is:
   12344345643543577768353645636456365464565886465756467475889576575675768686876878687869879
Min Number of [0] is:
   0
Min Number of [] is:
 */