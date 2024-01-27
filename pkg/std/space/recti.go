package space

type RectI struct {
	Pos  Vec2I
	Size Vec2I
}

func (r RectI) Contains(pos Vec2I) bool {
	return pos.X >= r.Pos.X && pos.X <= r.Pos.X+r.Size.X &&
		pos.Y >= r.Pos.Y && pos.Y <= r.Pos.Y+r.Size.Y
}
