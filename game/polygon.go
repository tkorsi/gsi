package game

type Polygon struct {
	x      int
	y      int
	width  int
	height int
	color  uint16
}

func NewPolygon(x, y, width, height int, color uint16) *Polygon {
	p := Polygon{x: x, y: y, width: width, height: height, color: color}
	return &p
}

func (p *Polygon) Draw(s *Screen) {
	for i := 0; i < p.width; i++ {
		for j := 0; j < p.height; j++ {
			s.RenderCell(p.x+i, p.y+j, &Cell{
				BackgroundColor: p.color,
				Character:       ' ',
			})
		}
	}
}

func (p *Polygon) Update(ev Event) {}

func (p *Polygon) Size() (int, int) {
	return p.width, p.height
}

func (p *Polygon) Position() (int, int) {
	return p.x, p.y
}

func (p *Polygon) SetPosition(x, y int) {
	p.x = x
	p.y = y
}

func (p *Polygon) SetSize(w, h int) {
	p.width = w
	p.height = h
}

func (p *Polygon) Color() uint16 {
	return p.color
}

func (p *Polygon) SetColor(color uint16) {
	p.color = color
}
