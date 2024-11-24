package method

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"image/color"
	"os"
)

type Method interface {
	I() []float64
	U() []float64
	Rp() []float64
	IRp() []float64
	T0() []float64
	T() []float64
	Calculate()
}

type MethodInfo struct {
	M     Method
	Name  string
	Color color.Color
}

func selectI(m MethodInfo) ([]float64, []float64) {
	return m.M.T(), m.M.I()
}

func selectU(m MethodInfo) ([]float64, []float64) {
	return m.M.T(), m.M.U()
}

func selectRp(m MethodInfo) ([]float64, []float64) {
	return m.M.T(), m.M.Rp()
}

func selectIRp(m MethodInfo) ([]float64, []float64) {
	return m.M.T(), m.M.IRp()
}

func selectT0(m MethodInfo) ([]float64, []float64) {
	return m.M.T(), m.M.T0()
}

func Plot(n int, onlyI bool, ms ...MethodInfo) error {
	if err := plotParam(n, "I(t)", selectI, ms...); err != nil {
		return fmt.Errorf("draw plot for I(t): %w", err)
	}

	if onlyI {
		return nil
	}

	if err := plotParam(n, "U(t)", selectU, ms...); err != nil {
		return fmt.Errorf("draw plot for U(t): %w", err)
	}

	if err := plotParam(n, "Rp(t)", selectRp, ms...); err != nil {
		return fmt.Errorf("draw plot for Rp(t): %w", err)
	}

	if err := plotParam(n, "I(t) * Rp(t)", selectIRp, ms...); err != nil {
		return fmt.Errorf("draw plot for I(t) * Rp(t): %w", err)
	}

	if err := plotParam(n, "T0(t)", selectT0, ms...); err != nil {
		return fmt.Errorf("draw plot for T0(t): %w", err)
	}

	return nil
}

func toXYs(x, y []float64) plotter.XYs {
	res := make(plotter.XYs, len(x))
	for i := range res {
		res[i].X = x[i]
		res[i].Y = y[i]
	}

	return res
}

func plotParam(n int, valueName string, selector func(MethodInfo) ([]float64, []float64), ms ...MethodInfo) error {
	p := plot.New()

	p.X.Label.Text = "t"
	p.Y.Label.Text = valueName

	for _, m := range ms {
		x, y := selector(m)
		l, err := plotter.NewLine(toXYs(x, y))
		if err != nil {
			return fmt.Errorf("create scatter: %w", err)
		}
		l.Color = m.Color
		p.Legend.Add(m.Name, l)
		p.Add(l)
	}

	directory := "All"
	if len(ms) == 1 {
		directory = ms[0].Name
	}

	filename := fmt.Sprintf("data/output/%d/%s/%s.pdf", n, directory, valueName)

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	w, err := p.WriterTo(512, 512, "pdf")
	if err != nil {
		return fmt.Errorf("create plot writer: %w", err)
	}

	if _, err := w.WriteTo(f); err != nil {
		return fmt.Errorf("write plot to file: %w", err)
	}

	return nil
}
