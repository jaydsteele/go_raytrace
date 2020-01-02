package geom

import "math/rand"

// lerp returns the linear interpolation between a and b by amount t.
// The amount t is usually a value between 0 and 1. If t=0 a will be returned;
// if t=1 b will be returned.
func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

const epsilon = 1e-5

// RandomInUnitSphere returns a random point in a unit radius sphere centered at the origin
func RandomInUnitSphere() Vec3 {
	result := Vec3{}
	for true {
		result = V3(rand.Float64(), rand.Float64(), rand.Float64()).Mul(2).Sub(V3Unit)
		result = result.Mul(2).Sub(V3Unit)
		if result.SqLen() < 1 {
			break
		}
	}
	return result
}
