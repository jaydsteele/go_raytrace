package scene

import "github.com/jaydsteele/go_raytrace/geom"

// HitRecord contains the information for a ray hit
type HitRecord struct {
	T      float64
	P      geom.Vec3
	Normal geom.Vec3
}

// Hitable defines the method for any hitable object in the scene
type Hitable interface {
	Hit(r geom.Ray, tMin, tMax float64, rec *HitRecord) bool
}
