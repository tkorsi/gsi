package main

import tl "github.com/tkorsi/gsi/game"

type CollRec struct {
	*tl.Polygon
	move bool
	px   int
	py   int
}

func NewCollRec(x, y, w, h int, color uint16, move bool) *CollRec {
	return &CollRec{
		Polygon: tl.NewPolygon(x, y, w, h, color),
		move:    move,
	}
}

func (r *CollRec) Update(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey && r.move {
		r.px, r.py = r.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			r.SetPosition(r.px+1, r.py)
		case tl.KeyArrowLeft:
			r.SetPosition(r.px-1, r.py)
		case tl.KeyArrowUp:
			r.SetPosition(r.px, r.py-1)
		case tl.KeyArrowDown:
			r.SetPosition(r.px, r.py+1)
		}
	}
}

func (r *CollRec) Collide(p tl.Physical) {
	// Check if it's a CollRec we're colliding with
	if _, ok := p.(*CollRec); ok && r.move {
		r.SetColor(tl.ColorBlue)
		r.SetPosition(r.px, r.py)
	}
}

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(60)
	l := tl.NewBaseLevel(tl.Cell{
		BackgroundColor: tl.ColorWhite,
	})
	l.AddEntity(NewCollRec(3, 3, 3, 3, tl.ColorRed, true))
	l.AddEntity(NewCollRec(7, 4, 3, 3, tl.ColorGreen, false))
	g.Screen().SetLevel(l)
	//g.Screen().AddEntity(tl.NewFpsText(0, 0, tl.ColorRed, tl.ColorDefault, 0.5))
	g.Start()
}
