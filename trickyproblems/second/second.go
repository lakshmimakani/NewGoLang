package main

import "fmt"

func main() {
	var i, j, rows, columns int

	var dSumMat [10][10]int

	fmt.Print("Enter the Matrix rows and Columns = ")
	fmt.Scan(&rows, &columns)

	if rows != columns {
		fmt.Println("Error: Matrix is not square")
		return
	}

	fmt.Println("Enter the Matrix Items to find the Diagonal Sum = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&dSumMat[i][j])
		}
	}
	sumMain := 0
	sumSecondary := 0
	for i = 0; i < rows; i++ {
		sumMain += dSumMat[i][i]
		sumSecondary += dSumMat[i][columns-i-1]
	}
	fmt.Println("The Sum of Matrix Main Diagonal Elements  = ", sumMain)
	fmt.Println("The Sum of Matrix Secondary Diagonal Elements  = ", sumSecondary)
}
