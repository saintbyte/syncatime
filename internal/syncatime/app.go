package app

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	serial "go.bug.st/serial.v1"
	"log"
)

const AppName string = "SyncAtime"

type App struct {
	app fyne.App
}

func windowSize(part float32) fyne.Size {
	return fyne.NewSize(800, 800)
}

func (mainApp *App) restoreWindow(window fyne.Window) {
	window.Resize(windowSize(0.9))
}
func NewApp() *App {
	return &App{}
}
func OpenText(w fyne.Window) func() {
	return func() {
		dialog.NewFileOpen(
			func(closer fyne.URIReadCloser, err error) {
				if err != nil {

				}
				fmt.Println(closer)
			},
			w,
		)
	}
}
func ClearText(textView *widget.Entry) func() {
	return func() {
		textView.SetText("")
	}
}
func NewTextView(w fyne.Window) *fyne.Container {
	textView := widget.NewMultiLineEntry()
	textView.Wrapping = fyne.TextWrapWord
	box := container.NewMax(container.NewVBox(
		container.NewMax(textView),
		container.NewVBox(
			widget.NewButton("Open", OpenText(w))),
		widget.NewButton("Clear", ClearText(textView)),
	))

	return container.NewMax(box)
}
func NewWorkArea(w fyne.Window) *container.Split {
	area := container.NewHSplit(
		NewTextView(w),
		widget.NewLabel("2"),
	)
	return area
}

func (mainApp *App) Run() {
	mainApp.app = app.New()
	w := mainApp.app.NewWindow(AppName)
	availablePorts, err := serial.GetPortsList()
	// https://pkg.go.dev/go.bug.st/serial.v1#section-documentation
	if err != nil {
		log.Fatalf("Error getting available ports: %v", err)
	}

	for _, port := range availablePorts {
		fmt.Println(port)
	}
	workingArea := NewWorkArea(w)
	timeline := container.NewHSplit(
		widget.NewLabel("3"),
		widget.NewLabel("4"),
	)
	w.SetContent(container.NewVSplit(
		//,
		workingArea,
		timeline,
	))
	mainApp.restoreWindow(w)
	w.ShowAndRun()
}
