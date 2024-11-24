package sdk

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/charmbracelet/log"
)

var fyneApp fyne.App
var fyneWindow fyne.Window
var fillFunc func(window fyne.Window)

func Window() fyne.Window {
	return defaultProvider.window()
}

func ShowError(err error) {
	log.Error("error occurred", "error", err)

	text := err.Error()

	var sdkErr *Error
	if errors.As(err, &sdkErr) {
		text = sdkErr.Human
	}

	defaultProvider.showError(text)
}

func Describe(label string, wrapped fyne.Widget) *fyne.Container {
	return defaultProvider.describe(label, wrapped)
}

func Result() ResultWidget {
	return defaultProvider.result()
}

func Layout(parameters []fyne.CanvasObject, results []fyne.CanvasObject, computeResults func()) *fyne.Container {
	return defaultProvider.layout(parameters, results, computeResults)
}

func SetFillWindowFunc(fill func(window fyne.Window)) {
	fillFunc = fill
}

func Run() {
	log.Info("running application")
	log.Info("setting content")
	fillFunc(fyneWindow)

	log.Info("running window")
	fyneWindow.ShowAndRun()
	log.Warn("exiting")
}

func Reload() {
	log.Debug("toggle reloading")

	log.Debug("setting content")
	fillFunc(fyneWindow)

	log.Debug("showing window")
	fyneWindow.Show()
}

func init() {
	log.SetLevel(log.DebugLevel)

	defaultProvider = newBlumbaProvider()

	//switch os.Getenv("USER") {
	//case "avknyazhev", "muhomorfus":
	//	defaultProvider = newDrynProvider()
	//case "polnaya_katuxa":
	//	defaultProvider = newBlumbaProvider()
	//default:
	//	log.Error("invalid environment")
	//	os.Exit(1)
	//}

	fyneApp = app.New()
	fyneApp.Settings().SetTheme(defaultProvider.Theme())
	fyneWindow = fyneApp.NewWindow("lab_04")
}
