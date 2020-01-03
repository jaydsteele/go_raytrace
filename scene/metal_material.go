package scene

import "github.com/jaydsteele/go_raytrace/geom"

// MetalMaterial provides rendering behavior for metal
type MetalMaterial struct {
	Albedo geom.Vec3
	Fuzz   float64
}

// Scatter method for MetalMaterial
func (m *MetalMaterial) Scatter(rayIn geom.Ray, rec HitRecord, attenuation *geom.Vec3, scattered *geom.Ray) bool {
	reflected := geom.Reflect(rayIn.Direction.Unit(), rec.Normal)
	*scattered = geom.Ray{
		Origin:    rec.P,
		Direction: reflected.Add(geom.RandomInUnitSphere().Mul(m.Fuzz)),
	}
	*attenuation = m.Albedo
	return scattered.Direction.Dot(rec.Normal) > 0
}
