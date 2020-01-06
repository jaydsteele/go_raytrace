package scene

import (
	"math"

	"github.com/jaydsteele/go_raytrace/geom"
)

// Camera defines a Camera in the scene
type Camera struct {
	Origin          geom.Vec3
	LowerLeftCorner geom.Vec3
	Horizontal      geom.Vec3
	Vertical        geom.Vec3
}

// MakeCamera makes a Camera object with default values
func MakeCamera(lookFrom, lookAt, vup geom.Vec3, vfov, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := lookFrom.Sub(lookAt).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)
	return Camera{
		LowerLeftCorner: lookFrom.Sub(u.Mul(halfWidth)).Sub(v.Mul(halfHeight)).Sub(w),
		Horizontal:      u.Mul(2 * halfWidth),
		Vertical:        v.Mul(2 * halfHeight),
		Origin:          lookFrom,
	}
}

// GetRay provides a Ray extending from the Camera at the specified (u,v)-coordinate in the view
func (c *Camera) GetRay(u, v float64) geom.Ray {
	return geom.Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeftCorner.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v)).Sub(c.Origin),
	}
}
