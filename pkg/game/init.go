package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"

	"github.com/candlestick-games/snake/pkg/std/collection"
	"github.com/candlestick-games/snake/pkg/std/rand"
	"github.com/candlestick-games/snake/pkg/std/space"
)

func (g *Game) Init() error {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Snake: Might and Magic")
	ebiten.SetFullscreen(!viper.GetBool("window"))
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g.resetSnake()

	return nil
}

func (g *Game) resetSnake() {
	g.gridCols = 16 * 2
	g.gridRows = 9 * 2

	g.placeWalls()

	g.snake = []space.Vec2I{
		{X: 3, Y: 1},
		{X: 2, Y: 1},
		{X: 1, Y: 1},
	}
	g.dir = space.Vec2I{X: 1, Y: 0}
	g.prevDir = g.dir

	g.food = collection.NewSet[space.Vec2I]()
	g.placeFood()

	g.gameOver = false
}

func (g *Game) placeWalls() {
	g.walls = make([][]bool, g.gridRows)
	for y := 0; y < g.gridRows; y++ {
		g.walls[y] = make([]bool, g.gridCols)
	}

	n := rand.Int((g.gridCols*g.gridRows)/20+1, (g.gridCols*g.gridRows)/10+1)
	for i := 0; i < n; i++ {
		w := space.NewVec2I(-1)
		for w.X < 0 || g.walls[w.Y][w.X] {
			w = space.RandomVec2I(0, g.gridCols-1, 0, g.gridRows-1)
		}

		g.walls[w.Y][w.X] = true
		m := rand.Int(0, 8)
		for j := 0; j < m; j++ {
			rw := g.randomWallNeighbour(w)
			g.walls[rw.Y][rw.X] = true
		}
	}
}

func (g *Game) randomWallNeighbour(w space.Vec2I) space.Vec2I {
	nw := space.NewVec2I(-1)
	for nw.X < 0 || nw == w {
		nw = space.RandomVec2I(max(w.X-1, 0), min(w.X+1, g.gridCols-1), max(w.Y-1, 0), min(w.Y+1, g.gridRows-1))
	}
	return nw
}

func (g *Game) placeFood() {
	for {
		pos := space.RandomVec2I(1, g.gridCols-1, 1, g.gridRows-1)

		if g.food.Has(pos) {
			continue
		}

		if g.walls[pos.Y][pos.X] {
			continue
		}

		var ok bool
		for _, snake := range g.snake {
			if snake == pos {
				ok = true
				break
			}
		}
		if ok {
			continue
		}

		g.food.Add(pos)
		return
	}
}
