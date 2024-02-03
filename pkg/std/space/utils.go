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

func RandomVec2IDir() Vec2I {
	switch rand.Int(0, 3) {
	case 0:
		return Vec2I{X: 1}
	case 1:
		return Vec2I{X: -1}
	case 2:
		return Vec2I{Y: 1}
	case 3:
		return Vec2I{Y: -1}
	default:
		panic("unreachable")
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

func Sign(n int) int {
	if n == 0 {
		return 0
	} else if n < 0 {
		return -1
	}
	return 1
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
