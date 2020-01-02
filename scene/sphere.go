package scene

import (
	"math"

	"github.com/jaydsteele/go_raytrace/geom"
)

// Sphere has a center and radius
type Sphere struct {
	Center geom.Vec3
	Radius float64
}

// Hit calculates whether this object has been hit by the Ray
func (s Sphere) Hit(r geom.Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.PointAtParameter(rec.T)
			rec.Normal = rec.P.Sub(s.Center).Div(s.Radius)
			return true
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.PointAtParameter(rec.T)
			rec.Normal = rec.P.Sub(s.Center).Div(s.Radius)
			return true
		}
	}
	return false
}
