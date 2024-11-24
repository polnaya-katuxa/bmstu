package app

import (
	"flag"
	"fmt"
	"lab_03/internal/model"
)

func Run() error {
	configName := flag.String("config", "config/config.json", "path to config file")
	plotName := flag.String("plot", "plot.png", "path to output plot")
	flag.Parse()

	m, err := model.FromFile(*configName)
	if err != nil {
		return fmt.Errorf("open model: %w", err)
	}

	if err := m.Compute().Print().SavePlot(*plotName).Error(); err != nil {
		return fmt.Errorf("output: %w", err)
	}

	return nil
}
