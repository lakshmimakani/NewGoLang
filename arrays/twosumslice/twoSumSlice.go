package main

import "fmt"

func findIndexesForSum(slice []int) (int, int) {
	indices := make(map[int]int)

	for i, num := range slice {
		complement := 10 - num
		if index, ok := indices[complement]; ok {
			return index, i
		}
		indices[num] = i
	}

	return -1, -1
}

func main() {
	// Sample slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Find indexes for sum 10
	index1, index2 := findIndexesForSum(slice)

	if index1 == -1 && index2 == -1 {
		fmt.Println("The two numbers don't exist in the slice.")
	} else {
		fmt.Printf("Indexes of two numbers that sum up to 10: %d, %d\n", index1, index2)
	}
}
