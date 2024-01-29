package game

import (
	"github.com/candlestick-games/snake/pkg/std/collection"
	"github.com/candlestick-games/snake/pkg/std/space"
)

type Game struct {
	boardBounds space.RectF
	boardOffset space.Vec2F
	cellSize    float64

	ticks uint

	snake    []space.Vec2I
	dir      space.Vec2I
	prevDir  space.Vec2I
	gameOver bool

	gridCols int
	gridRows int

	walls [][]bool

	food collection.Set[space.Vec2I]

	screenWidth  float64
	screenHeight float64

	quit bool
}

func (g *Game) Layout(_, _ int) (int, int) { panic("unreachable") }

func (g *Game) Shutdown() {}
