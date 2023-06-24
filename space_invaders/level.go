package spaceinvaders

import (
	g "github.com/tkorsi/gsi/game"
)

type Level struct {
	*g.Entity
	Init int
	End  int
}

const (
	LevelMaxWidth  = 100
	LevelMaxHeight = 37
)

func newLevel(screenWidth int, screenHeight int) *Level {
	width := ValueMinusPercent(screenWidth, 0.40)
	height := ValueMinusPercent(screenHeight, 0.15)

	if width > LevelMaxWidth {
		width = LevelMaxWidth
	}

	if height > LevelMaxHeight {
		height = LevelMaxHeight
	}

	centerX := screenWidth/2 - width/2
	centerY := screenHeight/2 - height/2
	init := centerX + 1
	end := centerX + width

	return &Level{g.NewEntityFromCanvas(centerX, centerY, createLevel(width, height)), init, end}
}

func createLevel(width, height int) g.Canvas {
	canvas := g.NewCanvas(width, height)

	for x, cell := range canvas {
		for y := range cell {
			fillTopBottom(x, y, height, canvas)
			fillSides(x, y, width, canvas)
		}
	}

	createCell(0, 0, canvas, '┌')
	createCell(width-1, 0, canvas, '┐')
	createCell(0, height-1, canvas, '└')
	createCell(width-1, height-1, canvas, '┘')

	return canvas
}

func fillTopBottom(x, y, height int, canvas g.Canvas) {
	if x > 0 && (y == 0 || y == height-1) {
		createCell(x, y, canvas, '─')
	}
}

func fillSides(x, y, width int, canvas g.Canvas) {
	if x == 0 || x == width-1 {
		createCell(x, y, canvas, '│')
	}
}

func createCell(x, y int, canvas g.Canvas, ch rune) {
	canvas[x][y] = g.Cell{BackgroundColor: g.ColorBlack, ForegroundColor: g.ColorWhite, Character: ch}
}

func ValueMinusPercent(val int, percentage float64) int {
	return int(float64(val) - float64(val)*percentage)
}
