package geom

// HitRecord contains the information for a ray hit
type HitRecord struct {
	t      float64
	p      Vec3
	Normal Vec3
}

// Hitable defines the method for any hitable object in the scene
type Hitable interface {
	Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool
}
