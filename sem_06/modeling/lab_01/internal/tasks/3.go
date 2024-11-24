package tasks

import "math"

type Task03 struct{}

func NewTask03() Task03 {
	return Task03{}
}

func (t Task03) Function(x, y float64) float64 {
	return math.Pow(x, 2) + math.Pow(y, 2)
}

func (t Task03) Picard1(x float64) float64 {
	return math.Pow(x, 3.0) / 3.0
}

func (t Task03) Picard2(x float64) float64 {
	return math.Pow(x, 7.0)/63.0 + math.Pow(x, 3.0)/3.0
}

func (t Task03) Picard3(x float64) float64 {
	return math.Pow(x, 15.0)/59535.0 + 2.0*math.Pow(x, 11.0)/2079.0 + math.Pow(x, 7.0)/63.0 + math.Pow(x, 3.0)/3.0
}

func (t Task03) Picard4(x float64) float64 {
	return math.Pow(x, 31.0)/109876902975.0 + 4.0*math.Pow(x, 27.0)/3341878155.0 + 662.0*math.Pow(x, 23.0)/1043812015.0 + 82.0*math.Pow(x, 19.0)/37328445.0 + 13.0*math.Pow(x, 15.0)/218295.0 + 2.0*math.Pow(x, 11.0)/2079.0 + math.Pow(x, 7.0)/63.0 + math.Pow(x, 3.0)/3.0
}

func (t Task03) Analytic(_ float64) float64 {
	return math.NaN()
}

func (t Task03) MinArg() float64 {
	return 0.0
}

func (t Task03) MinFunction() float64 {
	return 0.0
}
