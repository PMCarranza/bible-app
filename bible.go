package main

import (
	biblelibhc "bible-app/biblelibhc"
	"image/color"
	"log"
	"strings"

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
	fondo := canvas.NewRectangle(&color.RGBA{10, 23, 35, 5})
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

	w.Resize(fyne.NewSize(200, 200))

	//	bookQuery := "bcv:Genesis/050:025"
	//    bookQuery := "bcv:Genesis/010:016"
	//    bookQuery := "bcv:Genesis/010:026"
	//	bookQuery := "bcv:Revelation/011:006"
	//	bookQuery := "bcv:Revelation/011:017"
	//  bookQuery := "bcv:Genesis/001:001"
	//  bookQuery := "bcv:Genesis/040:010"
	//    bookQuery := "bcv:Psalms/023:001"
	//    bookQuery := "phrase:swore to Abraham"
	//    bookQuery := "seqphrase:the name of Jesus"
	//    bookQuery := "seqphrase:Christ"
	bookQuery := "booklist:"
	//    bookQuery :=
	queryResponse := biblelibhc.Query(bookQuery, 20)

	pageBody := widget.NewLabelWithStyle(queryResponse.Body.String(), fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	borderLayout := layout.NewBorderLayout(top, bottom, left, right)
	pageBody.Wrapping = fyne.TextWrapWord
	// pageBody.Hide()
	pageLayout := fyne.NewContainerWithLayout(borderLayout, top, bottom, left, right, underside, pageBody)
	w.SetContent(pageLayout)

	canvasSize := fyne.NewSize(800, 400)

	w.Resize(canvasSize)
	w.SetFixedSize(false)

	canvasTile := canvasSize.Width / 7

	popupPosition := fyne.NewPos(canvasTile, canvasSize.Height/2)

	popupSize := fyne.Size{Width: canvasSize.Width - 2*canvasTile, Height: 40}
	popupText := queryResponse.Body.String()
	//	popupText = strings.ReplaceAll(popupText, ",", "·")
	popupText = strings.ReplaceAll(popupText, ",", ".")
	popup := widget.NewTickerPopUpAtPosition(widget.NewLabel(popupText), w.Canvas(), popupPosition, popupSize)
	popup.Resize(popupSize)
	popup.Show()

	w.ShowAndRun()
	// q.ShowAndRun()

}

type tappableIcon struct {
	widget.Icon
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (t *tappableIcon) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}

func (t *tappableIcon) TappedSecondary(_ *fyne.PointEvent) {
}

// func secondMain() {
// 	p := app.New()
// 	q := p.NewWindow("Tappable")
// 	q.SetContent(newTappableIcon(theme.FyneLogo()))
// 	q.ShowAndRun()
// }
