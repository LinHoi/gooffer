package main

import "fmt"

func MaxProductAfterCuttingRope(length int64) int64 {
	if length <= 1 { return 0 }
	if length == 2 { return 1 }
	if length == 3 { return 2 }
	products := make([]int64, length+1)
	products[0],products[1],products[2],products[3] = 0,1,2,3

	for i := int64(4); i <= length ; i++ {
		maxProduct := int64(0)
		for j := int64(1); j <= i/2; j++ {
			product := products[j] * products[i-j]
			if maxProduct < product {
				maxProduct = product
			}
			products[i] = maxProduct
		}
	}
	return products[length]
}

func main() {
	for ropeLength := int64(0); ropeLength < 128 ; ropeLength ++ {
		res := MaxProductAfterCuttingRope(ropeLength)
		fmt.Printf("Max Product of rope with length:%d is %v \n",ropeLength, res)
	}
}
