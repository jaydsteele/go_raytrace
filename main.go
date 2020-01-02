package main

import (
	"fmt"
	"math"

	"github.com/jaydsteele/go_raytrace/geom"
)

func hitSphere(center geom.Vec3, radius float64, r geom.Ray) float64 {
	oc := r.Origin.Sub(center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1
	} else {
		return (-b - math.Sqrt(discriminant)) / (2 * a)
	}
}

func color(r geom.Ray) geom.Vec3 {
	t := hitSphere(geom.V3(0, 0, -1), 0.5, r)
	if t > 0 {
		N := r.PointAtParameter(t).Sub(geom.V3(0, 0, -1)).Unit()
		return geom.V3(N.X+1, N.Y+1, N.Z+1).Mul(0.5)
	}
	unitDirection := r.Direction.Unit()
	t = 0.5 * (unitDirection.Y + 1.0)
	return geom.V3Unit.Mul(1 - t).Add(geom.V3(0.5, 0.7, 1.0).Mul(t))
}

func main() {
	nx := 200
	ny := 100
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	lowerLeftCorner := geom.Vec3{X: -2.0, Y: -1.0, Z: -1.0}
	horizontal := geom.Vec3{X: 4.0, Y: 0.0, Z: 0.0}
	vertical := geom.Vec3{X: 0.0, Y: 2.0, Z: 0.0}
	origin := geom.V3Zero

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := geom.Ray{
				Origin:    origin,
				Direction: lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v)),
			}
			col := color(r)
			ir := int32(255.99 * col.R())
			ig := int32(255.99 * col.G())
			ib := int32(255.99 * col.B())
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
