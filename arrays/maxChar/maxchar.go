package main

import (
	"fmt"
)

func findMaxRepeatingChar(input string) int32 {
	charCount := make(map[int32]int)

	// Count occurrences of each character
	for _, char := range input {
		charCount[char]++
	}

	// Find the character with the maximum count
	maxCount := 0
	maxChar := int32(0)

	for char, count := range charCount {
		if count > maxCount {
			maxCount = count
			maxChar = char
		}
	}

	return maxChar
}

func main() {
	inputString := "programminggolanggg"

	maxRepeatingChar := findMaxRepeatingChar(inputString)

	if maxRepeatingChar != 0 {
		fmt.Printf("The character '%c' repeats the maximum number of times.\n", maxRepeatingChar)
	} else {
		fmt.Println("No characters in the string.")
	}
}
