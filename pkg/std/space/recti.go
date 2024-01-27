package space

import "image"

type RectI struct {
	Pos  Vec2I
	Size Vec2I
}

func (r RectI) Contains(pos Vec2I) bool {
	return pos.X >= r.Pos.X && pos.X <= r.Pos.X+r.Size.X &&
		pos.Y >= r.Pos.Y && pos.Y <= r.Pos.Y+r.Size.Y
}

func (r RectI) ToIR() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: r.Pos.X,
			Y: r.Pos.Y,
		},
		Max: image.Point{
			X: r.Pos.Y + r.Size.X,
			Y: r.Pos.Y + r.Size.Y,
		},
	}
}
