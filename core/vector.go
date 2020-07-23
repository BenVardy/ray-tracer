package core

import "math"

// Vector3 is the main vector class
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// Add adds two vectors together
func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Subtract subtracts one vector from another
func (a Vector3) Subtract(b Vector3) Vector3 {
	return Vector3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Smult multiples a vector by a scalar
func (a Vector3) Smult(n float64) Vector3 {
	return Vector3{n * a.X, n * a.Y, n * a.Z}
}

// Dot takes the dot product of two vectors
func (a Vector3) Dot(b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross takes the cross product of two vectors
func (a Vector3) Cross(b Vector3) Vector3 {
	return Vector3{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

// Length returns the mod of a vector
func (a Vector3) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Normalize returns a vector
func (a Vector3) Normalize() Vector3 {
	mod := a.Length()
	// Don't try to div by 0
	if mod == 0 {
		return a
	}

	return a.Smult(1 / mod)
}

// AsSlice gets the vector as a slice to index
func (a Vector3) AsSlice() []float64 {
	return []float64{a.X, a.Y, a.Z}
}
