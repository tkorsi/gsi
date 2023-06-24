package spaceinvaders

import g "github.com/tkorsi/gsi/game"

func CreateCanvas(fileContent []byte) g.Canvas {
	canvas := g.CanvasFromString(string(fileContent))
	return canvas
}
