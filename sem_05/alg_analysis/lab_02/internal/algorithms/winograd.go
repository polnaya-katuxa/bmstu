package algorithms

import (
	"errors"
	"lab_02/internal/matrix"
	"lab_02/internal/memory"
)

func rowCoefs(m matrix.Matrix) []int {
	coefs := make([]int, m.M)

	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N/2; j++ {
			coefs[i] = coefs[i] + m.Data[i][2*j]*m.Data[i][2*j+1]
		}
	}

	return coefs
}

func columnCoefs(m matrix.Matrix) []int {
	coefs := make([]int, m.N)

	for j := 0; j < m.N; j++ {
		for i := 0; i < m.M/2; i++ {
			coefs[j] = coefs[j] + m.Data[2*i][j]*m.Data[2*i+1][j]
		}
	}

	return coefs
}

func WinogradMulMatrix(m1 matrix.Matrix, m2 matrix.Matrix) (matrix.Matrix, error) {
	memory.MemoryInfo.Reset()
	memory.MemoryInfo.Add(memory.MemoryWinograd(m1.M, m2.N))
	defer memory.MemoryInfo.Done(memory.MemoryWinograd(m1.M, m2.N))

	if m1.N != m2.M {
		return matrix.Matrix{}, errors.New("mul sizes error")
	}

	res := matrix.CreateEmpty(m1.M, m2.N)

	rowCoefs := rowCoefs(m1)
	columnCoefs := columnCoefs(m2)

	for i := 0; i < res.M; i++ {
		for j := 0; j < res.N; j++ {
			res.Data[i][j] = res.Data[i][j] - rowCoefs[i] - columnCoefs[j]

			for k := 0; k < m1.N/2; k++ {
				res.Data[i][j] = res.Data[i][j] + (m1.Data[i][2*k]+m2.Data[2*k+1][j])*
					(m1.Data[i][2*k+1]+m2.Data[2*k][j])
			}
		}
	}

	if m1.N%2 != 0 {
		for i := 0; i < res.M; i++ {
			for j := 0; j < res.N; j++ {
				res.Data[i][j] = res.Data[i][j] + m1.Data[i][m1.N-1]*m2.Data[m1.N-1][j]
			}
		}
	}

	return res, nil
}
