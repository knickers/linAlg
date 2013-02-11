package linAlg

import (
	"math"
)

type Point struct {
	X, Y, Z float64
}

func (p *Point) Add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
		Z: p.Z + q.Z,
	}
}

func (p *Point) Sub(q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
		Z: p.Z - q.Z,
	}
}

type Vector struct {
	X, Y, Z float64
}

func (v *Vector) Length() float64 {
	l := v.X * v.X
	l += v.Y * v.Y
	l += v.Z * v.Z
	return math.Sqrt(l)
}

// Return this vector of length 1
func (v *Vector) Unit() Vector {
	l := v.Length()
	return Vector{
		X: v.X / l,
		Y: v.Y / l,
		Z: v.Z / l,
	}
}

// The dot product of two vectors
func (v *Vector) Dot(w Vector) float64 {
	d := v.X * w.X
	d += v.Y * w.Y
	d += v.Z * w.Z
	return d
}

// The cross product of two vectors
func (v *Vector) Cross(w Vector) Vector {
	return Vector{
		X: v.Y*w.Z - v.Z*w.Y,
		Y: v.Z*w.X - v.X*w.Z,
		Z: v.X*w.Y - v.Y*w.X,
	}
}

func (v *Vector) Add(w Vector) Vector {
	return Vector{
		X: v.X + w.X,
		Y: v.Y + w.Y,
		Z: v.Z + w.Z,
	}
}

func (v *Vector) Sub(w Vector) Vector {
	return Vector{
		X: v.X - w.X,
		Y: v.Y - w.Y,
		Z: v.Z - w.Z,
	}
}

// Return the vector that results from rotating this vector, radians amount,
// about axis.
func (v *Vector) Rotate(radians float64, axis Vector) Vector {
	return *v
}

// The accute angle between two vectors.
func (v *Vector) Angle(w Vector) float64 {
	return math.Acos(v.Dot(w) / v.Length() / w.Length())
}
