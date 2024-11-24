package sdk

import (
	"fyne.io/fyne/v2"
)

var defaultProvider Provider

type themeProvider interface {
	Theme() fyne.Theme
}

type ResultWidget interface {
	fyne.Widget

	SetText(text string)
}

type Provider interface {
	themeProvider

	window() fyne.Window

	showError(text string)

	describe(label string, wrapped fyne.CanvasObject) *fyne.Container
	result() ResultWidget
	layout(parameters []fyne.CanvasObject, results []fyne.CanvasObject, computeResults func()) *fyne.Container

	grid() *fyne.Container
}
