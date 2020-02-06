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