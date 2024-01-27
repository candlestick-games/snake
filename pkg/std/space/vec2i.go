package space

type Vec2I struct {
	X int
	Y int
}

func (v Vec2I) ToF() Vec2F {
	return Vec2F{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}

func (v Vec2I) Add(o Vec2I) Vec2I {
	return Vec2I{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vec2I) Sub(o Vec2I) Vec2I {
	return Vec2I{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vec2I) Mul(o Vec2I) Vec2I {
	return Vec2I{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v Vec2I) Div(o Vec2I) Vec2I {
	return Vec2I{
		X: v.X / o.X,
		Y: v.Y / o.Y,
	}
}

func (v Vec2I) Scale(s int) Vec2I {
	return Vec2I{
		X: v.X * s,
		Y: v.Y * s,
	}
}
