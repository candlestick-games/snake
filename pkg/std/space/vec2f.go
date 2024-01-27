package space

type Vec2F struct {
	X float64
	Y float64
}

func (v Vec2F) ToI() Vec2I {
	return Vec2I{
		X: int(v.X),
		Y: int(v.Y),
	}
}

func (v Vec2F) Add(o Vec2F) Vec2F {
	return Vec2F{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vec2F) Sub(o Vec2F) Vec2F {
	return Vec2F{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vec2F) Mul(o Vec2F) Vec2F {
	return Vec2F{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v Vec2F) Div(o Vec2F) Vec2F {
	return Vec2F{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}

func (v Vec2F) Scale(s float64) Vec2F {
	return Vec2F{
		X: v.X * s,
		Y: v.Y * s,
	}
}
