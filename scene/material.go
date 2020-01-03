package scene

import (
	"github.com/jaydsteele/go_raytrace/geom"
)

// Material interface defines a Scatter method for any Material
type Material interface {
	Scatter(rayIn geom.Ray, hitRecord HitRecord, attenuation *geom.Vec3, scattered *geom.Ray) bool
}
