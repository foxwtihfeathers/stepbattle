package camera

// GameCamera struct of camera ?? what a shit TODO:: redu
type GameCamera struct {
	X      int
	Y      int
	Width  int
	Height int
}

// MoveCamera - move and update camera coordinates
func (c *GameCamera) MoveCamera(targetX, targetY, mapWidth, mapHeight int) {
	x := targetX - c.Width/2
	y := targetY - c.Height/2

	if x < 0 {
		x = 0
	}

	if y < 0 {
		y = 0
	}

	if x > mapWidth-c.Width {
		x = mapWidth - c.Width
	}

	if y > mapHeight-c.Height {
		y = mapHeight - c.Height
	}

	c.X, c.Y = x, y
}

// ToCameraCoordinates convert coordinate on the map to coordinates on the viewport
func (c *GameCamera) ToCameraCoordinates(mapX, mapY int) (cameraX, cameraY int) {
	x, y := mapX-c.X, mapY-c.Y

	if x < 0 || y < 0 || x >= c.Width || y >= c.Height {
		return -1, -1
	}

	return x, y
}
