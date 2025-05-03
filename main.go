package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/go-vgo/robotgo"
)

func autoClick(button int8, amount float64, interval float64) { // 0 - Left; 1 - Right
	robotgo.MilliSleep(5000) // 5 seconds (FIRST SLEEP TO ALLOW USER TO PLACE CURSOR)

	for i := 0; float64(i) <= amount; i++ {
		if button == 0 {
			robotgo.Click("left")
			robotgo.MilliSleep(int(interval))
		} else if button == 1 {
			robotgo.Click("right")
			robotgo.MilliSleep(int(interval))
		}
	}
}

func autoType(in string, amount float64, interval float64) {
	robotgo.MilliSleep(5000) // 5 seconds (FIRST SLEEP TO ALLOW USER TO PLACE CURSOR)

	for i := 0; float64(i) <= amount; i++ {
		robotgo.TypeStr(in)
		robotgo.MilliSleep(int(interval))
	}
}

func main() {
	app := app.New()
	window := app.NewWindow("Autoclicker")

	// Widgets
	amountSlider := widget.NewSlider(0, 1000)
	amountSlider.Step = 1

	amountSliderLabel := widget.NewLabel(fmt.Sprintf("Amount of Clicks: %v", amountSlider.Value))
	amountSlider.OnChanged = func(f float64) {
		amountSliderLabel.SetText(fmt.Sprintf("Amount of Clicks: %v", f))
	}

	intervalSlider := widget.NewSlider(0, 10000)
	intervalSlider.Step = 1

	intervalSliderLabel := widget.NewLabel(fmt.Sprintf("Time between clicks (milliseconds): %v", intervalSlider.Value))
	intervalSlider.OnChanged = func(f float64) {
		intervalSliderLabel.SetText(fmt.Sprintf("Time between clicks (milliseconds): %v", f))
	}

	input := widget.NewEntry()
	input.SetPlaceHolder("Select a button/string")

	mouseButton := widget.NewButton("Press Mouse", func() { autoClick(0, amountSlider.Value, intervalSlider.Value) })
	kbdButton := widget.NewButton("Press Keyboard", func() { autoType(input.Text, amountSlider.Value, intervalSlider.Value) })

	window.SetContent(container.NewVBox(
		widget.NewLabel("version 0.1.0"),

		amountSliderLabel,
		amountSlider,

		intervalSliderLabel,
		intervalSlider,

		input,

		mouseButton,
		kbdButton,
	))

	window.ShowAndRun()
}
