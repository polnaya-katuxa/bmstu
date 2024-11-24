package random

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/hnakamur/randutil"
	"lab_04/pkg/sdk"
	"lab_04/pkg/sdk/advanced_widgets"
)

type Erlang struct {
	k      *advanced_widgets.NumericalEntry
	lambda *advanced_widgets.NumericalEntry
}

func NewErlang() *Erlang {
	return &Erlang{
		k:      advanced_widgets.NewIntEntry(9),
		lambda: advanced_widgets.NewFloatEntry(0.3),
	}
}

func (e *Erlang) Widgets() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewCard(
			"",
			"Параметры распределения Эрланга",
			container.NewVBox(
				sdk.Describe("λ", e.lambda),
				sdk.Describe("k", e.k),
			),
		),
	}
}

func (e *Erlang) Rand() (float64, error) {
	lambda, err := e.lambda.GetFloat()
	if err != nil {
		return 0.0, &sdk.Error{
			Human: "Неверный параметр: λ.",
			Raw:   err,
		}
	}

	k, err := e.k.GetInt()
	if err != nil {
		return 0.0, &sdk.Error{
			Human: "Неверный параметр: k.",
			Raw:   err,
		}
	}

	return randutil.Erlang(randutil.NewCryptoIntner(), lambda, k)
}
