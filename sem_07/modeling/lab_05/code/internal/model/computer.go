package model

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"lab_05/pkg/sdk"
	"lab_05/pkg/sdk/advanced_widgets"
)

type computer struct {
	name  string
	time  float64
	queue int
	busy  bool
}

func (c *computer) planEventComputerProcessed(e event) eventComputerProcessed {
	return eventComputerProcessed{
		baseEvent: baseEvent{ts: e.timestamp() + c.time},
		computer:  c,
	}
}

type ComputerCreator struct {
	name string
	base *advanced_widgets.NumericalEntry
}

func NewComputerCreator(name string, base int) *ComputerCreator {
	return &ComputerCreator{
		name: name,
		base: advanced_widgets.NewIntEntry(base),
	}
}

func (c *ComputerCreator) Widget() fyne.CanvasObject {
	return widget.NewCard(
		"",
		c.name,
		sdk.Describe("Время", c.base),
	)
}

func (c *ComputerCreator) create() (*computer, error) {
	base, err := c.base.GetInt()
	if err != nil {
		return nil, &sdk.Error{
			Human: "Неверное значение времени.",
			Raw:   err,
		}
	}

	return &computer{
		name:  c.name,
		time:  float64(base),
		queue: 0,
		busy:  false,
	}, nil
}
