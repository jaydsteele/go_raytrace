package scene

import (
	"math/rand"

	"github.com/jaydsteele/go_raytrace/geom"
)

// DialectricMaterial defines a glass material
type DialectricMaterial struct {
	RefIdx float64
}

// Scatter function for DialectricMaterial
func (m *DialectricMaterial) Scatter(rayIn geom.Ray, rec HitRecord, attenuation *geom.Vec3, scattered *geom.Ray) bool {
	outwardNormal := geom.Vec3{}
	reflected := geom.Reflect(rayIn.Direction, rec.Normal)
	var niOverNt float64
	*attenuation = geom.V3(1, 1, 1)
	refracted := geom.Vec3{}
	var reflectProb float64
	var cosine float64
	if rayIn.Direction.Dot(rec.Normal) > 0 {
		outwardNormal = rec.Normal.Neg()
		niOverNt = m.RefIdx
		cosine = m.RefIdx * rayIn.Direction.Dot(rec.Normal) / rayIn.Direction.Len()
	} else {
		outwardNormal = rec.Normal
		niOverNt = 1 / m.RefIdx
		cosine = -rayIn.Direction.Dot(rec.Normal) / rayIn.Direction.Len()
	}
	if geom.Refract(rayIn.Direction, outwardNormal, niOverNt, &refracted) {
		// *scattered = geom.Ray{Origin: rec.P, Direction: refracted}
		reflectProb = geom.Schlick(cosine, m.RefIdx)
	} else {
		reflectProb = 1
	}
	if rand.Float64() < reflectProb {
		*scattered = geom.Ray{Origin: rec.P, Direction: reflected}
	} else {
		*scattered = geom.Ray{Origin: rec.P, Direction: refracted}
	}
	return true
}
