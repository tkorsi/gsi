package game

import (
	"github.com/nsf/termbox-go"
)

type Screen struct {
	prevCanvas Canvas
	currCanvas Canvas
	width      int
	height     int
	fps        float64 // needs to be float for math
	level      Level
	Entities   []Drawable
	offsetx    int
	offsety    int
	delta      float64
}

func NewScreen() *Screen {
	s := Screen{}
	s.currCanvas = NewCanvas(10, 10)
	return &s
}

func (s *Screen) Update(ev Event) {
	if s.level != nil {
		s.level.Update(ev)
	}

	if ev.Type != EventNone {
		for _, e := range s.Entities {
			e.Update(ev)
		}
	}
}

func (s *Screen) Draw() {
	s.currCanvas = NewCanvas(s.width, s.height)
	if s.level != nil {
		s.level.DrawBackground(s)
		s.level.Draw(s)
	}
	for _, e := range s.Entities {
		e.Draw(s)
	}
	if !s.currCanvas.equals(&s.prevCanvas) {
		//termboxPixel(&s.currCanvas)
		termboxNormal(&s.currCanvas)
		termbox.Flush()
	}
	s.prevCanvas = s.currCanvas
}

func (s *Screen) resize(width, height int) {
	s.width = width
	s.height = height
	//s.height *= 2 // pixel
	canvas := NewCanvas(s.width, s.height)
	// Copy old data that fits
	for i := 0; i < min(s.width, len(s.currCanvas)); i++ {
		for j := 0; j < min(s.height, len(s.currCanvas[0])); j++ {
			canvas[i][j] = s.currCanvas[i][j]
		}
	}
	s.currCanvas = canvas
}

func (s *Screen) Size() (int, int) {
	return s.width, s.height
}

func (s *Screen) SetLevel(l Level) {
	s.level = l
}

func (s *Screen) Level() Level {
	return s.level
}

func (s *Screen) AddEntity(d Drawable) {
	s.Entities = append(s.Entities, d)
}

func (s *Screen) RemoveEntity(d Drawable) {
	for i, elem := range s.Entities {
		if elem == d {
			s.Entities = append(s.Entities[:i], s.Entities[i+1:]...)
			return
		}
	}
}

func termboxNormal(canvas *Canvas) {
	for i, col := range *canvas {
		for j, cell := range col {
			termbox.SetCell(i, j, cell.Character,
				termbox.Attribute(cell.ForegroundColor),
				termbox.Attribute(cell.BackgroundColor))
		}
	}

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

func (s *Screen) TimeDelta() float64 {
	return s.delta
}

func (s *Screen) SetFps(f float64) {
	s.fps = f
}

func (s *Screen) offset() (int, int) {
	return s.offsetx, s.offsety
}

func (s *Screen) setOffset(x, y int) {
	s.offsetx, s.offsety = x, y
}

func (s *Screen) RenderCell(x, y int, c *Cell) {
	newx := x + s.offsetx
	newy := y + s.offsety
	if newx >= 0 && newx < len(s.currCanvas) &&
		newy >= 0 && newy < len(s.currCanvas[0]) {
		renderCell(&s.currCanvas[newx][newy], c)
	}
}

func renderCell(old, new_ *Cell) {
	if new_.Character != 0 {
		old.Character = new_.Character
	}
	if new_.BackgroundColor != 0 {
		old.BackgroundColor = new_.BackgroundColor
	}
	if new_.ForegroundColor != 0 {
		old.ForegroundColor = new_.ForegroundColor
	}
}
