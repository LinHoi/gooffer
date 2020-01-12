package main

import "fmt"

func MovingAreaOfRobot(rows, cols, threshold int) (gridNum int){
	//if threshold >= getDigitSum(rows-1) + getDigitSum( cols -1) {return rows * cols}
	visitedMatrix := make([][]bool,rows)
	for i := range visitedMatrix{
		visitedMatrix[i] = make([]bool, cols)
	}
	gridNum = movingAreaOfRobot(0,0,rows,cols,threshold,visitedMatrix)
	return gridNum
}

func movingAreaOfRobot(row,col, rows, cols, threshold int , visitedMatrix [][]bool ) (gridNum int) {
	gridNum = 0
	if checkIn(row,col,rows,cols,threshold,visitedMatrix) {
		visitedMatrix[row][col] = true
		gridNum = 1 + movingAreaOfRobot(row-1,col,rows,cols,threshold,visitedMatrix) +
			movingAreaOfRobot(row+1,col,rows,cols,threshold,visitedMatrix) +
			movingAreaOfRobot(row,col-1,rows,cols,threshold,visitedMatrix) +
			movingAreaOfRobot(row,col+1,rows,cols,threshold,visitedMatrix)
	}
	return gridNum
}
func checkIn(row,col, rows, cols, threshold int , visitedMatrix [][]bool) bool {
	if  row >= 0 && row < rows &&
		col >= 0 && col < cols &&
		visitedMatrix[col][row] == false &&
		getDigitSum(col) + getDigitSum(row) <= threshold {

		return true
	}
	return false
}
func getDigitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num/10
	}
	return sum
}

func main() {
	res := MovingAreaOfRobot(3,3,2)
	fmt.Println(res)
	//TODO : stack overflow
}
