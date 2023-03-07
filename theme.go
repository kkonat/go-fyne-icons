package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed "3270Narrow.ttf"
var fontData []byte

var regularFont = &fyne.StaticResource{
	StaticName:    "3270Narrow.ttf",
	StaticContent: fontData}

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

// return bundled font resource
func (*myTheme) Font(s fyne.TextStyle) fyne.Resource {
	// if s.Monospace {
	// 	return theme.DefaultTheme().Font(s)
	// }
	// if s.Bold {
	// 	if s.Italic {
	// 		return theme.DefaultTheme().Font(s)
	// 	}
	// 	return resourceMplus1cBoldTtf
	// }
	// if s.Italic {
	// 	return theme.DefaultTheme().Font(s)
	// }
	return regularFont
}

func (*myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*myTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
