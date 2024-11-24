package app

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/charmbracelet/log"
	"lab_03/internal/rand"
	"lab_03/internal/rand/check"
	"strconv"
	"sync"
)

const (
	numbers  = 10
	variants = 3

	generateCount = 1000
)

var (
	errIntNotInRange = errors.New("value not in range from 0 to 9")
)

func validateInt(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if v < 0 || v > 9 {
		return errIntNotInRange
	}

	return nil
}

func (a *App) onValidationIntChanged(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Warn("changed validation state", "error", err)

	if err != nil {
		a.invalidEntries++
	} else {
		a.invalidEntries--
	}

	if a.invalidEntries == 0 {
		a.checkButton.Enable()
	} else {
		a.checkButton.Disable()
	}
}

type App struct {
	app    fyne.App
	window fyne.Window

	tabularLabels     [variants][10]*widget.Label
	tabularChecks     [variants]*widget.Label
	tabularGenerators [variants]*rand.Tabular

	computationalLabels     [variants][10]*widget.Label
	computationalChecks     [variants]*widget.Label
	computationalGenerators [variants]*rand.Computational

	generateButton *widget.Button

	checkEntries [10]*widget.Entry
	checkButton  *widget.Button
	check        *widget.Label

	mu             sync.Mutex
	invalidEntries int

	mode string
}

func New() *App {
	log.SetLevel(log.DebugLevel)

	var a App

	a.app = app.New()
	//a.app.Settings().SetTheme(theme.LightTheme())
	a.window = a.app.NewWindow("lab_03")
	a.window.Resize(fyne.NewSize(700, 400))

	for i := 0; i < numbers; i++ {
		for j := 0; j < variants; j++ {
			a.tabularLabels[j][i] = newLabelRight("")
		}
	}
	for i := 0; i < variants; i++ {
		a.tabularGenerators[i] = rand.NewTabular(i + 1)
		a.tabularChecks[i] = newLabelCenter("")
	}

	for i := 0; i < numbers; i++ {
		for j := 0; j < variants; j++ {
			a.computationalLabels[j][i] = newLabelRight("")
		}
	}
	for i := 0; i < variants; i++ {
		a.computationalGenerators[i] = rand.NewComputational(i+1, 124342, 22, 544, 3452, 435365)
		a.computationalChecks[i] = newLabelCenter("")
	}

	for i := 0; i < numbers; i++ {
		a.checkEntries[i] = widget.NewEntry()
		a.checkEntries[i].Text = "0"
		a.checkEntries[i].Validator = validateInt
		a.checkEntries[i].SetOnValidationChanged(a.onValidationIntChanged)
	}

	a.check = newLabelCenter("")

	a.checkButton = widget.NewButton("Проверить", a.checkButtonPressed)

	a.generateButton = widget.NewButton("Сгенерировать", a.generateButtonPressed)

	a.mode = check.Uniform

	return &a
}

func (a *App) setContent() {
	log.Info("setting content")

	tabularTable := container.NewGridWithColumns(variants)
	tabularTable.Add(newLabelHeader("1"))
	tabularTable.Add(newLabelHeader("2"))
	tabularTable.Add(newLabelHeader("3"))
	for i := 0; i < numbers; i++ {
		for j := 0; j < variants; j++ {
			tabularTable.Add(widget.NewCard("", "", a.tabularLabels[j][i]))
		}
	}
	for i := 0; i < variants; i++ {
		tabularTable.Add(a.tabularChecks[i])
	}
	tabularTableWithName := container.NewVBox(newLabelHeader("Табличный способ"), tabularTable)

	computationalTable := container.NewGridWithColumns(variants)
	computationalTable.Add(newLabelHeader("1"))
	computationalTable.Add(newLabelHeader("2"))
	computationalTable.Add(newLabelHeader("3"))
	for i := 0; i < numbers; i++ {
		for j := 0; j < variants; j++ {
			computationalTable.Add(widget.NewCard("", "", a.computationalLabels[j][i]))
		}
	}
	for i := 0; i < variants; i++ {
		computationalTable.Add(a.computationalChecks[i])
	}
	computationalTableWithName := container.NewVBox(newLabelHeader("Алгоритмический способ"), computationalTable)

	tables := container.NewGridWithColumns(2, tabularTableWithName, computationalTableWithName)
	body := container.NewVBox(tables, a.generateButton)

	checkInputGrid := container.NewGridWithColumns(1)
	for i := 0; i < numbers; i++ {
		checkInputGrid.Add(a.checkEntries[i])
	}

	checkBox := container.NewVBox(newLabelHeader("Критерий"), checkInputGrid, a.checkButton, a.check)

	configBox := container.NewVBox(newLabelHeader("Критерий"), widget.NewSelect([]string{check.Uniform, check.Correlation}, func(s string) {
		a.mode = s
	}))

	a.window.SetContent(container.NewBorder(nil, nil, configBox, checkBox, body))
}

func (a *App) Run() {
	a.setContent()
	log.Info("running application")
	a.window.ShowAndRun()
	log.Info("exiting")
}

func (a *App) checkButtonPressed() {
	log.Debug("check button pressed")

	sequence := make([]int, numbers)
	for i := range sequence {
		sequence[i], _ = strconv.Atoi(a.checkEntries[i].Text)
	}

	p := check.IsRandom(sequence, rand.Ranges[1], a.mode)
	a.check.SetText(formatFloat(p))

	log.Debug("check button pressed handler finished", "p", p)
}

func (a *App) generateButtonPressed() {
	log.Debug("generate button pressed")

	var tabularResults [variants][]int
	for i := 0; i < variants; i++ {
		for j := 0; j < generateCount; j++ {
			tabularResults[i] = append(tabularResults[i], a.tabularGenerators[i].Rand())
		}
	}
	for i := 0; i < variants; i++ {
		for j := 0; j < numbers; j++ {
			a.tabularLabels[i][j].SetText(formatInt(tabularResults[i][j]))
		}
		a.tabularChecks[i].SetText(formatFloat(check.IsRandom(tabularResults[i], rand.Ranges[i+1], a.mode)))
	}

	var computationalResults [variants][]int
	for i := 0; i < variants; i++ {
		for j := 0; j < generateCount; j++ {
			computationalResults[i] = append(computationalResults[i], a.computationalGenerators[i].Rand())
		}
	}
	for i := 0; i < variants; i++ {
		for j := 0; j < numbers; j++ {
			a.computationalLabels[i][j].SetText(formatInt(computationalResults[i][j]))
		}
		a.computationalChecks[i].SetText(formatFloat(check.IsRandom(computationalResults[i], rand.Ranges[i+1], a.mode)))
	}

	log.Debug("generate button pressed handler finished")
}

func newLabelRight(text string) *widget.Label {
	result := widget.NewLabel(text)
	result.Alignment = fyne.TextAlignTrailing
	return result
}

func newLabelCenter(text string) *widget.Label {
	result := widget.NewLabel(text)
	result.Alignment = fyne.TextAlignCenter
	return result
}

func newLabelHeader(text string) *widget.Label {
	result := widget.NewLabel(text)
	result.Alignment = fyne.TextAlignCenter
	result.TextStyle.Bold = true
	return result
}

func formatFloat(n float64) string {
	return fmt.Sprintf("%.4f", n)
}

func formatInt(n int) string {
	return fmt.Sprintf("%d", n)
}
