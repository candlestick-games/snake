package game

import (
	"github.com/candlestick-games/snake/pkg/std/collection"
	"github.com/candlestick-games/snake/pkg/std/space"
	"github.com/candlestick-games/snake/pkg/std/tick"
)

type Game struct {
	// Grid

	gridCols   int
	gridRows   int
	gridBounds space.RectI

	// Board

	boardBounds space.RectF
	boardOffset space.Vec2F
	cellSize    float64

	// Snake

	snake     []space.Vec2I
	dir       space.Vec2I
	prevDir   space.Vec2I
	startTime *tick.Timer
	pause     bool
	gameOver  bool

	walls [][]bool

	food collection.Set[space.Vec2I]

	// Core

	screenWidth  float64
	screenHeight float64

	ticker *tick.Ticker

	quit bool
}

func (g *Game) Layout(_, _ int) (int, int) { panic("unreachable") }

func (g *Game) Shutdown() {}

func (g *Game) isWall(pos space.Vec2I) bool {
	if !g.gridBounds.Contains(pos) {
		return true
	}
	return g.walls[pos.Y][pos.X]
}

func (g *Game) isSnake(pos space.Vec2I) bool {
	for _, p := range g.snake {
		if p == pos {
			return true
		}
	}
	return false
}

func (g *Game) randomUnoccupiedNeighbour(pos space.Vec2I) space.Vec2I { // TODO: This may fail
	nw := space.NewVec2I(-1)
	for nw.X < 0 || g.isWall(nw) {
		nw = pos.Add(space.RandomVec2IDir())
	}
	return nw
}
