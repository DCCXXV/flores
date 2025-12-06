package main

import (
	"math/rand"
	"time"
	"github.com/spf13/pflag"

	"github.com/DCCXXV/flores/drawable"
	"github.com/gdamore/tcell/v2"
)


func main() {
    speed := pflag.IntP("speed", "s", 100, "velocidad en ms")
    pflag.Parse();


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

	for {
		select {
		case <-quit:
			return
		default:
		}

		rf := rand.Intn(4)
		rs := rand.Intn(3) + 1
		ry := rand.Intn(h-10)
		rx := rand.Intn(w)

		// TODO: single flower generation mode with a base
		//drawBase(s, w/2, h-8, 2)
		for y := h; y > h-ry; y-- {
		    tile, _, _ := s.Get(rx, y)
			if tile == " " {
				s.SetContent(rx, y, '|', nil, style)
			}
			s.Show()
			time.Sleep(time.Duration(*speed) * time.Millisecond)
		}
		// TODO: if flower is taller d should be higher
		drawable.DrawFlower(s, rx, h-ry, rs, rf)
		s.Show()
		time.Sleep(time.Duration(*speed) * time.Millisecond)
	}
}
