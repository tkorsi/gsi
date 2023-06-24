package spaceinvaders

import (
	"time"

	g "github.com/tkorsi/gsi/game"
)

type Invaders struct {
	*g.Entity
	Game               *g.Game
	Level              *g.Level
	AlienLaserVelocity float64
	TimeDelta          float64
	RefreshSpeed       time.Duration
	Score              int
}

func NewEngine() *Invaders {
	e := Invaders{
		Entity:             g.NewEntity(0, 0, 1, 1),
		Game:               g.NewGame(),
		Level:              tl.NewBaseLevel(tl.Cell{Bg: tl.ColorBlack, Fg: tl.ColorWhite}),
		AlienLaserVelocity: 0.04,
		RefreshSpeed:       20,
		Score:              0,
	}

	e.Game.Screen().SetFps(60)
	e.Level.AddEntity(&e)

	return &e
}
