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
	g.cellSize = min(g.boardBounds.Size.X/cellCols, g.boardBounds.Size.Y/cellRows)
	g.boardOffset = space.Vec2F{
		X: (g.boardBounds.Size.X - g.cellSize*cellCols) / 2,
		Y: (g.boardBounds.Size.Y - g.cellSize*cellRows) / 2,
	}
}
