package main

import (
	"fmt"
	"strconv"
)
func PowerBase10(num int) int {
	res := 1
	for i:= 0; i< num; i++ {
		res = res *10
	}
	return res
}
func NumberOf1Between1AndN(n int) int {
	if n <= 0 {
		return 0
	}
	strN := strconv.Itoa(n)
	return NumberOf1(strN)
}

func NumberOf1(strN string) int {
	length := len(strN)
	var first int
	first = int(strN[0] - '0')
	//fmt.Println(first)

	if length == 1 && first == 0 {
		return 0
	}
	if length == 1 && first > 0 {
		return 1
	}

	numFirstDigit := 0
	if first > 1 {
		numFirstDigit = PowerBase10(length - 1)
	} else if first == 1 {
		numFirstDigit, _ = strconv.Atoi(strN[1:])
		numFirstDigit ++
	}

	numberOtherDigits := first * (length - 1) * PowerBase10(length-2)
	numbrRecursive := NumberOf1(strN[1:])
	return numFirstDigit + numberOtherDigits + numbrRecursive
}

func main() {
	nums := []int{
		1,
		10,
		100,
		244,
		4345,
		54556,
		35365765756,
		2355456576768878,

	}
	for _ , num := range nums {
		res := NumberOf1Between1AndN(num)
		fmt.Printf("Number of 1 between 1 and %d is %d\n",num,res)
	}
}
/*
Number of 1 between 1 and 1 is 1
Number of 1 between 1 and 10 is 2
Number of 1 between 1 and 100 is 21
Number of 1 between 1 and 244 is 155
Number of 1 between 1 and 4345 is 2375
Number of 1 between 1 and 54556 is 32416
Number of 1 between 1 and 35365765756 is 45896486756
Number of 1 between 1 and 2355456576768878 is 4602606964087678
 */