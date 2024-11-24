package sdk

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"lab_04/pkg/sdk/advanced_widgets"
)

type blumbaThemeProvider struct{}

func (p blumbaThemeProvider) Theme() fyne.Theme {
	return theme.DefaultTheme()
}

type blumbaProvider struct {
	blumbaThemeProvider
}

func newBlumbaProvider() *blumbaProvider {
	return &blumbaProvider{}
}

func (p *blumbaProvider) window() fyne.Window {
	return fyneWindow
}

func (p *blumbaProvider) describe(label string, wrapped fyne.Widget) *fyne.Container {
	l := widget.NewLabel(label)
	l.Alignment = fyne.TextAlignLeading
	l.TextStyle.Bold = false

	return container.NewGridWithColumns(2, l, wrapped)
}

func (p *blumbaProvider) result() ResultWidget {
	return advanced_widgets.NewLabelCard("")
}

func (p *blumbaProvider) layout(parameters []fyne.CanvasObject, results []fyne.CanvasObject, computeResults func()) *fyne.Container {
	box := container.NewVBox()
	for _, parameter := range parameters {
		box.Add(parameter)
	}

	box.Add(widget.NewButton("Вычислить", computeResults))
	box.Add(widget.NewSeparator())

	for _, result := range results {
		box.Add(result)
	}

	return box
}

func (p *blumbaProvider) showError(text string) {
	dialog.ShowError(errors.New(text), fyneWindow)
}

func (p *blumbaProvider) HideError() {

}
