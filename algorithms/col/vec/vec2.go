package vec

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Signed | constraints.Float
}

type V2D[T Number] struct {
	X T
	Y T
}

func New2D[T Number](x, y T) V2D[T] {
	return V2D[T]{
		X: x,
		Y: y,
	}
}

func Origin2[T Number]() V2D[T] {
	return New2D[T](0, 0)
}

func UnitNorth2[T Number]() V2D[T] {
	return New2D[T](0, 1)
}

func UnitWest2[T Number]() V2D[T] {
	return New2D[T](-1, 0)
}

func UnitSouth2[T Number]() V2D[T] {
	return New2D[T](0, -1)
}

func UnitEast2[T Number]() V2D[T] {
	return New2D[T](1, 0)
}

func (v V2D[T]) Dist(other V2D[T]) float64 {
	return math.Sqrt(
		math.Pow(float64(v.X-other.X), 2) +
			math.Pow(float64(v.Y-other.Y), 2))
}

func (v V2D[T]) DistOrigin() float64 {
	return v.Dist(Origin2[T]())
}

func (v V2D[T]) Add(other *V2D[T]) V2D[T] {
	v.X += other.X
	v.Y += other.Y

	return v
}

func (v V2D[T]) Sub(other V2D[T]) V2D[T] {
	v.X -= other.X
	v.Y -= other.Y

	return v
}

func (v V2D[T]) Mult(scalar T) V2D[T] {
	v.X *= scalar
	v.Y *= scalar

	return v
}

func (v V2D[T]) Div(scalar T) V2D[T] {
	v.X /= scalar
	v.Y /= scalar

	return v
}

func (v V2D[T]) Eq(other V2D[T]) bool {
	return v.X == other.X && v.Y == other.Y
}
