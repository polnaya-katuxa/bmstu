package model

import (
	"github.com/samber/lo"
	"math"
)

// newF создает функцию, описывающую пучок, направленный в точку (x0, z0).
func (model *Model) newF(f0, x0, z0 float64) func(x, z float64) float64 {
	return func(x, z float64) float64 {
		return f0 * math.Pow(math.E, -model.Beta*math.Pow(x-x0, 2)*math.Pow(z-z0, 2))

		// Сердцечко
		//x -= x0
		//z -= z0
		//
		//x /= 3
		//z /= 3
		//
		//if math.Pow(x*x+z*z-1, 3)-x*x*z*z*z < 0 {
		//	return f0
		//}
		//
		//return 0

		//x -= x0
		//z -= z0
		//
		//x /= 2
		//z /= 2
		//
		//r := math.Sqrt(x*x + z*z)
		//t := math.Atan(z / x)
		//
		//if r < (1+math.Sin(t))*(1+0.9*math.Cos(8*t))*(1+0.1*math.Cos(24*t)) {
		//	return f0
		//}
		//
		//return 0
	}
}

// newFAll создает функцию, описывающую все пучки в совокупности. Настраивается из конфига.
func (model *Model) newFAll() func(x, z float64) float64 {
	fs := lo.Map(model.Lasers, func(e Laser, _ int) func(x, z float64) float64 {
		return model.newF(e.F0, e.X, e.Z)
	})

	return func(x, z float64) float64 {
		res := 0.0
		for _, f := range fs {
			res += f(x, z)
		}

		return res
	}
}

func (model *Model) Compute() *Result {
	x := make([]float64, model.Steps.X)
	for i := range x {
		x[i] = float64(i) * model.step.X
	}

	z := make([]float64, model.Steps.Z)
	for i := range z {
		z[i] = float64(i) * model.step.Z
	}

	u := make([][]float64, model.Steps.X)
	for i := range u {
		u[i] = make([]float64, model.Steps.Z)
	}
	for n := range u {
		for m := range u[n] {
			u[n][m] = model.U0
		}
	}

	delta := math.MaxFloat64
	steps := 0

	for delta >= model.Eps {
		steps++

		u1 := make([][]float64, model.Steps.X)
		for i := range u1 {
			u1[i] = make([]float64, model.Steps.Z)
		}

		for n := 0; n < model.Steps.X; n++ {
			//u1[n][0] = u[n][0]
			//u1[n][model.Steps.Z-1] = u[n][model.Steps.Z-1]
			u1[n][0] = model.U0
			u1[n][model.Steps.Z-1] = model.U0
		}

		for m := 0; m < model.Steps.Z-1; m++ {
			system := NewSystem(model.Steps.X - 2)

			i := 0
			for n := 1; n < model.Steps.X-1; n++ {
				v := u[n][0]
				if m > 0 {
					v = u[n][m-1]
				}

				a := (u[n][m+1]+v-2*u[n][m])/math.Pow(model.step.Z, 2) + model.f(x[n], z[m])/model.Lambda

				tau := model.Tau / 2
				system[i].A = tau / model.C
				system[i].B = 2*tau/model.C + math.Pow(model.step.X, 2)
				system[i].C = tau / model.C
				system[i].D = a*math.Pow(model.step.X, 2)*tau/model.C + math.Pow(model.step.X, 2)*u[n][m]
				i++
			}

			start := model.leftBoundaryConditionsFixedM(u, m)
			end := model.rightBoundaryConditionsFixedM(u, m)

			res := system.runThrough(start, end)
			for n := range res {
				u1[n][m] = res[n]
			}
		}

		u2 := make([][]float64, model.Steps.X)
		for i := range u2 {
			u2[i] = make([]float64, model.Steps.Z)
		}

		for m := 0; m < model.Steps.Z; m++ {
			//u2[0][m] = u1[0][m]
			//u2[model.Steps.X-1][m] = u1[model.Steps.X-1][m]
			u2[0][m] = model.U0
			u2[model.Steps.X-1][m] = model.U0
		}

		for n := 0; n < model.Steps.X-1; n++ {
			system := NewSystem(model.Steps.Z - 2)

			i := 0
			for m := 1; m < model.Steps.Z-1; m++ {
				v := u1[0][m]
				if n > 0 {
					v = u1[n-1][m]
				}

				a := (u1[n+1][m]+v-2*u1[n][m])/math.Pow(model.step.X, 2) + model.f(x[n], z[m])/model.Lambda

				tau := model.Tau / 2
				system[i].A = tau / model.C
				system[i].B = 2*tau/model.C + math.Pow(model.step.Z, 2)
				system[i].C = tau / model.C
				system[i].D = a*math.Pow(model.step.Z, 2)*tau/model.C + math.Pow(model.step.Z, 2)*u1[n][m]
				i++
			}

			start := model.leftBoundaryConditionsFixedN(u1, n)
			end := model.rightBoundaryConditionsFixedN(u1, n)

			res := system.runThrough(start, end)
			for m := range res {
				u2[n][m] = res[m]
			}
		}

		delta = 0.0
		for n := range u2 {
			for m := range u2[n] {
				// Абсолютная погрешность.
				curDelta := math.Abs((u[n][m] - u2[n][m]) / u2[n][m])
				if curDelta > delta {
					delta = curDelta
				}
			}
		}

		u = u2
	}

	return &Result{
		X: x,
		Z: z,
		U: u,
	}
}

