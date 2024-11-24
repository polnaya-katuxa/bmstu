package advanced_widgets

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type NumericalEntry struct {
	widget.Entry
}

func NewNumericalEntry() *NumericalEntry {
	entry := &NumericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func NewIntEntry(v int) *NumericalEntry {
	entry := &NumericalEntry{}
	entry.ExtendBaseWidget(entry)
	entry.Validator = validateInt
	entry.SetText(fmt.Sprint(v))
	return entry
}

func NewFloatEntry(v float64) *NumericalEntry {
	entry := &NumericalEntry{}
	entry.ExtendBaseWidget(entry)
	entry.Validator = validateFloat
	entry.SetText(fmt.Sprint(v))
	return entry
}

func (e *NumericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' || r == '-' {
		e.Entry.TypedRune(r)
	}
}

func (e *NumericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *NumericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func (e *NumericalEntry) GetInt() (int, error) {
	v, err := strconv.Atoi(e.Text)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

func (e *NumericalEntry) GetFloat() (float64, error) {
	v, err := strconv.ParseFloat(e.Text, 64)
	if err != nil {
		return 0.0, err
	}

	return v, nil
}

func validateFloat(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	return nil
}

func validateInt(s string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	return nil
}
