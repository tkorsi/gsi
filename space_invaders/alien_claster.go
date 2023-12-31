package spaceinvaders

import (
	"math/rand"
	"time"

	g "github.com/tkorsi/gsi/game"
)

type AlienCluster struct {
	Aliens                   [][]*Alien
	Lasers                   []*Laser
	TimeToMove               float64
	WaitingTime              float64
	WaitingTimeToMoveNextRow float64
	CurrentRowMoving         int
	Direction                int
	MoveSize                 int
	IsMoving                 bool
	IsMovingDown             bool
	IsAllDead                bool
	ReachedEndLevel          bool
}

func NewAlienCluster() *AlienCluster {
	lineSize := 13

	alienCluster := AlienCluster{
		TimeToMove:               0,
		WaitingTime:              1,
		WaitingTimeToMoveNextRow: 0.07,
		CurrentRowMoving:         -1,
		Direction:                3,
		MoveSize:                 3,
		IsMoving:                 false,
		ReachedEndLevel:          false,
		IsMovingDown:             false,
	}

	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Strong, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Medium, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Medium, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Basic, lineSize))
	alienCluster.Aliens = append(alienCluster.Aliens, CreateAliensLine(Basic, lineSize))

	return &alienCluster
}

func (alienCluster *AlienCluster) UpdateAliensPositions(timeDelta float64, level *Level) {
	alienCluster.prepareMovement()

	if alienCluster.canMove(timeDelta) {
		alienCluster.move(level)
	}
}

func (alienCluster *AlienCluster) prepareMovement() {
	if alienCluster.CurrentRowMoving == -1 {
		alienCluster.CurrentRowMoving = len(alienCluster.Aliens) - 1
		alienCluster.IsMoving = false
	}

	if alienCluster.ReachedEndLevel {
		alienCluster.changeDirection()
	}
}

func (alienCluster *AlienCluster) changeDirection() {
	alienCluster.Direction *= -1
	alienCluster.ReachedEndLevel = false
}

func (alienCluster *AlienCluster) canMove(timeDelta float64) bool {
	alienCluster.TimeToMove += timeDelta
	if alienCluster.isWaitingToMoveFirstRow() || alienCluster.isWaitingToMoveNextRow() {
		return false
	}
	return true
}

func (alienCluster *AlienCluster) isWaitingToMoveFirstRow() bool {
	return alienCluster.IsMoving == false && alienCluster.TimeToMove < alienCluster.WaitingTime
}

func (alienCluster *AlienCluster) isWaitingToMoveNextRow() bool {
	return alienCluster.IsMoving && alienCluster.TimeToMove < alienCluster.WaitingTimeToMoveNextRow
}

func (alienCluster *AlienCluster) move(level *Level) {
	alienCluster.TimeToMove = 0
	alienCluster.IsMoving = true

	if alienCluster.IsMovingDown {
		alienCluster.moveDown()
	} else {
		alienCluster.moveSideways(level)
	}

	alienCluster.CurrentRowMoving -= 1
}

func (alienCluster *AlienCluster) moveDown() {
	row := alienCluster.getCurrentLine()

	for _, alien := range row {
		x, y := alien.Position()
		alien.SetPosition(x, y+alienCluster.MoveSize)
	}

	finishedMovingRows := alienCluster.CurrentRowMoving == 0
	if finishedMovingRows {
		alienCluster.IsMovingDown = false
	}
}

func (alienCluster *AlienCluster) moveSideways(level *Level) {
	row := alienCluster.getCurrentLine()

	for _, alien := range row {
		x, y := alien.Position()
		w, _ := alien.Size()
		x = x + alienCluster.Direction

		alien.SetPosition(x, y)

		if alien.IsAlive {
			alienCluster.checkIfReachedEndOfLevel(x, w, level)
		}
	}
}

func (alienCluster *AlienCluster) getCurrentLine() []*Alien {
	row := alienCluster.Aliens[alienCluster.CurrentRowMoving]
	return row
}

func (alienCluster *AlienCluster) checkIfReachedEndOfLevel(x int, w int, level *Level) {
	if alienCluster.CurrentRowMoving == 0 && (x+w >= level.End-alienCluster.MoveSize || x <= level.Init+alienCluster.MoveSize) {
		alienCluster.ReachedEndLevel = true
		alienCluster.IsMovingDown = true
	}
}

func (alienCluster *AlienCluster) RemoveDeadAliensAndGetPoints(level *g.BaseLevel) int {
	hasAtLeastOneAlive := false
	points := 0

	for _, alienRow := range alienCluster.Aliens {
		for _, alien := range alienRow {
			if alien.IsAlive == false && alien.IsRendered == true {
				points += alien.Points
				alien.IsRendered = false

				level.RemoveEntity(alien)
			}

			if alien.IsAlive == true {
				hasAtLeastOneAlive = true
			}
		}
	}

	alienCluster.IsAllDead = !hasAtLeastOneAlive

	return points
}

func (alienCluster *AlienCluster) Shoot() {
	if alienCluster.canShoot() {
		alien := alienCluster.selectRandomAlien()

		if alien == nil {
			return
		}

		x, y := alien.Position()
		width, _ := alien.Size()
		alienGunPosition := x + (width-1)/2
		distanceToAlien := y + 2

		laser := NewAlienLaser(alienGunPosition, distanceToAlien)
		alienCluster.Lasers = append(alienCluster.Lasers, laser)
	}
}

func (alienCluster *AlienCluster) canShoot() bool {
	return len(alienCluster.Lasers) == 0
}

func (alienCluster *AlienCluster) selectRandomAlien() *Alien {
	var shooterAlien *Alien

	rowSize := len(alienCluster.Aliens) - 1
	col := alienCluster.selectRandomAlienColumn()

	for row := rowSize; row >= 0; row-- {
		alien := alienCluster.Aliens[row][col]

		if alien.IsAlive {
			shooterAlien = alien
			break
		}
	}
	return shooterAlien
}

func (alienCluster *AlienCluster) selectRandomAlienColumn() int {
	rand.Seed(time.Now().UnixNano())
	col := rand.Intn(len(alienCluster.Aliens[0]))
	return col
}

func (alienCluster *AlienCluster) IsAllAliensDead() bool {
	return alienCluster.IsAllDead
}
