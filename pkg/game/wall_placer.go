package game

import (
	"github.com/charmbracelet/log"

	"github.com/candlestick-games/snake/pkg/std/rand"
	"github.com/candlestick-games/snake/pkg/std/space"
)

func (g *Game) placeWalls() {
	// Fill walls
	g.walls = make([][]bool, g.gridRows)
	for y := 0; y < g.gridRows; y++ {
		g.walls[y] = make([]bool, g.gridCols)
		for x := 0; x < g.gridCols; x++ {
			g.walls[y][x] = true
		}
	}

	// Carve rooms
	const numberOfRoom = 8
	const maxIterations = 100

	i := 0
	var rooms []space.RectI
	for len(rooms) != numberOfRoom {
		if i > maxIterations {
			break
		}
		i++

		size := space.Vec2I{
			X: rand.IntWithSkew(2, g.gridCols, 0.5),
			Y: rand.IntWithSkew(2, g.gridRows, 0.4),
		}
		pos := space.RandomVec2I(0, g.gridCols-size.X, 0, g.gridRows-size.Y)

		room := space.RectI{
			Pos:  pos,
			Size: size,
		}
		if !room.Inside(g.gridBounds) {
			log.Error("Room out of bounds", "room", room, "grid", g.gridBounds)
			continue
		}

		roomBound := room.
			Move(space.NewVec2I(-1)).
			Grow(space.NewVec2I(2)).
			Clamp(g.gridBounds)

		intercepted := false
		for _, placedRoom := range rooms {
			if roomBound.Intercepts(placedRoom) {
				intercepted = true
				break
			}
		}
		if intercepted {
			continue
		}
		rooms = append(rooms, room)

		for y := pos.Y; y < pos.Y+size.Y; y++ {
			for x := pos.X; x < pos.X+size.X; x++ {
				g.walls[y][x] = false
			}
		}
	}

	// Place random walls (clusters in rooms)
	for _, room := range rooms {
		if rand.Bool(0.2) {
			continue
		}
		if room.Size.X < 2 || room.Size.Y < 2 {
			continue
		}

		size := space.Vec2I{
			X: rand.IntWithSkew(1, room.Size.X-1, 0.5),
			Y: rand.IntWithSkew(1, room.Size.Y-1, 0.5),
		}

		x1, y1, x2, y2 := room.ToCoords()
		pos := space.RandomVec2I(x1, x2-size.X-1, y1, y2-size.Y-1)
		cluster := space.RectI{
			Pos:  pos,
			Size: size,
		}

		if !cluster.Inside(room) {
			log.Error("Cluster out of bounds", "cluster", cluster, "room", room)
			continue
		}

		for y := pos.Y; y < pos.Y+size.Y; y++ {
			for x := pos.X; x < pos.X+size.X; x++ {
				g.walls[y][x] = true
			}
		}
	}

	// Remove random walls (clusters in rooms)
	for _, room := range rooms {
		if rand.Bool(0.3) {
			continue
		}

		size := space.Vec2I{
			X: rand.IntWithSkew(1, max(room.Size.X-1, 1), 0.2),
			Y: rand.IntWithSkew(1, max(room.Size.Y-1, 1), 0.2),
		}

		roomOutside := room.Move(space.NewVec2I(-1)).Grow(space.NewVec2I(2)).Clamp(g.gridBounds)

		x1, y1, x2, y2 := roomOutside.ToCoords()
		pos := space.RandomVec2I(x1, x2-size.X-1, y1, y2-size.Y-1)
		cluster := space.RectI{
			Pos:  pos,
			Size: size,
		}

		if !cluster.Inside(roomOutside) || !cluster.Inside(g.gridBounds) {
			log.Error("Cluster out of bounds", "cluster", cluster, "room-outside", roomOutside,
				"grid", g.gridBounds)
			continue
		}

		for y := pos.Y; y < pos.Y+size.Y; y++ {
			for x := pos.X; x < pos.X+size.X; x++ {
				g.walls[y][x] = false
			}
		}
	}

	// Carve corridors
	for _, fromRoom := range rooms {
		for _, toRoom := range rooms {
			p1 := fromRoom.Center()
			p2 := toRoom.Center()

			for p1 != p2 {
				g.walls[p1.Y][p1.X] = false

				dir := p2.Sub(p1)
				a := dir.Abs()
				if a.X != 0 {
					dir.X = space.Sign(dir.X)
					dir.Y = 0
				} else {
					dir.X = 0
					dir.Y = space.Sign(dir.Y)
				}

				p1 = p1.Add(dir)
			}
			g.walls[p2.Y][p2.X] = false
		}
	}

	// Validate that there are empty cells & all cells connected
	// TODO: Implement
	hasEmpty := false
	connected := make([][]bool, g.gridRows)
	for y := 0; y < g.gridRows; y++ {
		connected[y] = make([]bool, g.gridCols)
		for x := 0; x < g.gridCols; x++ {
			connected[y][x] = false
			if !hasEmpty && !g.walls[y][x] {
				hasEmpty = true
				connected[y][x] = true
			}
		}
	}
	if !hasEmpty {
		log.Error("No empty cells generated")
		g.placeWalls()
		return
	}

	// Remove dead ends (group of cells with only 1 entrance)
	// TODO: Implement
}
