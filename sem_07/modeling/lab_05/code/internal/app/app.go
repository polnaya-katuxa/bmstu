package app

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/charmbracelet/log"
	"lab_05/internal/model"
	"lab_05/pkg/sdk"
	"lab_05/pkg/sdk/advanced_widgets"
)

type App struct {
	widgets widgets

	model *model.Model
}

type widgets struct {
	count *advanced_widgets.NumericalEntry

	queueLen sdk.ResultWidget
}

func New() *App {
	var a App

	a.model = model.NewModel(
		model.NewClientCreator("Клиент", 10, 2),
		[]*model.OperatorCreator{
			model.NewOperatorCreator("Оператор 1", 20, 5),
			model.NewOperatorCreator("Оператор 2", 40, 10),
			model.NewOperatorCreator("Оператор 3", 40, 20),
		},
		[]*model.ComputerCreator{
			model.NewComputerCreator("Компьютер 1", 15),
			model.NewComputerCreator("Компьютер 2", 30),
		},
		map[int]int{0: 0, 1: 0, 2: 1},
	)

	a.createWidgets()

	return &a
}

func (a *App) createWidgets() {
	a.widgets.count = advanced_widgets.NewIntEntry(300)

	a.widgets.queueLen = sdk.Result()

	sdk.SetFillWindowFunc(a.setContent)
}

func (a *App) setContent(window fyne.Window) {
	params := []fyne.CanvasObject{
		sdk.Describe("Количество заявок", a.widgets.count),
	}
	params = append(params, a.model.Widget())

	results := []fyne.CanvasObject{
		sdk.Describe("Вероятность отказа", a.widgets.queueLen),
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

	count, err := a.widgets.count.GetInt()
	if err != nil {
		sdk.ShowError(&sdk.Error{Human: "Неверное значение процента возвращаемых заявок.", Raw: err})
		return
	}

	result, err := a.model.Compute(count)
	if err != nil {
		sdk.ShowError(err)
		return
	}

	a.widgets.queueLen.SetText(fmt.Sprintf("%.2f%%", result.DeclinePercent))
}

func (a *App) Run() {
	sdk.Run()
}
