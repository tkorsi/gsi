package spaceinvaders

import (
	_ "embed"

	g "github.com/tkorsi/gsi/game"
)

type Alien struct {
	*g.Entity
	IsAlive    bool
	IsRendered bool
	Points     int
}

type AlienType struct {
	Source []byte
	Points int
}

var (
	//go:embed files/alien_basic.txt
	alienBasicBytes []byte

	//go:embed files/alien_medium.txt
	alienMediumBytes []byte

	//go:embed files/alien_strong.txt
	alienStrongBytes []byte

	Basic  = AlienType{Source: alienBasicBytes, Points: 10}
	Medium = AlienType{Source: alienMediumBytes, Points: 20}
	Strong = AlienType{Source: alienStrongBytes, Points: 30}
)

func NewAlien(alienType AlienType) *Alien {
	canvas := CreateCanvas(alienType.Source)
	return &Alien{Entity: g.NewEntityFromCanvas(0, 0, canvas), IsAlive: true, Points: alienType.Points}
}

func CreateAliensLine(alienType AlienType, lineSize int) []*Alien {
	aliens := make([]*Alien, lineSize)
	for i := 0; i < lineSize; i++ {
		aliens[i] = NewAlien(alienType)
	}

	return aliens
}

func SetPositionAndRenderAliens(aliens [][]*Alien, baseLevel *g.BaseLevel, level *Level) {
	initialX, initialY, space := calcInitialPositionAndSpace(aliens, level)

	for index, line := range aliens {
		x := initialX

		for _, alien := range line {
			_, height := alien.Size()
			y := initialY + height*(index+1) - 2

			alien.SetPosition(x, y)
			alien.IsRendered = true

			baseLevel.AddEntity(alien)

			x += space
		}
	}
}

func calcInitialPositionAndSpace(aliens [][]*Alien, level *Level) (int, int, int) {
	lineSize := len(aliens[0])
	alienW, _ := aliens[0][0].Size()
	space := alienW + 1

	arenaX, arenaY := level.Position()
	arenaW, _ := level.Size()

	totalWidth := lineSize * space
	x := arenaX + arenaW/2 - totalWidth/2

	return x, arenaY, space
}

func (alien *Alien) Collide(collision g.Physical) {
	if _, ok := collision.(*Laser); ok {
		laser := collision.(*Laser)

		if laser.IsFromPlayer {
			laser.HasHit = true
			alien.IsAlive = false
		}
	}
}
