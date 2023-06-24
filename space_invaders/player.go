package spaceinvaders

import (
	_ "embed"
	"time"

	g "github.com/tkorsi/gsi/game"
)

type Player struct {
	*g.Entity
	Level   *Level
	Lasers  []*Laser
	IsAlive bool
}

//go:embed files/player.txt
var playerBytes []byte

func NewPlayer(level *Level) *Player {
	c := CreateCanvas(playerBytes)
	x, y := setPlayerPosition(level, c)

	return &Player{Entity: g.NewEntityFromCanvas(x, y, c), Level: level, IsAlive: true}
}

func setPlayerPosition(level *Level, playerCanvas g.Canvas) (int, int) {
	levelX, levelY := level.Position()
	levelW, levelH := level.Size()

	x := levelX + levelW/2 - len(playerCanvas)/2
	y := levelY + levelH - len(playerCanvas[0])

	return x, y
}

func (p *Player) Update(ev g.Event) {
	if p.IsAlive == false {
		return
	}

	if ev.Type == g.EventKey {
		x, y := p.Position()
		playerWidth, _ := p.Size()

		switch ev.Key {
		case g.KeyArrowLeft:
			if x > p.Level.Init {
				x = x - 1
				p.SetPosition(x, y)
			}
		case g.KeyArrowRight:
			if x < p.Level.End-playerWidth-1 {
				x = x + 1
				p.SetPosition(x, y)
			}
		case g.KeySpace:
			p.shoot()
		}
	}
}

func (p *Player) shoot() {
	x, y := p.Position()

	if p.isReloading() {
		return
	}

	playerWidth, _ := p.Size()
	playerGunPosition := x + (playerWidth-1)/2
	distanceToPlayer := y - 1

	laser := NewPlayerLaser(playerGunPosition, distanceToPlayer)
	p.Lasers = append(p.Lasers, laser)
}

func (p *Player) isReloading() bool {
	return len(p.Lasers) > 0
}

func (p *Player) Collide(collision g.Physical) {
	if _, ok := collision.(*Laser); ok {
		laser := collision.(*Laser)

		laser.HasHit = true
		p.IsAlive = false
	}
}

func (p *Player) IsDead() bool {
	return p.IsAlive == false
}

func (p *Player) animatePlayerEndGame(level *g.BaseLevel) {
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			level.RemoveEntity(p)
		} else {
			level.AddEntity(p)
		}

		time.Sleep(450 * time.Millisecond)
	}
}
