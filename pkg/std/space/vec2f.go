package space

import "fmt"

type Vec2F struct {
	X float64
	Y float64
}

func NewVec2F(points ...float64) Vec2F {
	switch len(points) {
	case 0:
		return Vec2F{}
	case 1:
		return Vec2F{X: points[0], Y: points[0]}
	case 2:
		return Vec2F{X: points[0], Y: points[1]}
	default:
		panic(fmt.Errorf("invalid number of points: %d", len(points)))
	}
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

func (v Vec2F) ScaleX(s float64) Vec2F {
	return Vec2F{
		X: v.X * s,
		Y: v.Y,
	}
}

func (v Vec2F) ScaleY(s float64) Vec2F {
	return Vec2F{
		X: v.X,
		Y: v.Y * s,
	}
}
