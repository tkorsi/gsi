package spaceinvaders

import g "github.com/tkorsi/gsi/game"

type Laser struct {
	*g.Polygon
	Direction     int
	IsFromPlayer  bool
	IsNew         bool
	HasHit        bool
	HitAlienLaser bool
	Points        int
}

func NewPlayerLaser(playerGunPosition int, y int) *Laser {
	return &Laser{
		Polygon:      g.NewPolygon(playerGunPosition, y, 1, 1, g.ColorRed),
		Direction:    1,
		IsNew:        true,
		IsFromPlayer: true,
	}
}

func NewAlienLaser(alienGunPosition int, y int) *Laser {
	return &Laser{
		Polygon:      g.NewPolygon(alienGunPosition, y, 1, 1, g.ColorGreen),
		Direction:    -1,
		IsNew:        true,
		IsFromPlayer: false,
		Points:       5,
	}
}

func (laser *Laser) Collide(collision g.Physical) {
	if laser.IsFromPlayer == false {
		return
	}

	if laserCollide, isLaser := collision.(*Laser); isLaser {
		laser.HasHit = true
		laser.HitAlienLaser = true
		laser.Points = laserCollide.Points

		laserCollide.HasHit = true
	}
}
