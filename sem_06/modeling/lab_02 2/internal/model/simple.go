package model

import (
	"lab_02/internal/interpolation"
)

type SimpleModel struct {
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

func (m *SimpleModel) Rp(I float64) float64 {
	return -m.Rk
}

func (m *SimpleModel) T0(I float64) float64 {
	return m.IToT0.Get(I)
}

func (m *SimpleModel) DIDT(U, I float64) float64 {
	return U / m.Lk
}

func (m *SimpleModel) DUDT(I float64) float64 {
	return -I / m.Ck
}
