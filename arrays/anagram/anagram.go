package main

import "fmt"

func Anagram(s string, t string) bool {

	string1 := len(s)
	string2 := len(t)
	fmt.Println("Letter 1 =", s, "\nLetter 2 =", t)
	fmt.Println("Is it an Anagram ?")
	if string1 != string2 {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < string1; i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < string2; i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < string1; i++ {
		if anagramMap[string(s[i])] != 0 { // if this condition satisfies return false
			return false
		}
	}
	// In the end return true
	return true
}

func main() {
	fmt.Println("Golang program to to check if two strings are anagram")

	output := Anagram("listen", "silent")
	fmt.Println(output)

	output = Anagram("man", "name")
	fmt.Println(output)

	// print the result using the function fmt.Println ()
}
