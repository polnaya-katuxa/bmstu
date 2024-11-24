package model

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/interp"
	"os"
)

type Range struct {
	Min, Max float64
}

type Steps struct {
	Time   int
	Radius int
}

type Step struct {
	Time   float64
	Radius float64
}

type CParams struct {
	A, B, C, M float64
}

type F0Params struct {
	Const bool
	Value float64
}

type Mode struct {
	Impulse   bool
	Frequency float64
	ZeroC     bool
}

type Model struct {
	Lambda   interp.Predictor
	K        interp.Predictor
	N        float64
	Radius   Range
	Time     Range
	Steps    Steps
	T0       float64
	TObj     float64
	Sigma    float64
	FMax     float64
	TMax     float64
	Alpha    float64
	Eps1     float64
	Eps2     float64
	CParams  CParams
	F0Params F0Params
	Mode     Mode

	baseP  float64
	baseF  float64
	step   Step
	period float64
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

	m := &Model{
		Lambda:   lambda,
		K:        k,
		N:        c.N,
		Radius:   c.Radius,
		Time:     c.Time,
		Steps:    c.Steps,
		T0:       c.T0,
		TObj:     c.TObj,
		Sigma:    c.Sigma,
		FMax:     c.FMax,
		TMax:     c.TMax,
		Alpha:    c.Alpha,
		Eps1:     c.Eps1,
		Eps2:     c.Eps2,
		CParams:  c.CParams,
		F0Params: c.F0,
		Mode:     c.Mode,
	}

	m.computeSteps()
	m.computePeriod()

	return m, nil
}

func (m *Model) computeRadiusStep() float64 {
	return (m.Radius.Max - m.Radius.Min) / float64(m.Steps.Radius)
}

func (m *Model) computeTimeStep() float64 {
	return (m.Time.Max - m.Time.Min) / float64(m.Steps.Time)
}

func (m *Model) computeSteps() {
	m.step = Step{
		Time:   m.computeTimeStep(),
		Radius: m.computeRadiusStep(),
	}
}

func (m *Model) computePeriod() {
	if m.Mode.Impulse {
		m.period = 1e6 / m.Mode.Frequency
	}
}

type configTable struct {
	X []float64
	Y []float64
}

type config struct {
	Lambda  configTable
	K       configTable
	N       float64
	Radius  Range
	Time    Range
	Steps   Steps
	T0      float64
	TObj    float64
	Sigma   float64
	FMax    float64
	TMax    float64
	Alpha   float64
	Eps1    float64
	Eps2    float64
	CParams CParams
	F0      F0Params
	Mode    Mode
}
