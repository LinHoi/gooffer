package main

import "fmt"

func hasPath(matrix [][]rune, visited [][]bool, targetString string) bool {
	if matrix == nil || targetString == "" {
		return false
	}
	pathLength := 0
	//row , col := len(matrix) , len(matrix[0])
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
	if pathLength == len(targetString){
		return true
	}
	hasPath := false
	rows := len(matrix)
	cols := len(matrix[0])
	if  row >= 0 &&
		row <= rows &&
		col >= 0 &&
		col <= cols &&
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
	visited := [][]bool{
		{false,false,false,false},
		{false,false,false,false},
		{false,false,false,false},
	}
	fmt.Println(len(matrix), len(visited[0]))
	targetStrings := []string{
		"bfce",
		"abfc",
		"abtgscfcjdeh",
		//"sfetgdfgv",
	}
	fmt.Println("Matrix M: ")
	for i:= range matrix {
		for j := range  matrix[i] {
			fmt.Printf("%v ", string(matrix[i][j]))
		}
		fmt.Printf("\n")
	}
	for _,targetString := range targetStrings {
		result := hasPath(matrix, visited, targetString)
		if result == true {
			fmt.Printf("Matrix M has a path of %s \n", targetString)
		} else {
			fmt.Printf("Matrix M do not has a path of %s \n", targetString)
		}
	}


}