package methods

import "lab_02/internal/model"

type RungeKutta4 struct {
	TStart float64
	TEnd   float64
	H      float64
	I0     float64
	U0     float64
	M      model.Model
	result
}

func (r *RungeKutta4) Calculate() {
	I := r.I0
	U := r.U0

	for t := r.TStart; t <= r.TEnd; t += r.H {
		r.valuesT = append(r.valuesT, t)
		r.valuesI = append(r.valuesI, I)
		r.valuesU = append(r.valuesU, U)
		r.valuesT0 = append(r.valuesT0, r.M.T0(I))
		r.valuesRp = append(r.valuesRp, r.M.Rp(I))
		r.valuesIRp = append(r.valuesIRp, I*r.M.Rp(I))

		k1 := r.H * r.M.DIDT(U, I)
		q1 := r.H * r.M.DUDT(I)

		k2 := r.H * r.M.DIDT(U+q1/2.0, I+k1/2.0)
		q2 := r.H * r.M.DUDT(I+k1/2.0)

		k3 := r.H * r.M.DIDT(U+q2/2.0, I+k2/2.0)
		q3 := r.H * r.M.DUDT(I+k2/2.0)

		k4 := r.H * r.M.DIDT(U+q3/2.0, I+k3/2.0)
		q4 := r.H * r.M.DUDT(I+k3/2.0)

		I = I + (k1+2.0*k2+2.0*k3+k4)/6.0
		U = U + (q1+2.0*q2+2.0*q3+q4)/6.0
	}
}
