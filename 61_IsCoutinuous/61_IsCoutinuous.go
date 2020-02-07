package main

import "fmt"

const (
	pockerNumber  = 5
)
func IsCoutinuous(pokers []int) bool {
	countMap := make([]int,15)
	for i := range pokers {
		countMap[pokers[i]]++
	}

	index := 1
	for countMap[index] == 0 {
		index ++
	}
	startIndex := index
	for index < (startIndex + pockerNumber) {
		if countMap[index] >= 2 {
			return false
		} else if countMap[index] == 1 {
			index ++
		} else {
				if countMap[0] == 0 {
					return false
				} else {
					countMap[0] --
					index ++
				}
		}
	}
	return true
}

func main() {
	pokers := [][]int{
		{0,0,1,2,3,},
		{2,4,6,8,0},
		{10,11,12,13,0},
		{1,1,1,1,1},
		{0,1,2,3,4},
		{13,13,13,13,13},
	}
	for _, poker := range pokers {
		res := IsCoutinuous(poker)
		if res {
			fmt.Printf("%v is continuous\n",poker)
		} else {
			fmt.Printf("%v is not continuous\n",poker)
		}
	}
}
/* output
[0 0 1 2 3] is continuous
[2 4 6 8 0] is not continuous
[10 11 12 13 0] is continuous
[1 1 1 1 1] is not continuous
[0 1 2 3 4] is continuous
[13 13 13 13 13] is not continuous
 */