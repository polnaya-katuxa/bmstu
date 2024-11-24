package model

import (
	"fmt"
	"github.com/mazznoer/colorgrad"
	"github.com/samber/lo"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"os"
)

const (
	margin = 5
)

var gradient = colorgrad.Warm()

type Result struct {
	Time        []float64
	Radius      []float64
	Z           []float64
	Temperature [][]float64

	err error
}

func (r *Result) getLimits() (float64, float64) {
	max := lo.MaxBy(r.Temperature, func(a []float64, b []float64) bool {
		return lo.Max(a) < lo.Max(b)
	})

	min := lo.MinBy(r.Temperature, func(a []float64, b []float64) bool {
		return lo.Max(a) < lo.Max(b)
	})

	return lo.Min(min), lo.Max(max)
}

func (r *Result) PlotByTimeMinRadius(name string) *Result {
	if r.err != nil {
		return r
	}

	p := plot.New()

	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Temperature"

	min, max := r.getLimits()
	p.Y.Min = min - margin
	p.Y.Max = max + margin

	colors := gradient.Colors(uint(len(r.Radius)))

	temp := lo.Map(r.Temperature, func(e []float64, _ int) float64 {
		return e[0]
	})

	l, err := plotter.NewLine(toXYs(r.Time, temp))
	if err != nil {
		r.err = fmt.Errorf("create scatter: %w", err)
		return r
	}
	l.Color = colors[0]

	l.Width = 2
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

func (r *Result) PlotByTime(name string) *Result {
	if r.err != nil {
		return r
	}

	p := plot.New()

	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Temperature"

	min, max := r.getLimits()
	p.Y.Min = min - margin
	p.Y.Max = max + margin

	colors := gradient.Colors(uint(len(r.Radius)))

	for i := range r.Radius {
		temp := lo.Map(r.Temperature, func(e []float64, _ int) float64 {
			return e[i]
		})

		l, err := plotter.NewLine(toXYs(r.Time, temp))
		if err != nil {
			r.err = fmt.Errorf("create scatter: %w", err)
			return r
		}
		l.Color = colors[i]

		p.Add(l)
	}

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

func (r *Result) PlotByZ(name string) *Result {
	if r.err != nil {
		return r
	}

	p := plot.New()

	p.X.Label.Text = "Z"
	p.Y.Label.Text = "T"

	min, max := r.getLimits()
	p.Y.Min = min - margin
	p.Y.Max = max + margin

	colors := gradient.Colors(uint(len(r.Time)))

	for i := range r.Time {
		l, err := plotter.NewLine(toXYs(r.Z, r.Temperature[i]))
		if err != nil {
			r.err = fmt.Errorf("create scatter: %w", err)
			return r
		}

		l.Color = colors[i]
		l.Width = 2

		p.Add(l)
		p.Legend.Add(fmt.Sprintf("t = %.0f", r.Time[i]), l)
	}

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
