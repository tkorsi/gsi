package game

type Level interface {
	DrawBackground(*Screen)
	AddEntity(Drawable)
	RemoveEntity(Drawable)
	Draw(*Screen)
	Update(Event)
}

type BaseLevel struct {
	Entities []Drawable
	bg       Cell
	offsetx  int
	offsety  int
}

func NewBaseLevel(bg Cell) *BaseLevel {
	level := BaseLevel{Entities: make([]Drawable, 0), bg: bg}
	return &level
}

func (bl *BaseLevel) Update(ev Event) {
	for _, e := range bl.Entities {
		e.Update(ev)
	}

	colls := make([]Physical, 0)
	dynamics := make([]DynamicPhysical, 0)
	for _, e := range bl.Entities {
		if p, ok := interface{}(e).(Physical); ok {
			colls = append(colls, p)
		}
		if p, ok := interface{}(e).(DynamicPhysical); ok {
			dynamics = append(dynamics, p)
		}

	}
	jobs := make(chan DynamicPhysical, len(dynamics))
	results := make(chan int, len(dynamics))
	for w := 0; w <= len(dynamics)/3; w++ {
		go checkCollisionsWorker(colls, jobs, results)
	}
	for _, p := range dynamics {
		jobs <- p
	}
	close(jobs)
	for r := 0; r < len(dynamics); r++ {
		<-results
	}
}

func (bl *BaseLevel) DrawBackground(s *Screen) {
	for i, row := range s.currCanvas {
		for j := range row {
			s.currCanvas[i][j] = bl.bg
		}
	}
}

func (bl *BaseLevel) Draw(s *Screen) {
	offx, offy := s.offset()
	s.setOffset(bl.offsetx, bl.offsety)
	for _, e := range bl.Entities {
		e.Draw(s)
	}
	s.setOffset(offx, offy)
}

func checkCollisionsWorker(ps []Physical, jobs <-chan DynamicPhysical, results chan<- int) {
	for p := range jobs {
		for _, c := range ps {
			if c == p {
				continue
			}
			px, py := p.Position()
			cx, cy := c.Position()
			pw, ph := p.Size()
			cw, ch := c.Size()
			if px < cx+cw && px+pw > cx &&
				py < cy+ch && py+ph > cy {
				p.Collide(c)
			}
		}
		results <- 1
	}
}

func (l *BaseLevel) AddEntity(d Drawable) {
	l.Entities = append(l.Entities, d)
}

func (l *BaseLevel) RemoveEntity(d Drawable) {
	for i, elem := range l.Entities {
		if elem == d {
			l.Entities = append(l.Entities[:i], l.Entities[i+1:]...)
			return
		}
	}
}

func (l *BaseLevel) Offset() (int, int) {
	return l.offsetx, l.offsety
}

func (l *BaseLevel) SetOffset(x, y int) {
	l.offsetx, l.offsety = x, y
}
