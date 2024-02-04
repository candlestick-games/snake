package space

import "image"

type RectI struct {
	Pos  Vec2I
	Size Vec2I
}

func NewRectI(x1, y1, x2, y2 int) RectI {
	return RectI{
		Pos: Vec2I{
			X: x1,
			Y: y1,
		},
		Size: Vec2I{
			X: x2 - x1,
			Y: y2 - y1,
		},
	}
}

func (r RectI) ToCoords() (int, int, int, int) {
	return r.Pos.X, r.Pos.Y, r.Pos.X + r.Size.X, r.Pos.Y + r.Size.Y
}

func (r RectI) Contains(pos Vec2I) bool {
	return pos.X >= r.Pos.X && pos.X < r.Pos.X+r.Size.X &&
		pos.Y >= r.Pos.Y && pos.Y < r.Pos.Y+r.Size.Y
}

func (r RectI) Inside(o RectI) bool {
	return o.Contains(r.Pos) && o.Contains(r.Pos.Add(r.Size.Add(NewVec2I(-1))))
}

func (r RectI) Intercepts(o RectI) bool {
	return (r.Pos.X <= o.Pos.X+o.Size.X && r.Pos.X+r.Size.X >= o.Pos.X) &&
		(r.Pos.Y <= o.Pos.Y+o.Size.Y && r.Pos.Y+r.Size.Y >= o.Pos.Y)
}

func (r RectI) Move(d Vec2I) RectI {
	return RectI{
		Pos:  r.Pos.Add(d),
		Size: r.Size,
	}
}

func (r RectI) Grow(d Vec2I) RectI {
	return RectI{
		Pos:  r.Pos,
		Size: r.Size.Add(d),
	}
}

func (r RectI) Clamp(o RectI) RectI {
	rX1, rY1, rX2, rY2 := r.ToCoords()
	oX1, oY1, oX2, oY2 := o.ToCoords()
	return NewRectI(
		min(max(rX1, oX1), oX2),
		min(max(rY1, oY1), oY2),
		max(min(rX2, oX2), oX1),
		max(min(rY2, oY2), oY1),
	)
}

func (r RectI) Center() Vec2I {
	return Vec2I{
		X: r.Pos.X + r.Size.X/2,
		Y: r.Pos.Y + r.Size.Y/2,
	}
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
