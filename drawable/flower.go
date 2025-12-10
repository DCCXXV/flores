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
	"hibiscus":  5,
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
	switch t {
	case Flowers["daisy"]:
		s.SetContent(x, y, '@', nil, Yellow)
		horizontalPetal := strings.Repeat("==", d)
		s.PutStrStyled(x+1, y, horizontalPetal, White)
		s.PutStrStyled(x-len(horizontalPetal), y, horizontalPetal, White)
		for i := 1; i < d+1; i++ {
			s.PutStrStyled(x-i, y-i, "\\", White)
			s.PutStrStyled(x+i, y-i, "/", White)
			s.PutStrStyled(x-i, y+i, "/", White)
			s.PutStrStyled(x+i, y+i, "\\", White)
		}
	case Flowers["rose"]:
		petalColors := []tcell.Style{Red, LightBlue, Pale}
		randomColor := petalColors[random.IntN(len(petalColors))]

		s.PutStrStyled(x-1, y, utils.GenerateRandomString(3), randomColor)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomString(3), randomColor)
		s.PutStrStyled(x-2, y-2, utils.GenerateRandomString(5), randomColor)
		s.PutStrStyled(x-2, y-3, utils.GenerateRandomString(5), randomColor)
		s.PutStrStyled(x-2, y-4, utils.GenerateRandomString(5), randomColor)
	case Flowers["sunflower"]:
		if d <= 1 {
			s.PutStrStyled(x, y-1, "#", Brown)
			s.PutStrStyled(x-1, y-1, "o", Yellow)
			s.PutStrStyled(x+1, y-1, "o", Yellow)
			s.PutStrStyled(x-1, y, "ooo", Yellow)
			s.PutStrStyled(x-1, y-2, "ooo", Yellow)
		} else if d == 2 {
			s.PutStrStyled(x, y, "o", Brown)
			s.PutStrStyled(x-1, y-1, "ooo", Brown)
			s.PutStrStyled(x, y-2, "o", Brown)
			s.PutStrStyled(x-2, y-2, utils.GenerateRandomString(2), Yellow)
			s.PutStrStyled(x+1, y-2, utils.GenerateRandomString(2), Yellow)
			s.PutStrStyled(x-2, y, utils.GenerateRandomString(2), Yellow)
			s.PutStrStyled(x+1, y, utils.GenerateRandomString(2), Yellow)
			s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(3), Yellow)
			s.PutStrStyled(x-1, y-3, utils.GenerateRandomString(3), Yellow)
			s.PutStrStyled(x-3, y-1, utils.GenerateRandomString(2), Yellow)
			s.PutStrStyled(x+2, y-1, utils.GenerateRandomString(2), Yellow)
		} else {
			s.PutStrStyled(x-1, y, "@@@", Brown)
			s.PutStrStyled(x-2, y-1, "@@@@@", Brown)
			s.PutStrStyled(x-1, y-2, "@@@", Brown)

			s.PutStrStyled(x-4, y+1, "&  &&@", Yellow)
			s.PutStrStyled(x-1, y+2, "&  &", Yellow)
			s.PutStrStyled(x-2, y-3, "&&%%&", Yellow)
			s.PutStrStyled(x-1, y-4, "& & ", Yellow)
			s.PutStrStyled(x-3, y, "%&", Yellow)
			s.PutStrStyled(x+2, y, "&&&", Yellow)
			s.PutStrStyled(x-3, y-1, "&", Yellow)
			s.PutStrStyled(x+3, y-1, "&", Yellow)
			s.PutStrStyled(x-4, y-2, "&$&", Yellow)
			s.PutStrStyled(x+2, y-2, "&#&", Yellow)
		}
	case Flowers["lavender"]:
		s.PutStrStyled(x-1, y, "\\|/", DarkGreen)
		s.PutStrStyled(x-2, y-1, "\\ | /", DarkGreen)
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
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomTile("o", "O", "0"), Violet)
		s.PutStrStyled(x+1, y-1, utils.GenerateRandomTile("o", "O", "0"), Violet)
		s.PutStrStyled(x-1, y-2, "o o", Violet)
		s.PutStrStyled(x-1, y-3, "o o", Violet)
		s.PutStrStyled(x-1, y-4, " 0 ", Violet)
	case Flowers["poppy"]:
		s.PutStrStyled(x, y, utils.GenerateRandomTile(), Brown)
		s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(3), Red)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomString(3), Red)
		s.PutStrStyled(x-2, y, utils.GenerateRandomString(2), Red)
		s.PutStrStyled(x+1, y, utils.GenerateRandomString(2), Red)
	case Flowers["hibiscus"]:
		petalColors := []tcell.Style{Red, Violet, Blue}
		randomColor := petalColors[random.IntN(len(petalColors))]

		s.PutStrStyled(x, y, "|", White)
		s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(3), randomColor)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomTile(), randomColor)
		s.PutStrStyled(x, y-1, "%", Yellow)
		s.PutStrStyled(x+1, y-1, utils.GenerateRandomTile(), randomColor)
		s.PutStrStyled(x-2, y, utils.GenerateRandomString(2), randomColor)
		s.PutStrStyled(x+1, y, utils.GenerateRandomString(2), randomColor)
	case Flowers["bluet"]:
		s.PutStrStyled(x, y-4, "_", Blue)
		s.PutStrStyled(x-2, y-3, "_/|\\_", Blue)

		s.PutStrStyled(x-3, y-2, "/_", Blue)
		s.PutStrStyled(x-1, y-2, "\\", LightBlue)
		s.PutStrStyled(x, y-2, "@", Yellow)
		s.PutStrStyled(x+1, y-2, "/", LightBlue)
		s.PutStrStyled(x+2, y-2, "_\\", Blue)

		s.PutStrStyled(x-3, y-1, "\\_", Blue)
		s.PutStrStyled(x-1, y-1, "/", LightBlue)
		s.PutStrStyled(x, y-1, "|", Blue)
		s.PutStrStyled(x+1, y-1, "\\", LightBlue)
		s.PutStrStyled(x+2, y-1, "_/", Blue)

		s.PutStrStyled(x-1, y, "\\_/", Blue)
	}
}
