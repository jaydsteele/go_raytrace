package geom

import (
	"math"
	"math/rand"
)

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

// Reflect a vector according to the specified normal
func Reflect(v, n Vec3) Vec3 {
	return v.Sub(n.Mul(2 * v.Dot(n)))
}

// Refract a vector according to the specified parameters
func Refract(v, n Vec3, niOverNt float64, refracted *Vec3) bool {
	uv := v.Unit()
	dt := uv.Dot(n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		*refracted = uv.Sub(n.Mul(dt)).Mul(niOverNt).Sub(n.Mul(math.Sqrt(discriminant)))
		return true
	}
	return false
}

// Schlick is an approximation for glass reflectivity
func Schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1+r0)*math.Pow(1-cosine, 5)
}
