package engine

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Unpack() (x float64, y float64) {
	return v.X, v.Y
}

func (v *Vector) Set(x float64, y float64) {
	v.X = x
	v.Y = y
}

func (v *Vector) Translate(dx float64, dy float64) {
	v.X += dx
	v.Y += dy
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector) Sub(other Vector) Vector {
	return Vector{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector) Mul(factor float64) Vector {
	return Vector{
		X: v.X * factor,
		Y: v.Y * factor,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Distance(other Vector) float64 {
	return other.Sub(v).Length()
}

func (v Vector) Normalized() Vector {
	length := v.Length()
	return Vector{
		X: v.X / length,
		Y: v.Y / length,
	}
}
