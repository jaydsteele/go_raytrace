package geom

import "math"

// Sphere has a center and radius
type Sphere struct {
	Center Vec3
	Radius float64
}

// Hit calculates whether this object has been hit by the Ray
func (s Sphere) Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.Normal = rec.p.Sub(s.Center).Div(s.Radius)
			return true
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.Normal = rec.p.Sub(s.Center).Div(s.Radius)
			return true
		}
	}
	return false
}
