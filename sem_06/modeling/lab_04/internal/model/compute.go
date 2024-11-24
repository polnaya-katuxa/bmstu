package model

import (
	"fmt"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/interp"
	"math"
)

func (m *Model) c(T float64) float64 {
	if m.Mode.ZeroC {
		return 0.0
	}

	a2, b2, c2, m2 := m.CParams.A, m.CParams.B, m.CParams.C, m.CParams.M
	return a2 + b2*math.Pow(T, m2) - c2/math.Pow(T, 2)
}

func (m *Model) f(T float64) float64 {
	return 4 * m.K.Predict(T) * math.Pow(m.N, 2) * m.Sigma * (math.Pow(T, 4) - math.Pow(m.T0, 4))
}

func (m *Model) F0(t float64) float64 {
	if m.F0Params.Const {
		return m.F0Params.Value
	}

	if m.Mode.Impulse {
		t -= math.Floor(t/m.period) * m.period
	}

	return (m.FMax / m.TMax) * t * math.Exp(1-(t/m.TMax))
}

func (m *Model) Compute() *Result {
	r := make([]float64, m.Steps.Radius+1)
	for i := range r {
		r[i] = m.Radius.Min + float64(i)*m.step.Radius
	}

	t := make([]float64, m.Steps.Time)
	for i := range t {
		t[i] = m.Time.Min + float64(i)*m.step.Time
	}

	T := make([][]float64, m.Steps.Time)
	for i := range T {
		T[i] = make([]float64, m.Steps.Radius+1)
	}
	for i := range T[0] {
		T[0][i] = m.TObj
	}

	for ct := 0; ct < m.Steps.Time-1; ct++ {
		system := NewSystem(m.Steps.Radius - 1)

		m.baseP = 4 * m.N * m.N * m.Sigma
		m.baseF = m.baseP * math.Pow(m.T0, 4)

		tmpT := T[ct]
		iterations := 0

		for {
			for n := 0; n < len(system); n++ {
				c := m.c(tmpT[n])
				f := m.f(tmpT[n])

				stepTime := m.step.Time

				system[n].A = r[n] * stepTime * m.Lambda.Predict(tmpT[n])
				system[n].C = m.step.Radius*stepTime*m.Lambda.Predict(tmpT[n]) + r[n]*stepTime*m.Lambda.Predict(tmpT[n+1])
				system[n].B = system[n].C + r[n]*math.Pow(m.step.Radius, 2)*c + r[n]*stepTime*m.Lambda.Predict(tmpT[n])
				system[n].D = -f*r[n]*math.Pow(m.step.Radius, 2)*stepTime + r[n]*math.Pow(m.step.Radius, 2)*c*T[ct][n]
			}

			start := m.leftBoundaryConditions(tmpT, t, ct)
			end := m.rightBoundaryConditions(tmpT, t, ct)

			newT := system.runThrough(start, end)

			c := 0

			for i := 0; i <= m.Steps.Radius; i++ {
				if math.Abs((tmpT[i]-newT[i])/newT[i]) < m.Eps1 {
					c++
				}
			}

			tmpT = newT
			iterations++

			if c == m.Steps.Radius+1 {
				break
			}

			f1 := m.Radius.Min*m.F0(t[ct+1]) - m.Radius.Max*m.Alpha*(tmpT[len(tmpT)-1]-m.T0)

			tmp := make([]float64, m.Steps.Radius+1)
			for i := range tmp {
				tmp[i] = m.K.Predict(tmpT[i]) * (math.Pow(tmpT[i], 4) - math.Pow(m.T0, 4)) * r[i]
			}

			interpolation := new(interp.AkimaSpline)
			interpolation.Fit(r, tmp)

			f2 := 4 * m.N * m.N * m.Sigma * trapezoidal(m.Radius.Min, m.Radius.Max, m.Steps.Radius, interpolation)

			if math.Abs((f1-f2)/f1) < m.Eps2 {
				break
			}
		}

		fmt.Printf("t = %.0f, iterations = %d\n", t[ct], iterations)
		T[ct+1] = tmpT
	}

	return &Result{
		Radius: r,
		Z: lo.Map(r, func(e float64, _ int) float64 {
			return e / m.Radius.Max
		}),
		Temperature: T,
		Time:        t,
	}
}

// leftBoundaryConditions расчет коэффициентов левого краевого условия вида M0 * y0 + K0 * y1 = P0.
func (m *Model) leftBoundaryConditions(T []float64, t []float64, ct int) BoundaryConditions {
	lambda := m.Lambda.Predict(T[0])

	return BoundaryConditions{
		M: lambda,
		K: -lambda,
		P: m.F0(t[ct+1]) * m.step.Radius,
	}

	//return BoundaryConditions{
	//	M: 1,
	//	K: 0,
	//	P: 300,
	//}
}

// rightBoundaryConditions расчет коэффициентов правого краевого условия вида MN * y{N-1} + KN * yN = PN.
func (m *Model) rightBoundaryConditions(T []float64, _ []float64, _ int) BoundaryConditions {
	lambda := m.Lambda.Predict(T[m.Steps.Radius-1])

	return BoundaryConditions{
		M: lambda,
		K: -lambda - m.Alpha*m.step.Radius,
		P: -m.Alpha * m.step.Radius * m.T0,
	}

	//return BoundaryConditions{
	//	M: 0,
	//	K: 1,
	//	P: 1000,
	//}
}

func trapezoidal(start, end float64, steps int, f interp.Predictor) float64 {
	step := (end - start) / float64(steps)
	x := start + step

	result := f.Predict(start) + f.Predict(end)

	for x < end {
		result += 2 * f.Predict(x)
		x += step
	}

	return result * step / 2
}
