package main

import (
	"fmt"
	"os"

	"github.com/tkorsi/gsi/game"
	tl "github.com/tkorsi/gsi/game"
)

type MovingText struct {
	*tl.Text
}

func (m *MovingText) Update(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey {
		x, y := m.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			x += 1
		case tl.KeyArrowLeft:
			x -= 1
		case tl.KeyArrowUp:
			y -= 1
		case tl.KeyArrowDown:
			y += 1
		}
		m.SetPosition(x, y)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide text")
		return
	}
	g := game.NewGame()
	g.Screen().AddEntity(&MovingText{tl.NewText(0, 0, os.Args[1], tl.ColorWhite, tl.ColorBlue)})
	g.Start()
}
