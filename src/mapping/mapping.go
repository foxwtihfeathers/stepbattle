package mapping

// TODO:: tile or map - who must be exported?

type tile struct {
	Blocked    bool
	BlockSight bool
}

// Map - 2D map with width and height of Tiles,
type Map struct {
	Width  int
	Height int
	Tiles  [][]*tile
}

// InitializeMap set up a map with all the border
// TODO:: We heen build maps more dynamically
func (m *Map) InitializeMap() {
	m.Tiles = make([][]*tile, m.Width)
	for i := range m.Tiles {
		m.Tiles[i] = make([]*tile, m.Height)
	}

	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			if x == 0 || x == m.Width-1 || y == 0 || y == m.Height-1 {
				m.Tiles[x][y] = &tile{true, true}
			} else {
				m.Tiles[x][y] = &tile{false, false}
			}
		}
	}
}

// IsBlocked check if movement blocked
func (m *Map) IsBlocked(x, y int) bool {
	if m.Tiles[x][y].Blocked {
		return true
	} else {
		return false
	}
}
