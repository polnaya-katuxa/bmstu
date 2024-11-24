package model

type System []*EquationCoefficients

func NewSystem(size int) System {
	system := System(make([]*EquationCoefficients, size))
	for i := range system {
		system[i] = new(EquationCoefficients)
	}

	return system
}

func (m System) runThrough(start, end BoundaryConditions) []float64 {
	result := make([]float64, len(m)+2)

	k := RunThroughCoefficient{
		Alpha: -(start.K / start.M),
		Beta:  start.P / start.M,
	}

	coefficients := []RunThroughCoefficient{k}

	for i := 0; i < len(m); i++ {
		k = RunThroughCoefficient{
			Alpha: m[i].C / (m[i].B - m[i].A*k.Alpha),
			Beta:  (m[i].A*k.Beta + m[i].D) / (m[i].B - m[i].A*k.Alpha),
		}

		coefficients = append(coefficients, k)
	}

	result[len(result)-1] = (end.P - end.M*k.Beta) / (end.K + end.M*k.Alpha)

	for i := len(m); i >= 0; i-- {
		result[i] = coefficients[i].Alpha*result[i+1] + coefficients[i].Beta
	}

	return result
}

type EquationCoefficients struct {
	A, B, C, D float64
}

type BoundaryConditions struct {
	M, P, K float64
}

type RunThroughCoefficient struct {
	Alpha, Beta float64
}
