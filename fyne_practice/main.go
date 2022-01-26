package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	//New App command
	myapp := app.New()

	// windows with title
	window := myapp.NewWindow("miniERP V0.0.1")

	// Windows resize of app
	window.Resize(fyne.NewSize(600, 400))

	// label call
	label1 := widget.NewLabel("Welcome to new miniERP")

	//set contain
	window.SetContent(label1)
	// running and show
	window.ShowAndRun()
}
