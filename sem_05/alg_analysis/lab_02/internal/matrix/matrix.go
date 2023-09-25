package matrix

import (
	"errors"
	"fmt"
	"lab_02/internal/memory"
	"math/rand"
)

const (
	maxNum   = 1000
	interval = 500
)

type Matrix struct {
	M    int
	N    int
	Data [][]int
}

func CreateEmpty(m, n int) Matrix {
	matr := make([][]int, m)

	for i := 0; i < m; i++ {
		matr[i] = make([]int, n)
	}

	return Matrix{
		m,
		n,
		matr,
	}
}

func FillRandom(m Matrix) {
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			m.Data[i][j] = rand.Intn(maxNum+1) - interval
		}
	}
}

func AreEqual(m1, m2 Matrix) bool {
	if m1.M != m2.M || m1.N != m2.N {
		return false
	}

	for i := 0; i < m1.M; i++ {
		for j := 0; j < m1.N; j++ {
			if m1.Data[i][j] != m2.Data[i][j] {
				return false
			}
		}
	}

	return true
}

func Copy(m Matrix) Matrix {
	copyM := CreateEmpty(m.M, m.N)

	for i := 0; i < m.M; i++ {
		copy(copyM.Data[i], m.Data[i])
	}

	return copyM
}

func inputArr(arr []int) error {
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func Input(m, n int) (Matrix, error) {
	res := CreateEmpty(m, n)

	for i := 0; i < m; i++ {
		if err := inputArr(res.Data[i]); err != nil {
			return Matrix{}, err
		}
	}

	return res, nil
}

func Print(m Matrix) {
	fmt.Println()
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			fmt.Printf("%3d ", m.Data[i][j])
		}
		fmt.Println()
	}
}

func Mul(m1, m2 Matrix) (Matrix, error) {
	memory.MemoryInfo.Reset()
	memory.MemoryInfo.Add(memory.MemoryUsual(m1.M, m2.N))
	defer memory.MemoryInfo.Done(memory.MemoryUsual(m1.M, m2.N))

	if m1.N != m2.M {
		return Matrix{}, errors.New("mul sizes error")
	}

	res := CreateEmpty(m1.M, m2.N)

	for i := 0; i < res.M; i++ {
		for j := 0; j < res.N; j++ {
			for k := 0; k < m1.N; k++ {
				res.Data[i][j] = res.Data[i][j] + m1.Data[i][k]*m2.Data[k][j]
			}
		}
	}

	return res, nil
}
