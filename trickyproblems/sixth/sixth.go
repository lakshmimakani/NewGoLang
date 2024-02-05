package main

import (
	"fmt"
	"sort"
)

// Pair represents an element and its frequency
type Pair struct {
	val       int
	frequency int
}

// ByFrequency implements sort.Interface for []Pair based on the frequency field.
type ByFrequency []Pair

func (a ByFrequency) Len() int { return len(a) }
func (a ByFrequency) Less(i, j int) bool { // Define custom sorting
	if a[i].frequency == a[j].frequency {
		return a[i].val < a[j].val // If frequencies are the same, sort by natural order
	}
	return a[i].frequency > a[j].frequency // Sort by frequency in descending order
}
func (a ByFrequency) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func sortByFrequency(arr []int) []int {
	// Step 1: Count the frequency of each element
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}

	// Step 2: Create an array of pairs (element, frequency)
	var pairs []Pair
	for key, value := range freqMap {
		pairs = append(pairs, Pair{key, value})
	}

	// Step 3: Sort the pairs based on frequency and natural order
	sort.Sort(ByFrequency(pairs))

	// Step 4: Create the sorted array based on the sorted pairs
	sortedArr := make([]int, 0, len(arr))
	for _, pair := range pairs {
		for i := 0; i < pair.frequency; i++ {
			sortedArr = append(sortedArr, pair.val)
		}
	}

	return sortedArr
}

func main() {
	arr := []int{2, 3, 2, 4, 5, 12, 2, 3, 3, 3, 12}
	sortedArr := sortByFrequency(arr)
	fmt.Println("Sorted array by frequency:", sortedArr)
}
