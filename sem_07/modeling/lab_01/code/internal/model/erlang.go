package model

import (
	"fmt"
	"math"
)

const eps = 1e-4

type ErlangDistribution struct {
	k      int
	lambda float64
}

func NewErlang(k int, lambda float64) (*ErlangDistribution, error) {
	if k < 1 {
		return nil, fmt.Errorf("incorrect params")
	}

	if lambda < eps {
		return nil, fmt.Errorf("incorrect params")
	}

	return &ErlangDistribution{
		k:      k,
		lambda: lambda,
	}, nil
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func (e *ErlangDistribution) Density(x float64) float64 {
	if x < 0 {
		return 0
	}
	res := math.Pow(x, float64(e.k-1)) * math.Exp(-x/e.lambda) / math.Pow(e.lambda, float64(e.k))

	return res / float64(factorial(e.k-1))
}

func (e *ErlangDistribution) Distribution(x float64) float64 {
	if x < 0 {
		return 0
	}

	sum := 0.0
	for i := 0; i < e.k; i++ {
		a1 := math.Exp(-e.lambda * x)
		a2 := math.Pow(e.lambda*x, float64(i))
		sum += (a1 * a2) / float64(factorial(i))
	}

	return 1 - sum
}

func (e *ErlangDistribution) RightBound() float64 {
	sb := 0.0
	for e.Distribution(sb) < 0.999 {
		sb++
	}
	return sb + 1
}

func (e *ErlangDistribution) LeftBound() float64 {
	return 0.0
}

func (e *ErlangDistribution) ComputeDistribution(step, left, right float64) Result {
	x := make([]float64, 0)
	y := make([]float64, 0)

	for i := left; i < right; i += step {
		x = append(x, i)
		y = append(y, e.Distribution(i))
	}

	return Result{
		Name: fmt.Sprintf("k = %d, lambda = %.2f", e.k, e.lambda),
		X:    x,
		Y:    y,
	}
}

func (e *ErlangDistribution) ComputeDensity(step, left, right float64) Result {
	x := make([]float64, 0)
	y := make([]float64, 0)

	for i := left; i < right; i += step {
		x = append(x, i)
		y = append(y, e.Density(i))
	}

	return Result{
		Name: fmt.Sprintf("k = %d, lambda = %.2f", e.k, e.lambda),
		X:    x,
		Y:    y,
	}
}
