package main

import (
	"errors"
	"fmt"
	"lab_02/internal/algorithms"
	"lab_02/internal/matrix"
	"log"
)

//заменить операцию x = x + k на x += k;
//заменить умножение на 2 на побитовый сдвиг;
//предвычислять некоторые слагаемые для алгоритма.

func input() (matrix.Matrix, error) {
	var r, c int

	fmt.Println("\nВведите количество строк в матрице: ")
	if _, err := fmt.Scanf("%d", &r); err != nil || r <= 0 {
		return matrix.Matrix{}, errors.New("input size error")
	}
	fmt.Println("\nВведите количество столбцов в матрице: ")
	if _, err := fmt.Scanf("%d", &c); err != nil || c <= 0 {
		return matrix.Matrix{}, errors.New("input size error")
	}

	fmt.Print("\nInput matrix.\n" +
		"Enter several integers > 0, separated by space, for a row.\n" +
		"Press enter, when row ends.\n")

	m, err := matrix.Input(r, c)
	if err != nil {
		return matrix.Matrix{}, errors.New("input matrix error")
	}

	return m, nil
}

func inputTwoMatrix() (matrix.Matrix, matrix.Matrix, error) {
	m1, err := input()
	if err != nil {
		return matrix.Matrix{}, matrix.Matrix{}, errors.New("input error")
	}

	m2, err := input()
	if err != nil {
		return matrix.Matrix{}, matrix.Matrix{}, errors.New("input error")
	}

	if m1.N != m2.M {
		return matrix.Matrix{}, matrix.Matrix{}, errors.New("incorrect matrix for mul")
	}

	return m1, m2, nil
}

func main() {
	m1, m2, err := inputTwoMatrix()
	if err != nil {
		log.Fatalln(err)
	}

	var num int
	algs := []func(matrix.Matrix, matrix.Matrix) (matrix.Matrix, error){
		matrix.Mul,
		algorithms.WinogradMulMatrix,
		algorithms.WinogradBetterMulMatrix,
	}

	for num != 4 {
		fmt.Print("\nMenu:\n" +
			"1. Usual multiplication\n" +
			"2. Winograd\n" +
			"3. Improved Winograd\n" +
			"4. Exit\n" +
			"5. Reenter matrix\n" +
			"\nEnter the chosen number: ")
		fmt.Scan(&num)
		if num < 1 || num > 5 {
			fmt.Println("\nTry again.")
			continue
		}

		if num > 0 && num < 4 {
			fmt.Println("\nMatrix A:")
			matrix.Print(m1)
			fmt.Println("\nMatrix B:")
			matrix.Print(m2)

			res, err := algs[num-1](m1, m2)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("\nMatrix C = AB:")
			matrix.Print(res)
		} else if num == 4 {
			break
		} else {
			copyM1 := matrix.Copy(m1)
			copyM2 := matrix.Copy(m2)

			m1, m2, err = inputTwoMatrix()
			if err != nil {
				fmt.Println("\nTry again.")
				m1 = matrix.Copy(copyM1)
				m2 = matrix.Copy(copyM2)
				continue
			}
		}
	}
}
