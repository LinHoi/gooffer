package main

import "fmt"

func MaxProfit(prices []int) int {
	length := len(prices)
	lowPrice := make([]int,length)
	lowPrice[0] = prices[0]
	maxProfit := 0
	for i := 1; i < length; i++ {
		if prices[i] < lowPrice[i-1] {
			lowPrice[i] = prices[i]
		}else {
			lowPrice[i] = lowPrice[i-1]
		}
		if (prices[i] - lowPrice[i]) > maxProfit {
			maxProfit = prices[i] - lowPrice[i]
		}
	}
	return maxProfit
}

func main() {
	prices := [][]int{
		{9,11,8,5,7,12,16,14},
		{1,2,3,4,5,6,7,8,9,10},
		{10,9,8,7,6,2,1,0},
	}
	for _, price := range prices {
		res := MaxProfit(price)
		fmt.Printf("max profit in %v is %d\n",price, res)
	}
}
/* output
max profit in [9 11 8 5 7 12 16 14] is 11
max profit in [1 2 3 4 5 6 7 8 9 10] is 9
max profit in [10 9 8 7 6 2 1 0] is 0
 */