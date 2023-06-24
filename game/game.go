package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

type Canvas [][]Cell

func NewCanvas(width, heigh int) Canvas {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]Cell, heigh)
	}
	return canvas
}

func (canvas *Canvas) equals(other *Canvas) bool {
	first := *canvas
	second := *other
	if second == nil {
		return false
	}
	if len(first) != len(second) {
		return false
	}
	if len(first[0]) != len(second[0]) {
		return false
	}
	for i := range first {
		for j := range first[i] {
			equal := first[i][j].equals(&(second[i][j]))
			if !equal {
				return false
			}
		}
	}
	return true
}

func CanvasFromString(str string) Canvas {
	lines := strings.Split(str, "\n")
	runes := make([][]rune, len(lines))
	width := 0
	for i := range lines {
		runes[i] = []rune(lines[i])
		width = max(width, len(runes[i]))
	}
	height := len(runes)
	canvas := make(Canvas, width)
	for i := 0; i < width; i++ {
		canvas[i] = make([]Cell, height)
		for j := 0; j < height; j++ {
			if i < len(runes[j]) {
				canvas[i][j] = Cell{Character: runes[j][i]}
			}
		}
	}
	return canvas
}

type Cell struct {
	ForegroundColor uint16
	BackgroundColor uint16
	Character       rune
}

func (c *Cell) equals(c2 *Cell) bool {
	return c.ForegroundColor == c2.ForegroundColor &&
		c.BackgroundColor == c2.BackgroundColor &&
		c.Character == c2.Character
}

type Game struct {
	screen *Screen
	debug  bool
	input  *input
	logs   []string
}

func NewGame() *Game {
	g := Game{
		screen: NewScreen(),
		input:  newInput(),
		logs:   make([]string, 0),
	}
	return &g
}

func (g *Game) Screen() *Screen {
	return g.screen
}

func (g *Game) dumpLogs() {
	if g.debug {
		fmt.Println("=== Logs: ===")
		for _, l := range g.logs {
			fmt.Println(l)
		}
		fmt.Println("=============")
	}
}

func (g *Game) SetEndKey(key Key) {
	g.input.endKey = termbox.Key(key)
}

func (g *Game) SetScreen(s *Screen) {
	g.screen = s
	g.screen.resize(termbox.Size())
}

func (g *Game) Start() {
	err := termbox.Init()
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	if err != nil {
		panic(err)
	}

	defer g.dumpLogs()
	defer termbox.Close()
	g.screen.resize(termbox.Size())

	// Init input
	g.input.start()
	defer g.input.stop()
	clock := time.Now()

mainloop:
	for {
		update := time.Now()
		g.screen.delta = update.Sub(clock).Seconds()
		clock = update

		select {
		case ev := <-g.input.eventQ:
			if ev.Key == g.input.endKey {
				break mainloop
			}
			// else if EventType(ev.Type) == EventResize {
			// 	g.screen.resize(ev.Width, ev.Height)
			// } else if EventType(ev.Type) == EventError {
			// 	g.Log(ev.Err.Error())
			// }
			g.screen.Update(convertEvent(ev))
		default:
			g.screen.Update(Event{Type: EventNone})
		}

		g.screen.Draw()
		time.Sleep(time.Duration((time.Until(update).Seconds()*1000.0)+1000.0/g.screen.fps) * time.Millisecond)
	}
}

type Drawable interface {
	Update(Event)
	Draw(*Screen)
}

type Event struct {
	Type      EventType
	Key       Key // The key pressed
	Character rune
	Err       error
}

func convertEvent(ev termbox.Event) Event {
	return Event{
		Type:      EventType(ev.Type),
		Key:       Key(ev.Key),
		Character: ev.Ch,
		Err:       ev.Err,
	}
}

type (
	Key       uint16
	EventType uint8
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const (
	EventKey EventType = iota
	EventNone
)

type Physical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
}

// DynamicPhysical represents something that can process its own collisions.
// Implementing this is an optional addition to Drawable.
type DynamicPhysical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
	Collide(Physical)     // Handle collisions with another Physical
}
