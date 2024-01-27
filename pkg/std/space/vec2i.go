package space

import "fmt"

type Vec2I struct {
	X int
	Y int
}

func NewVec2I(points ...int) Vec2I {
	switch len(points) {
	case 0:
		return Vec2I{}
	case 1:
		return Vec2I{X: points[0], Y: points[0]}
	case 2:
		return Vec2I{X: points[0], Y: points[1]}
	default:
		panic(fmt.Errorf("invalid number of points: %d", len(points)))
	}
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

func (v Vec2I) Shrink(s int) Vec2I {
	return Vec2I{
		X: v.X / s,
		Y: v.Y / s,
	}
}
