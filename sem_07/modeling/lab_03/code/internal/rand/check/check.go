package check

import (
	"github.com/samber/lo"
	"gonum.org/v1/gonum/stat"
	"lab_03/internal/rand"
	"math"
	"sort"
)

const (
	Uniform     = "Колмогорова-Смирнова"
	Correlation = "Сериальной корреляции"
)

// IsRandom проверяет, насколько соответствует набор случайных чисел равномерному
// распределению. Возвращает вероятность — вероятность того, что функция Колмогорова
// будет меньше той, которая посчитается на основе заданного набора чисел.
// Функция Колмогорова — максимальная разность значений фактической функции распределения
// и эмпирической функции распределения.
//
// - values — значения выборки.
// - start, end — индексы начала и конца возможного диапазона значений, не включая конец.
func IsRandom(values []int, r rand.Range, mode string) float64 {
	if mode == Uniform {
		return checkUniform(values, r)
	}

	return checkCorrelation(values, r)
}

func checkUniform(values []int, r rand.Range) float64 {
	data := lo.Map(values, func(v int, _ int) float64 {
		return float64(v)
	})
	sort.Float64s(data)

	return kolmogorovSmirnov(data, generateIdealUniform(r.Start, r.End))
}

func checkCorrelation(values []int, r rand.Range) float64 {
	size := 5
	correlations := make([]float64, 0, len(values)/size)

	data := lo.Map(values, func(v int, _ int) float64 {
		return float64(v)
	})

	for i := size; i+size <= len(values); i += size {
		correlations = append(correlations, stat.Correlation(data[i-size:i], data[i:i+size], nil))
	}

	result := math.Abs(stat.Mean(correlations, nil))
	if math.IsNaN(result) || math.IsInf(result, 1) || math.IsInf(result, -1) {
		return 0.0
	}

	return 1.0 - result
}

// generateIdealUniform создает идеальное равномерное распределение для заданного диапазона.
// Пример:
// generateIdealUniform(10, 15) = [10, 11, 12, 13, 14].
func generateIdealUniform(start, end int) []float64 {
	result := make([]float64, 0, end-start)
	for i := 0; i < end-start; i++ {
		result = append(result, float64(start+i))
	}

	return result
}

func kolmogorovSmirnov(data1, data2 []float64) float64 {
	ks := stat.KolmogorovSmirnov(data1, nil, data2, nil)

	n1, n2 := len(data1), len(data2)
	gamma := math.Sqrt(float64(n1*n2) / float64(n1+n2))
	lambda := (gamma + 0.12 + 0.11/gamma) * ks

	p, sign, k := 0.0, 1.0, 1.0
	for i := 0; i < 101; i++ {
		p += sign * math.Exp(-2*lambda*lambda*k*k)
		sign, k = -sign, k+1
	}
	p *= 2
	if p < 0 {
		p = 0
	} else if p > 1 {
		p = 1
	}

	return 1.0 - p
}
