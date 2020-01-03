package scene

import "github.com/jaydsteele/go_raytrace/geom"

// Camera defines a Camera in the scene
type Camera struct {
	Origin          geom.Vec3
	LowerLeftCorner geom.Vec3
	Horizontal      geom.Vec3
	Vertical        geom.Vec3
}

// MakeCamera makes a Camera object with default values
func MakeCamera() Camera {
	return Camera{
		LowerLeftCorner: geom.V3(-2, -1, -1),
		Horizontal:      geom.V3(4, 0, 0),
		Vertical:        geom.V3(0, 2, 0),
		Origin:          geom.V3Zero,
	}
}

// GetRay provides a Ray extending from the Camera at the specified (u,v)-coordinate in the view
func (c *Camera) GetRay(u, v float64) geom.Ray {
	return geom.Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeftCorner.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v)).Sub(c.Origin),
	}
}
