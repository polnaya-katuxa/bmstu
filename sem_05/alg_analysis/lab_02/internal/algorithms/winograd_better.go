package algorithms

import (
	"errors"
	"lab_02/internal/matrix"
	"lab_02/internal/memory"
)

func betterRowCoefs(m matrix.Matrix) []int {
	coefs := make([]int, m.M)

	halfN := m.N >> 1

	for i := 0; i < m.M; i++ {
		for j := 0; j < halfN; j++ {
			coefs[i] += m.Data[i][j<<1] * m.Data[i][(j<<1)+1]
		}
	}

	return coefs
}

func betterColumnCoefs(m matrix.Matrix) []int {
	coefs := make([]int, m.N)

	halfM := m.M >> 1

	for j := 0; j < m.N; j++ {
		for i := 0; i < halfM; i++ {
			coefs[j] += m.Data[i<<1][j] * m.Data[(i<<1)+1][j]
		}
	}

	return coefs
}

func WinogradBetterMulMatrix(m1 matrix.Matrix, m2 matrix.Matrix) (matrix.Matrix, error) {
	memory.MemoryInfo.Reset()
	memory.MemoryInfo.Add(memory.MemoryWinogradImproved(m1.M, m2.N))
	defer memory.MemoryInfo.Done(memory.MemoryWinogradImproved(m1.M, m2.N))

	if m1.N != m2.M {
		return matrix.Matrix{}, errors.New("mul sizes error")
	}

	res := matrix.CreateEmpty(m1.M, m2.N)

	rowCoefs := betterRowCoefs(m1)
	columnCoefs := betterColumnCoefs(m2)

	isOdd := m1.N%2 != 0
	halfN := m1.N >> 1
	preN := m1.N - 1

	for i := 0; i < res.M; i++ {
		for j := 0; j < res.N; j++ {
			res.Data[i][j] -= rowCoefs[i] + columnCoefs[j]

			for k := 0; k < halfN; k++ {
				res.Data[i][j] += (m1.Data[i][k<<1] + m2.Data[(k<<1)+1][j]) *
					(m1.Data[i][(k<<1)+1] + m2.Data[k<<1][j])
			}

			if isOdd {
				res.Data[i][j] += m1.Data[i][preN] * m2.Data[preN][j]
			}
		}
	}

	return res, nil
}
