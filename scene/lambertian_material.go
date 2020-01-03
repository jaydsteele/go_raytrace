package scene

import (
	"github.com/jaydsteele/go_raytrace/geom"
)

// LambertianMaterial defines a diffuse material
type LambertianMaterial struct {
	Albedo geom.Vec3
}

// Scatter method for LambertianMaterial
func (m *LambertianMaterial) Scatter(rayIn geom.Ray, rec HitRecord, attenuation *geom.Vec3, scattered *geom.Ray) bool {
	target := rec.P.Add(rec.Normal).Add(geom.RandomInUnitSphere())
	*scattered = geom.Ray{
		Origin:    rec.P,
		Direction: target.Sub(rec.P),
	}
	*attenuation = m.Albedo
	return true
}
