package main

import (
	_ "embed"
	"strings"
	_ "unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:linkname icons fyne.io/fyne/v2/theme.icons
var icons map[fyne.ThemeIconName]fyne.Resource

func main() {
	app := app.New()
	app.Settings().SetTheme(&myTheme{})
	w := app.NewWindow("Fyne icons browser")
	w.Resize(fyne.NewSize(840, 460))
	w.SetContent(iconScreen())
	w.ShowAndRun()
}

func iconScreen() fyne.CanvasObject {

	txt := widget.NewLabel("Click on an icon to see its name..")

	gridWrap := container.NewGridWrap(fyne.Size{Width: 160, Height: 64})
	iconsScroll := container.NewVScroll(gridWrap)

	for _, icon := range icons {
		btn := widget.NewButtonWithIcon("", icon, nil)
		btn.Alignment = widget.ButtonAlignCenter

		btn.OnTapped = func() {
			txt.SetText(btn.Icon.Name())
		}
		l := widget.NewLabel(strings.Split(icon.Name(), ".")[0])
		c := container.NewVBox(btn, container.NewCenter(l))

		gridWrap.Add(c)
	}
	return container.NewBorder(txt, nil, nil, nil, iconsScroll)
}
