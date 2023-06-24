package spaceinvaders

import g "github.com/tkorsi/gsi/game"

type GameOver struct {
	*g.Entity
	EnteredZone bool
}

func CreateGameOver(level *Level, player *Player) *GameOver {
	_, heroH := player.Size()
	arenaX, arenaY := level.Position()
	arenaW, arenaH := level.Size()

	x := arenaX
	y := arenaY + arenaH - heroH
	w := arenaW
	h := heroH

	return &GameOver{Entity: g.NewEntity(x, y, w, h)}
}

func (gameOver *GameOver) Collide(collision g.Physical) {
	if _, ok := collision.(*Alien); ok && collision.(*Alien).IsAlive {
		gameOver.EnteredZone = true
	}
}
