package interpolation

import (
	"errors"
	"fmt"
	"os"
)

var ErrNotEnoughPairs = errors.New("not enough pair for interpolation")

type Pair struct {
	X, Y float64
}

type Interpolation struct {
	table []Pair
}

func New(table []Pair) (*Interpolation, error) {
	if len(table) < 2 {
		return nil, fmt.Errorf("new interpolation: %w", ErrNotEnoughPairs)
	}

	return &Interpolation{
		table: table,
	}, nil
}

func FromFile(filename string) (*Interpolation, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open file: %w")
	}

	var n int
	if _, err := fmt.Fscan(file, &n); err != nil {
		return nil, fmt.Errorf("read n: %w", err)
	}

	pairs := make([]Pair, n)
	for i := range pairs {
		if _, err := fmt.Fscan(file, &pairs[i].X, &pairs[i].Y); err != nil {
			return nil, fmt.Errorf("read x and y: %w", err)
		}
	}

	ip, err := New(pairs)
	if err != nil {
		return nil, fmt.Errorf("create interpolation from pairs: %w")
	}

	return ip, nil
}

func (ip *Interpolation) Get(x float64) float64 {
	if x <= ip.table[0].X {
		return ip.table[0].Y
	}

	if x >= ip.table[len(ip.table)-1].X {
		return ip.table[len(ip.table)-1].Y
	}

	for i := 1; i < len(ip.table); i++ {
		if x >= ip.table[i-1].X && x <= ip.table[i].X {
			x0, y0 := ip.table[i-1].X, ip.table[i-1].Y
			x1, y1 := ip.table[i].X, ip.table[i].Y

			return y0 + ((y1-y0)/(x1-x0))*(x-x0)
		}
	}

	return 0
}
