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

func memoryPancake(arr []int) int {
	a := sizeOf[[]int]()

	vars := 2 * sizeOf[int]()

	getMaxFunc := sizeOf[[]int]() + 3*sizeOf[int]()
	flipFunc := sizeOf[[]int]() + 2*sizeOf[int]()

	return a + vars + getMaxFunc + flipFunc
}

func memoryQuick(arr []int) int {
	n := sizeOf[int]()

	a := sizeOf[[]int]()

	arrs := sizeOfArray[int](len(arr)) + 2*sizeOf[[]int]()

	pivot := sizeOf[int]()

	loop := sizeOf[int]()

	lens := 3 * sizeOf[int]()

	loops := 3 * sizeOf[int]()

	return n + a + arrs + pivot + loop + lens + loops
}

func memoryBead(arr []int) int {
	n := sizeOf[int]()

	a := sizeOf[[]int]()

	max := sizeOf[int]()

	m := sizeOf[[][]int]() + len(arr)*sizeOfArray[int](getIndMax(arr))

	loop1 := sizeOf[int]()
	loop2 := 2 * sizeOf[int]()
	loop3 := 4 * sizeOf[int]()
	loop4 := 3 * sizeOf[int]()

	return n + a + max + m + loop1 + loop2 + loop3 + loop4
}
