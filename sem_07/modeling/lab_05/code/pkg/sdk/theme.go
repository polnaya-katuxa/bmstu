package sdk

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type glamourTheme struct {
	fyne.Theme
}

func (t *glamourTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	m := map[fyne.ThemeColorName]color.Color{
		theme.ColorNameBackground:      color.NRGBA{R: 255, G: 210, B: 170, A: 255},
		theme.ColorNameButton:          color.NRGBA{R: 255, G: 200, B: 160, A: 255},
		theme.ColorNameInputBackground: color.NRGBA{R: 255, G: 200, B: 160, A: 255},
		theme.ColorNameInputBorder:     color.NRGBA{R: 0, G: 0, B: 0, A: 10},
		theme.ColorNamePrimary:         color.NRGBA{R: 255, G: 160, B: 160, A: 255},
		theme.ColorNameMenuBackground:  color.NRGBA{R: 255, G: 160, B: 160, A: 255},
	}

	if c, ok := m[n]; ok {
		return c
	}

	return t.Theme.Color(n, v)
}

func (t *glamourTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}
