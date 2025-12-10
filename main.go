package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/spf13/pflag"

	"github.com/DCCXXV/flores/drawable"
	"github.com/gdamore/tcell/v2"
)

func main() {
	floresSeed := rand.Uint64()

	speed := pflag.Uint64P("speed", "s", 50, "time in ms of sleep between actions")
	single := pflag.BoolP("single", "i", false, "single flower in a flowerpot (it is recommended to change speed")
	straight := pflag.BoolP("straight", "t", false, "straight flower(s)")
	seed := pflag.Uint64P("seed", "e", floresSeed, "use a seed instead of randomizing")
	setFlowers := pflag.StringSliceP("flowers", "f", nil, "set specific flowers to draw")
	listFlowers := pflag.BoolP("list", "l", false, "list available flowers")
	//trueColor := pflag.BoolP("true-color", "c", false, "use true color")

	pflag.Parse()

	if *listFlowers {
		fmt.Println("Available flowers:")
		for flower, _ := range drawable.Flowers {
			fmt.Println("-", flower)
		}
		return
	}

	// probably a better way to do it
	random := rand.New(rand.NewPCG(*seed, *seed))

	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()
	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	w, h := s.Size()

	quit := make(chan any)

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Rune() == 'q' || ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(quit)
					return
				}
			}
		}
	}()

	interruptibleSleep := func(d time.Duration) bool {
		select {
		case <-quit:
			return true
		case <-time.After(d):
			return false
		}
	}

	if *single {
		drawable.DrawBase(s, w/2, h-6)
		for {
			select {
			case <-quit:
				return
			default:
			}

			rf, rs, _, _ := randomize(random, *setFlowers, h, w)

			x := w / 2
			for y := h - 6; y > h/2; y-- {
				if !*straight {
					x += random.IntN(3) - 1
				}
				s.SetContent(x, y, '|', nil, style)
				s.Show()
				if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
					return
				}
			}

			drawable.DrawFlower(s, random, x, h/2, rs, rf)
			s.Show()

			if interruptibleSleep(20 * time.Duration(*speed) * time.Millisecond) {
				return
			}

			for y := h/2 - 8; y <= h-6; y++ {
				s.PutStr(w/2-12, y, "                        ")
			}
		}
	} else {
		for {
			select {
			case <-quit:
				return
			default:
			}

			rf, rs, ry, rx := randomize(random, *setFlowers, h, w)

			for y := h; y > h-ry; y-- {
				tile, _, _ := s.Get(rx, y)
				if tile == " " {

					if !*straight {
						rx += random.IntN(3) - 1
					}

					// leaf
					if random.IntN(15) == 0 {
						if random.IntN(2) == 0 { // direction of the leaf
							go func() {
								x := rx
								if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
									return
								}
								s.SetContent(x+1, y, '<', nil, style)
								if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
									return
								}
								s.SetContent(x+2, y, '=', nil, style)
								if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
									return
								}
								s.SetContent(x+3, y, '>', nil, style)
							}()
						} else {
							go func() {
								x := rx
								if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
									return
								}
								s.SetContent(x-1, y, '>', nil, style)
								if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
									return
								}
								s.SetContent(x-2, y, '=', nil, style)
								if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
									return
								}
								s.SetContent(x-3, y, '<', nil, style)
							}()
						}
					}
					s.SetContent(rx, y, '|', nil, style)
				}
				s.Show()
				if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
					return
				}
			}
			// TODO: if flower is taller d should be higher
			drawable.DrawFlower(s, random, rx, h-ry, rs, rf)
			s.Show()
			if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
				return
			}
		}
	}
}

func randomize(random *rand.Rand, flowers []string, h, w int) (int, int, int, int) {
	rf := random.IntN(drawable.TotalFlowers)

	if flowers != nil {
		newFlowers := []int{}
		for _, flower := range flowers {
			newFlowers = append(newFlowers, drawable.Flowers[flower])
		}
		aux := random.IntN(len(flowers))
		rf = newFlowers[aux]
	}

	rs := random.IntN(3) + 1
	ry := random.IntN(h - 10)
	rx := random.IntN(w)

	return rf, rs, ry, rx
}
