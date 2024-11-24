package advanced_widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type LabelCard struct {
	widget.BaseWidget
	Text *widget.Card
}

func NewLabelCard(title string) *LabelCard {
	if title == "" {
		title = " "
	}

	l := &LabelCard{
		Text: widget.NewCard("", title, nil),
	}
	l.ExtendBaseWidget(l)

	return l
}

func (l *LabelCard) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, nil, l.Text)
	return widget.NewSimpleRenderer(c)
}

func (l *LabelCard) SetText(text string) {
	l.Text.SetSubTitle(text)
}
