package app

import (
	"flag"
	"fmt"
	"lab_05/internal/model"
)

func Run() error {
	configName := flag.String("config", "config/config.json", "path to config file")
	flag.Parse()

	m, err := model.FromFile(*configName)
	if err != nil {
		return fmt.Errorf("open model: %w", err)
	}

	if err := m.Compute().ToFile("out/data.txt").Python().Error(); err != nil {
		return fmt.Errorf("output: %w", err)
	}

	return nil
}
