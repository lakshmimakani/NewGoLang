package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 4, "c": 5, "d": 6}

	mergedMap := mergeMaps(map1, map2)
	fmt.Println("Merged Map:", mergedMap)

	highestKey, highestValue := findHighestValue(mergedMap)
	fmt.Printf("Key with the highest value: %s (value: %d)\n", highestKey, highestValue)
}

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] += value
	}

	for key, value := range map2 {
		mergedMap[key] += value
	}

	return mergedMap
}

func findHighestValue(inputMap map[string]int) (string, int) {
	var highestKey string
	var highestValue int

	for key, value := range inputMap {
		if value > highestValue {
			highestKey = key
			highestValue = value
		}
	}

	return highestKey, highestValue
}
