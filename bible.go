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
	cell := canvas.NewRectangle(&color.Transparent)
	cell.SetMinSize(fyne.NewSize(30, 30))
	return cell
}

func fondo() fyne.CanvasObject {
	fondo := canvas.NewRectangle(&color.RGBA{147, 220, 245, 5})
	fondo.SetMinSize(fyne.NewSize(120, 120))
	return fondo
}

func main() {
	a := app.New()

	w := a.NewWindow("Bible App")

	underside := fondo()

	top := cell()
	bottom := cell()
	left := cell()
	right := cell()

	middle := widget.NewLabelWithStyle("test\ntest\ntest\ntest", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	borderLayout := layout.NewBorderLayout(top, bottom, left, right)

	pageLayout := fyne.NewContainerWithLayout(borderLayout, top, bottom, left, right, underside, middle)
	w.SetContent(pageLayout)

	w.Resize(fyne.NewSize(600, 600))
	w.SetFixedSize(false)
	w.ShowAndRun()
}
