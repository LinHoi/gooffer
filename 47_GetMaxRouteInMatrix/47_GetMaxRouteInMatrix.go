package main

import "fmt"

func GetMaxRouteInMatrix(matrix [][]int, maxRow,MaxCol,row,col int) int {
	if row == maxRow {
		sum := 0
		for c := col; c <= MaxCol; c++ {
			sum += matrix[row][c]
		}
		return sum
	}
	if  col == MaxCol {
		sum := 0
		for r := row; r <= maxRow; r++ {
			sum += matrix[r][col]
		}
		return sum
	}
	right := GetMaxRouteInMatrix(matrix, maxRow,MaxCol, row,col+1)
	down  := GetMaxRouteInMatrix(matrix,maxRow,MaxCol,row+1, col)
	if right > down {
		return matrix[row][col] + right
	} else {
		return matrix[row][col] + down
	}
}

func main() {

	matrixes := [][][]int{
		{
			{1, 24, 5, 65, 6, 87, 99, 678},
			{234, 656, 677, 57, 77, 455, 566, 33},
			{12, 34, 4, -3, 34, -234, 34, 324},
			{-34, 435, 3452, 3, 45, 6, 34, 465},
		},

		{
			{0},
		},

		{
			{1,0,2,-4,3,2,5},
		},

		{
			{1,4,4,-5},
			{3,4,-5,3},
		},
	}
	for _, matrix := range matrixes {
		maxRow := len(matrix) - 1
		maxCol := len(matrix[0]) - 1
		res := GetMaxRouteInMatrix(matrix, maxRow, maxCol, 0, 0)
		fmt.Printf("Max Sum of Route in matrix is %d \n", res)
	}
}
/*
Max Sum of Route in matrix is 5577
Max Sum of Route in matrix is 0
Max Sum of Route in matrix is 9
Max Sum of Route in matrix is 7
 */