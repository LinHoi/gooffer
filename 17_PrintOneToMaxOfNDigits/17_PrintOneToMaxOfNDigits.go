package main

import "fmt"

func PrintOneToMaxOfNDigits(n int64) {
	if n <= 0 {
		return
	}
	number := make([]rune, n)

	for i := 0; i < 10; i++ {
		number[0] = rune(i)
		PrintOneToMaxOfNDigitsRecursively(number, n, 0)
	}
}

func PrintOneToMaxOfNDigitsRecursively(number []rune, length int64, index int64) {
	if index == (length - 1) {
		PrintNumber(number)
		return
	}
	for i := 0; i < 10; i++ {
		number[index+1] = rune(i)
		PrintOneToMaxOfNDigitsRecursively(number, length, index+1)
	}

}

func PrintNumber(number []rune) {
	isBeginning0 := true
	length := len(number)
	for i := 0; i < length; i++ {
		if isBeginning0 && number[i] != rune(0) {
			isBeginning0 = false
		}
		if !isBeginning0 {
			fmt.Printf("%v", number[i])
		}
	}
	fmt.Println()
}

func main() {
	var n int64 = 3
	PrintOneToMaxOfNDigits(n)
}
