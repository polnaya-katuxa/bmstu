package task1

import (
	"fmt"
	"image/color"
	"lab_02/internal/app/method"
	"lab_02/internal/interpolation"
	"lab_02/internal/methods"
	"lab_02/internal/model"
	"math"
)

const (
	maxI = 791.126266
)

func Run() error {
	ms, err := initMethods()
	if err != nil {
		return fmt.Errorf("init methods: %w", err)
	}

	for i := range ms {
		ms[i].M.Calculate()
	}

	for i := range ms {
		if err := method.Plot(1, false, ms[i]); err != nil {
			return fmt.Errorf("plot MethodInfo %s: %w", ms[i].Name, err)
		}
	}

	if err := method.Plot(1, false, ms...); err != nil {
		return fmt.Errorf("plot all methods: %w", err)
	}

	for i := range ms {
		delta := math.Abs(maxI - max(ms[i].M.I()))
		fmt.Printf("Rate for %s method = %.3f\n", ms[i].Name, delta/maxI)
	}

	return nil
}

func max(arr []float64) float64 {
	m := arr[0]
	for _, el := range arr {
		if el > m {
			m = el
		}
	}

	return m
}

func initModel() (*model.FullModel, error) {
	IToT0, err := interpolation.FromFile("data/input/I_to_T0.txt")
	if err != nil {
		return nil, fmt.Errorf("open file I to T0: %w", err)
	}

	IToM, err := interpolation.FromFile("data/input/I_to_M.txt")
	if err != nil {
		return nil, fmt.Errorf("open file I to M: %w", err)
	}

	TToSigma, err := interpolation.FromFile("data/input/T_to_Sigma.txt")
	if err != nil {
		return nil, fmt.Errorf("open file T to Sigma: %w", err)
	}

	return &model.FullModel{
		R:        0.35,
		Lp:       12,
		Ck:       268e-7,
		Rk:       0.25,
		Lk:       187e-7,
		Tw:       2000,
		IToT0:    IToT0,
		IToM:     IToM,
		TToSigma: TToSigma,
	}, nil
}

func initMethods() ([]method.MethodInfo, error) {
	m, err := initModel()
	if err != nil {
		return nil, fmt.Errorf("init model: %w", err)
	}

	start := 0.0
	end := 80e-6
	I0 := 2.0
	U0 := 1400.0

	return []method.MethodInfo{
		{
			M: &methods.Euler{
				TStart: start,
				TEnd:   end,
				H:      1e-8,
				I0:     I0,
				U0:     U0,
				M:      m,
			},
			Name:  "Euler",
			Color: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			M: &methods.RungeKutta2{
				TStart: start,
				TEnd:   end,
				H:      2e-8,
				I0:     I0,
				U0:     U0,
				Alpha:  0.5,
				M:      m,
			},
			Name:  "Runge-Kutta 2",
			Color: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		},
		{
			M: &methods.RungeKutta4{
				TStart: start,
				TEnd:   end,
				H:      1e-7,
				I0:     I0,
				U0:     U0,
				M:      m,
			},
			Name:  "Runge-Kutta 4",
			Color: color.RGBA{R: 0, G: 0, B: 255, A: 255},
		},
	}, nil
}
