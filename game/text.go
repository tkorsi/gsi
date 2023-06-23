package game

type Text struct {
	x               int
	y               int
	foregroundColor uint16
	backgroundColor uint16
	text            []rune
	canvas          []Cell
}

func NewText(x, y int, text string, fg, bg uint16) *Text {
	str := []rune(text)
	c := make([]Cell, len(str))
	for i := range c {
		c[i] = Cell{Character: str[i], ForegroundColor: fg, BackgroundColor: bg}
	}
	return &Text{
		x:               x,
		y:               y,
		foregroundColor: fg,
		backgroundColor: bg,
		text:            str,
		canvas:          c,
	}
}

func (t *Text) Update(ev Event) {}

func (t *Text) Draw(s *Screen) {
	w, _ := t.Size()
	for i := 0; i < w; i++ {
		s.RenderCell(t.x+i, t.y, &t.canvas[i])
	}
}

func (t *Text) Position() (int, int) {
	return t.x, t.y
}

func (t *Text) Size() (int, int) {
	return len(t.text), 1
}

func (t *Text) SetPosition(x, y int) {
	t.x = x
	t.y = y
}

func (t *Text) Text() string {
	return string(t.text)
}

func (t *Text) SetText(text string) {
	t.text = []rune(text)
	c := make([]Cell, len(t.text))
	for i := range c {
		c[i] = Cell{Character: t.text[i], ForegroundColor: t.foregroundColor, BackgroundColor: t.backgroundColor}
	}
	t.canvas = c
}

func (t *Text) Color() (uint16, uint16) {
	return t.foregroundColor, t.backgroundColor
}

func (t *Text) SetColor(fg, bg uint16) {
	t.foregroundColor = fg
	t.backgroundColor = bg
	for i := range t.canvas {
		t.canvas[i].ForegroundColor = fg
		t.canvas[i].BackgroundColor = bg
	}
}
