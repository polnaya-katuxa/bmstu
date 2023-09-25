package main

import (
	"aa/internal/algorithms"
	"fmt"
)

func main() {
	var s1, s2 string
	algorithms.Print = true
	fmt.Print("Enter 2 strings.\ns1: ")

	fmt.Scanln(&s1)

	fmt.Print("s2: ")
	fmt.Scanln(&s2)

	var num int
	algs := []func(string, string) int{
		algorithms.Levenshtein,
		algorithms.DamerauLevenshtein,
		algorithms.RecursiveDamerauLevenshtein,
		algorithms.RecursiveDamerauLevenshteinCached,
	}

	for num != 5 {
		fmt.Print("\nAlgorythmss of finding the distance:\n" +
			"1. Levenshtein algorythm\n" +
			"2. Damerau-Levenstein algorythm\n" +
			"3. Recursive Damerau-Levenshtein algorythm\n" +
			"4. Recursive Damerau-Levenshtein algorythm with cash\n" +
			"5. Exit\n" +
			"6. Reenter strings\n" +
			"\nEnter the chosen number: ")
		fmt.Scanln(&num)
		if num < 1 || num > 6 {
			fmt.Println("\nTry again.")
			continue
		}

		if num > 0 && num < 5 {
			fmt.Println("\nDistance: ", algs[num-1](s1, s2))
		} else if num == 5 {
			break
		} else {
			fmt.Print("Enter 2 strings.\ns1: ")

			fmt.Scanln(&s1)

			fmt.Print("s2: ")
			fmt.Scanln(&s2)
		}
	}
}
