package methods

import "lab_02/internal/model"

type RungeKutta2 struct {
	TStart float64
	TEnd   float64
	H      float64
	I0     float64
	U0     float64
	Alpha  float64
	M      model.Model
	result
}

func (r *RungeKutta2) Calculate() {
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

		I = I + ((1-r.Alpha)*k1 + r.Alpha*k2)
		U = U + ((1-r.Alpha)*q1 + r.Alpha*q2)
	}
}
