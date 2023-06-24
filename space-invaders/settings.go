/*
settings screen
*/
package spaceinvaders

import (
	"fmt"
	"strconv"

	g "github.com/tkorsi/gsi/game"
)

type Settings struct {
	Title       string
	PressToQuit string
	Score       *FpsText
}

func NewSettigns(level *Level, baseLevel *g.BaseLevel) *Settings {
	settings := Settings{Title: "Go Invaders", PressToQuit: "Press Enter to quit"}
	settings.drawTitle(level, baseLevel)
	settings.drawPressToQuit(level, baseLevel)
	settings.drawScore(level, baseLevel)

	return &settings
}

func (s *Settings) drawTitle(level *Level, baseLevel *g.BaseLevel) {
	levelX, levelY := level.Position()
	x := levelX + 1
	y := levelY - 1

	title := g.NewText(x, y, s.Title, g.ColorWhite, g.ColorBlack)
	baseLevel.AddEntity(title)
}

func (s *Settings) drawPressToQuit(level *Level, baseLevel *g.BaseLevel) {
	levelX, levelY := level.Position()
	levelW, levelH := level.Size()

	x := levelX + levelW - len(s.PressToQuit) - 1
	y := levelY + levelH

	title := g.NewText(x, y, s.PressToQuit, g.ColorWhite, g.ColorBlack)
	baseLevel.AddEntity(title)
}

func (s *Settings) drawScore(level *Level, baseLevel *g.BaseLevel) {
	levelX, levelY := level.Position()
	_, levelH := level.Size()

	x := levelX + 1
	y := levelY + levelH

	s.Score = NewFpsText(x, y, g.ColorWhite, g.ColorBlack, 60)
	baseLevel.AddEntity(s.Score)
}

func (s *Settings) UpdateScore(score int) {
	s.Score.SetText(s.getScoreText(score))
}

func (s *Settings) getScoreText(score int) string {
	txtScore := fmt.Sprintf("Score: %4d", score)
	return txtScore
}

type FpsText struct {
	*g.Text
	time   float64
	update float64
}

func NewFpsText(x, y int, fg, bg uint16, update float64) *FpsText {
	return &FpsText{
		Text:   g.NewText(x, y, "", fg, bg),
		time:   0,
		update: update,
	}
}

// Draw updates the framerate on the FpsText and draws it to the Screen s.
func (f *FpsText) Draw(s *g.Screen) {
	f.time += s.TimeDelta()
	if f.time > f.update {
		fps := strconv.FormatFloat(1.0/s.TimeDelta(), 'f', 10, 64)
		f.SetText(fps)
		f.time -= f.update
	}
	f.Text.Draw(s)
}

func cubeIndex(x int, points [5]int) int {
	n := 0
	for _, p := range points {
		if x <= p {
			break
		} else {
			n++
		}
	}
	return n
}

func RgbTo256Color(r, g, b int) uint16 {
	cubepoints := [5]int{47, 115, 155, 195, 235}
	r256 := cubeIndex(r, cubepoints)
	g256 := cubeIndex(g, cubepoints)
	b256 := cubeIndex(b, cubepoints)
	return uint16(r256*36 + g256*6 + b256 + 17)
}
