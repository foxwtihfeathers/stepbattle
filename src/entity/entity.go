package entity

import (
	blt "bearlibterminal"
)

// GameEntity is a general purpose object.
type GameEntity struct {
	X     int
	Y     int
	Layer int
	Char  string
	Color string
}

// Move by amout dx, dy
func (e *GameEntity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

// Draw entity on screen
func (e *GameEntity) Draw(mapX, mapY int) {
	blt.Layer(e.Layer)
	blt.Color(blt.ColorFromName(e.Color))
	blt.Print(mapX, mapY, e.Char) // TODO:: set up complicate char
}

// Clear remove entity from the screen
func (e *GameEntity) Clear(mapX, mapY int) {
	blt.Layer(e.Layer)
	blt.Print(mapX, mapY, " ")
}
