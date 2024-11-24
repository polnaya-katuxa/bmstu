package table

import (
	"fmt"
	"github.com/Arafatk/glot"
	"github.com/jedib0t/go-pretty/v6/table"
	"math"
)

type Table struct {
	argName  string
	funcName string

	arguments []float64

	analytic []float64

	picard1 []float64
	picard2 []float64
	picard3 []float64
	picard4 []float64

	euler []float64
}

func (t *Table) initTable(task Task, maxArg, step float64) {
	n := int((maxArg-task.MinArg())/step) + 1
	t.arguments = make([]float64, 0, n)

	t.analytic = make([]float64, 0, n)

	t.picard1 = make([]float64, 0, n)
	t.picard2 = make([]float64, 0, n)
	t.picard3 = make([]float64, 0, n)
	t.picard4 = make([]float64, 0, n)

	t.euler = make([]float64, 0, n)
}

func Generate(task Task, argName, funcName string, maxArg, step float64) Table {
	t := Table{
		argName:  argName,
		funcName: funcName,
	}
	t.initTable(task, maxArg, step)

	eulerFunction := task.MinFunction()
	for arg := task.MinArg(); arg <= maxArg; arg += step {
		t.arguments = append(t.arguments, arg)

		t.analytic = append(t.analytic, task.Analytic(arg))

		t.picard1 = append(t.picard1, task.Picard1(arg))
		t.picard2 = append(t.picard2, task.Picard2(arg))
		t.picard3 = append(t.picard3, task.Picard3(arg))
		t.picard4 = append(t.picard4, task.Picard4(arg))

		t.euler = append(t.euler, eulerFunction)
		eulerFunction += step * task.Function(arg, eulerFunction)
	}

	return t
}

func (t *Table) Print(n int) {
	tw := table.NewWriter()

	tw.AppendHeader(table.Row{t.argName, "analytic", "euler", "picard 1", "picard 2", "picard 3", "picard 4"})
	tw.SetStyle(table.StyleDouble)

	for i := 0; i < len(t.arguments); i += n {
		tw.AppendRow(table.Row{
			fmt.Sprintf("%.7f", t.arguments[i]),
			fmt.Sprintf("%.7f", t.analytic[i]),
			fmt.Sprintf("%.7f", t.euler[i]),
			fmt.Sprintf("%.7f", t.picard1[i]),
			fmt.Sprintf("%.7f", t.picard2[i]),
			fmt.Sprintf("%.7f", t.picard3[i]),
			fmt.Sprintf("%.7f", t.picard4[i]),
		})
	}

	fmt.Println(tw.Render())
}

func (t *Table) Plot(name string) error {
	p, err := glot.NewPlot(2, false, false)
	if err != nil {
		return err
	}

	if !math.IsNaN(t.analytic[0]) {
		if err := p.AddPointGroup("analytic", "lines", [][]float64{t.arguments, t.analytic}); err != nil {
			return err
		}
	}

	if err := p.AddPointGroup("euler", "lines", [][]float64{t.arguments, t.euler}); err != nil {
		return err
	}

	if err := p.AddPointGroup("picard 1", "lines", [][]float64{t.arguments, t.picard1}); err != nil {
		return err
	}

	if err := p.AddPointGroup("picard 2", "lines", [][]float64{t.arguments, t.picard2}); err != nil {
		return err
	}
	if err := p.AddPointGroup("picard 3", "lines", [][]float64{t.arguments, t.picard3}); err != nil {
		return err
	}

	if err := p.AddPointGroup("picard 4", "lines", [][]float64{t.arguments, t.picard4}); err != nil {
		return err
	}

	p.SetFormat("png")
	if err := p.SavePlot(name); err != nil {
		return err
	}

	fmt.Scanln()

	return nil
}
