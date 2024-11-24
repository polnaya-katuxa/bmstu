package model

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/interp"
	"math"
	"os"
)

const (
	Eps = 1e-8
)

type Model struct {
	Lambda  interp.Predictor
	K       interp.Predictor
	N       float64
	R0      float64
	RMax    float64
	Steps   int
	T0      float64
	Sigma   float64
	F0      float64
	Alpha   float64
	Eps1    float64
	Eps2    float64
	Epsilon float64
	baseP   float64
	baseF   float64
	step    float64
}

func FromFile(name string) (*Model, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	var c config
	if err := json.NewDecoder(file).Decode(&c); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}

	lambda := new(interp.AkimaSpline)
	if err := lambda.Fit(c.Lambda.X, c.Lambda.Y); err != nil {
		return nil, fmt.Errorf("fit interpolation table to lambda: %w", err)
	}

	k := new(interp.AkimaSpline)
	if err := k.Fit(c.K.X, c.K.Y); err != nil {
		return nil, fmt.Errorf("fit interpolation table to k: %w", err)
	}

	return &Model{
		Lambda:  lambda,
		K:       k,
		N:       c.N,
		R0:      c.R0,
		RMax:    c.RMax,
		Steps:   c.Steps,
		T0:      c.T0,
		Sigma:   c.Sigma,
		F0:      c.F0,
		Alpha:   c.Alpha,
		Eps1:    c.Eps1,
		Eps2:    c.Eps2,
		Epsilon: c.Epsilon,
	}, nil
}

func (m *Model) computeStep() float64 {
	return (m.RMax - m.R0) / float64(m.Steps)
}

func (m *Model) Compute() *Result {
	m.step = m.computeStep()
	invRMax := 1 / m.RMax

	r := make([]float64, m.Steps+1)
	for i := 0; i <= m.Steps; i++ {
		r[i] = m.R0 + float64(i)*m.step
	}

	z := make([]float64, m.Steps+1)
	for i, e := range r {
		z[i] = e * invRMax
	}

	result := make([]float64, m.Steps+1)
	for i := range r {
		result[i] = m.T0
	}

	irsqrh := 1 / (m.RMax * m.RMax * m.step) // TODO naming

	m.baseP = 4 * m.N * m.N * m.Sigma
	m.baseF = m.baseP * math.Pow(m.T0, 4)

	system := NewSystem(m.Steps - 1)

	f1 := 0.0
	f2 := 0.0

	iterations := 0

	for {
		for i := 0; i < len(system); i++ {
			zPrev := (z[i] + z[i+1]) / 2         // z{i-1/2}
			zNext := (z[i+2] + z[i+1]) / 2       // z{i+1/2}
			v := (zNext*zNext - zPrev*zPrev) / 2 // единичный объем

			XPrev := (m.Lambda.Predict(result[i]) + m.Lambda.Predict(result[i+1])) / 2   // ae{i-1/2}
			XNext := (m.Lambda.Predict(result[i+2]) + m.Lambda.Predict(result[i+1])) / 2 // ae{i-1/2}

			k := m.K.Predict(result[i+1])
			p := m.baseP * k * math.Pow(result[i+1], 3)
			f := m.baseF * k

			//a := irsqrh * zPrev * XPrev
			//system[i].A = a
			//fmt.Printf("%e, %e, %e, %e\n", 1.0-1.0, system[i].A, a, system[i].A-a)
			system[i].A = irsqrh * zPrev * XPrev
			system[i].C = irsqrh * zNext * XNext
			system[i].B = system[i].A + system[i].C + p*v
			system[i].D = f * v
		}

		start := m.leftBoundaryConditions(z, result)
		end := m.rightBoundaryConditions(z, result)

		resultNew := system.runThrough(start, end)
		c := 0

		for i := 0; i <= m.Steps; i++ {
			if math.Abs((result[i]-resultNew[i])/resultNew[i]) < m.Eps1 {
				c++
			}
		}

		result = resultNew
		iterations++

		if c == m.Steps+1 {
			f1 = m.R0*m.F0 - m.RMax*m.Alpha*(resultNew[len(resultNew)-1]-m.T0)

			tmp := make([]float64, m.Steps+1)
			for i := range tmp {
				tmp[i] = m.K.Predict(resultNew[i]) * (math.Pow(resultNew[i], 4) - math.Pow(m.T0, 4)) * r[i]
			}

			interpolation := new(interp.AkimaSpline)
			interpolation.Fit(r, tmp)

			f2 = 4 * m.N * m.N * m.Sigma * trapezoidal(m.R0, m.RMax, m.Steps, interpolation)

			if math.Abs(f1) > Eps || math.Abs(f2) < Eps {
				break
			}

			if math.Abs((f1-f2)/f1) < m.Eps2 {
				break
			}
		}
	}

	return &Result{
		Radius:      r,
		Temperature: result,
		F1:          f1,
		F2:          f2,
		Iterations:  iterations,
	}
}

