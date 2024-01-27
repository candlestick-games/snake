package space

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/candlestick-games/snake/pkg/std/rand"
)

func RandomVec2I(minX, maxX, minY, maxY int) Vec2I {
	return Vec2I{
		X: rand.Int(minX, maxX),
		Y: rand.Int(minY, maxY),
	}
}

func RandomVec2F(minX, maxX, minY, maxY float64) Vec2F {
	return Vec2F{
		X: rand.Float(minX, maxX),
		Y: rand.Float(minY, maxY),
	}
}

func RandomNormalVec2F(minX, maxX, minY, maxY float64) Vec2F {
	return RandomVec2F(0, 1, 0, 1)
}

func CursorPos() Vec2I {
	pos := Vec2I{}
	pos.X, pos.Y = ebiten.CursorPosition()
	return pos
}

func ImageSize(img *ebiten.Image) Vec2I {
	bounds := img.Bounds()
	return Vec2I{
		X: bounds.Max.X,
		Y: bounds.Max.Y,
	}
}
