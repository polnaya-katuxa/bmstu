package algorithms

import (
	"fmt"
	"math"
)

var Print = false

func PrintMatrix(m [][]int) {
	fmt.Println()
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%3d ", m[i][j])
		}
		fmt.Println()
	}
}

func min(operands ...int) int {
	minOperand := operands[0]
	for _, v := range operands {
		if v < minOperand {
			minOperand = v
		}
	}

	return minOperand
}

func Levenshtein(s1, s2 string) int {
	MemoryInfo.Reset()
	MemoryInfo.Add(memoryLevenshtein(s1, s2))
	defer MemoryInfo.Done(memoryLevenshtein(s1, s2))

	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	m := make([][]int, len(s1Rune)+1)
	for i := range m {
		m[i] = make([]int, len(s2Rune)+1)
		m[i][0] = i
	}
	for j := range m[0] {
		m[0][j] = j
	}

	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			insertOpt := m[i][j-1] + 1
			deleteOpt := m[i-1][j] + 1
			replaceOpt := m[i-1][j-1]

			if s1Rune[i-1] != s2Rune[j-1] {
				replaceOpt += 1
			}

			m[i][j] = min(insertOpt, deleteOpt, replaceOpt)
		}
	}

	if Print {
		PrintMatrix(m)
	}

	return m[len(m)-1][len(m[0])-1]
}

func DamerauLevenshtein(s1, s2 string) int {
	MemoryInfo.Reset()
	MemoryInfo.Add(memoryDamerauLevenshtein(s1, s2))
	defer MemoryInfo.Done(memoryDamerauLevenshtein(s1, s2))

	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	m := make([][]int, len(s1Rune)+1)
	for i := range m {
		m[i] = make([]int, len(s2Rune)+1)
		m[i][0] = i
	}
	for j := range m[0] {
		m[0][j] = j
	}

	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			insertOpt := m[i][j-1] + 1
			deleteOpt := m[i-1][j] + 1
			replaceOpt := m[i-1][j-1]
			substituteOpt := math.MaxInt

			if s1Rune[i-1] != s2Rune[j-1] {
				replaceOpt += 1
			}

			if i > 1 && j > 1 {
				if s1Rune[i-1] == s2Rune[j-2] && s1Rune[i-2] == s2Rune[j-1] {
					substituteOpt = m[i-2][j-2] + 1
				}
			}

			m[i][j] = min(insertOpt, deleteOpt, replaceOpt, substituteOpt)
		}
	}

	if Print {
		PrintMatrix(m)
	}

	return m[len(m)-1][len(m[0])-1]
}

func RecursiveDamerauLevenshtein(s1, s2 string) int {
	MemoryInfo.Reset()
	MemoryInfo.Add(memoryRecursiveDamerauLevenshtein(s1, s2))
	defer MemoryInfo.Done(memoryRecursiveDamerauLevenshtein(s1, s2))

	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	return recursiveDamerauLevenshtein(s1Rune, s2Rune)
}

func recursiveDamerauLevenshtein(s1, s2 []rune) int {
	MemoryInfo.Add(memoryRecursiveDamerauLevenshteinInner())
	defer MemoryInfo.Done(memoryRecursiveDamerauLevenshteinInner())

	l1 := len(s1)
	l2 := len(s2)

	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 > 0 && l2 == 0 {
		return l1
	} else if l1 == 0 && l2 > 0 {
		return l2
	} else {
		e := 0
		if s1[l1-1] != s2[l2-1] {
			e += 1
		}

		if l1 > 1 && l2 > 1 &&
			(s1[l1-1] == s2[l2-2] && s1[l1-2] == s2[l2-1]) {
			return min(recursiveDamerauLevenshtein(s1, s2[:l2-1])+1,
				recursiveDamerauLevenshtein(s1[:l1-1], s2)+1,
				recursiveDamerauLevenshtein(s1[:l1-1], s2[:l2-1])+e,
				recursiveDamerauLevenshtein(s1[:l1-2], s2[:l2-2])+1)
		} else {
			return min(recursiveDamerauLevenshtein(s1, s2[:l2-1])+1,
				recursiveDamerauLevenshtein(s1[:l1-1], s2)+1,
				recursiveDamerauLevenshtein(s1[:l1-1], s2[:l2-1])+e)
		}
	}
}

func RecursiveDamerauLevenshteinCached(s1, s2 string) int {
	MemoryInfo.Reset()
	MemoryInfo.Add(memoryRecursiveDamerauLevenshteinCached(s1, s2))
	defer MemoryInfo.Done(memoryRecursiveDamerauLevenshteinCached(s1, s2))

	s1Rune := []rune(s1)
	s2Rune := []rune(s2)

	cache := make([][]int, len(s1)+1)
	for i := range cache {
		cache[i] = make([]int, len(s2)+1)
		cache[i][0] = i
	}
	for j := range cache[0] {
		cache[0][j] = j
	}

	for i := 1; i < len(cache); i++ {
		for j := 1; j < len(cache[i]); j++ {
			cache[i][j] = math.MaxInt
		}
	}

	res := recursiveDamerauLevenshteinCached(s1Rune, s2Rune, cache)

	if Print {
		PrintMatrix(cache)
	}

	return res
}

func recursiveDamerauLevenshteinCached(s1, s2 []rune, cache [][]int) int {
	MemoryInfo.Add(memoryRecursiveDamerauLevenshteinInner() + sizeOf[[][]int]()) // тут вроде можно эту же, память одинаковая
	defer MemoryInfo.Done(memoryRecursiveDamerauLevenshteinInner() + sizeOf[[][]int]())

	l1 := len(s1)
	l2 := len(s2)

	if cache[l1][l2] == math.MaxInt {
		if l1 == 0 && l2 == 0 {
			cache[l1][l2] = 0
		} else if l1 > 0 && l2 == 0 {
			cache[l1][l2] = l1
		} else if l1 == 0 && l2 > 0 {
			cache[l1][l2] = l2
		} else {
			e := 0
			if s1[l1-1] != s2[l2-1] {
				e += 1
			}

			if l1 > 1 && l2 > 1 &&
				(s1[l1-1] == s2[l2-2] && s1[l1-2] == s2[l2-1]) {
				cache[l1][l2] = min(recursiveDamerauLevenshteinCached(s1, s2[:l2-1], cache)+1,
					recursiveDamerauLevenshteinCached(s1[:l1-1], s2, cache)+1,
					recursiveDamerauLevenshteinCached(s1[:l1-1], s2[:l2-1], cache)+e,
					recursiveDamerauLevenshteinCached(s1[:l1-2], s2[:l2-2], cache)+1)
			} else {
				cache[l1][l2] = min(recursiveDamerauLevenshteinCached(s1, s2[:l2-1], cache)+1,
					recursiveDamerauLevenshteinCached(s1[:l1-1], s2, cache)+1,
					recursiveDamerauLevenshteinCached(s1[:l1-1], s2[:l2-1], cache)+e)
			}
		}
	}
	return cache[l1][l2]
}
