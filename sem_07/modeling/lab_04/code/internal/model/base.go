package model

import "fyne.io/fyne/v2"

type BaseModel struct {
	generatorTimer rand
	processorTimer rand
}

func NewBaseModel(g, p rand) *BaseModel {
	return &BaseModel{
		generatorTimer: g,
		processorTimer: p,
	}
}

func (m *BaseModel) Widgets() []fyne.CanvasObject {
	result := m.generatorTimer.Widgets()
	return append(result, m.processorTimer.Widgets()...)
}

type rand interface {
	Rand() (float64, error)
	Widgets() []fyne.CanvasObject
}
