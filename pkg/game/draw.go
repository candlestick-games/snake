package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/candlestick-games/snake/assets"
	"github.com/candlestick-games/snake/pkg/std/debugger"
	"github.com/candlestick-games/snake/pkg/std/pencil"
	"github.com/candlestick-games/snake/pkg/std/space"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Gray)
	board := screen.SubImage(g.boardBounds.ToIR()).(*ebiten.Image)
	board.Clear()

	// Cells
	cellSize := space.NewVec2F(g.cellSize)
	cellOffset := g.boardBounds.Pos.Add(g.boardOffset)
	for y := 0; y < cellRows; y++ {
		for x := 0; x < cellCols; x++ {
			if g.walls[y][x] {
				continue
			}

			img := assets.Image(assets.Floor)

			op := &ebiten.DrawImageOptions{}
			op.GeoM = space.ImgResizeTo(op.GeoM, img, cellSize)

			pos := space.NewVec2I(x, y)
			op.GeoM = space.Translate(op.GeoM, pos.ToF().Mul(cellSize).Add(cellOffset))

			screen.DrawImage(img, op)
		}
	}

	// Walls
	for y := 0; y < cellRows; y++ {
		for x := 0; x < cellCols; x++ {
			if !g.walls[y][x] {
				continue
			}

			img := assets.Image(assets.Wall)

			op := &ebiten.DrawImageOptions{}
			op.GeoM = space.ImgResizeTo(op.GeoM, img, cellSize)

			pos := space.NewVec2I(x, y)
			op.GeoM = space.Translate(op.GeoM, pos.ToF().Mul(cellSize).Add(cellOffset))

			screen.DrawImage(img, op)
		}
	}

	// Food
	for pos := range g.food {
		img := assets.Image(assets.Apple)

		op := &ebiten.DrawImageOptions{}
		op.GeoM = space.ImgResizeTo(op.GeoM, img, cellSize)
		op.GeoM = space.Translate(op.GeoM, pos.ToF().Mul(cellSize).Add(cellOffset))

		screen.DrawImage(img, op)
	}

	// Snake
	for i, pos := range g.snake {
		var img *ebiten.Image
		angle := 0.0
		verticalFlip := false
		horizontalFlip := false

		switch i {
		case 0:
			switch g.prevDir {
			case space.Vec2I{X: 1}:
				img = assets.Image(assets.SnakeHeadSide)
			case space.Vec2I{X: -1}:
				img = assets.Image(assets.SnakeHeadSide)
				angle = -math.Pi
				verticalFlip = true
			case space.Vec2I{Y: 1}:
				img = assets.Image(assets.SnakeHeadTop)
			case space.Vec2I{Y: -1}:
				img = assets.Image(assets.SnakeHeadTop)
				angle = -math.Pi
			default:
				panic("unreachable")
			}
		case len(g.snake) - 1:
			diff := g.snake[i-1].Sub(pos)
			switch diff {
			case space.Vec2I{X: 1}:
				img = assets.Image(assets.SnakeTail)
			case space.Vec2I{X: -1}:
				img = assets.Image(assets.SnakeTail)
				horizontalFlip = true
			case space.Vec2I{Y: 1}:
				img = assets.Image(assets.SnakeTail)
				angle = math.Pi / 2
			case space.Vec2I{Y: -1}:
				img = assets.Image(assets.SnakeTail)
				angle = -math.Pi / 2
			default:
				panic("unreachable")
			}
		default:
			prev := g.snake[i-1]
			next := g.snake[i+1]

			img = assets.Image(assets.SnakeBody)

			switch {
			// Vertical
			case prev.X == pos.X && next.X == pos.X:
				img = assets.Image(assets.SnakeBody)
				angle = math.Pi / 2
				if prev.Y > next.Y {
					angle *= -1
				}
			// Horizontal
			case prev.Y == pos.Y && next.Y == pos.Y:
				img = assets.Image(assets.SnakeBody)
				if prev.X > next.X {
					horizontalFlip = true
				}
			// Turn
			default:
				img = assets.Image(assets.SnakeBodyTurn)

				d1 := prev.Sub(pos)
				d2 := next.Sub(pos)

				switch {
				case d1 == space.Vec2I{Y: 1} && d2 == space.Vec2I{X: -1}:
					angle = math.Pi / 2
				case d1 == space.Vec2I{Y: -1} && d2 == space.Vec2I{X: -1}:
					angle = -math.Pi / 2
					horizontalFlip = true
				case d1 == space.Vec2I{X: 1} && d2 == space.Vec2I{Y: -1}:
					verticalFlip = true
				case d1 == space.Vec2I{X: -1} && d2 == space.Vec2I{Y: 1}:
					horizontalFlip = true
				case d1 == space.Vec2I{Y: 1} && d2 == space.Vec2I{X: 1}:
					// No changes
				case d1 == space.Vec2I{X: 1} && d2 == space.Vec2I{Y: 1}:
					angle = -math.Pi / 2
					verticalFlip = true
				case d1 == space.Vec2I{X: -1} && d2 == space.Vec2I{Y: -1}:
					angle = -math.Pi / 2
					horizontalFlip = true
				case d1 == space.Vec2I{Y: -1} && d2 == space.Vec2I{X: 1}:
					angle = -math.Pi / 2
				default:
					panic("unreachable")
				}
			}
		}

		op := &ebiten.DrawImageOptions{}

		imgCenter := space.ImageSize(img).ToF().Scale(0.5)
		op.GeoM = space.Rotate(op.GeoM, imgCenter, angle)
		op.GeoM = space.Flip(op.GeoM, imgCenter, horizontalFlip, verticalFlip)

		op.GeoM = space.ImgResizeTo(op.GeoM, img, cellSize)
		op.GeoM = space.Translate(op.GeoM, pos.ToF().Mul(cellSize).Add(cellOffset))

		screen.DrawImage(img, op)
	}

	if g.gameOver {
		moveTo := g.snake[0].ToF().Add(g.dir.ToF())
		pencil.StrokeRectV(
			screen,
			moveTo.Mul(cellSize).Add(cellOffset),
			cellSize,
			2,
			colornames.Red,
		)
	}

	debugger.Draw(screen)
}
