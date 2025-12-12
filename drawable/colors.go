package drawable

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

type ColorScheme struct {
	White, Red, Blue, Green, Yellow, Violet, Brown tcell.Style
	LightBlue, DarkGreen, Pale                     tcell.Style
}

var (
	colors *ColorScheme
	once   sync.Once
	style  = tcell.StyleDefault
)

func InitColors(trueColor bool) {
	once.Do(func() {
		colors = newColorScheme(trueColor)
	})
}

func GetColors() *ColorScheme {
	return colors
}

func newColorScheme(trueColor bool) *ColorScheme {
	scheme := &ColorScheme{}
	if trueColor {
		scheme.White = style.Foreground(tcell.ColorWhite.TrueColor())
		scheme.Red = style.Foreground(tcell.ColorRed.TrueColor())
		scheme.Blue = style.Foreground(tcell.ColorRoyalBlue.TrueColor())
		scheme.Green = style.Foreground(tcell.ColorLawnGreen.TrueColor())
		scheme.Yellow = style.Foreground(tcell.ColorYellow.TrueColor())
		scheme.Violet = style.Foreground(tcell.ColorViolet.TrueColor())
		scheme.Brown = style.Foreground(tcell.ColorSandyBrown.TrueColor())

		scheme.LightBlue = style.Foreground(tcell.ColorLightBlue.TrueColor())
		scheme.DarkGreen = style.Foreground(tcell.ColorDarkGreen.TrueColor())
		scheme.Pale = style.Foreground(tcell.ColorSeashell.TrueColor())
	} else {
		scheme.White = style.Foreground(tcell.ColorWhite)
		scheme.Red = style.Foreground(tcell.ColorRed)
		scheme.Blue = style.Foreground(tcell.ColorRoyalBlue)
		scheme.Green = style.Foreground(tcell.ColorGreen)
		scheme.Yellow = style.Foreground(tcell.ColorYellow)
		scheme.Violet = style.Foreground(tcell.ColorViolet)
		scheme.Brown = style.Foreground(tcell.ColorSandyBrown)

		scheme.LightBlue = style.Foreground(tcell.ColorLightBlue)
		scheme.DarkGreen = style.Foreground(tcell.ColorDarkGreen)
		scheme.Pale = style.Foreground(tcell.ColorSeashell)
	}
	return scheme
}
