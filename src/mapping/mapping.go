package mapping

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO:: tile or map - who must be exported?

type tile struct {
	Blocked    bool
	BlockSight bool
	Visited    bool
	X          int
	Y          int
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
	rand.Seed(time.Now().UTC().UnixNano())
}

func (t *tile) isWall() bool {
	if t.BlockSight && t.Blocked {
		return true
	} else {
		return false
	}
}

func (t *tile) isVisited() bool {
	return t.Visited
}

// IsBlocked check if movement blocked
func (m *Map) IsBlocked(x, y int) bool {
	if m.Tiles[x][y].Blocked {
		return true
	} else {
		return false
	}
}

// GenerateArena - create large empty room, with wall ringing the outside edge
func (m *Map) GenerateArena() {
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			if x == 0 || x == m.Width-1 || y == 0 || y == m.Height-1 {
				m.Tiles[x][y] = &tile{true, true, false, x, y}
			} else {
				m.Tiles[x][y] = &tile{false, false, false, x, y}
			}
		}
	}
}

// GenerateCavern - return valid starting position for the player
func (m *Map) GenerateCavern(seedWallChance, iter int) (int, int) {
	//Fill with random
	fmt.Printf("%d", seedWallChance)
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			state := rand.Intn(100)
			if state < (100 - seedWallChance) {
				m.Tiles[x][y] = &tile{true, true, false, x, y}
			} else {
				m.Tiles[x][y] = &tile{false, false, false, x, y}
			}

		}
	}

	for i := 0; i < iter; i++ {
		for x := 0; x < m.Width; x++ {
			for y := 0; y < m.Height-1; y++ {
				wallOneAway := m.countWallsNStepsAway(1, x, y)

				wallTwoAway := m.countWallsNStepsAway(2, x, y)

				if wallOneAway >= 5 || wallTwoAway <= 2 {
					m.Tiles[x][y].Blocked = true
					m.Tiles[x][y].BlockSight = true
				} else {
					m.Tiles[x][y].Blocked = false
					m.Tiles[x][y].BlockSight = false
				}
			}
		}

	}
	return 0, 0
}

// countWallsNStepsAway - return the number of wall tiles that are within n spaces of the given tile
func (m *Map) countWallsNStepsAway(n int, x int, y int) int {
	wallCount := 0

	for r := -n; r <= n; r++ {
		for c := -n; c <= n; c++ {
			if x+r >= m.Width || x+r <= 0 || y+c >= m.Height || y+c <= 0 {
				// Check if the current coordinates would be off the map. Off map coordinates count as a wall.
				wallCount++
			} else if m.Tiles[x+r][y+c].Blocked && m.Tiles[x+r][y+c].BlockSight {
				wallCount++
			}
		}
	}

	return wallCount
}
