package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type uniformParams struct {
	A float64
	B float64
}

type uniformConfig struct {
	Name   string
	Params []uniformParams
}

type erlangParams struct {
	K      int
	Lambda float64
}

type erlangConfig struct {
	Name   string
	Params []erlangParams
}

type hyperExponentialParams struct {
	Lambdas       []float64
	Probabilities []float64
}

type hyperExponentialConfig struct {
	Name   string
	Params []hyperExponentialParams
}

type Config struct {
	Step             float64
	Uniform          uniformConfig
	Erlang           erlangConfig
	HyperExponential hyperExponentialConfig
}

func FromFile(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	var c Config
	if err := json.NewDecoder(file).Decode(&c); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}

	return &c, nil
}
