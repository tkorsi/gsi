package game

import "github.com/nsf/termbox-go"

type Screen struct {
	prevCanvas Canvas
	currCanvas Canvas
	width      int
	height     int
	fps        int
}

func NewScreen() *Screen {
	s := Screen{}
	s.currCanvas = NewCanvas(10, 10)
	return &s
}

// func (s *Screen) Update(ev Event) {
// 	if ev.Type != EventNone {

// 	}
// }

func (s *Screen) Draw() {
	s.currCanvas = NewCanvas(s.width, s.height)
	if !s.currCanvas.equals(&s.prevCanvas) {
		termboxPixel(&s.currCanvas)
		termbox.Flush()
	}
	s.prevCanvas = s.currCanvas
}

func termboxPixel(c *Canvas) {
	for i, col := range *c {
		for j := 0; j < len(col); j += 2 {
			cellBack := col[j]
			cellFront := col[j+1]
			termj := j / 2
			char := '\u2584'
			if cellFront.BackgroundColor == 0 {
				char = 0
			}
			termbox.SetCell(i, termj, char,
				termbox.Attribute(cellFront.BackgroundColor),
				termbox.Attribute(cellBack.BackgroundColor))
		}
	}
}
