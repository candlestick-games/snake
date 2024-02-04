package game

import (
	"slices"

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
	const maxIterations = 256

	i := 0
	var rooms []space.RectI
	for len(rooms) != numberOfRoom {
		if i > maxIterations {
			log.Debug("Max iterations reached")
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
			log.Debug("Room out of bounds", "room", room, "grid", g.gridBounds)
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
			log.Debug("Cluster out of bounds", "cluster", cluster, "room", room)
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
			log.Debug("Cluster out of bounds", "cluster", cluster, "room-outside", roomOutside,
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
	floodFillStart := space.NewVec2I(-1)
	connected := make([][]bool, g.gridRows)
	for y := 0; y < g.gridRows; y++ {
		connected[y] = make([]bool, g.gridCols)
		for x := 0; x < g.gridCols; x++ {
			connected[y][x] = g.walls[y][x]
			if floodFillStart.X < 0 && !g.walls[y][x] {
				floodFillStart = space.NewVec2I(x, y)
			}
		}
	}
	if floodFillStart.X < 0 {
		log.Debug("No empty cells generated")
		g.placeWalls()
		return
	}

	g.floodFill(floodFillStart, connected)
	fullyConnected := slices.ContainsFunc(connected, func(cols []bool) bool {
		return slices.Contains(cols, false)
	})
	if fullyConnected {
		log.Debug("Not fully connected")
		g.placeWalls()
		return
	}

	// Remove simple dead ends
	for y := 0; y < g.gridRows; y++ {
		for x := 0; x < g.gridCols; x++ {
			if g.walls[y][x] {
				continue
			}

			pos := space.NewVec2I(x, y)
			if g.countEntrances(pos) != 1 {
				continue
			}
			log.Debug("Has a dead end")

			dirs := []space.Vec2I{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
			rand.Shuffle(dirs)

			removed := false
			for _, dir := range dirs {
				w := pos.Add(dir)
				if !g.gridBounds.Contains(w) || !g.walls[w.Y][w.X] {
					continue
				}

				if g.countEntrances(w) >= 2 {
					g.walls[w.Y][w.X] = false
					removed = true
					break
				}
			}
			if !removed {
				log.Debug("Failed to remove a dead end")
				g.placeWalls()
				return
			}
		}
	}

	// Remove complex dead ends
	// TODO: Implement if whole area has only one entrance
}

func (g *Game) floodFill(start space.Vec2I, connected [][]bool) {
	if start.X < 0 || start.X >= g.gridCols ||
		start.Y < 0 || start.Y >= g.gridRows ||
		connected[start.Y][start.X] {
		return
	}

	connected[start.Y][start.X] = true

	g.floodFill(start.Add(space.NewVec2I(1, 0)), connected)
	g.floodFill(start.Add(space.NewVec2I(-1, 0)), connected)
	g.floodFill(start.Add(space.NewVec2I(0, 1)), connected)
	g.floodFill(start.Add(space.NewVec2I(0, -1)), connected)
}

func (g *Game) countEntrances(pos space.Vec2I) int {
	entrances := 0
	if !g.isWall(pos.Add(space.NewVec2I(1, 0))) {
		entrances++
	}
	if !g.isWall(pos.Add(space.NewVec2I(-1, 0))) {
		entrances++
	}
	if !g.isWall(pos.Add(space.NewVec2I(0, 1))) {
		entrances++
	}
	if !g.isWall(pos.Add(space.NewVec2I(0, -1))) {
		entrances++
	}
	return entrances
}
