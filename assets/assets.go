package assets

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/candlestick-games/snake/pkg/std/loader"
)

var (
	assetFS fs.FS

	assetImageCache    = map[ImageID]*ebiten.Image{}
	assetFontCache     = map[FontID]*opentype.Font{}
	assetFontFaceCache = map[FontFaceID]font.Face{}
)

type (
	ImageID    string
	FontID     string
	FontFaceID struct {
		Font FontID
		Size float64
	}
)

const (
	SnakeHeadSide ImageID = "img/snake-head-side.png"
	SnakeHeadTop  ImageID = "img/snake-head-top.png"
	SnakeBody     ImageID = "img/snake-body.png"
	SnakeBodyTurn ImageID = "img/snake-body-turn.png"
	SnakeTail     ImageID = "img/snake-tail.png"

	Apple ImageID = "img/apple.png"

	Floor ImageID = "img/floor.png"
	Wall  ImageID = "img/wall.png"
)

func Image(id ImageID) *ebiten.Image {
	if value, ok := assetImageCache[id]; ok {
		return value
	}
	value := loader.LoadImage(assetFS, string(id))
	assetImageCache[id] = value
	return value
}

const BackTo1982Font FontID = "font/BACKTO1982.TTF"

func Font(id FontID) *opentype.Font {
	if value, ok := assetFontCache[id]; ok {
		return value
	}
	value := loader.LoadFont(assetFS, string(id))
	assetFontCache[id] = value
	return value
}

var (
	RegularText = FontFaceID{
		Font: BackTo1982Font,
		Size: 32,
	}
	BigText = FontFaceID{
		Font: BackTo1982Font,
		Size: 48,
	}
)

func FontFace(id FontFaceID) font.Face {
	if value, ok := assetFontFaceCache[id]; ok {
		return value
	}
	value := loader.NewFontFace(Font(id.Font), id.Size)
	assetFontFaceCache[id] = value
	return value
}
