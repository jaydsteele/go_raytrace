package geom

import (
	"math"
)

// Vec3 is a generic 3-tuple representation
type Vec3 struct {
	X, Y, Z float64
}

var (
	// V3Zero is the zero vector (0,0,0)
	V3Zero = Vec3{0, 0, 0}
	// V3Unit is the unit vector (1,1,1)
	V3Unit = Vec3{1, 1, 1}
	// V3UnitX is the x-axis unit vector (1,0,0)
	V3UnitX = Vec3{1, 0, 0}
	// V3UnitY is the y-axis unit vector (0,1,0)
	V3UnitY = Vec3{0, 1, 0}
	// V3UnitZ is the z-axis unit vector (0,0,1)
	V3UnitZ = Vec3{0, 0, 1}
)

// V3 is shorthand for Vec3{X: x, Y: y, Z: z}
func V3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
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

// Add returns the vector v+w.
func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Sub returns the vector v-w.
func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Mul returns the vector v*s.
func (v Vec3) Mul(s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

// Div returns the vector v/s.
func (v Vec3) Div(s float64) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

// Neg returns the negated vector of v.
func (v Vec3) Neg() Vec3 {
	return v.Mul(-1)
}

// Dot returns the dot (a.k.a. scalar) product of v and w.
func (v Vec3) Dot(w Vec3) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Cross returns the cross product of v and w.
func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}

// CompMul returns the component-wise multiplication of two vectors.
func (v Vec3) CompMul(w Vec3) Vec3 {
	return Vec3{v.X * w.X, v.Y * w.Y, v.Z * w.Z}
}

// CompDiv returns the component-wise division of two vectors.
func (v Vec3) CompDiv(w Vec3) Vec3 {
	return Vec3{v.X / w.X, v.Y / w.Y, v.Z / w.Z}
}

// SqDist returns the square of the euclidian distance between two vectors.
func (v Vec3) SqDist(w Vec3) float64 {
	return v.Sub(w).SqLen()
}

// Dist returns the euclidian distance between two vectors.
func (v Vec3) Dist(w Vec3) float64 {
	return v.Sub(w).Len()
}

// SqLen returns the square of the length (euclidian norm) of a vector.
func (v Vec3) SqLen() float64 {
	return v.Dot(v)
}

// Len returns the length (euclidian norm) of a vector.
func (v Vec3) Len() float64 {
	return float64(math.Sqrt(float64(v.SqLen())))
}

// Norm returns the normalized vector of a vector.
func (v Vec3) Norm() Vec3 {
	return v.Div(v.Len())
}

// Reflect returns the reflection vector of v given a normal n.
func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Sub(n.Mul(2 * v.Dot(n)))
}

// Lerp returns the linear interpolation between v and w by amount t.
// The amount t is usually a value between 0 and 1. If t=0 v will be
// returned; if t=1 w will be returned.
func (v Vec3) Lerp(w Vec3, t float64) Vec3 {
	return Vec3{lerp(v.X, w.X, t), lerp(v.Y, w.Y, t), lerp(v.Z, w.Z, t)}
}

// Unit returns the unit vector of the Vec3
func (v Vec3) Unit() Vec3 {
	return v.Div(v.Len())
}
