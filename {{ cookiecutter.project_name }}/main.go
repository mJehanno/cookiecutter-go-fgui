package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("{{ cookiecutter.project_name }}")
	w.Resize(fyne.NewSize(200, 100))
	w.ShowAndRun()
}
