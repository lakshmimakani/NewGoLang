package main

import "fmt"

func findDuplicates(slice []int) []int {
	duplicates := []int{}
	encountered := map[int]bool{}

	for _, v := range slice {
		if encountered[v] == true {
			duplicates = append(duplicates, v)
		} else {
			encountered[v] = true
		}
	}

	return duplicates
}

func main() {
	slice := []int{1, 2, 3, 4, 2, 5, 6, 3, 7, 8, 9, 1, 5}
	duplicates := findDuplicates(slice)
	fmt.Println("Duplicates:", duplicates)
}
