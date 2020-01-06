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
	U, V, W         geom.Vec3
	LensRadius      float64
}

// MakeCamera makes a Camera object
func MakeCamera(lookFrom, lookAt, vup geom.Vec3, vfov, aspect, aperture, focusDist float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := lookFrom.Sub(lookAt).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)
	return Camera{
		LowerLeftCorner: lookFrom.Sub(u.Mul(halfWidth * focusDist)).Sub(v.Mul(halfHeight * focusDist)).Sub(w.Mul(focusDist)),
		Horizontal:      u.Mul(2 * halfWidth * focusDist),
		Vertical:        v.Mul(2 * halfHeight * focusDist),
		Origin:          lookFrom,
		LensRadius:      aperture / 2,
		U:               u,
		V:               v,
		W:               w,
	}
}

// GetRay provides a Ray extending from the Camera at the specified (u,v)-coordinate in the view
func (c *Camera) GetRay(u, v float64) geom.Ray {
	rd := geom.RandomInUnitDisk().Mul(c.LensRadius)
	offset := c.U.Mul(rd.X).Add(c.V.Mul(rd.Y))
	return geom.Ray{
		Origin:    c.Origin.Add(offset),
		Direction: c.LowerLeftCorner.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v)).Sub(c.Origin).Sub(offset),
	}
}
