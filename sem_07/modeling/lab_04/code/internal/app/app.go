package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/charmbracelet/log"
	"lab_04/internal/model"
	"lab_04/internal/random"
	"lab_04/pkg/sdk"
	"lab_04/pkg/sdk/advanced_widgets"
	"strconv"
)

const (
	limitProcessed = 1000
)

type App struct {
	widgets widgets

	baseModel *model.BaseModel
	models    map[string]methodModel
}

type widgets struct {
	returnPercentEntry *advanced_widgets.NumericalEntry
	method             *widget.Select

	queueLen sdk.ResultWidget
}

func New() *App {
	var a App

	a.baseModel = model.NewBaseModel(random.NewUniform(), random.NewErlang())
	a.models = map[string]methodModel{
		"пошаговый":  model.NewTimeModel(a.baseModel),
		"событийный": model.NewEventModel(a.baseModel),
	}

	a.createWidgets()

	return &a
}

func (a *App) createWidgets() {
	a.widgets.returnPercentEntry = advanced_widgets.NewIntEntry(10)

	keys := make([]string, 0, len(a.models))
	for k := range a.models {
		keys = append(keys, k)
	}
	a.widgets.method = widget.NewSelect(keys, func(string) {})
	a.widgets.method.SetSelectedIndex(0)

	a.widgets.queueLen = sdk.Result()

	sdk.SetFillWindowFunc(a.setContent)
}

func (a *App) setContent(window fyne.Window) {
	params := []fyne.CanvasObject{
		sdk.Describe("Алгоритм протяжки времени", a.widgets.method),
		sdk.Describe("Процент возвращаемых заявок", a.widgets.returnPercentEntry),
	}
	params = append(params, a.baseModel.Widgets()...)

	results := []fyne.CanvasObject{
		sdk.Describe("Размер очереди", a.widgets.queueLen),
	}

	content := sdk.Layout(
		params,
		results,
		a.onComputeButtonPressed,
	)

	window.SetContent(content)
}

func (a *App) onComputeButtonPressed() {
	log.Debug("compute button pressed")

	percent, err := a.widgets.returnPercentEntry.GetInt()
	if err != nil {
		sdk.ShowError(&sdk.Error{Human: "Неверное значение процента возвращаемых заявок.", Raw: err})
		return
	}

	if percent < 0 || percent > 100 {
		sdk.ShowError(&sdk.Error{Human: "Процент должен быть в диапазоне [0; 100].", Raw: err})
		return
	}

	selected := a.widgets.method.Selected
	log.Debug("use selected method", "method", selected)
	result, err := a.models[selected].Compute(limitProcessed, percent)
	if err != nil {
		sdk.ShowError(err)
		return
	}

	a.widgets.queueLen.SetText(strconv.Itoa(result.QueueLen))
}

func (a *App) Run() {
	sdk.Run()
}

type methodModel interface {
	Compute(limit, percent int) (*model.Result, error)
}
