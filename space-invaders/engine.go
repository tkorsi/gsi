package spaceinvaders

import (
	"time"

	g "github.com/tkorsi/gsi/game"
)

type Invaders struct {
	*g.Entity
	Game               *g.Game
	Level              *g.BaseLevel
	AlienCluster       *AlienCluster
	AlienLaserVelocity float64
	TimeDelta          float64
	RefreshSpeed       time.Duration
	Score              int
	Started            bool
	ScreenSizeNotOK    bool
}

func NewEngine() *Invaders {
	e := Invaders{
		Entity:             g.NewEntity(0, 0, 1, 1),
		Game:               g.NewGame(),
		Level:              g.NewBaseLevel(g.Cell{BackgroundColor: g.ColorBlack, ForegroundColor: g.ColorWhite}),
		AlienLaserVelocity: 0.04,
		RefreshSpeed:       20,
		Score:              0,
	}

	e.Game.Screen().SetFps(60)
	e.Level.AddEntity(&e)

	return &e
}

func (e *Invaders) Start() {
	e.Game.Start()
}

func (e *Invaders) Update(ev g.Event) {
	if e.Started == false && e.ScreenSizeNotOK == false && ev.Type == g.EventKey && ev.Key == g.KeyEnter {
		go e.initializeGame()
	}
}

func (e *Invaders) initializeGame() {
	e.Started = true

	e.initHero()
	e.initAliens()
	e.initGameOverZone()
	e.gameLoop()
}

func (e *Invaders) initLevel() {
	screenWidth, screenHeight := e.getScreenSize()
	e.Level = xN(screenWidth, screenHeight)
	invaders.Level.AddEntity(invaders.level)
}

func (invaders *Invaders) initHud() {
	invaders.Hud = NewHud(invaders.level, invaders.Level)
}

func (invaders *Invaders) getScreenSize() (int, int) {
	screenWidth, screenHeight := invaders.Game.Screen().Size()

	for screenWidth == 0 && screenHeight == 0 {
		time.Sleep(100 * time.Millisecond)
		screenWidth, screenHeight = invaders.Game.Screen().Size()
	}

	return screenWidth, screenHeight
}

func (invaders *Invaders) initHero() {
	invaders.Hero = NewHero(invaders.level)
	invaders.Level.AddEntity(invaders.Hero)
}

func (invaders *Invaders) initAliens() {
	invaders.AlienCluster = NewAlienCluster()
	SetPositionAndRenderAliens(invaders.AlienCluster.Aliens, invaders.Level, invaders.level)
}

func (invaders *Invaders) initGameOverZone() {
	invaders.GameOverZone = CreateGameOverZone(invaders.level, invaders.Hero)
	invaders.Level.AddEntity(invaders.GameOverZone)
}

func (invaders *Invaders) gameLoop() {
	for {
		if invaders.Hero.IsDead() || invaders.AlienCluster.IsAllAliensDead() {
			invaders.Hero.animateHeroEndGame(invaders.Level)
			invaders.Started = false
			break
		}

		invaders.updateLaserPositions()
		invaders.RemoveDeadAliensAndIncrementScore()
		invaders.updateAlienClusterPosition()
		invaders.updateScore()
		invaders.verifyGameOverZone()

		time.Sleep(invaders.RefreshSpeed * time.Millisecond)
	}

	if invaders.Hero.IsDead() {
		ShowGameOverScreen(invaders)
	}

	if invaders.AlienCluster.IsAllAliensDead() {
		invaders.initializeGame()
	}
}

func (invaders *Invaders) updateScore() {
	invaders.Hud.UpdateScore(invaders.Score)
}

func (invaders *Invaders) updateAlienClusterPosition() {
	invaders.AlienCluster.UpdateAliensPositions(invaders.Game.Screen().TimeDelta(), invaders.level)
	invaders.AlienCluster.Shoot()
}

func (invaders *Invaders) RemoveDeadAliensAndIncrementScore() {
	points := invaders.AlienCluster.RemoveDeadAliensAndGetPoints(invaders.Level)
	invaders.addScore(points)
}

func (invaders *Invaders) updateLaserPositions() {
	invaders.updateHeroLasers()
	invaders.updateAlienLasers()
	invaders.removeLasers()
}

func (invaders *Invaders) updateHeroLasers() {
	invaders.updateLasers(invaders.Hero.Lasers)
}

func (invaders *Invaders) updateAlienLasers() {
	invaders.TimeDelta += invaders.Game.Screen().TimeDelta()

	if invaders.TimeDelta >= invaders.AlienLaserVelocity {
		invaders.TimeDelta = 0
		invaders.updateLasers(invaders.AlienCluster.Lasers)
	}
}

func (invaders *Invaders) updateLasers(lasers []*Laser) {
	for _, laser := range lasers {
		if laser.IsNew {
			invaders.renderNewLaser(laser)
			continue
		}

		x, y := laser.Position()
		laser.SetPosition(x, y-laser.Direction)
	}
}

func (invaders *Invaders) renderNewLaser(laser *Laser) {
	laser.IsNew = false
	invaders.Level.AddEntity(laser)
}

func (invaders *Invaders) removeLasers() {
	_, levelY := invaders.level.Position()
	_, levelH := invaders.level.Size()

	upperLimit := levelY
	bottomLimit := levelY + levelH - 1

	invaders.Hero.Lasers = invaders.removeLaserOf(invaders.Hero.Lasers, upperLimit)
	invaders.AlienCluster.Lasers = invaders.removeLaserOf(invaders.AlienCluster.Lasers, bottomLimit)
}

func (invaders *Invaders) removeLaserOf(lasers []*Laser, levelLimit int) []*Laser {
	for index, laser := range lasers {
		_, y := laser.Position()
		isEndOflevel := y == levelLimit

		if isEndOflevel || laser.HasHit {
			invaders.Level.RemoveEntity(laser)

			if laser.HitAlienLaser {
				invaders.addScore(laser.Points)
			}

			if index < len(lasers) {
				lasers = append(lasers[:index], lasers[index+1:]...)
			}
		}
	}

	return lasers
}

func (invaders *Invaders) addScore(points int) {
	invaders.Score += points
}

func (invaders *Invaders) verifyGameOverZone() {
	if invaders.GameOverZone.EnteredZone {
		invaders.Hero.IsAlive = false
	}
}
