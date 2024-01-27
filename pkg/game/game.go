package game

import (
	"github.com/candlestick-games/snake/pkg/std/collection"
	"github.com/candlestick-games/snake/pkg/std/space"
)

const (
	cellCols = 16 * 2
	cellRows = 9 * 2
)

type Game struct {
	boardBounds space.RectF
	boardOffset space.Vec2F
	cellSize    float64

	snake    []space.Vec2I
	dir      space.Vec2I
	prevDir  space.Vec2I
	gameOver bool
	ticks    uint

	walls [cellRows][cellCols]bool

	food collection.Set[space.Vec2I]

	screenWidth  float64
	screenHeight float64

	quit bool
}

func (g *Game) Layout(_, _ int) (int, int) { panic("unreachable") }

func (g *Game) Shutdown() {}
