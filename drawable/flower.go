package drawable

import (
	"math/rand/v2"
	"strings"

	"github.com/DCCXXV/flores/utils"
	"github.com/gdamore/tcell/v2"
)

var Flowers = map[string]int{
	"daisy":     0,
	"rose":      1,
	"sunflower": 2,
	"lavender":  3,
	"poppy":     4,
	"cherry":    5,
	"bluet":     6,
}

var TotalFlowers = len(Flowers)

// DrawFlower draws a flower where:
//   - s is the screen
//   - random is the random number generator
//   - x,y -> x,y-1 is where the stem ends
//   - d is the dimension/size (optional for some flowers)
//   - t is the type of flower :3
func DrawFlower(s tcell.Screen, random *rand.Rand, x, y, d, t int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	switch t {
	case 0: // daisy
		s.SetContent(x, y, '@', nil, style.Foreground(tcell.ColorYellow))
		horizontalPetal := strings.Repeat("==", d)
		s.PutStrStyled(x+1, y, horizontalPetal, style)
		s.PutStrStyled(x-len(horizontalPetal), y, horizontalPetal, style)
		for i := 1; i < d+1; i++ {
			s.PutStrStyled(x-i, y-i, "\\", style)
			s.PutStrStyled(x+i, y-i, "/", style)
			s.PutStrStyled(x-i, y+i, "/", style)
			s.PutStrStyled(x+i, y+i, "\\", style)
		}
	case 1: // rose
		petalColors := []tcell.Color{tcell.ColorRed, tcell.ColorLightSteelBlue, tcell.ColorSeashell}
		randomColor := petalColors[random.IntN(len(petalColors))]
		style = tcell.StyleDefault.Foreground(randomColor)
		s.PutStrStyled(x-1, y, utils.GenerateRandomString(3), style)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomString(3), style)
		s.PutStrStyled(x-2, y-2, utils.GenerateRandomString(5), style)
		s.PutStrStyled(x-2, y-3, utils.GenerateRandomString(5), style)
		s.PutStrStyled(x-2, y-4, utils.GenerateRandomString(5), style)
	case 2: // sunflower
		style = tcell.StyleDefault.Foreground(tcell.ColorMaroon)
		if d <= 1 {
			s.PutStrStyled(x, y-1, "#", style)
			style = tcell.StyleDefault.Foreground(tcell.ColorYellow)
			s.PutStrStyled(x-1, y-1, "o", style)
			s.PutStrStyled(x+1, y-1, "o", style)
			s.PutStrStyled(x, y, "o", style)
			s.PutStrStyled(x, y-2, "o", style)
		} else if d == 2 {
			s.PutStrStyled(x, y, "o", style)
			s.PutStrStyled(x-1, y-1, "ooo", style)
			s.PutStrStyled(x, y-2, "o", style)
			style = tcell.StyleDefault.Foreground(tcell.ColorYellow)
			s.PutStrStyled(x-2, y-2, utils.GenerateRandomString(2), style)
			s.PutStrStyled(x+1, y-2, utils.GenerateRandomString(2), style)
			s.PutStrStyled(x-2, y, utils.GenerateRandomString(2), style)
			s.PutStrStyled(x+1, y, utils.GenerateRandomString(2), style)
			s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(3), style)
			s.PutStrStyled(x-1, y-3, utils.GenerateRandomString(3), style)
			s.PutStrStyled(x-3, y-1, utils.GenerateRandomString(2), style)
			s.PutStrStyled(x+2, y-1, utils.GenerateRandomString(2), style)
		} else {
			s.PutStrStyled(x-1, y, "@@@", style)
			s.PutStrStyled(x-2, y-1, "@@@@@", style)
			s.PutStrStyled(x-1, y-2, "@@@", style)
			style = tcell.StyleDefault.Foreground(tcell.ColorYellow)
			s.PutStrStyled(x-4, y+1, "&  &&@", style)
			s.PutStrStyled(x-1, y+2, "&  &", style)
			s.PutStrStyled(x-2, y-3, "&&%%&", style)
			s.PutStrStyled(x-1, y-4, "& & ", style)
			s.PutStrStyled(x-3, y, "%&", style)
			s.PutStrStyled(x+2, y, "&&&", style)
			s.PutStrStyled(x-3, y-1, "&", style)
			s.PutStrStyled(x+3, y-1, "&", style)
			s.PutStrStyled(x-4, y-2, "&$&", style)
			s.PutStrStyled(x+2, y-2, "&#&", style)
		}
	case 3: // lavender
		style = tcell.StyleDefault.Foreground(tcell.ColorPurple)
		s.PutStrStyled(x-1, y, "\\|/", style.Foreground(tcell.ColorDarkGreen))
		s.PutStrStyled(x-2, y-1, "\\ | /", style.Foreground(tcell.ColorDarkGreen))
		/*
			s.PutStrStyled(x-3, y-1, "()o()Oo", style)
			s.PutStrStyled(x-2, y-2, "o O o", style)
			s.PutStrStyled(x-3, y-3, "()oO o", style)
			s.PutStrStyled(x-1, y-4, "o O", style)
			s.PutStrStyled(x-1, y-5, "Ooo", style)
			s.PutStrStyled(x-1, y-6, "o O", style)
			s.PutStrStyled(x-1, y-7, "o o", style)
			s.PutStrStyled(x-1, y-8, " o ", style)
		*/
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomTile("o", "O", "0"), style)
		s.PutStrStyled(x+1, y-1, utils.GenerateRandomTile("o", "O", "0"), style)
		s.PutStrStyled(x-1, y-2, "o o", style)
		s.PutStrStyled(x-1, y-3, "o o", style)
		s.PutStrStyled(x-1, y-4, " 0 ", style)
	case 4: // poppy
		s.PutStrStyled(x, y, utils.GenerateRandomTile(), style.Foreground(tcell.ColorDarkRed))

		petalColors := []tcell.Color{tcell.ColorRed}
		randomColor := petalColors[random.IntN(len(petalColors))]
		style := style.Foreground(randomColor)
		s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(3), style)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomString(3), style)
		s.PutStrStyled(x-2, y, utils.GenerateRandomString(2), style)
		s.PutStrStyled(x+1, y, utils.GenerateRandomString(2), style)
	case 5: // cherry
		s.PutStrStyled(x, y, utils.GenerateRandomTile(), style.Foreground(tcell.ColorWhite))

		petalColors := []tcell.Color{tcell.ColorPink, tcell.ColorPlum, tcell.ColorLightPink}
		randomColor := petalColors[random.IntN(len(petalColors))]
		style := style.Foreground(randomColor)

		s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(3), style)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomString(3), style)
		s.PutStrStyled(x-2, y, utils.GenerateRandomString(2), style)
		s.PutStrStyled(x+1, y, utils.GenerateRandomString(2), style)
	case 6: // bluet
		style := style.Foreground(tcell.ColorRoyalBlue)

		s.PutStrStyled(x, y-4, "_", style)
		s.PutStrStyled(x-2, y-3, "_/|\\_", style)

		s.PutStrStyled(x-3, y-2, "/_", style)
		s.PutStrStyled(x-1, y-2, "\\", style.Foreground(tcell.ColorLightBlue))
		s.PutStrStyled(x, y-2, "@", style.Foreground(tcell.ColorYellow))
		s.PutStrStyled(x+1, y-2, "/", style.Foreground(tcell.ColorLightBlue))
		s.PutStrStyled(x+2, y-2, "_\\", style)

		s.PutStrStyled(x-3, y-1, "\\_", style)
		s.PutStrStyled(x-1, y-1, "/", style.Foreground(tcell.ColorLightBlue))
		s.PutStrStyled(x, y-1, "|", style)
		s.PutStrStyled(x+1, y-1, "\\", style.Foreground(tcell.ColorLightBlue))
		s.PutStrStyled(x+2, y-1, "_/", style)

		s.PutStrStyled(x-1, y, "\\_/", style)
	}
}
