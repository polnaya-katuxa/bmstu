package model

import (
	"lab_02/internal/interpolation"
	"math"
)

type FullModel struct {
	R        float64
	Lp       float64
	Ck       float64
	Rk       float64
	Lk       float64
	Tw       float64
	IToT0    *interpolation.Interpolation
	IToM     *interpolation.Interpolation
	TToSigma *interpolation.Interpolation
}

func (m *FullModel) DIDT(U, I float64) float64 {
	return (U - I*(m.Rk+m.Rp(I))) / m.Lk
}

func (m *FullModel) DUDT(I float64) float64 {
	return -I / m.Ck
}

func (m *FullModel) Rp(I float64) float64 {
	return m.Lp / (2.0 * math.Pi * math.Pow(m.R, 2) * m.Integral(0, 1, I))
}

func (m *FullModel) T0(I float64) float64 {
	return m.IToT0.Get(I)
}

func (m *FullModel) T(I, z float64) float64 {
	T0 := m.T0(I)
	M := m.IToM.Get(I)

	return T0 + (m.Tw-T0)*math.Pow(z, M)
}

func (m *FullModel) Sigma(I, z float64) float64 {
	return m.TToSigma.Get(m.T(I, z))
}

func (m *FullModel) Integral(a, b, I float64) float64 {
	res := 0.0
	h := (b - a) / n
	res += m.Sigma(I, a)*a + m.Sigma(I, b)*b
	for z := a + h; z < b; z += h {
		res += 2.0 * m.Sigma(I, z) * z
	}

	return h * res / 2.0
}
