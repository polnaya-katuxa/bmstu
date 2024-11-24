package app

import (
	"flag"
	"fmt"
	"lab_01/internal/model"
	"path/filepath"
)

var configName = flag.String("config", "config/config.json", "path to config file")

type Model interface {
	ComputeDensity(step, left, right float64) model.Result
	ComputeDistribution(step, left, right float64) model.Result
	RightBound() float64
	LeftBound() float64
}

func Run() error {
	flag.Parse()

	c, err := model.FromFile(*configName)
	if err != nil {
		return fmt.Errorf("open model: %w", err)
	}

	if err := runUniform(c); err != nil {
		return fmt.Errorf("run uniform: %w", err)
	}

	if err := runErlang(c); err != nil {
		return fmt.Errorf("run erlang: %w", err)
	}

	if err := runHyperExponential(c); err != nil {
		return fmt.Errorf("run hyper exponential: %w", err)
	}

	return nil
}

func runUniform(c *model.Config) error {
	models := make([]Model, len(c.Uniform.Params))
	var err error

	for i := range models {
		if models[i], err = model.NewUniform(c.Uniform.Params[i].A, c.Uniform.Params[i].B); err != nil {
			return err
		}
	}

	return compute(models, c.Step, c.Uniform.Name)
}

func runErlang(c *model.Config) error {
	models := make([]Model, len(c.Erlang.Params))
	var err error

	for i := range models {
		if models[i], err = model.NewErlang(c.Erlang.Params[i].K, c.Erlang.Params[i].Lambda); err != nil {
			return err
		}
	}

	return compute(models, c.Step, c.Erlang.Name)
}

func runHyperExponential(c *model.Config) error {
	models := make([]Model, len(c.HyperExponential.Params))
	var err error

	for i := range models {
		if models[i], err = model.NewHyperExponential(c.HyperExponential.Params[i].Lambdas, c.HyperExponential.Params[i].Probabilities); err != nil {
			return err
		}
	}

	return compute(models, c.Step, c.HyperExponential.Name)
}

func compute(models []Model, step float64, name string) error {
	resultsDensity := make([]model.Result, len(models))
	resultsDistribution := make([]model.Result, len(models))

	right := models[0].RightBound()
	for _, m := range models {
		if b := m.RightBound(); b > right {
			right = b
		}
	}

	left := models[0].LeftBound()
	for _, m := range models {
		if b := m.LeftBound(); b < left {
			left = b
		}
	}

	for i, m := range models {
		resultsDensity[i] = m.ComputeDensity(step, left, right)
		resultsDistribution[i] = m.ComputeDistribution(step, left, right)
	}

	if err := model.Plot(filepath.Join("out", name, "density"), name+" density", resultsDensity); err != nil {
		return fmt.Errorf("plot density: %w", err)
	}

	if err := model.Plot(filepath.Join("out", name, "distribution"), name+" distribution", resultsDistribution); err != nil {
		return fmt.Errorf("plot distribution: %w", err)
	}

	return nil
}
