package model

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lab_05/pkg/sdk"
	"lab_05/pkg/sdk/advanced_widgets"
	"math/rand"
)

type operator struct {
	name        string
	averageTime float64
	delta       float64
	busy        bool
	toComputer  *computer
}

func (o *operator) planEventOperatorProcessed(e event) eventOperatorProcessed {
	a := o.averageTime - o.delta
	b := o.averageTime + o.delta

	period := rand.Float64()*(b-a) + a

	return eventOperatorProcessed{
		baseEvent:  baseEvent{ts: e.timestamp() + period},
		operator:   o,
		toComputer: o.toComputer,
	}
}

type OperatorCreator struct {
	name  string
	base  *advanced_widgets.NumericalEntry
	delta *advanced_widgets.NumericalEntry
}

func NewOperatorCreator(name string, base, delta int) *OperatorCreator {
	return &OperatorCreator{
		name:  name,
		base:  advanced_widgets.NewIntEntry(base),
		delta: advanced_widgets.NewIntEntry(delta),
	}
}

func (c *OperatorCreator) Widget() fyne.CanvasObject {
	l := widget.NewLabel("±")
	l.Alignment = fyne.TextAlignCenter

	line := container.NewGridWithColumns(3, c.base, l, c.delta)

	return widget.NewCard(
		"",
		c.name,
		sdk.Describe("Интервал", line),
	)
}

func (c *OperatorCreator) create() (*operator, error) {
	base, err := c.base.GetInt()
	if err != nil {
		return nil, &sdk.Error{
			Human: "Неверное значение времени.",
			Raw:   err,
		}
	}

	delta, err := c.delta.GetInt()
	if err != nil {
		return nil, &sdk.Error{
			Human: "Неверный значение отклонения.",
			Raw:   err,
		}
	}

	return &operator{
		name:        c.name,
		averageTime: float64(base),
		delta:       float64(delta),
		busy:        false,
	}, nil
}
