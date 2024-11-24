package app

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lab_02/internal/model"
	"log"
	"strconv"
	"sync"
)

var (
	ErrInvalidFloatInRussian = errors.New("Неверное число")
)

func validateMatrixValue(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return ErrInvalidFloatInRussian
	}

	return nil
}

func (a *App) onValidationMatrixValueChanged(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Println("changed validation state", err)

	if err != nil {
		a.invalidEntries++
	} else {
		a.invalidEntries--
	}

	if a.invalidEntries == 0 {
		a.computeButton.Enable()
	} else {
		a.computeButton.Disable()
	}
}

type App struct {
	app    fyne.App
	window fyne.Window

	matrix  *fyne.Container
	entries []*widget.Entry

	incrementButton *widget.Button
	decrementButton *widget.Button
	n               int
	nLabel          *widget.Label

	errorLabel *widget.Label

	computeButton *widget.Button

	mu             sync.Mutex
	invalidEntries int
}

func (a *App) resizeMatrix(n int) {
	if n < 1 {
		return
	}
	a.n = n
	a.entries = make([]*widget.Entry, 0, a.n*a.n)
	a.nLabel.SetText(fmt.Sprint(a.n))

	a.matrix.RemoveAll()
	a.matrix = container.NewGridWithColumns(a.n)

	for i := 0; i < a.n*a.n; i++ {
		e := widget.NewEntry()
		e.SetText("0.0")
		e.Validator = validateMatrixValue
		e.SetOnValidationChanged(a.onValidationMatrixValueChanged)

		a.matrix.Add(e)
		a.entries = append(a.entries, e)
	}
}

func New() *App {
	var a App

	a.app = app.New()
	a.window = a.app.NewWindow("lab_02")
	a.window.SetFixedSize(true)
	a.window.Resize(fyne.NewSize(1100, 700))

	a.matrix = container.NewGridWithColumns(0)
	a.nLabel = widget.NewLabel("0")
	a.n = 0

	a.nLabel.Alignment = fyne.TextAlignCenter
	a.incrementButton = widget.NewButton("+", func() {
		log.Println("incrementing number of states")
		a.resizeMatrix(a.n + 1)
		a.show()
	})
	a.decrementButton = widget.NewButton("-", func() {
		log.Println("decrementing number of states")
		a.resizeMatrix(a.n - 1)
		a.show()
	})
	a.resizeMatrix(3)

	a.computeButton = widget.NewButton("Рассчитать", func() {
		log.Println("start computing")
		a.compute()
	})

	a.errorLabel = widget.NewLabel("")
	a.errorLabel.Hidden = true
	a.errorLabel.TextStyle.Bold = true

	return &a
}

func (a *App) setContent() {
	leftMenu := container.NewVBox(
		widget.NewLabel("Размерность матрицы"),
		container.NewGridWithColumns(3, a.decrementButton, a.nLabel, a.incrementButton),
		widget.NewSeparator(),
		a.computeButton,
		a.errorLabel,
	)

	body := container.NewGridWithColumns(1, a.matrix)
	a.window.SetContent(container.NewBorder(nil, nil, leftMenu, nil, body))
}

type minSizedCard struct {
	*widget.Card
	width, height float32
}

func (m minSizedCard) MinSize() fyne.Size {
	return fyne.NewSize(m.width, m.height)
}

func squareLabel(text string) fyne.CanvasObject {
	return minSizedCard{
		Card:   widget.NewCard("", "", widget.NewLabel(text)),
		width:  120,
		height: 60,
	}
}

func (a *App) compute() {
	a.errorLabel.Hide()

	elements := make([]float64, a.n*a.n)
	for i, e := range a.entries {
		elements[i], _ = strconv.ParseFloat(e.Text, 64)
	}

	m := model.New(elements, a.n)
	probabilities, err := m.Probabilities()
	if err != nil {
		a.errorLabel.SetText("Не удается\nрешить систему.")
		a.errorLabel.Show()
		log.Println(err)
		return
	}

	times, err := m.Times()
	if err != nil {
		a.errorLabel.SetText("Не удается\nрешить систему.")
		a.errorLabel.Show()
		log.Println(err)
		return
	}

	window := a.app.NewWindow("Результат")
	window.SetFixedSize(false)

	grid := container.NewGridWithColumns(a.n + 1)

	grid.Add(squareLabel("P"))
	for i := 0; i < a.n; i++ {
		grid.Add(squareLabel(fmt.Sprintf("%.4f", probabilities[i])))
	}

	grid.Add(squareLabel("t"))
	for i := 0; i < a.n; i++ {
		grid.Add(squareLabel(fmt.Sprintf("%.4f", times[i])))
	}

	window.SetContent(grid)
	window.Show()
}

func (a *App) Run() {
	a.setContent()
	log.Println("running application")
	a.window.ShowAndRun()
}

func (a *App) show() {
	a.setContent()
	a.window.Show()
}
