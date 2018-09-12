package main

import (
	"fmt"

	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/desktop"
	"github.com/fyne-io/fyne/examples/apps"
	"github.com/fyne-io/fyne/layout"
	"github.com/fyne-io/fyne/theme"
	W "github.com/fyne-io/fyne/widget"
)

func myApp(app fyne.App) {
	w := app.NewWindow("My Fyne App")
	w.Show()
}

func main() {
	app := desktop.NewApp()
	welcome(app)
}

func appButton(app fyne.App, label string, onClick func(fyne.App)) *W.Button {
	return &W.Button{Text: label, OnTapped: func() {
		onClick(app)
	}}
}

func blogApp(app fyne.App) {
	apps.Blog(app)
}

func calcApp(app fyne.App) {
	apps.Calculator(app)
}

func clockApp(app fyne.App) {
	apps.Clock(app)
}

func fractalApp(app fyne.App) {
	apps.Fractal(app)
}

func lifeApp(app fyne.App) {
	apps.Life(app)
}

func layoutApp(app fyne.App) {
	apps.Layout(app)
}

func canvasApp(app fyne.App) {
	apps.Canvas(app)
}

func welcome(app fyne.App) {
	w := app.NewWindow("Examples")
	w.SetContent(&W.List{Children: []fyne.CanvasObject{
		&W.Label{Text: "Fyne Examples!"},

		W.NewGroup("Apps", []fyne.CanvasObject{
			appButton(app, "Blog", blogApp),
			appButton(app, "Calculator", calcApp),
			appButton(app, "Clock", clockApp),
			appButton(app, "Fractal", fractalApp),
			appButton(app, "Life", lifeApp),
		}...),

		W.NewGroup("Demos", []fyne.CanvasObject{
			appButton(app, "Canvas", canvasApp),
			appButton(app, "Layout", layoutApp),
			&W.Entry{},
			&W.Check{Text: "Check", OnChanged: func(on bool) { fmt.Println("checked", on) }},
		}...),

		layout.NewSpacer(),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			&W.Button{Text: "Dark", OnTapped: func() {
				fyne.GetSettings().SetTheme("dark")
			}},
			&W.Button{Text: "Light", OnTapped: func() {
				fyne.GetSettings().SetTheme("light")
			}},
		),
		&W.Button{Text: "Quit", Icon: theme.CancelIcon(), OnTapped: func() {
			app.Quit()
		}},
	}})
	w.Show()
}
