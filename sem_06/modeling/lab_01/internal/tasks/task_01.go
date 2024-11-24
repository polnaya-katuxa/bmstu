package tasks

import "math"

type Task01 struct{}

func NewTask01() Task01 {
	return Task01{}
}

func (t Task01) Function(y, x float64) float64 {
	return x + math.Pow(y, 2.0)
}

func (t Task01) Picard1(y float64) float64 {
	return math.Pow(y, 3.0)/3.0 + y + 1.0
}

func (t Task01) Picard2(y float64) float64 {
	return math.Pow(y, 4.0)/12.0 + math.Pow(y, 3.0)/3.0 + math.Pow(y, 2.0)/2.0 + y + 1.0
}

func (t Task01) Picard3(y float64) float64 {
	return math.Pow(y, 5.0)/60.0 + math.Pow(y, 4.0)/12.0 + math.Pow(y, 3.0)/2.0 + math.Pow(y, 2.0)/2.0 + y + 1.0
}

func (t Task01) Picard4(y float64) float64 {
	return math.Pow(y, 6.0)/360.0 + math.Pow(y, 5.0)/60.0 + math.Pow(y, 4.0)/8.0 + math.Pow(y, 3.0)/2.0 + math.Pow(y, 2.0)/2.0 + y + 1.0
}

func (t Task01) Analytic(y float64) float64 {
	return 3*math.Exp(y) - math.Pow(y, 2.0) - 2.0*y - 2.0
}

func (t Task01) MinArg() float64 {
	return 0.0
}

func (t Task01) MinFunction() float64 {
	return 1.0
}
