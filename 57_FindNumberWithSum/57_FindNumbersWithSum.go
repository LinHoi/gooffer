package main

import (
	"fmt"
	"errors"
	"strconv"
)

func FindNumbersWithSum(nums []int,sum int) (ans [2]int, err error){
	low ,high := 0, len(nums)-1
	for nums[low] + nums[high] != sum && low < high {
		if nums[low] + nums[high] > sum {
			high --
		} else if nums[low] + nums[high] < sum {
				low ++
		}
	}
	if nums[low] + nums[high] == sum {
		ans[0], ans[1] = nums[low], nums[high]
		return ans, nil
	} else {
		return ans, errors.New("can't find numbers with such sum: "+strconv.Itoa(sum))
	}
}

func main() {
	nums := []int{1,2,3,4,5,5,6,7,8,9,9,10,11,12,12,23,45,67,100}
	sums := []int{1,2,3,6,10,21,23,31,46,77,101,1000,0}
	for _, sum := range sums {
		res, err := FindNumbersWithSum(nums,sum)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("sum of %d and %d is %d\n",res[0],res[1], sum)
		}
	}
}
/* output
sum of 1 and 1 is 2
sum of 1 and 2 is 3
sum of 1 and 5 is 6
sum of 1 and 9 is 10
sum of 9 and 12 is 21
sum of 11 and 12 is 23
sum of 8 and 23 is 31
sum of 1 and 45 is 46
sum of 10 and 67 is 77
sum of 1 and 100 is 101
can't find numbers with such sum: 1000
can't find numbers with such sum: 0
 */
