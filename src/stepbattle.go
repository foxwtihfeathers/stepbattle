package main

import (
	blt "bearlibterminal"
	"strconv"

	"./camera"
	"./entity"
	"./mapping"
)

const (
	windowSizeX = 100
	windowSizeY = 30
	mapWidth    = windowSizeX
	mapHeight   = windowSizeY
	title       = "Step"
	font        = "default"
	fontSize    = 30
)

var (
	player     *entity.GameEntity
	entities   []*entity.GameEntity
	levelmap   *mapping.Map
	gameCamera *camera.GameCamera
)

func initWindow() {
	blt.Open()
	size := "size=" + strconv.Itoa(windowSizeX) + "x" + strconv.Itoa(windowSizeY)
	title := "title='" + title + "'"
	window := "window: " + size + "," + title
	fontSize := "size=" + strconv.Itoa(fontSize)
	font := "font: " + font + ", " + fontSize
	blt.Set(window + "; " + font)
	blt.Clear()
}

func initUnit() {
	player = &entity.GameEntity{X: 1, Y: 1, Layer: 1, Char: "@", Color: "white"}
	npc := &entity.GameEntity{X: 10, Y: 10, Layer: 0, Char: "N", Color: "red"}
	entities = append(entities, player, npc)
}

func initLevel() {
	levelmap = &mapping.Map{Width: mapWidth, Height: mapHeight}
	levelmap.InitializeMap()
	levelmap.GenerateArena()
	levelmap.GenerateCavern(50, 5)
}

func initCamera() {
	gameCamera = &camera.GameCamera{X: 1, Y: 1, Width: windowSizeX, Height: windowSizeY}
}

func init() {
	initWindow()
	initLevel()
	initUnit()
	initCamera()
}

func renderEntities() {
	for _, e := range entities {
		cameraX, cameraY := gameCamera.ToCameraCoordinates(e.X, e.Y)

		e.Draw(cameraX, cameraY)

	}
}

func renderMap() {
	for x := 0; x < gameCamera.Width; x++ {
		for y := 0; y < gameCamera.Height; y++ {
			mapX, mapY := gameCamera.X+x, gameCamera.Y+y
			if levelmap.Tiles[mapX][mapY].Blocked == true {
				blt.Color(blt.ColorFromName("gray"))
				blt.Print(x, y, "#")
			} else {
				blt.Color(blt.ColorFromName("brown"))
				blt.Print(x, y, ".")
			}
		}
	}
}

func renderAll() {
	// Very controversial decision, better create method?
	gameCamera.MoveCamera(player.X, player.Y, mapWidth, mapHeight)
	renderMap()
	renderEntities()
}

func handleKey(key int) {

	var (
		dx, dy int
	)
	switch key {
	case blt.TK_KP_1:
		dx, dy = -1, 1
	case blt.TK_KP_2:
		dx, dy = 0, 1
	case blt.TK_KP_3:
		dx, dy = 1, 1
	case blt.TK_KP_4:
		dx, dy = -1, 0
	case blt.TK_KP_5:
		dx, dy = 0, 0
	case blt.TK_KP_6:
		dx, dy = 1, 0
	case blt.TK_KP_7:
		dx, dy = -1, -1
	case blt.TK_KP_8:
		dx, dy = 0, -1
	case blt.TK_KP_9:
		dx, dy = 1, -1
	}

	if 0 > (player.X+dx) || (player.X+dx) > windowSizeX-1 {
		dx = 0
	}
	if 0 > (player.Y+dy) || (player.Y+dy) > windowSizeY-1 {
		dy = 0
	}

	if !levelmap.IsBlocked(player.X+dx, player.Y+dy) {
		player.Move(dx, dy)
	}
}

func main() {
	renderAll()

	for {
		blt.Refresh()

		key := blt.Read()

		for _, e := range entities {
			mapX, mapY := gameCamera.ToCameraCoordinates(e.X, e.Y)
			e.Clear(mapX, mapY)
		}

		if key != blt.TK_CLOSE {
			handleKey(key)
		} else {
			break
		}
		renderAll()
	}

	blt.Close()
}
