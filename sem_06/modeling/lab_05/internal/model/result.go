package model

import (
	"fmt"
	"github.com/Arafatk/glot"
	"log"
	"os"
	"os/exec"
)

type Result struct {
	X []float64
	Z []float64
	U [][]float64

	err error
}

func (r *Result) Plot(name string) *Result {
	plot, err := glot.NewPlot(3, true, false)
	if err != nil {
		r.err = err
		return r
	}

	x := make([]float64, 0, len(r.X)*len(r.Z))
	z := make([]float64, 0, len(r.X)*len(r.Z))
	u := make([]float64, 0, len(r.X)*len(r.Z))

	for i := 0; i < len(r.X); i++ {
		for j := 0; j < len(r.Z); j++ {
			x = append(x, r.X[i])
			z = append(z, r.Z[j])
			u = append(u, r.U[i][j])
		}
	}

	if err := plot.AddPointGroup("u", "dots", [][]float64{x, z, u}); err != nil {
		r.err = err
		return r
	}

	if err := plot.SavePlot(name); err != nil {
		r.err = err
		return r
	}

	log.Println("Saving plot...")
	fmt.Scanln()

	return r
}

func (r *Result) ToFile(name string) *Result {
	file, err := os.Create(name)
	if err != nil {
		r.err = err
		return r
	}

	for _, e := range r.X {
		fmt.Fprintf(file, "%f ", e)
	}
	fmt.Fprintln(file)

	for _, e := range r.Z {
		fmt.Fprintf(file, "%f ", e)
	}
	fmt.Fprintln(file)

	for n := range r.U {
		for m := range r.U[n] {
			fmt.Fprintf(file, "%f ", r.U[n][m])
		}
		fmt.Fprintln(file)
	}

	return r
}

func (r *Result) Python() *Result {
	cmd := exec.Command("python3", "scripts/plot.py")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		r.err = err
	}

	return r
}

func (r *Result) Error() error {
	return r.err
}
