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

func CountOfIntergers(digits int) int {
	if digits == 0 {
		return 0
	}
	if digits == 1 {
		return 10
	}
	return 9 * PowerBase10(digits - 1)
}

func DigitAtIntegerList(index int) int {
	if index < 0 {
		return -1
	}

	digits := 1
	for {
		countOfUnitsDigits := digits * CountOfIntergers(digits)
		if index < countOfUnitsDigits {
			return digitAtIntegerList(index, digits)
		}
		index -= countOfUnitsDigits
		digits ++
	}
}

func beginNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return PowerBase10(digits-1)
}
func digitAtIntegerList(index int,digits int) int {
	beginNumberInDigits := beginNumber(digits)
	targetNumber := beginNumberInDigits + index / digits
	index = index % digits
	for index > 0 {
		targetNumber,_ = strconv.Atoi(strconv.Itoa(targetNumber)[1:])
		index --
	}
	res := int((strconv.Itoa(targetNumber)[0]) - '0')
	return res
}

func main() {
	indexs := []int{
		0,
		1,
		2,
		3,
		10,
		11,
		12,
		13,
		15,
		45,
		667,
		7999,
		575778,
		57577788,
		1243465465765,
	}

	for _,index := range indexs {
		res := DigitAtIntegerList(index)
		fmt.Printf("digit at index %d from integer List is %d\n",index,res)
	}
}
/*
digit at index 0 from integer List is 0
digit at index 1 from integer List is 1
digit at index 2 from integer List is 2
digit at index 3 from integer List is 3
digit at index 10 from integer List is 1
digit at index 11 from integer List is 0
digit at index 12 from integer List is 1
digit at index 13 from integer List is 1
digit at index 15 from integer List is 2
digit at index 45 from integer List is 7
digit at index 667 from integer List is 2
digit at index 7999 from integer List is 2
digit at index 575778 from integer List is 4
digit at index 57577788 from integer List is 8
digit at index 1243465465765 from integer List is 8
 */