package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func cell() fyne.CanvasObject {
	cell := canvas.NewRectangle(&color.Black)
	cell.SetMinSize(fyne.NewSize(30, 30))
	return cell
}

func main() {
	a := app.New()

	w := a.NewWindow("Bible App")

	top := cell()
	bottom := cell()
	left := cell()
	right := cell()

	middle := widget.NewLabelWithStyle("test\ntest\ntest\ntest", fyne.TextAlignLeading, fyne.TextStyle{})

	borderLayout := layout.NewBorderLayout(top, bottom, left, right)

	pageLayout := fyne.NewContainerWithLayout(borderLayout, top, bottom, left, right, middle)
	w.SetContent(pageLayout)

	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(false)
	w.ShowAndRun()
}
