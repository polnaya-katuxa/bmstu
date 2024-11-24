package app

import (
	"flag"
	"fmt"
	"lab_04/internal/model"
)

func Run() error {
	configName := flag.String("config", "config/config.json", "path to config file")
	flag.Parse()

	m, err := model.FromFile(*configName)
	if err != nil {
		return fmt.Errorf("open model: %w", err)
	}

	if err := m.Compute().PlotByTime("out/T(t).pdf").PlotByZ("out/T(z).pdf").PlotByTimeMinRadius("out/T(0,t).pdf").Error(); err != nil {
		return fmt.Errorf("output: %w", err)
	}

	return nil
}
