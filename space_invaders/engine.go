package spaceinvaders

import (
	"time"

	g "github.com/tkorsi/gsi/game"
)

type Invaders struct {
	*g.Entity
	Game               *g.Game
	BaseLevel          *g.BaseLevel
	Level              *Level
	AlienCluster       *AlienCluster
	Settings           *Settings
	Player             *Player
	GameOver           *GameOver
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
		BaseLevel:          g.NewBaseLevel(g.Cell{BackgroundColor: g.ColorBlack, ForegroundColor: g.ColorWhite}),
		AlienLaserVelocity: 0.04,
		RefreshSpeed:       20,
		Score:              0,
	}

	e.Game.Screen().SetFps(60)
	e.BaseLevel.AddEntity(&e)

	return &e
}

func (e *Invaders) Start() {
	go ShowTitleScreen(e)
	e.Game.Start()
}

func (e *Invaders) Update(ev g.Event) {
	if e.Started == false && e.ScreenSizeNotOK == false && ev.Type == g.EventKey && ev.Key == g.KeyEnter {
		go e.initializeGame()
	}
}

func (e *Invaders) initializeGame() {
	e.Started = true

	e.initPlayer()
	e.initAliens()
	e.initGameOver()
	e.gameLoop()
}

func (e *Invaders) initLevel() {
	screenWidth, screenHeight := e.getScreenSize()
	e.Level = newLevel(screenWidth, screenHeight)
	e.BaseLevel.AddEntity(e.Level)
}

func (e *Invaders) initSettings() {
	e.Settings = NewSettigns(e.Level, e.BaseLevel)
}

func (e *Invaders) getScreenSize() (int, int) {
	screenWidth, screenHeight := e.Game.Screen().Size()

	for screenWidth == 0 && screenHeight == 0 {
		time.Sleep(100 * time.Millisecond)
		screenWidth, screenHeight = e.Game.Screen().Size()
	}

	return screenWidth, screenHeight
}

func (e *Invaders) initPlayer() {
	e.Player = NewPlayer(e.Level)
	e.BaseLevel.AddEntity(e.Player)
}

func (e *Invaders) initAliens() {
	e.AlienCluster = NewAlienCluster()
	SetPositionAndRenderAliens(e.AlienCluster.Aliens, e.BaseLevel, e.Level)
}

func (e *Invaders) initGameOver() {
	e.GameOver = CreateGameOver(e.Level, e.Player)
	e.BaseLevel.AddEntity(e.GameOver)
}

func (e *Invaders) gameLoop() {
	for {
		if e.Player.IsDead() || e.AlienCluster.IsAllAliensDead() {
			e.Player.animatePlayerEndGame(e.BaseLevel)
			e.Started = false
			break
		}

		e.updateLaserPositions()
		e.RemoveDeadAliensAndIncrementScore()
		e.updateAlienClusterPosition()
		e.updateScore()
		e.verifyGameOver()

		time.Sleep(e.RefreshSpeed * time.Millisecond)
	}

	if e.Player.IsDead() {
		ShowGameOverScreen(e)
	}

	if e.AlienCluster.IsAllAliensDead() {
		e.initializeGame()
	}
}

func (e *Invaders) updateScore() {
	e.Settings.UpdateScore(e.Score)
}

func (e *Invaders) updateAlienClusterPosition() {
	e.AlienCluster.UpdateAliensPositions(e.Game.Screen().TimeDelta(), e.Level)
	e.AlienCluster.Shoot()
}

func (e *Invaders) RemoveDeadAliensAndIncrementScore() {
	points := e.AlienCluster.RemoveDeadAliensAndGetPoints(e.BaseLevel)
	e.addScore(points)
}

func (e *Invaders) updateLaserPositions() {
	e.updatePlayerLasers()
	e.updateAlienLasers()
	e.removeLasers()
}

func (e *Invaders) updatePlayerLasers() {
	e.updateLasers(e.Player.Lasers)
}

func (e *Invaders) updateAlienLasers() {
	e.TimeDelta += e.Game.Screen().TimeDelta()

	if e.TimeDelta >= e.AlienLaserVelocity {
		e.TimeDelta = 0
		e.updateLasers(e.AlienCluster.Lasers)
	}
}

func (e *Invaders) updateLasers(lasers []*Laser) {
	for _, laser := range lasers {
		if laser.IsNew {
			e.renderNewLaser(laser)
			continue
		}

		x, y := laser.Position()
		laser.SetPosition(x, y-laser.Direction)
	}
}

func (e *Invaders) renderNewLaser(laser *Laser) {
	laser.IsNew = false
	e.BaseLevel.AddEntity(laser)
}

func (e *Invaders) removeLasers() {
	_, levelY := e.Level.Position()
	_, levelH := e.Level.Size()

	upperLimit := levelY
	bottomLimit := levelY + levelH - 1

	e.Player.Lasers = e.removeLaserOf(e.Player.Lasers, upperLimit)
	e.AlienCluster.Lasers = e.removeLaserOf(e.AlienCluster.Lasers, bottomLimit)
}

func (e *Invaders) removeLaserOf(lasers []*Laser, levelLimit int) []*Laser {
	for index, laser := range lasers {
		_, y := laser.Position()
		isEndOflevel := y == levelLimit

		if isEndOflevel || laser.HasHit {
			e.BaseLevel.RemoveEntity(laser)

			if laser.HitAlienLaser {
				e.addScore(laser.Points)
			}

			if index < len(lasers) {
				lasers = append(lasers[:index], lasers[index+1:]...)
			}
		}
	}

	return lasers
}

func (e *Invaders) addScore(points int) {
	e.Score += points
}

func (e *Invaders) verifyGameOver() {
	if e.GameOver.EnteredZone {
		e.Player.IsAlive = false
	}
}