// leftBoundaryConditions расчет коэффициентов левого краевого условия вида M0 * y0 + K0 * y1 = P0.
func (m *Model) leftBoundaryConditions1(z, result []float64) BoundaryConditions {
	step := m.computeStep()

	lambda := m.Lambda.Predict(result[0])

	return BoundaryConditions{
		M: lambda/(m.RMax*step) + m.Sigma*m.Epsilon*math.Pow(m.T0, 3),
		P: -lambda / (m.RMax * step),
		K: m.F0,
	}
}

func (m *Model) leftBoundaryConditions(z, result []float64) BoundaryConditions {
	step := m.computeStep()

	zLeft := (z[0] + z[1]) / 2 // z{1/2}

	XLeft := (m.Lambda.Predict(result[0]) + m.Lambda.Predict(result[1])) / 2 // ae{1/2}

	k0 := m.K.Predict(result[0])
	k1 := m.K.Predict(result[1])

	p0 := m.baseP * k0 * math.Pow(result[0], 3)
	p1 := m.baseP * k1 * math.Pow(result[1], 3)

	pLeft := (p0 + p1) / 2 // p{1/2}

	baseF := m.baseP * math.Pow(m.T0, 4)
	f0 := baseF * k0
	f1 := baseF * k1

	fLeft := (f0 + f1) / 2 // f{1/2}

	return BoundaryConditions{
		M: (XLeft/(math.Pow(m.RMax, 2)*step)+step*pLeft/8)*zLeft + step*p0*z[0]/4,
		P: z[0]*(m.F0-m.Sigma*m.Epsilon*math.Pow(result[0], 4))/m.RMax + step*(f0*z[0]+fLeft*zLeft)/4,
		K: (step*pLeft/8 - XLeft/(math.Pow(m.RMax, 2)*step)) * zLeft,
	}
}

// rightBoundaryConditions расчет коэффициентов правого краевого условия вида MN * y{N-1} + KN * yN = PN.
func (m *Model) rightBoundaryConditions(z, result []float64) BoundaryConditions {
	step := m.computeStep()

	zRight := (z[m.Steps] + z[m.Steps-1]) / 2 // z{N-1/2}

	XRight := (m.Lambda.Predict(result[m.Steps]) + m.Lambda.Predict(result[m.Steps-1])) / 2 // ae{N-1/2}

	kPre := m.K.Predict(result[m.Steps-1])
	kLast := m.K.Predict(result[m.Steps])

	pPre := m.baseP * kPre * math.Pow(result[m.Steps-1], 3)
	pLast := m.baseP * kLast * math.Pow(result[m.Steps], 3)

	pRight := (pPre + pLast) / 2

	fPre := m.baseF * kPre
	fLast := m.baseF * kLast

	fRight := (fPre + fLast) / 2

	return BoundaryConditions{
		M: (-XRight/(math.Pow(m.RMax, 2)*step) + step*pRight/8) * zRight,
		P: (m.Alpha*m.T0/m.RMax)*z[m.Steps] + step*(fLast*z[m.Steps]+fRight*zRight)/4,
		K: XRight*zRight/(math.Pow(m.RMax, 2)*step) + z[m.Steps]*m.Alpha/m.RMax + step*(pRight*zRight/2+pLast*z[m.Steps])/4,
	}
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

type configTable struct {
	X []float64
	Y []float64
}

type config struct {
	Lambda  configTable
	K       configTable
	N       float64
	R0      float64
	RMax    float64
	Steps   int
	T0      float64
	Sigma   float64
	F0      float64
	Alpha   float64
	Eps1    float64
	Eps2    float64
	Epsilon float64
}
