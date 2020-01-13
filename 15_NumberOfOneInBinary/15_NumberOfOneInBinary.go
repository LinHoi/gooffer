package main

import "fmt"

func NumberOfOneInBinary(num int) int {
	count := 0
	for num != 0 {
		count ++
		num = (num-1) & num
	}
	return count
}

func main(){
	for i := 0; i <= 1024; i++ {
		count := NumberOfOneInBinary(i)
		fmt.Printf("binary number %b has %d one\n",i,count)
	}
}