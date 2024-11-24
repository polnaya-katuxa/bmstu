package model

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Model struct {
	lambdas *mat.Dense
	states  int
}

func New(elements []float64, states int) *Model {
	return &Model{lambdas: mat.NewDense(states, states, elements), states: states}
}

func (m *Model) Probabilities() ([]float64, error) {
	a := mat.NewDense(m.states, m.states, nil)
	a.Copy(m.lambdas.T())

	// Идем по диагонали, вычитаем все значения из данного столбца (все строки).
	// То есть в уравнении получаем P_i' = ... - P_i * (L_i0 + L_i1 + ...).
	// Для каждого P_i вычитаем сумму L_i0 + L_i1 + ...
	for j := 0; j < m.states; j++ {
		sum := 0.0
		for i := 0; i < m.states; i++ {
			sum += a.At(i, j)
		}

		a.Set(j, j, a.At(j, j)-sum)
	}

	// Первое уравнение заменим на P_0 + ... = 1.
	for j := 0; j < m.states; j++ {
		a.Set(0, j, 1)
	}

	// Решаем систему уравнений. Нужна матрица коэффициентов A, и ветор b.
	// Ax = b.
	b := mat.NewDense(m.states, 1, nil)

	// Первое уравнение заменим на P_0 + ... = 1.
	b.Set(0, 0, 1)

	x := mat.NewDense(m.states, 1, nil)

	// Решение линейного уравнения.
	var qr mat.QR
	qr.Factorize(a)

	err := qr.SolveTo(x, false, b)
	if err != nil {
		return nil, fmt.Errorf("solve linear equation: %w", err)
	}

	return mat.Col(nil, 0, x), nil
}

func (m *Model) Times() ([]float64, error) {
	sums := make([]float64, m.states)
	for j := 0; j < m.states; j++ {
		for i := 0; i < m.states; i++ {
			sums[j] += m.lambdas.At(i, j)
		}
		sums[j] -= m.lambdas.At(j, j)
	}

	probabilities, err := m.Probabilities()
	if err != nil {
		return nil, fmt.Errorf("probabilities: %w", err)
	}

	for i := range probabilities {
		probabilities[i] /= sums[i]
	}

	return probabilities, nil
}
