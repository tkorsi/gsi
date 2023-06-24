package spacee

import (
	_ "embed"
	"fmt"

	g "github.com/tkorsi/gsi/game"
)

var (
	//go:embed files/tige.txt
	tigeScreenFile []byte

	//go:embed files/game_over.txt
	gameOverScreenFile []byte
)

func ShowTigeScreen(e *Invaders) {
	prepareScreen(e)
	showTige(e)

	if checkLevelSizeNotOk(e) {
		e.ScreenSizeNotOK = true
		showMaximizeScreen(e)
		return
	}

	showPressToInit(e, 0)
}

func checkLevelSizeNotOk(e *Invaders) bool {
	w, h := e.Level.Size()

	if w < 100 || h < 37 {
		return true
	}

	return false
}

func ShowGameOverScreen(e *Invaders) {
	prepareScreen(e)
	showGameOver(e)
	showScore(e)
	showPressToInit(e, 2)
}

func prepareScreen(e *Invaders) {
	e.BaseLevel = g.NewBaseLevel(g.Cell{BackgroundColor: g.ColorBlack, ForegroundColor: g.ColorWhite})
	e.Game.Screen().SetLevel(e.BaseLevel)
	e.BaseLevel.AddEntity(e)

	e.initlevel()
	e.initHud()
}

func showTige(e *Invaders) {
	showCanvas(e, tigeScreenFile)
}

func showGameOver(e *Invaders) {
	showCanvas(e, gameOverScreenFile)
}

func showCanvas(e *Invaders, file []byte) {
	canvas := CreateCanvas(file)

	levelX, levelY := e.level.Position()
	levelW, levelH := e.level.Size()

	x := levelX + levelW/2 - len(canvas)/2
	y := levelY + levelH/2 + -len(canvas[0]) - 1

	e.baseLevel.AddEntity(g.NewEntityFromCanvas(x, y, canvas))
}

func showScore(e *Invaders) {
	score := fmt.Sprintf("SCORE: %4d ", e.Score)
	showCenterText(score, 0, e)
}

func showPressToInit(e *Invaders, topPadding int) {
	showCenterText("Press ENTER to start", topPadding, e)
}

func showMaximizeScreen(e *Invaders) {
	showCenterText("Maximize the console and run the game again", 0, e)
}

func showCenterText(text string, topPadding int, e *Invaders) {
	levelX, levelY := e.level.Position()
	levelW, levelH := e.level.Size()

	x := levelX + levelW/2 - len(text)/2
	y := levelY + levelH/2 + topPadding

	e.baseLevel.AddEntity(g.NewText(x, y, text, g.ColorWhite, g.ColorBlack))
}
