package vec

import (
	"math"
)

type Vec2 struct {
	X float64
	Y float64
}

func New2D(x, y float64) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

func Origin2() Vec2 {
	return New2D(0, 0)
}

func UnitNorth2() Vec2 {
	return New2D(0, 1)
}

func UnitWest2() Vec2 {
	return New2D(-1, 0)
}

func UnitSouth2() Vec2 {
	return New2D(0, -1)
}

func UnitEast2() Vec2 {
	return New2D(1, 0)
}

func (v Vec2) Dist(other Vec2) float64 {
	return math.Sqrt(math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2))
}

func (v Vec2) DistOrigin() float64 {
	return v.Dist(Origin2())
}

func (v Vec2) Add(other *Vec2) Vec2 {
	v.X += other.X
	v.Y += other.Y

	return v
}

func (v Vec2) Sub(other Vec2) Vec2 {
	v.X -= other.X
	v.Y -= other.Y

	return v
}

func (v Vec2) Mult(scalar float64) Vec2 {
	v.X *= scalar
	v.Y *= scalar

	return v
}

func (v Vec2) Div(scalar float64) Vec2 {
	v.X /= scalar
	v.Y /= scalar

	return v
}
