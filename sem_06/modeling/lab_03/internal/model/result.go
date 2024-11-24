package model

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"os"
)

type Result struct {
	Radius      []float64
	Temperature []float64
	F1, F2      float64
	Iterations  int
	err         error
}

func (r *Result) Print() *Result {
	if r.err != nil {
		return r
	}

	fmt.Printf("f1 = %e\n", r.F1)
	fmt.Printf("f2 = %e\n", r.F2)
	fmt.Printf("iterations = %d\n", r.Iterations)
	fmt.Printf("first = %f\n", r.Temperature[0])

	return r
}

func (r *Result) SavePlot(name string) *Result {
	if r.err != nil {
		return r
	}

	p := plot.New()

	p.X.Label.Text = "Radius"
	p.Y.Label.Text = "Temperature"

	l, err := plotter.NewLine(toXYs(r.Radius, r.Temperature))
	if err != nil {
		r.err = fmt.Errorf("create scatter: %w", err)
		return r
	}

	p.Add(l)
	p.Add(plotter.NewGrid())

	f, err := os.Create(name)
	if err != nil {
		r.err = fmt.Errorf("create file: %w", err)
		return r
	}
	defer f.Close()

	w, err := p.WriterTo(512, 512, "pdf")
	if err != nil {
		r.err = fmt.Errorf("create plot writer: %w", err)
		return r
	}

	if _, err := w.WriteTo(f); err != nil {
		r.err = fmt.Errorf("write plot to file: %w", err)
		return r
	}

	return r
}

func (r *Result) Error() error {
	return r.err
}

func toXYs(x, y []float64) plotter.XYs {
	res := make(plotter.XYs, len(x))
	for i := range res {
		res[i].X = x[i]
		res[i].Y = y[i]
	}

	return res
}
