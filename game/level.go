package game

type Level interface {
	DrawBackground(*Screen)
	AddEntitiy(Drawable)
	RemveEntity(Drawable)
	Draw(*Screen)
	Update(Event)
}
