package game

type Canvas [][]Cell

func NewCanvas(width, heigh int) Canvas {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]Cell, heigh)
	}
	return canvas
}

func (canvas *Canvas) equals(other *Canvas) bool {
	first := *canvas
	second := *other
	if second == nil {
		return false
	}
	if len(first) != len(second) {
		return false
	}
	if len(first[0]) != len(second[0]) {
		return false
	}
	for i := range first {
		for j := range first[i] {
			equal := first[i][j].equals(&(second[i][j]))
			if !equal {
				return false
			}
		}
	}
	return true
}

type Cell struct {
	ForegroundColor uint16
	BackgroundColor uint16
	Character       rune
}

func (c *Cell) equals(c2 *Cell) bool {
	return c.ForegroundColor == c2.ForegroundColor &&
		c.BackgroundColor == c2.BackgroundColor &&
		c.Character == c2.Character
}

type Game struct {
	debug bool
	logs  []string
}

func NewGame() *Game {
	g := Game{
		logs: make([]string, 0),
	}
	return &g
}
