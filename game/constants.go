package game

// Took from `https://github.com/eiannone/keyboard/blob/master/keyboard_common.go`

const (
	KeyF1 Key = 0xFFFF - iota
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgup
	KeyPgdn
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	key_min
	MouseLeft
	MouseMiddle
	MouseRight
	MouseRelease
	MouseWheelUp
	MouseWheelDown
)

const (
	ColorDefault uint16 = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)
