package game

import "github.com/candlestick-games/snake/pkg/std/space"

func (g *Game) LayoutF(screenWidth, screenHeight float64) (float64, float64) {
	if g.screenWidth != screenWidth || g.screenHeight != screenHeight {
		g.resize(screenWidth, screenHeight)
	}
	return screenWidth, screenHeight
}

func (g *Game) resize(screenWidth, screenHeight float64) {
	g.screenWidth = screenWidth
	g.screenHeight = screenHeight

	g.boardBounds = space.NewRectF(32, 32, screenWidth-32, screenHeight-32)
	g.cellSize = min(g.boardBounds.Size.X/float64(g.gridCols), g.boardBounds.Size.Y/float64(g.gridRows))
	g.boardOffset = space.Vec2F{
		X: (g.boardBounds.Size.X - g.cellSize*float64(g.gridCols)) / 2,
		Y: (g.boardBounds.Size.Y - g.cellSize*float64(g.gridRows)) / 2,
	}
}
