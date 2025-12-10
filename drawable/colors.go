package drawable

import (
	"github.com/gdamore/tcell/v2"
)

var style = tcell.StyleDefault

// main colors
var White = style.Foreground(tcell.ColorWhite)
var Red = style.Foreground(tcell.ColorRed)
var Blue = style.Foreground(tcell.ColorRoyalBlue)
var Green = style.Foreground(tcell.ColorGreen)
var Yellow = style.Foreground(tcell.ColorYellow)
var Violet = style.Foreground(tcell.ColorViolet)
var Brown = style.Foreground(tcell.ColorSandyBrown)

// shades
var LightBlue = style.Foreground(tcell.ColorLightSteelBlue)
var DarkGreen = style.Foreground(tcell.ColorDarkGreen)
var Pale = style.Foreground(tcell.ColorSeashell)

func TrueColorOn() {
	// main colors
	White = style.Foreground(tcell.ColorWhite.TrueColor())
	Red = style.Foreground(tcell.ColorRed.TrueColor())
	Blue = style.Foreground(tcell.ColorRoyalBlue.TrueColor())
	Green = style.Foreground(tcell.ColorGreen.TrueColor())
	Yellow = style.Foreground(tcell.ColorYellow.TrueColor())
	Violet = style.Foreground(tcell.ColorViolet.TrueColor())
	Brown = style.Foreground(tcell.ColorSandyBrown.TrueColor())

	// shades
	LightBlue = style.Foreground(tcell.ColorLightSteelBlue.TrueColor())
	DarkGreen = style.Foreground(tcell.ColorDarkGreen.TrueColor())
	Pale = style.Foreground(tcell.ColorSeashell.TrueColor())
}
