package game

import (
	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/candlestick-games/snake/pkg/std/debugger"
	"github.com/candlestick-games/snake/pkg/std/space"
)

func (g *Game) Update() error {
	debugger.Update()
	g.ticker.Update()

	// Quit
	{
		if ebiten.IsWindowBeingClosed() {
			g.quit = true
		}
		if g.quit {
			return ebiten.Termination
		}
	}

	// Fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	// Game over
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.resetSnake()
		}
		return nil
	}

	// Pause game
	if !g.startTime.IsStarted() && inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.pause = !g.pause
	}
	if g.pause {
		return nil
	}

	// Snake controls & movement
	g.handleControls()
	g.moveSnake()

	return nil
}

func (g *Game) handleControls() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp) && g.prevDir.Y == 0:
		g.dir = space.Vec2I{X: 0, Y: -1}
	case ebiten.IsKeyPressed(ebiten.KeyRight) && g.prevDir.X == 0:
		g.dir = space.Vec2I{X: 1, Y: 0}
	case ebiten.IsKeyPressed(ebiten.KeyDown) && g.prevDir.Y == 0:
		g.dir = space.Vec2I{X: 0, Y: 1}
	case ebiten.IsKeyPressed(ebiten.KeyLeft) && g.prevDir.X == 0:
		g.dir = space.Vec2I{X: -1, Y: 0}
	}
}

func (g *Game) moveSnake() {
	if !g.ticker.Every(20) {
		return
	}
	if g.startTime.IsStarted() && !g.startTime.Wait() {
		return
	}

	head := g.snake[0]
	newHead := head.Add(g.dir)

	if newHead.X < 0 || newHead.X >= g.gridCols || newHead.Y < 0 || newHead.Y >= g.gridRows {
		g.gameOver = true
		log.Debug("Bounds collision", "head", newHead)
		return
	}

	if g.walls[newHead.Y][newHead.X] {
		g.gameOver = true
		log.Debug("Wall collision", "head", newHead)
		return
	}

	for i := 0; i < len(g.snake)-1; i++ {
		if g.snake[i] != newHead {
			continue
		}

		g.gameOver = true
		log.Debug("Self collision", "head", newHead)
		return
	}

	g.prevDir = g.dir

	tail := g.snake[len(g.snake)-1]
	copy(g.snake[1:], g.snake[:len(g.snake)-1])
	g.snake[0] = newHead

	for foodPos := range g.food {
		if newHead == foodPos {
			g.snake = append(g.snake, tail)
			g.food.Remove(foodPos)
			g.placeFood()
		}
	}
}
