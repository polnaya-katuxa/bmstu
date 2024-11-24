package table

type Task interface {
	// Function argument, function
	Function(float64, float64) float64

	MinArg() float64
	MinFunction() float64

	Picard1(float64) float64
	Picard2(float64) float64
	Picard3(float64) float64
	Picard4(float64) float64

	Analytic(float64) float64
}
