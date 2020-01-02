package geom

// Ray has an origin and direction and helper methods
type Ray struct {
	Origin, Direction Vec3
}

// PointAtParameter returns the parameterized Vec3 point along the Ray given by t
func (r *Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