// leftBoundaryConditionsFixedM расчет коэффициентов левого краевого условия вида M0 * y0 + K0 * y1 = P0.
func (model *Model) leftBoundaryConditionsFixedM(u [][]float64, m int) BoundaryConditions {
	switch model.BoundaryVariant {
	case 1:
		panic("boundary conditions not implemented")
	case 2:
		return BoundaryConditions{
			M: -model.K.Predict(u[0][m]),
			K: model.K.Predict(u[0][m]),
			P: model.Alpha1 * model.step.X * (u[0][m] - model.U0),
		}
	case 3:
		//return BoundaryConditions{
		//	M: 1,
		//	K: 0,
		//	P: model.U0,
		//}

		return BoundaryConditions{
			M: model.Lambda,
			K: -model.Lambda,
			P: model.S * model.step.X,
		}
	default:
		panic("boundary conditions not implemented")
	}
}

// rightBoundaryConditionsFixedM расчет коэффициентов правого краевого условия вида MN * y{N-1} + KN * yN = PN.
func (model *Model) rightBoundaryConditionsFixedM(u [][]float64, m int) BoundaryConditions {
	switch model.BoundaryVariant {
	case 1:
		panic("boundary conditions not implemented")
	case 2:
		return BoundaryConditions{
			M: model.K.Predict(u[len(u)-1][m]),
			K: -model.K.Predict(u[len(u)-1][m]),
			P: model.Alpha2 * model.step.X * (u[len(u)-1][m] - model.U0),
		}
	case 3:
		return BoundaryConditions{
			M: 0,
			K: 1,
			P: model.U0,
		}
	default:
		panic("boundary conditions not implemented")
	}
}

// leftBoundaryConditionsFixedN расчет коэффициентов левого краевого условия вида M0 * y0 + K0 * y1 = P0.
func (model *Model) leftBoundaryConditionsFixedN(u [][]float64, n int) BoundaryConditions {
	switch model.BoundaryVariant {
	case 1:
		panic("boundary conditions not implemented")
	case 2:
		return BoundaryConditions{
			M: -model.K.Predict(u[n][0]),
			K: model.K.Predict(u[n][0]),
			P: model.Alpha3 * model.step.Z * (u[n][0] - model.U0),
		}
	case 3:
		return BoundaryConditions{
			M: 1,
			K: 0,
			P: model.U0,
		}

		return BoundaryConditions{
			M: 1,
			K: -1,
			P: 0,
		}
	default:
		panic("boundary conditions not implemented")
	}
}

// rightBoundaryConditionsFixedN расчет коэффициентов правого краевого условия вида MN * y{N-1} + KN * yN = PN.
func (model *Model) rightBoundaryConditionsFixedN(u [][]float64, n int) BoundaryConditions {
	switch model.BoundaryVariant {
	case 1:
		panic("boundary conditions not implemented")
	case 2:
		return BoundaryConditions{
			M: model.K.Predict(u[n][len(u[n])-1]),
			K: -model.K.Predict(u[n][len(u[n])-1]),
			P: model.Alpha4 * model.step.Z * (u[n][len(u[n])-1] - model.U0),
		}
	case 3:
		return BoundaryConditions{
			M: 0,
			K: 1,
			P: model.U0,
		}

		return BoundaryConditions{
			M: 1,
			K: -1,
			P: 0,
		}
	default:
		panic("boundary conditions not implemented")
	}
}
