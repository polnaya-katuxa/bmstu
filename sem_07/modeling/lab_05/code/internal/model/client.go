package model

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lab_05/pkg/sdk"
	"lab_05/pkg/sdk/advanced_widgets"
	"math/rand"
)

type client struct {
	name        string
	averageTime float64
	delta       float64
}

func (c *client) planEventClientArrived(e event) eventClientArrived {
	a := c.averageTime - c.delta
	b := c.averageTime + c.delta

	period := rand.Float64()*(b-a) + a

	return eventClientArrived{
		baseEvent{ts: e.timestamp() + period},
	}
}

type ClientCreator struct {
	name  string
	base  *advanced_widgets.NumericalEntry
	delta *advanced_widgets.NumericalEntry
}

func NewClientCreator(name string, base, delta int) *ClientCreator {
	return &ClientCreator{
		name:  name,
		base:  advanced_widgets.NewIntEntry(base),
		delta: advanced_widgets.NewIntEntry(delta),
	}
}

func (c *ClientCreator) Widget() fyne.CanvasObject {
	l := widget.NewLabel("±")
	l.Alignment = fyne.TextAlignCenter

	line := container.NewGridWithColumns(3, c.base, l, c.delta)

	return widget.NewCard(
		"",
		c.name,
		sdk.Describe("Интервал", line),
	)
}

func (c *ClientCreator) create() (*client, error) {
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

	return &client{
		name:        c.name,
		averageTime: float64(base),
		delta:       float64(delta),
	}, nil
}
