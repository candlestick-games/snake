package pencil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/candlestick-games/snake/pkg/std/space"
)

func Line(dst *ebiten.Image, x1, y1, x2, y2, thickness float64, clr color.Color) {
	if thickness <= 0 {
		return
	}

	vector.StrokeLine(
		dst,
		float32(x1),
		float32(y1),
		float32(x2),
		float32(y2),
		float32(thickness),
		clr,
		false,
	)
}

func FillRect(dst *ebiten.Image, x, y, width, height float64, clr color.Color) {
	if width == 0 || height == 0 {
		return
	}

	vector.DrawFilledRect(
		dst,
		float32(x),
		float32(y),
		float32(width),
		float32(height),
		clr,
		false,
	)
}

func FillRectV(dst *ebiten.Image, pos space.Vec2F, size space.Vec2F, clr color.Color) {
	FillRect(dst, pos.X, pos.Y, size.X, size.Y, clr)
}

func StrokeRect(dst *ebiten.Image, x, y, width, height, thickness float64, clr color.Color) {
	if width == 0 || height == 0 || thickness == 0 {
		return
	}

	vector.StrokeRect(
		dst,
		float32(x),
		float32(y),
		float32(width),
		float32(height),
		float32(thickness),
		clr,
		false,
	)
}

func StrokeRectV(dst *ebiten.Image, pos space.Vec2F, size space.Vec2F, thickness float64, clr color.Color) {
	StrokeRect(dst, pos.X, pos.Y, size.X, size.Y, thickness, clr)
}

func StrokeRectR(dst *ebiten.Image, rect space.RectF, thickness float64, clr color.Color) {
	StrokeRect(dst, rect.Pos.X, rect.Pos.Y, rect.Size.X, rect.Size.Y, thickness, clr)
}

func StrokeCircle(dst *ebiten.Image, cx, cy, r, thickness float64, clr color.Color) {
	if r == 0 || thickness == 0 {
		return
	}

	vector.StrokeCircle(
		dst,
		float32(cx),
		float32(cy),
		float32(r),
		float32(thickness),
		clr,
		false,
	)
}
