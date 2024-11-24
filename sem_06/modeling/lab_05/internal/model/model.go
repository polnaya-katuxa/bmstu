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
	X    int
	Z    int
	Time int
}

type Step struct {
	X float64
	Z float64
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

type Laser struct {
	X, Z float64
	F0   float64
}

type Model struct {
	F0     float64
	Beta   float64
	Lasers []Laser

	ZMax   float64
	XMax   float64
	U0     float64
	Eps    float64
	Tau    float64
	C      float64
	Lambda float64

	f func(x, z float64) float64

	Steps Steps
	step  Step

	K                              interp.Predictor
	Alpha1, Alpha2, Alpha3, Alpha4 float64

	BoundaryVariant int

	S float64
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

	k := new(interp.AkimaSpline)
	if err := k.Fit(c.K.X, c.K.Y); err != nil {
		return nil, fmt.Errorf("fit interpolation table to k: %w", err)
	}

	m := &Model{
		F0:              c.F0,
		Beta:            c.Beta,
		Lasers:          c.Lasers,
		ZMax:            c.ZMax,
		XMax:            c.XMax,
		U0:              c.U0,
		Eps:             c.Eps,
		Tau:             c.Tau,
		C:               c.C,
		Lambda:          c.Lambda,
		Steps:           c.Steps,
		K:               k,
		Alpha1:          c.Alpha1,
		Alpha2:          c.Alpha2,
		Alpha3:          c.Alpha4,
		Alpha4:          c.Alpha4,
		BoundaryVariant: c.BoundaryVariant,
		S:               c.S,
	}

	m.computeSteps()
	m.f = m.newFAll()

	return m, nil
}

func (model *Model) computeSteps() {
	model.step = Step{
		X: model.XMax / float64(model.Steps.X),
		Z: model.ZMax / float64(model.Steps.Z),
	}
}

type configTable struct {
	X []float64
	Y []float64
}

type config struct {
	F0     float64
	Beta   float64
	Lasers []Laser

	ZMax   float64
	XMax   float64
	U0     float64
	Eps    float64
	Tau    float64
	C      float64
	Lambda float64

	Steps Steps

	K configTable

	Alpha1, Alpha2, Alpha3, Alpha4 float64

	BoundaryVariant int

	S float64
}
