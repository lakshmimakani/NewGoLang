package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Use two pointers: one for iterating through the array, and one for keeping track of the unique elements.
	// The unique elements will be placed at the beginning of the array.
	uniqueIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	// The length of the array without duplicates will be (uniqueIndex + 1).
	return uniqueIndex + 1
}

func main() {
	// Example usage:
	arr := []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 6}
	length := removeDuplicates(arr)
	fmt.Println("Array after removing duplicates:", arr[:length]) // Print the array without duplicates
}
