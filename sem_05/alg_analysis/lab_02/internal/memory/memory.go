package memory

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

// получение полного размера матрицы (заголовок +
//заголовки массивов + элементы)
// sizeOfMatrix[int](10, 10)
func sizeOfMatrix[T any](m, n int) int {
	return sizeOf[[][]T]() + m*sizeOf[[]T]() + m*n*sizeOf[T]()
}

func MemoryUsual(m, n int) int {
	args := 3 * sizeOf[[][]int]()

	res := sizeOf[[][]int]()

	loop := 3 * sizeOf[int]()

	create := 3*sizeOf[int]() + sizeOfMatrix[int](m, n) + sizeOf[[][]int]()

	return args + res + loop + create
}

func MemoryWinograd(m, n int) int {
	args := 3 * sizeOf[[][]int]()

	res := sizeOf[[][]int]()

	create := 3*sizeOf[int]() + sizeOfMatrix[int](m, n) + sizeOf[[][]int]()

	coefs := 2 * sizeOf[[]int]()

	coefR := sizeOf[[]int]() + sizeOf[[][]int]() + sizeOfArray[int](m) + 2*sizeOf[int]()
	coefC := sizeOf[[]int]() + sizeOf[[][]int]() + sizeOfArray[int](n) + 2*sizeOf[int]()

	loop := 5 * sizeOf[int]()

	return args + res + create + coefs + coefR + coefC + loop
}

func MemoryWinogradImproved(m, n int) int {
	args := 3 * sizeOf[[][]int]()

	res := sizeOf[[][]int]()

	create := 3*sizeOf[int]() + sizeOfMatrix[int](m, n) + sizeOf[[][]int]()

	coefs := 2 * sizeOf[[]int]()

	coefR := sizeOf[[]int]() + sizeOf[[][]int]() + sizeOfArray[int](m) + 2*sizeOf[int]()
	coefC := sizeOf[[]int]() + sizeOf[[][]int]() + sizeOfArray[int](n) + 2*sizeOf[int]()

	extra := 3 * sizeOf[int]()

	loop := 3 * sizeOf[int]()

	return args + res + create + coefs + coefR + coefC + loop + extra
}
