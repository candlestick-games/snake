package space

type RectF struct {
	Pos  Vec2F
	Size Vec2F
}

func (r RectF) Contains(pos Vec2F) bool {
	return pos.X >= r.Pos.X && pos.X <= r.Pos.X+r.Size.X &&
		pos.Y >= r.Pos.Y && pos.Y <= r.Pos.Y+r.Size.Y
}
