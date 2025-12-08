package drawable

import (
	"strings"

	"github.com/DCCXXV/flores/utils"
	"github.com/gdamore/tcell/v2"
)

/*
 * x,y -> x,y-1 is where the tail ends
 * d -> dimension/size (optional for some flowers)
 * t -> type of flower :3
 */
func DrawFlower(s tcell.Screen, x, y, d, t int) {
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
		style = tcell.StyleDefault.Foreground(tcell.ColorRed)
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
	}
}
