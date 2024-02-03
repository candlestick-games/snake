package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/viper"

	"github.com/candlestick-games/snake/pkg/std/collection"
	"github.com/candlestick-games/snake/pkg/std/space"
	"github.com/candlestick-games/snake/pkg/std/tick"
)

func (g *Game) Init() error {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Snake: Might and Magic")
	ebiten.SetFullscreen(!viper.GetBool("window"))
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g.ticker = tick.NewTicker()

	g.resetSnake()

	return nil
}

func (g *Game) resetSnake() {
	g.gridCols = 16 * 2
	g.gridRows = 9 * 2

	g.placeWalls()

	g.placeSnake()

	g.food = collection.NewSet[space.Vec2I]()
	g.placeFood()

	g.gameOver = false
	g.startTime = g.ticker.StartTimer(60 * 3)
}

func (g *Game) placeSnake() {
	g.snake = make([]space.Vec2I, 0, 3)

	snakeHead := space.NewVec2I(-1)
	for snakeHead.X < 0 || g.isOccupied(snakeHead) {
		snakeHead = space.RandomVec2I(0, g.gridCols, 0, g.gridRows)
	}
	g.snake = append(g.snake, snakeHead)

	snakeBody := space.NewVec2I(-1)
	for snakeBody.X < 0 || g.isSnake(snakeBody) {
		snakeBody = g.randomUnoccupiedNeighbour(snakeHead)
	}
	g.snake = append(g.snake, snakeBody)

	snakeTail := space.NewVec2I(-1)
	for snakeTail.X < 0 || g.isSnake(snakeTail) {
		snakeTail = g.randomUnoccupiedNeighbour(snakeBody)
	}
	g.snake = append(g.snake, snakeTail)

	g.dir = snakeHead.Sub(snakeBody)
	g.prevDir = g.dir

	if g.isOccupied(snakeHead.Add(g.dir)) {
		g.placeSnake()
	}
}

func (g *Game) placeFood() {
	for {
		pos := space.RandomVec2I(0, g.gridCols-1, 0, g.gridRows-1)

		if g.food.Has(pos) {
			continue
		}

		if g.walls[pos.Y][pos.X] {
			continue
		}

		var ok bool
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
