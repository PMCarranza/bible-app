package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Bible App")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Check it out Joel! \n On my to catch up with you!"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	w.ShowAndRun()
}
