package model

import (
	"fmt"
	"math"
)

type HyperExponentialDistribution struct {
	lambdas       []float64
	probabilities []float64
}

func NewHyperExponential(lambdas, ps []float64) (*HyperExponentialDistribution, error) {
	return &HyperExponentialDistribution{
		lambdas:       lambdas,
		probabilities: ps,
	}, nil
}

func (h *HyperExponentialDistribution) Density(x float64) float64 {
	if x < 0 {
		return 0
	}

	res := 0.0
	for i := range h.lambdas {
		res += h.probabilities[i] * h.lambdas[i] * math.Exp(-h.lambdas[i]*x)
	}

	return res
}

func (h *HyperExponentialDistribution) Distribution(x float64) float64 {
	if x < 0 {
		return 0
	}

	sum := 0.0
	for i := range h.lambdas {
		sum += h.probabilities[i] * math.Exp(-h.lambdas[i]*x)
	}

	return 1 - sum
}

func (h *HyperExponentialDistribution) RightBound() float64 {
	sb := 0.0
	for h.Distribution(sb) < 0.999 {
		sb++
	}
	return sb + 1
}

func (h *HyperExponentialDistribution) LeftBound() float64 {
	return 0.0
}

func (h *HyperExponentialDistribution) ComputeDistribution(step, left, right float64) Result {
	x := make([]float64, 0)
	y := make([]float64, 0)

	for i := left; i < right; i += step {
		x = append(x, i)
		y = append(y, h.Distribution(i))
	}

	return Result{
		Name: fmt.Sprintf("lambdas = %v, probabilities = %v", h.lambdas, h.probabilities),
		X:    x,
		Y:    y,
	}
}

func (h *HyperExponentialDistribution) ComputeDensity(step, left, right float64) Result {
	x := make([]float64, 0)
	y := make([]float64, 0)

	for i := left; i < right; i += step {
		x = append(x, i)
		y = append(y, h.Density(i))
	}

	return Result{
		Name: fmt.Sprintf("lambdas = %v, probabilities = %v", h.lambdas, h.probabilities),
		X:    x,
		Y:    y,
	}
}
