package main

import (
	"fmt"
	"math"
	"math/rand"

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
	numSamples := 100
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)

	world := scene.HitableList{}
	world.Add(scene.Sphere{Center: geom.V3(0, 0, -1), Radius: 0.5})
	world.Add(scene.Sphere{Center: geom.V3(0, -100.5, -1), Radius: 100})

	cam := scene.MakeCamera()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			col := geom.V3(0, 0, 0)
			for s := 0; s < numSamples; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)

				r := cam.GetRay(u, v)
				// p := r.PointAtParameter(2)
				col = col.Add(color(r, &world))
			}
			col = col.Div(float64(numSamples))
			ir := int32(255.99 * col.R())
			ig := int32(255.99 * col.G())
			ib := int32(255.99 * col.B())
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
