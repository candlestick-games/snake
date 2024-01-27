package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/spf13/viper"
	"golang.org/x/image/colornames"

	"github.com/candlestick-games/snake/pkg/std/collection"
	"github.com/candlestick-games/snake/pkg/std/debugger"
	"github.com/candlestick-games/snake/pkg/std/pencil"
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
	canMove  bool
	gameOver bool
	ticks    uint

	food collection.Set[space.Vec2I]

	screenWidth  float64
	screenHeight float64

	quit bool
}

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

func (g *Game) placeFood() {
	for {
		pos := space.RandomVec2I(1, cellCols-1, 1, cellRows-1)

		ok := g.food.Has(pos)
		if ok {
			continue
		}
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

func (g *Game) Update() error {
	debugger.Update()

	// Quit handler
	{
		if ebiten.IsWindowBeingClosed() {
			g.quit = true
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.quit = true
		}
		if g.quit {
			return ebiten.Termination
		}
	}

	// Game ticks
	g.ticks++

	// Snake controls & movement
	g.HandleControls()
	g.MoveSnake()

	return nil
}

func (g *Game) HandleControls() {
	if !g.canMove {
		return
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyUp) && g.dir.Y == 0:
		g.dir = space.Vec2I{X: 0, Y: -1}
		g.canMove = false
	case inpututil.IsKeyJustPressed(ebiten.KeyRight) && g.dir.X == 0:
		g.dir = space.Vec2I{X: 1, Y: 0}
		g.canMove = false
	case inpututil.IsKeyJustPressed(ebiten.KeyDown) && g.dir.Y == 0:
		g.dir = space.Vec2I{X: 0, Y: 1}
		g.canMove = false
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.dir.X == 0:
		g.dir = space.Vec2I{X: -1, Y: 0}
		g.canMove = false
	}
}

func (g *Game) MoveSnake() {
	if g.gameOver || g.ticks%20 != 0 {
		return
	}

	g.canMove = true

	head := g.snake[0]
	newHead := head.Add(g.dir)

	if newHead.X < 0 || newHead.X >= cellCols || newHead.Y < 0 || newHead.Y >= cellRows {
		g.gameOver = true
		return
	}

	tail := g.snake[len(g.snake)-1]
	copy(g.snake[1:], g.snake[:len(g.snake)-1])
	g.snake[0] = newHead

	for i, s := range g.snake {
		if i != 0 && s == newHead {
			g.gameOver = true
			return
		}
	}

	for foodPos := range g.food {
		if newHead == foodPos {
			g.snake = append(g.snake, tail)
			g.food.Remove(foodPos)
			g.placeFood()
		}
	}
}

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

	debugger.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	panic("unreachable")
}

func (g *Game) Resize(screenWidth, screenHeight float64) {
	g.screenWidth = screenWidth
	g.screenHeight = screenHeight

	g.boardBounds = space.NewRectF(32, 32, screenWidth-32, screenHeight-32)
	g.cellSize = min(g.boardBounds.Size.X/cellCols, g.boardBounds.Size.Y/cellRows)
	g.boardOffset = space.Vec2F{
		X: (g.boardBounds.Size.X - g.cellSize*cellCols) / 2,
		Y: (g.boardBounds.Size.Y - g.cellSize*cellRows) / 2,
	}
}

func (g *Game) LayoutF(screenWidth, screenHeight float64) (float64, float64) {
	if g.screenWidth != screenWidth || g.screenHeight != screenHeight {
		g.Resize(screenWidth, screenHeight)
	}
	return screenWidth, screenHeight
}

func (g *Game) Shutdown() {}
