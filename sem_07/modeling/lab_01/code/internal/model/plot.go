package model

import (
	"fmt"
	"github.com/mazznoer/colorgrad"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"os"
	"path/filepath"
)

const (
	margin = 0
)

var gradient = colorgrad.Turbo()

func toXYs(x, y []float64) plotter.XYs {
	res := make(plotter.XYs, len(x))
	for i := range res {
		res[i].X = x[i]
		res[i].Y = y[i]
	}

	return res
}

func plotOne(filename string, title string, result Result, lim limits) error {
	p := plot.New()

	l, err := plotter.NewLine(toXYs(result.X, result.Y))
	if err != nil {
		return err
	}

	l.Color = gradient.At(0.3)
	p.Add(l)
	p.Legend.Add(result.Name, l)

	p.Title.Text = title

	p.X.Label.Text = "X"
	p.X.Min = lim.minX
	p.X.Max = lim.maxX
	p.Y.Label.Text = "Y"
	p.Y.Min = lim.minY
	p.Y.Max = lim.maxY

	p.Add(plotter.NewGrid())

	wt, err := p.WriterTo(512, 512, "pdf")
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := wt.WriteTo(f); err != nil {
		return err
	}

	return nil
}

func plotAll(filename string, title string, results []Result, lim limits) error {
	p := plot.New()

	colors := gradient.Colors(uint(len(results)))

	for i, result := range results {
		l, err := plotter.NewLine(toXYs(result.X, result.Y))
		if err != nil {
			return err
		}

		l.Color = colors[i]
		p.Add(l)
		p.Legend.Add(result.Name, l)
	}

	p.Title.Text = title

	p.X.Label.Text = "X"
	p.X.Min = lim.minX
	p.X.Max = lim.maxX
	p.Y.Label.Text = "Y"
	p.Y.Min = lim.minY
	p.Y.Max = lim.maxY

	p.Add(plotter.NewGrid())

	wt, err := p.WriterTo(512, 512, "pdf")
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := wt.WriteTo(f); err != nil {
		return err
	}

	return nil
}

type limits struct {
	minX, maxX float64
	minY, maxY float64
}

func findLimits(results []Result) limits {
	l := limits{
		minX: results[0].X[0],
		maxX: results[0].X[len(results[0].X)-1],
		minY: results[0].Y[0],
		maxY: results[0].Y[len(results[0].Y)-1],
	}

	for _, r := range results {
		if r.X[0] < l.minX {
			l.minX = r.X[0]
		}

		if r.X[len(r.X)-1] > l.maxX {
			l.maxX = r.X[len(r.X)-1]
		}

		for i := range r.Y {
			if r.Y[i] < l.minY {
				l.minY = r.Y[i]
			}

			if r.Y[i] > l.maxY {
				l.maxY = r.Y[i]
			}
		}
	}

	l.minX -= margin
	l.maxX += margin
	l.minY -= margin
	l.maxY += margin

	return l
}

func Plot(dirname string, title string, results []Result) error {
	l := findLimits(results)

	if err := plotAll(filepath.Join(dirname, "all.pdf"), title, results, l); err != nil {
		return err
	}

	for i := range results {
		if err := plotOne(filepath.Join(dirname, fmt.Sprintf("%d.pdf", i)), title, results[i], l); err != nil {
			return err
		}
	}

	return nil
}
