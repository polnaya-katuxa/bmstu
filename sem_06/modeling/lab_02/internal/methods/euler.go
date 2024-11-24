package methods

import "lab_02/internal/model"

type Euler struct {
	TStart float64
	TEnd   float64
	H      float64
	I0     float64
	U0     float64
	M      model.Model
	result
}

func (e *Euler) Calculate() {
	I := e.I0
	U := e.U0

	for t := e.TStart; t <= e.TEnd; t += e.H {
		e.valuesT = append(e.valuesT, t)
		e.valuesI = append(e.valuesI, I)
		e.valuesU = append(e.valuesU, U)
		e.valuesT0 = append(e.valuesT0, e.M.T0(I))
		e.valuesRp = append(e.valuesRp, e.M.Rp(I))
		e.valuesIRp = append(e.valuesIRp, I*e.M.Rp(I))

		IOld := I
		I = I + e.H*e.M.DIDT(U, I)
		U = U + e.H*e.M.DUDT(IOld)
	}
}
