package main

import "fmt"

func rotateLeft(slice []int, n int) []int {
	length := len(slice)
	if length == 0 {
		return slice
	}
	// Calculate the effective number of rotations
	effectiveRotations := n % length
	// No rotation needed
	if effectiveRotations == 0 {
		return slice
	}
	// Rotate the slice using slicing
	return append(slice[effectiveRotations:], slice[:effectiveRotations]...)
}

func main() {
	// Sample slice
	slice := []int{1, 2, 3, 4, 5}

	// Rotate the slice to the left by 3 positions
	rotatedSlice := rotateLeft(slice, 3)

	// Print the rotated slice
	fmt.Println("Rotated slice:", rotatedSlice)
}
