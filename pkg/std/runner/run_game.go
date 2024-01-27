package runner

import (
	"github.com/hajimehoshi/ebiten/v2"

	_ "github.com/candlestick-games/snake/pkg/std/logger"
)

type GameRunner interface {
	Init() error
	ebiten.Game
	Shutdown()
}

type gracefulGameRunner struct {
	game GameRunner
}

func (g *gracefulGameRunner) Init() error {
	return g.game.Init()
}

func (g *gracefulGameRunner) Run() error {
	return ebiten.RunGame(g.game)
}

func (g *gracefulGameRunner) Shutdown() {
	g.game.Shutdown()
}

func RunGame(game GameRunner) {
	RunGraceful(&gracefulGameRunner{
		game: game,
	})
}
