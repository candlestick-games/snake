package space

import "image"

type RectF struct {
	Pos  Vec2F
	Size Vec2F
}

func NewRectF(x1, y1, x2, y2 float64) RectF {
	return RectF{
		Pos: Vec2F{
			X: x1,
			Y: y1,
		},
		Size: Vec2F{
			X: x2 - x1,
			Y: y2 - y1,
		},
	}
}

func (r RectF) Contains(pos Vec2F) bool {
	return pos.X >= r.Pos.X && pos.X <= r.Pos.X+r.Size.X &&
		pos.Y >= r.Pos.Y && pos.Y <= r.Pos.Y+r.Size.Y
}

func (r RectF) ToI() RectI {
	return RectI{
		Pos:  r.Pos.ToI(),
		Size: r.Size.ToI(),
	}
}

func (r RectF) ToIR() image.Rectangle {
	return r.ToI().ToIR()
}
