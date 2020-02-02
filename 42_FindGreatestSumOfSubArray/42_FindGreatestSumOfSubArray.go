package main

import "fmt"

func FindGreatestSumOfSubArray(num []int) int {
	f := make([]int, len(num))
	f[0] = num[0]
	for i := 1; i < len(num); i++ {
		if f[i-1] >= 0 {
			f[i] = f[i-1] + num[i]
		} else {
			f[i] = num[i]
		}
	}

	max := f[0]
	for i := 1; i < len(num); i++ {
		if f[i] > max {
			max = f[i]
		}
	}
	return max
}


func main() {
	nums := [][]int {
		{1,2,3,4,-4,5,-3,-5,6,-99,55,-43},
		{1,2,3,4,5,6,7,8,9,0},
		{-1,-2,-3,-4,-5,-6,9,0,4},
		{1},
	}
	for _, num := range nums {
		res := FindGreatestSumOfSubArray(num)
		fmt.Printf("greatest sum of sub array in %v is %d\n",num,res)
	}
}
/* output
greatest sum of sub array in [1 2 3 4 -4 5 -3 -5 6 -99 55 -43] is 55
greatest sum of sub array in [1 2 3 4 5 6 7 8 9 0] is 45
greatest sum of sub array in [-1 -2 -3 -4 -5 -6 9 0 4] is 13
greatest sum of sub array in [1] is 1

*/