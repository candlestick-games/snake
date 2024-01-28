package assets

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/candlestick-games/snake/pkg/std/loader"
)

var (
	assetFS    fs.FS
	assetCache = map[ID]any{}
)

type ID string

const (
	SnakeHeadSide ID = "img/snake-head-side.png"
	SnakeHeadTop  ID = "img/snake-head-top.png"
	SnakeBody     ID = "img/snake-body.png"
	SnakeBodyTurn ID = "img/snake-body-turn.png"
	SnakeTail     ID = "img/snake-tail.png"
)

func Image(id ID) *ebiten.Image {
	if value, ok := assetCache[id]; ok {
		return value.(*ebiten.Image)
	}
	value := loader.LoadImage(assetFS, string(id))
	assetCache[id] = value
	return value
}
