package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/candlestick-games/snake/pkg/std/debugger"
	"github.com/candlestick-games/snake/pkg/std/pencil"
	"github.com/candlestick-games/snake/pkg/std/space"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Gray)
	board := screen.SubImage(g.boardBounds.ToIR()).(*ebiten.Image)
	board.Clear()

	pad := space.NewVec2F(4)
	cellSize := space.NewVec2F(g.cellSize).Sub(pad)
	cellOffset := g.boardBounds.Pos.Add(g.boardOffset).Add(pad.Scale(0.5))
	for y := 0; y < cellRows; y++ {
		for x := 0; x < cellCols; x++ {
			pencil.FillRectV(
				screen,
				space.Vec2F{
					X: float64(x) * g.cellSize,
					Y: float64(y) * g.cellSize,
				}.Add(cellOffset),
				cellSize,
				colornames.Aliceblue,
			)
		}
	}

	for y := 0; y < cellRows; y++ {
		for x := 0; x < cellCols; x++ {
			if !g.walls[y][x] {
				continue
			}

			pos := space.NewVec2I(x, y)
			pencil.FillRectV(
				screen,
				pos.ToF().Mul(cellSize.Add(pad)).Add(cellOffset),
				cellSize,
				colornames.Black,
			)
		}
	}

	for _, pos := range g.snake {
		pencil.FillRectV(
			screen,
			pos.ToF().Mul(cellSize.Add(pad)).Add(cellOffset),
			cellSize,
			colornames.Green,
		)
	}

	for pos := range g.food {
		pencil.FillRectV(
			screen,
			pos.ToF().Mul(cellSize.Add(pad)).Add(cellOffset),
			cellSize,
			colornames.Red,
		)
	}

	moveTo := g.snake[0].ToF().Add(g.dir.ToF().Scale(0.5))
	pencil.FillRectV(
		screen,
		moveTo.Mul(cellSize.Add(pad)).Add(cellOffset).Add(cellSize.Scale(0.25)),
		cellSize.Scale(0.5),
		colornames.Orange,
	)

	debugger.Draw(screen)
}
