package main

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/h2non/bimg"
)

func main() {
	a := app.New()
	w := a.NewWindow("Irene")

	w.Resize(fyne.NewSize(1920/2, 1080/2))

	ScaleText := widget.NewEntry()
	ScaleText.PlaceHolder = "Scale eg. (2 for Double the size)"
	ScaleText.TextStyle.Bold = true
	page_1 := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		widget.NewLabel("Drag Files into window to process"),
		ScaleText,
		layout.NewSpacer(),
	)

	w.SetContent(page_1)
	w.SetOnDropped(func(p fyne.Position, u []fyne.URI) {
		for _, uri := range u {

			scale, err := strconv.ParseFloat(ScaleText.Text, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			buffer, err := bimg.Read(uri.Path())
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			image, err := bimg.NewImage(buffer).Size()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if scale <= 1 {
				NewWidth, NewHeight := int(float64(image.Width)*scale), int(float64(image.Height)*scale)
				newImage, _ := bimg.NewImage(buffer).Resize(NewWidth, NewHeight)
				bimg.Write(uri.Path()+ScaleText.Text+".png", newImage)
			}

			if scale >= 1 {
				newImage, _ := bimg.NewImage(buffer).Enlarge(image.Width*int(scale), image.Height*int(scale))
				bimg.Write(uri.Path()+ScaleText.Text+".png", newImage)
			}

		}
	})

	w.SetOnClosed(a.Quit)

	w.ShowAndRun()
}
