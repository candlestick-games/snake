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
	g.placeWalls()

	g.snake = []space.Vec2I{
		{X: 3, Y: 1},
		{X: 2, Y: 1},
		{X: 1, Y: 1},
	}
	g.dir = space.Vec2I{X: 1, Y: 0}

	g.food = collection.NewSet[space.Vec2I]()
	g.placeFood()

	g.gameOver = false
}

func (g *Game) placeWalls() {
	for y := 0; y < cellRows; y++ {
		for x := 0; x < cellCols; x++ {
			g.walls[y][x] = false
		}
	}

	n := rand.Int((cellCols*cellRows)/20+1, (cellCols*cellRows)/10+1)
	for i := 0; i < n; i++ {
		w := space.NewVec2I(-1)
		for w.X < 0 || g.walls[w.Y][w.X] {
			w = space.RandomVec2I(0, cellCols-1, 0, cellRows-1)
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
		nw = space.RandomVec2I(max(w.X-1, 0), min(w.X+1, cellCols-1), max(w.Y-1, 0), min(w.Y+1, cellRows-1))
	}
	return nw
}

func (g *Game) placeFood() {
	for {
		pos := space.RandomVec2I(1, cellCols-1, 1, cellRows-1)

		if g.food.Has(pos) {
			continue
		}

		if g.walls[pos.Y][pos.Y] {
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
