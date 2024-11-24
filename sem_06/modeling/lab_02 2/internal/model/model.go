package model

const (
	n = 1000
)

type Model interface {
	DIDT(U, I float64) float64
	DUDT(I float64) float64

	Rp(I float64) float64
	T0(I float64) float64
}
