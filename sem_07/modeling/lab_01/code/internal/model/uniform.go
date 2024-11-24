package model

import (
	"fmt"
)

type UniformDistribution struct {
	a float64
	b float64
}

func NewUniform(a, b float64) (*UniformDistribution, error) {
	if a > b {
		return nil, fmt.Errorf("incorrect params")
	}
	return &UniformDistribution{
		a: a,
		b: b,
	}, nil
}

func (u *UniformDistribution) Distribution(x float64) float64 {
	if u.a <= x && u.b > x {
		return (x - u.a) / (u.b - u.a)
	}
	if x < u.a {
		return 0.0
	}
	return 1.0
}

func (u *UniformDistribution) RightBound() float64 {
	return u.b + ((u.b - u.a) / 2.0) + 1
}

func (u *UniformDistribution) LeftBound() float64 {
	return u.a - ((u.b - u.a) / 2.0) - 1
}

func (u *UniformDistribution) Density(x float64) float64 {
	if u.a <= x && x <= u.b {
		return 1.0 / (u.b - u.a)
	}
	return 0.0
}

func (u *UniformDistribution) ComputeDistribution(step, left, right float64) Result {
	x := make([]float64, 0)
	y := make([]float64, 0)

	for i := left; i < right; i += step {
		x = append(x, i)
		y = append(y, u.Distribution(i))
	}

	return Result{
		Name: fmt.Sprintf("a = %.2f, b = %.2f", u.a, u.b),
		X:    x,
		Y:    y,
	}
}

func (u *UniformDistribution) ComputeDensity(step, left, right float64) Result {
	x := make([]float64, 0)
	y := make([]float64, 0)

	for i := left; i < right; i += step {
		x = append(x, i)
		y = append(y, u.Density(i))
	}

	return Result{
		Name: fmt.Sprintf("a = %.2f, b = %.2f", u.a, u.b),
		X:    x,
		Y:    y,
	}
}
