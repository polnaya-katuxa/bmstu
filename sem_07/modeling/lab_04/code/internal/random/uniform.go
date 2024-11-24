package random

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lab_04/pkg/sdk"
	"lab_04/pkg/sdk/advanced_widgets"
	"math/rand"
)

type Uniform struct {
	a *advanced_widgets.NumericalEntry
	b *advanced_widgets.NumericalEntry
}

func NewUniform() *Uniform {
	return &Uniform{
		a: advanced_widgets.NewFloatEntry(2),
		b: advanced_widgets.NewFloatEntry(10),
	}
}

func (u *Uniform) Widgets() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewCard(
			"",
			"Параметры равномерного распределения",
			container.NewVBox(
				sdk.Describe("a", u.a),
				sdk.Describe("b", u.b),
			),
		),
	}
}

func (u *Uniform) Rand() (float64, error) {
	a, err := u.a.GetFloat()
	if err != nil {
		return 0.0, &sdk.Error{
			Human: "Неверный параметр: a.",
			Raw:   err,
		}
	}

	b, err := u.b.GetFloat()
	if err != nil {
		return 0.0, &sdk.Error{
			Human: "Неверный параметр: b.",
			Raw:   err,
		}
	}

	return rand.Float64()*(b-a) + a, nil
}
