package methods

type result struct {
	valuesI   []float64
	valuesU   []float64
	valuesRp  []float64
	valuesIRp []float64
	valuesT0  []float64
	valuesT   []float64
}

func (r result) I() []float64 {
	return r.valuesI
}

func (r result) U() []float64 {
	return r.valuesU
}

func (r result) Rp() []float64 {
	return r.valuesRp
}

func (r result) IRp() []float64 {
	return r.valuesIRp
}

func (r result) T0() []float64 {
	return r.valuesT0
}

func (r result) T() []float64 {
	return r.valuesT
}
