package main

import (
	"fmt"
	"math"
)

const (
	maxValueInDice = 6
)
func ProbabilityOfSumInDices(diceNumber int) {
	minSum := 1* diceNumber
	offset := minSum
	maxSum := maxValueInDice * diceNumber
	probabilities := make([]int, maxSum-minSum+1)

	Probability(offset,diceNumber, probabilities)

	total := math.Pow(float64(maxValueInDice),float64(diceNumber))
	for i := range probabilities {
		fmt.Printf("%d:%f; ",i+offset, float64(probabilities[i])/total)
	}
	fmt.Println()
}

func Probability(offset, diceNumber int, probabilities []int) {
	for i := 1; i <= maxValueInDice; i++ {
		probability(offset,diceNumber,i, probabilities)
	}
}

func probability(offset int, diceNumber int, sum int, probabilities []int) {
	if diceNumber == 1 {
		probabilities[sum - offset] ++
	} else {
			for i := 1; i<= maxValueInDice; i++ {
				probability(offset,diceNumber-1, sum+i, probabilities)
			}
	}
}

func main() {
	diceNumbers := []int{1,2,3,4,5,6}
	for _, diceNumber := range diceNumbers {
		ProbabilityOfSumInDices(diceNumber)
	}
}
/*
1:0.166667; 2:0.166667; 3:0.166667; 4:0.166667; 5:0.166667; 6:0.166667;
2:0.027778; 3:0.055556; 4:0.083333; 5:0.111111; 6:0.138889; 7:0.166667; 8:0.138889; 9:0.111111; 10:0.083333; 11:0.055556; 12:0.027778;
3:0.004630; 4:0.013889; 5:0.027778; 6:0.046296; 7:0.069444; 8:0.097222; 9:0.115741; 10:0.125000; 11:0.125000; 12:0.115741; 13:0.097222; 14:0.069444; 15:0.046296; 16:0.027778; 17:0.013889; 18:0.004630;
4:0.000772; 5:0.003086; 6:0.007716; 7:0.015432; 8:0.027006; 9:0.043210; 10:0.061728; 11:0.080247; 12:0.096451; 13:0.108025; 14:0.112654; 15:0.108025; 16:0.096451; 17:0.080247; 18:0.061728; 19:0.043210; 20:0.027006; 21:0.015432; 22:0.007716; 23:0.003086; 24:0.000772;
5:0.000129; 6:0.000643; 7:0.001929; 8:0.004501; 9:0.009002; 10:0.016204; 11:0.026363; 12:0.039223; 13:0.054012; 14:0.069444; 15:0.083719; 16:0.094522; 17:0.100309; 18:0.100309; 19:0.094522; 20:0.083719; 21:0.069444; 22:0.054012; 23:0.039223; 24:0.026363; 25:0.016204; 26:0.009002; 27:0.004501; 28:0.001929; 29:0.000643; 30:0.000129;
6:0.000021; 7:0.000129; 8:0.000450; 9:0.001200; 10:0.002701; 11:0.005401; 12:0.009774; 13:0.016204; 14:0.024884; 15:0.035708; 16:0.048161; 17:0.061214; 18:0.073538; 19:0.083719;
 */