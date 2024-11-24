package tasks

import "math"

type Task02 struct{}

func NewTask02() Task02 {
	return Task02{}
}

func (t Task02) Function(y, x float64) float64 {
	return math.Pow(y, 3.0) + 2.0*x*y
}

func (t Task02) Picard1(y float64) float64 {
	return math.Pow(y, 4.0)/4.0 + math.Pow(y, 2.0)/2.0 + 1.0/2.0
}

func (t Task02) Picard2(y float64) float64 {
	return math.Pow(y, 6.0)/12.0 + math.Pow(y, 4.0)/2.0 + math.Pow(y, 2.0)/2.0 + 1.0/2.0
}

func (t Task02) Picard3(y float64) float64 {
	return math.Pow(y, 8.0)/48.0 + math.Pow(y, 6.0)/6.0 + math.Pow(y, 4.0)/2.0 + math.Pow(y, 2.0)/2.0 + 1.0/2.0
}

func (t Task02) Picard4(y float64) float64 {
	return math.Pow(y, 10.0)/240.0 + math.Pow(y, 8.0)/24.0 + math.Pow(y, 6.0)/6.0 + math.Pow(y, 4.0)/2.0 + math.Pow(y, 2.0)/2.0 + 1.0/2.0
}

func (t Task02) Analytic(y float64) float64 {
	return math.Exp(math.Pow(y, 2.0)) - math.Pow(y, 2.0)/2.0 - 1.0/2.0
}

func (t Task02) MinArg() float64 {
	return 0.0
}

func (t Task02) MinFunction() float64 {
	return 0.5
}
