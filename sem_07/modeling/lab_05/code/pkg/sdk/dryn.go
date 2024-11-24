package sdk

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type drynThemeProvider struct{}

func (p drynThemeProvider) Theme() fyne.Theme {
	return &glamourTheme{Theme: theme.LightTheme()}
}

type drynProvider struct {
	drynThemeProvider

	errorLabel *widget.Label
}

func newDrynProvider() *drynProvider {
	return &drynProvider{
		errorLabel: widget.NewLabel(""),
	}
}

func (p *drynProvider) window() fyne.Window {
	return fyneWindow
}

func (p *drynProvider) describe(label string, wrapped fyne.CanvasObject) *fyne.Container {
	l := widget.NewLabel(label)
	l.Alignment = fyne.TextAlignLeading
	l.TextStyle.Bold = true

	return container.NewVBox(l, wrapped)
}

func (p *drynProvider) result() ResultWidget {
	l := widget.NewLabel("")
	l.TextStyle.Monospace = true
	l.Alignment = fyne.TextAlignCenter

	return l
}

func (p *drynProvider) layout(parameters []fyne.CanvasObject, results []fyne.CanvasObject, computeResults func()) *fyne.Container {
	parametersBox := container.NewVBox()
	for _, parameter := range parameters {
		parametersBox.Add(parameter)
	}

	parametersBox.Add(widget.NewButton("Получить результат", computeResults))
	parametersCard := widget.NewCard("", "", parametersBox)

	resultsBox := container.NewVBox()
	for _, result := range results {
		resultsBox.Add(result)
	}
	resultsCard := widget.NewCard("", "", resultsBox)

	return container.NewBorder(nil, p.errorLabel, nil, resultsCard, parametersCard)
}

func (p *drynProvider) showError(text string) {
	p.errorLabel.SetText(text)
}

func (p *drynProvider) HideError() {
	p.errorLabel.SetText("")
}

func (p *drynProvider) grid() *fyne.Container {
	return container.NewGridWithColumns(3)
}
