package main

import (
	"fmt"
	"math"

	"github.com/jaydsteele/go_raytrace/geom"
	"github.com/jaydsteele/go_raytrace/scene"
)

func color(r geom.Ray, world scene.Hitable) geom.Vec3 {
	rec := scene.HitRecord{}
	if world.Hit(r, 0.0, math.MaxFloat64, &rec) {
		return rec.Normal.Add(geom.V3Unit).Mul(0.5)
	}
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)
	return geom.V3Unit.Mul(1 - t).Add(geom.V3(0.5, 0.7, 1.0).Mul(t))
}

func main() {
	nx := 400
	ny := 200
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	lowerLeftCorner := geom.Vec3{X: -2.0, Y: -1.0, Z: -1.0}
	horizontal := geom.Vec3{X: 4.0, Y: 0.0, Z: 0.0}
	vertical := geom.Vec3{X: 0.0, Y: 2.0, Z: 0.0}
	origin := geom.V3Zero

	world := scene.HitableList{}
	world.Add(scene.Sphere{Center: geom.V3(0, 0, -1), Radius: 0.5})
	world.Add(scene.Sphere{Center: geom.V3(0, -100.5, -1), Radius: 100})

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := geom.Ray{
				Origin:    origin,
				Direction: lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v)),
			}

			col := color(r, &world)
			ir := int32(255.99 * col.R())
			ig := int32(255.99 * col.G())
			ib := int32(255.99 * col.B())
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
