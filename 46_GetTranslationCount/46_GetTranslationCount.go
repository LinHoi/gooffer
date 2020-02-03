package main

import (
	"fmt"
	"strconv"
)

func GetTranslationCount(num int) int {
	numStr := strconv.Itoa(num)
	return getTranslationCount(numStr)
}

func getTranslationCount(numStr string) int {
	if len(numStr) == 0 {
		return 1
	}
	if len(numStr) == 1 {
		return 1
	}
	if numStr[0] == '1' {
		return getTranslationCount(numStr[1:]) + getTranslationCount(numStr[2:])
	} else if numStr[0] == '2' && numStr[1] < '6' && numStr[1] >= '0' {
		return getTranslationCount(numStr[1:]) + getTranslationCount(numStr[2:])
	} else {
		return getTranslationCount(numStr[1:])
	}
}

func main() {
	nums := []int{
		12,
		345,
		12434,
		1424344,
		2423432555,
		232154354654,
		255326456575888,
		254364576657577878,
		1223341242414124124,
	}
	for _, num := range nums {
		res := GetTranslationCount(num)
		fmt.Printf("Translation Count of %d is %d \n", num, res)
	}
}
/* output
Translation Count of 12 is 2
Translation Count of 345 is 1
Translation Count of 12434 is 3
Translation Count of 1424344 is 4
Translation Count of 2423432555 is 8
Translation Count of 232154354654 is 6
Translation Count of 255326456575888 is 2
Translation Count of 254364576657577878 is 2
Translation Count of 1223341242414124124 is 540
 */