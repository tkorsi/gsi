package game

type Game struct {
	debug bool
	logs  []string
}

func NewGame() *Game {
	g := Game{
		logs: make([]string, 0),
	}
	return &g
}
