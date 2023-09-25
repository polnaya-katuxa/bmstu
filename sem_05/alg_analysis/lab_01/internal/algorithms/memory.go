package algorithms

import (
	"unsafe"
)

var MemoryInfo Metrics

type Metrics struct {
	current int
	max     int
}

func (m *Metrics) Reset() {
	m.current = 0
	m.max = 0
}

func (m *Metrics) Add(v int) {
	m.current += v
	if m.current > m.max {
		m.max = m.current
	}
}

func (m *Metrics) Done(v int) {
	m.current -= v
}

func (m *Metrics) Max() int64 {
	return int64(m.max)
}

// sizeOf[int]()
func sizeOf[T any]() int {
	var v T
	return int(unsafe.Sizeof(v))
}

// sizeOfArray[int](10)
func sizeOfArray[T any](n int) int {
	return sizeOf[[]T]() + n*sizeOf[T]()
}

func memoryLevenshtein(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)

	s01 := sizeOf[string]() + n1*sizeOf[byte]()
	s02 := sizeOf[string]() + n2*sizeOf[byte]()

	n1r := len([]rune(s1))
	n2r := len([]rune(s2))

	sr01 := sizeOfArray[rune](n1r)
	sr02 := sizeOfArray[rune](n2r)

	res := sizeOf[int]()

	loop1 := 2 * sizeOf[int]()

	m := sizeOf[[][]int]() + sizeOfArray[int](n2r+1)*(n1r+1)

	loop2 := 5 * sizeOf[int]()

	minFunc := sizeOfArray[int](3) + sizeOf[int]()*3

	return s01 + s02 + sr01 + sr02 + res + m + loop1 + loop2 + minFunc
}

func memoryDamerauLevenshtein(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)

	s01 := sizeOf[string]() + n1*sizeOf[byte]()
	s02 := sizeOf[string]() + n2*sizeOf[byte]()

	n1r := len([]rune(s1))
	n2r := len([]rune(s2))

	sr01 := sizeOfArray[rune](n1r)
	sr02 := sizeOfArray[rune](n2r)

	res := sizeOf[int]()

	loop1 := 2 * sizeOf[int]()

	m := sizeOf[[][]int]() + sizeOfArray[int](n2r+1)*(n1r+1)

	loop2 := 6 * sizeOf[int]()

	minFunc := sizeOfArray[int](4) + sizeOf[int]()*3

	return s01 + s02 + sr01 + sr02 + res + m + loop1 + loop2 + minFunc
}

func memoryRecursiveDamerauLevenshtein(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)

	s01 := sizeOf[string]() + n1*sizeOf[byte]()
	s02 := sizeOf[string]() + n2*sizeOf[byte]()

	n1r := len([]rune(s1))
	n2r := len([]rune(s2))

	sr01 := sizeOfArray[rune](n1r)
	sr02 := sizeOfArray[rune](n2r)

	res := sizeOf[int]()

	return s01 + s02 + sr01 + sr02 + res
}

func memoryRecursiveDamerauLevenshteinInner() int {
	sr01 := sizeOf[[]rune]()
	sr02 := sizeOf[[]rune]()

	res := sizeOf[int]()

	vars := 3 * sizeOf[int]()

	minFunc := sizeOfArray[int](4) + sizeOf[int]()*3

	return sr01 + sr02 + res + vars + minFunc
}

func memoryRecursiveDamerauLevenshteinCached(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)

	s01 := sizeOf[string]() + n1*sizeOf[byte]()
	s02 := sizeOf[string]() + n2*sizeOf[byte]()

	n1r := len([]rune(s1))
	n2r := len([]rune(s2))

	sr01 := sizeOfArray[rune](n1r)
	sr02 := sizeOfArray[rune](n2r)

	res := sizeOf[int]()

	loop1 := 2 * sizeOf[int]()

	m := sizeOf[[][]int]() + sizeOfArray[int](n2r+1)*(n1r+1)

	loop2 := 2 * sizeOf[int]()

	return s01 + s02 + sr01 + sr02 + res + loop1 + loop2 + m
}
