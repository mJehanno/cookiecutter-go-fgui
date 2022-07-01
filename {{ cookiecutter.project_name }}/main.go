package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

func main() {
	a := app.New()
	w := a.NewWindow("{{ cookiecutter.project_name }}")

	if desk, ok := a.(desktop.App); ok {
		menu := fyne.NewMenu("{{ cookiecutter.project_name }}", nil)
		desk.SetSystemTrayMenu(m)
	}

	w.Resize(fyne.NewSize(200, 100))
	w.ShowAndRun()
}
