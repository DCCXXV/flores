package main

import (
	"math/rand"
	"time"

	"github.com/spf13/pflag"

	"github.com/DCCXXV/flores/drawable"
	"github.com/gdamore/tcell/v2"
)

func main() {
	speed := pflag.IntP("speed", "s", 25, "time in ms of sleep between actions")
	single := pflag.Bool("single", false, "single flower in a flowerpot")
	straight := pflag.Bool("straight", false, "straight flower(s)")

	pflag.Parse()

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
				if ev.Rune() == 'q' || ev.Key() == tcell.KeyEscape {
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
			rf := rand.Intn(4)
			rs := rand.Intn(3) + 1
			x := w / 2
			for y := h - 6; y > h/2; y-- {
				if !*straight {
					x += rand.Intn(3) - 1
				}
				s.SetContent(x, y, '|', nil, style)
				s.Show()
				if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
					return
				}
			}

			drawable.DrawFlower(s, w/2, h/2, rs, rf)
			s.Show()

			if interruptibleSleep(20 * time.Duration(*speed) * time.Millisecond) {
				return
			}

			for y := h/2 - 8; y < h-6; y++ {
				s.PutStr(w/2-8, y, "                ")
			}
		}
	} else {
		for {
			select {
			case <-quit:
				return
			default:
			}

			rf := rand.Intn(4)
			rs := rand.Intn(3) + 1
			ry := rand.Intn(h - 10)
			rx := rand.Intn(w)

			for y := h; y > h-ry; y-- {
				tile, _, _ := s.Get(rx, y)
				if tile == " " {

					if !*straight {
						rx += rand.Intn(3) - 1
					}

					// leaf
					if rand.Intn(15) == 0 {
						if rand.Intn(2) == 0 {
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
			drawable.DrawFlower(s, rx, h-ry, rs, rf)
			s.Show()
			if interruptibleSleep(time.Duration(*speed) * time.Millisecond) {
				return
			}
		}
	}
}
