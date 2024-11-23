package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("App started...")
	//creating app, window and adjusting settings
	a := app.New()
	w := a.NewWindow("Deployment app")
	w.Resize(fyne.NewSize(800, 600))
	//setup whole UI
	SetupUI(w)

	w.ShowAndRun()
}

func SetupUI(w fyne.Window) {
	// Set window content
	w.SetContent(
		container.NewVBox(
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Portfolio!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Weather app!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Notes Frontend!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Notes Api!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Playground Frontend!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
			container.NewBorder(
				layout.NewSpacer(),
				nil,
				widget.NewLabel("Playground Api!"),
				container.NewBorder(
					nil,nil,
					widget.NewButton("Deploy!", func() {
						w.Close()
					}),
					nil,
				),
			),
		),
	)

	// Set window icon
	iconData, err := os.ReadFile("icon.png")
	if err == nil {
		w.SetIcon(fyne.NewStaticResource("icon.png", iconData))
	}
}