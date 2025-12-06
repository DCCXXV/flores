package drawable
import (
   	"github.com/gdamore/tcell/v2"
)

func DrawBase(s tcell.Screen, x, y int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorSandyBrown)
	s.PutStrStyled(x-3, y, "_______", style)
	s.PutStrStyled(x-5, y+1, "(___      )", style)
	s.PutStrStyled(x-4, y+2, "|       |", style)
	s.PutStrStyled(x-4, y+3, "|       |", style)
	s.PutStrStyled(x-3, y+4, "\\     /", style)
	s.PutStrStyled(x-2, y+5, "_____", style)
}
