package geom

import (
	"errors"
	"math"
)

// Vec3 is a generic 3-tuple representation
type Vec3 struct {
	X, Y, Z float64
}

// R returns the first coordinate of the Vec3, a convenience when working with colors
func (v *Vec3) R() float64 {
	return v.X
}

// G returns the second coordinate of the Vec3, a convenience when working with colors
func (v *Vec3) G() float64 {
	return v.Y
}

// B returns the third coordinate of the Vec3, a convenience when working with colors
func (v *Vec3) B() float64 {
	return v.Z
}

// Negate the Vec3 and return the new version
func (v *Vec3) Negate(a *Vec3, b *Vec3) Vec3 {
	return Vec3{-a.X, -a.Y, -b.Z}
}

// ElementAt returns the tuple value at the specified index
func (v *Vec3) ElementAt(index int) (float64, error) {
	if index == 0 {
		return v.X, nil
	}
	if index == 1 {
		return v.Y, nil
	}
	if index == 2 {
		return v.Z, nil
	}
	return 0, errors.New("Index must be >=0, <=2")
}

// Add the specified Vec3 to this Vec3
func (v *Vec3) Add(w *Vec3) {
	v.X += w.X
	v.Y += w.Y
	v.Z += w.Z
}

// Subtract the specified Vec3 from this Vec3
func (v *Vec3) Subtract(w *Vec3) {
	v.X -= w.X
	v.Y -= w.Y
	v.Z -= w.Z
}

// Multiply the specified Vec3 to this Vec3
func (v *Vec3) Multiply(w *Vec3) {
	v.X *= w.X
	v.Y *= w.Y
	v.Z *= w.Z
}

// Divide the specified Vec3 from this Vec3
func (v *Vec3) Divide(w *Vec3) {
	v.X /= w.X
	v.Y /= w.Y
	v.Z /= w.Z
}

// Scale this Vec3 by the specified value
func (v *Vec3) Scale(s float64) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}

// Length of this Vec3
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared of this Vec3
func (v *Vec3) LengthSquared() float64 {
	return Dot(v, v)
}

// MakeUnitVector converts this Vec3 into a unit Vector
func (v *Vec3) MakeUnitVector() {
	v.Scale(1.0 / v.Length())
}

// Dot product
func Dot(a, b *Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z + b.Z
}

// Cross product
func Cross(v1, v2 *Vec3) Vec3 {
	return Vec3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		-(v1.X*v2.Z - v1.Z*v2.X),
		v1.X*v2.Y - v1.Y*v2.X,
	}
}
