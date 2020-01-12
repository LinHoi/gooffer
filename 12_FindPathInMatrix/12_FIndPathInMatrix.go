package main

import "fmt"

func hasPath(matrix [][]rune, targetString string) bool {
	if matrix == nil || targetString == "" {
		return false
	}
	pathLength := 0

	//init visited Matrix
	rows , cols := len(matrix) , len(matrix[0])
	visited := make([][]bool,rows)
	for row := 0; row <  rows; row++ {
		visited[row] = make([]bool,cols)
	}

	for rowNum := range matrix {
		for colNum := range matrix[rowNum] {
			if hasPathCore(matrix, targetString,rowNum, colNum, pathLength, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix [][]rune,targetString string, row int, col int, pathLength int, visited [][]bool) bool {
	if pathLength == len(targetString) {
		return true
	}
	hasPath := false
	rows := len(matrix)
	cols := len(matrix[0])
	if  row >= 0 &&
		row < rows &&
		col >= 0 &&
		col < cols &&
		matrix[row][col] == rune(targetString[pathLength]) &&
		visited[row][col] == false {
		pathLength ++
		visited[row][col] = true
		hasPath = hasPathCore(matrix,targetString,row,col-1,pathLength,visited)||
			      hasPathCore(matrix,targetString,row,col+1,pathLength,visited)||
			      hasPathCore(matrix,targetString,row-1,col,pathLength,visited)||
			      hasPathCore(matrix,targetString,row+1,col,pathLength,visited)
		if !hasPath {
			pathLength --
			visited[row][col] = false
		}

	}
	return hasPath
}

func main() {
	matrix := [][]rune{
		{'a','b','t','g'},
		{'c','f','c','s'},
		{'j','d','e','h'},
	}

	targetStrings := []string{
		"bfceh",
		"abtg",
		"abtgscfcjdeh",
		"sfetgdfgv",
		"qwrretfsdgddc",
		"acjdehsgtbfc",
	}
	fmt.Println("Matrix M: ")
	for i:= range matrix {
		for j := range  matrix[i] {
			fmt.Printf("%v ", string(matrix[i][j]))
		}
		fmt.Printf("\n")
	}
	for _,targetString := range targetStrings {
		result := hasPath(matrix, targetString)
		if result == true {
			fmt.Printf("Matrix M has  path : %s \n", targetString)
		} else {
			fmt.Printf("Matrix M does't have  path : %s \n", targetString)
		}
	}
}