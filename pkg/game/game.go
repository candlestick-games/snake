package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/spf13/viper"

	"github.com/candlestick-games/snake/pkg/std/debugger"
)

type Game struct {
	quit bool
}

func (g *Game) Init() error {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Snake: Might and Magic")
	ebiten.SetFullscreen(!viper.GetBool("window"))
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return nil
}

func (g *Game) Update() error {
	debugger.Update()

	if ebiten.IsWindowBeingClosed() {
		g.quit = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.quit = true
	}
	if g.quit {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	debugger.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	panic("unreachable")
}

func (g *Game) LayoutF(screenWidth, screenHeight float64) (float64, float64) {
	return screenWidth, screenHeight
}

func (g *Game) Shutdown() {}
