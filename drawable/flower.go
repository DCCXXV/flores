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
	colors := GetColors()
	switch t {
	case Flowers["daisy"]:
		s.SetContent(x, y, '@', nil, colors.Yellow)
		horizontalPetal := strings.Repeat("==", d)
		s.PutStrStyled(x+1, y, horizontalPetal, colors.White)
		s.PutStrStyled(x-len(horizontalPetal), y, horizontalPetal, colors.White)
		for i := 1; i < d+1; i++ {
			s.PutStrStyled(x-i, y-i, "\\", colors.White)
			s.PutStrStyled(x+i, y-i, "/", colors.White)
			s.PutStrStyled(x-i, y+i, "/", colors.White)
			s.PutStrStyled(x+i, y+i, "\\", colors.White)
		}
	case Flowers["rose"]:
		petalColors := []tcell.Style{colors.Red, colors.LightBlue, colors.Pale}
		randomColor := petalColors[random.IntN(len(petalColors))]

		s.PutStrStyled(x-1, y, utils.GenerateRandomString(random, 3), randomColor)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomString(random, 3), randomColor)
		s.PutStrStyled(x-2, y-2, utils.GenerateRandomString(random, 5), randomColor)
		s.PutStrStyled(x-2, y-3, utils.GenerateRandomString(random, 5), randomColor)
		s.PutStrStyled(x-2, y-4, utils.GenerateRandomString(random, 5), randomColor)
	case Flowers["sunflower"]:
		if d <= 1 {
			s.PutStrStyled(x, y-1, "#", colors.Brown)
			s.PutStrStyled(x-1, y-1, "o", colors.Yellow)
			s.PutStrStyled(x+1, y-1, "o", colors.Yellow)
			s.PutStrStyled(x-1, y, "ooo", colors.Yellow)
			s.PutStrStyled(x-1, y-2, "ooo", colors.Yellow)
		} else if d == 2 {
			s.PutStrStyled(x, y, "o", colors.Brown)
			s.PutStrStyled(x-1, y-1, "ooo", colors.Brown)
			s.PutStrStyled(x, y-2, "o", colors.Brown)
			s.PutStrStyled(x-2, y-2, utils.GenerateRandomString(random, 2), colors.Yellow)
			s.PutStrStyled(x+1, y-2, utils.GenerateRandomString(random, 2), colors.Yellow)
			s.PutStrStyled(x-2, y, utils.GenerateRandomString(random, 2), colors.Yellow)
			s.PutStrStyled(x+1, y, utils.GenerateRandomString(random, 2), colors.Yellow)
			s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(random, 3), colors.Yellow)
			s.PutStrStyled(x-1, y-3, utils.GenerateRandomString(random, 3), colors.Yellow)
			s.PutStrStyled(x-3, y-1, utils.GenerateRandomString(random, 2), colors.Yellow)
			s.PutStrStyled(x+2, y-1, utils.GenerateRandomString(random, 2), colors.Yellow)
		} else {
			s.PutStrStyled(x-1, y, "@@@", colors.Brown)
			s.PutStrStyled(x-2, y-1, "@@@@@", colors.Brown)
			s.PutStrStyled(x-1, y-2, "@@@", colors.Brown)

			s.PutStrStyled(x-4, y+1, "&  &&@", colors.Yellow)
			s.PutStrStyled(x-1, y+2, "&  &", colors.Yellow)
			s.PutStrStyled(x-2, y-3, "&&%%&", colors.Yellow)
			s.PutStrStyled(x-1, y-4, "& & ", colors.Yellow)
			s.PutStrStyled(x-3, y, "%&", colors.Yellow)
			s.PutStrStyled(x+2, y, "&&&", colors.Yellow)
			s.PutStrStyled(x-3, y-1, "&", colors.Yellow)
			s.PutStrStyled(x+3, y-1, "&", colors.Yellow)
			s.PutStrStyled(x-4, y-2, "&$&", colors.Yellow)
			s.PutStrStyled(x+2, y-2, "&#&", colors.Yellow)
		}
	case Flowers["lavender"]:
		s.PutStrStyled(x-1, y, "\\|/", colors.DarkGreen)
		s.PutStrStyled(x-2, y-1, "\\ | /", colors.DarkGreen)
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
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomTile(random, "o", "O", "0"), colors.Violet)
		s.PutStrStyled(x+1, y-1, utils.GenerateRandomTile(random, "o", "O", "0"), colors.Violet)
		s.PutStrStyled(x-1, y-2, "o o", colors.Violet)
		s.PutStrStyled(x-1, y-3, "o o", colors.Violet)
		s.PutStrStyled(x-1, y-4, " 0 ", colors.Violet)
	case Flowers["poppy"]:
		s.PutStrStyled(x, y, utils.GenerateRandomTile(random), colors.Brown)
		s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(random, 3), colors.Red)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomString(random, 3), colors.Red)
		s.PutStrStyled(x-2, y, utils.GenerateRandomString(random, 2), colors.Red)
		s.PutStrStyled(x+1, y, utils.GenerateRandomString(random, 2), colors.Red)
	case Flowers["hibiscus"]:
		petalColors := []tcell.Style{colors.Red, colors.Violet, colors.Blue}
		randomColor := petalColors[random.IntN(len(petalColors))]

		s.PutStrStyled(x, y, "|", colors.White)
		s.PutStrStyled(x-1, y+1, utils.GenerateRandomString(random, 3), randomColor)
		s.PutStrStyled(x-1, y-1, utils.GenerateRandomTile(random), randomColor)
		s.PutStrStyled(x, y-1, "%", colors.Yellow)
		s.PutStrStyled(x+1, y-1, utils.GenerateRandomTile(random), randomColor)
		s.PutStrStyled(x-2, y, utils.GenerateRandomString(random, 2), randomColor)
		s.PutStrStyled(x+1, y, utils.GenerateRandomString(random, 2), randomColor)
	case Flowers["bluet"]:
		s.PutStrStyled(x, y-4, "_", colors.Blue)
		s.PutStrStyled(x-2, y-3, "_/|\\_", colors.Blue)

		s.PutStrStyled(x-3, y-2, "/_", colors.Blue)
		s.PutStrStyled(x-1, y-2, "\\", colors.LightBlue)
		s.PutStrStyled(x, y-2, "@", colors.Yellow)
		s.PutStrStyled(x+1, y-2, "/", colors.LightBlue)
		s.PutStrStyled(x+2, y-2, "_\\", colors.Blue)

		s.PutStrStyled(x-3, y-1, "\\_", colors.Blue)
		s.PutStrStyled(x-1, y-1, "/", colors.LightBlue)
		s.PutStrStyled(x, y-1, "|", colors.Blue)
		s.PutStrStyled(x+1, y-1, "\\", colors.LightBlue)
		s.PutStrStyled(x+2, y-1, "_/", colors.Blue)

		s.PutStrStyled(x-1, y, "\\_/", colors.Blue)
	}
}
