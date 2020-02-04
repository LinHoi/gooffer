package main

import "fmt"

func Min(a,b,c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func GetUglyNumber(index int) int {
	if index <= 0 {
		return 0
	}
	uglyNumbers := make([]int, index)
	uglyNumbers[0] = 1
	multiply2Index, multiply3Index, multiply5Index := 0, 0, 0
	nextUglyIndex := 1
	for nextUglyIndex < index {
		min := Min(uglyNumbers[multiply2Index]*2,uglyNumbers[multiply3Index]*3, uglyNumbers[multiply5Index]*5)
		uglyNumbers[nextUglyIndex] = min
		nextUglyIndex ++

		for uglyNumbers[multiply2Index]*2 <= min {
			multiply2Index ++
		}
		for uglyNumbers[multiply3Index]*3 <= min {
			multiply3Index ++
		}
		for uglyNumbers[multiply5Index]*5 <= min {
			multiply5Index ++
		}

	}
	return uglyNumbers[index-1]
}

func main() {
	indexes := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		10,
		34,
		244,
		1500,
		2530,
	}
	for _, index := range indexes {
		res := GetUglyNumber(index)
		fmt.Printf("the %d th ugly number is %d\n",index,res)
	}
}
/* output
the 1 th ugly number is 1
the 2 th ugly number is 2
the 3 th ugly number is 3
the 4 th ugly number is 4
the 5 th ugly number is 5
the 6 th ugly number is 6
the 7 th ugly number is 8
the 10 th ugly number is 12
the 34 th ugly number is 100
the 244 th ugly number is 34992
the 1500 th ugly number is 859963392
the 2530 th ugly number is 59719680000
 */