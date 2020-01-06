package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/jaydsteele/go_raytrace/geom"
	"github.com/jaydsteele/go_raytrace/scene"
)

func color(r geom.Ray, world scene.Hitable, depth int) geom.Vec3 {
	rec := scene.HitRecord{}
	if world.Hit(r, 0.001, math.MaxFloat64, &rec) {
		scattered := geom.Ray{}
		attenuation := geom.Vec3{}
		if depth < 50 && rec.Material.Scatter(r, rec, &attenuation, &scattered) {
			return attenuation.CompMul(color(scattered, world, depth+1))
		}
		return geom.V3Zero
	}
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)
	return geom.V3Unit.Mul(1 - t).Add(geom.V3(0.5, 0.7, 1.0).Mul(t))
}

func randomScene() scene.Hitable {
	world := scene.HitableList{}
	world.Add(scene.Sphere{
		Center: geom.V3(0, -1000, 0),
		Radius: 1000,
		Material: &scene.LambertianMaterial{
			Albedo: geom.V3(.5, .5, .5),
		},
	})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			center := geom.V3(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())
			if center.Sub(geom.V3(4, .2, 0)).Len() > .9 {
				chooseMat := rand.Float64()
				if chooseMat < .8 { // diffuse
					world.Add(scene.Sphere{
						Center: center,
						Radius: .2,
						Material: &scene.LambertianMaterial{
							Albedo: geom.V3(
								rand.Float64()*rand.Float64(),
								rand.Float64()*rand.Float64(),
								rand.Float64()*rand.Float64()),
						},
					})
				} else if chooseMat < .95 { // metal
					world.Add(scene.Sphere{
						Center: center,
						Radius: .2,
						Material: &scene.MetalMaterial{
							Albedo: geom.V3(
								0.5*(1+rand.Float64()),
								0.5*(1+rand.Float64()),
								0.5*(1+rand.Float64())),
							Fuzz: 0.5 * rand.Float64(),
						},
					})
				} else {
					world.Add(scene.Sphere{
						Center:   center,
						Radius:   .2,
						Material: &scene.DialectricMaterial{RefIdx: 1.5},
					})
				}
			}
		}
	}

	world.Add(scene.Sphere{
		Center:   geom.V3(0, 1, 0),
		Radius:   1,
		Material: &scene.DialectricMaterial{RefIdx: 1.5},
	})
	world.Add(scene.Sphere{
		Center: geom.V3(-4, 1, 0),
		Radius: .2,
		Material: &scene.LambertianMaterial{
			Albedo: geom.V3(0.4, 0.2, 0.1),
		},
	})
	world.Add(scene.Sphere{
		Center: geom.V3(4, 1, 0),
		Radius: 1,
		Material: &scene.MetalMaterial{
			Albedo: geom.V3(0.7, 0.6, 0.5),
		},
	})

	return &world
}

func main() {
	nx := 400
	ny := 200
	numSamples := 50
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)

	world := randomScene()

	lookFrom := geom.V3(13, 2, 3)
	lookAt := geom.V3(0, 0, 0)
	distToFocus := 10.0
	aperture := 0.1
	cam := scene.MakeCamera(lookFrom, lookAt, geom.V3(0, 1, 0), 20, float64(nx)/float64(ny), aperture, distToFocus)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := geom.V3(0, 0, 0)
			for s := 0; s < numSamples; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				col = col.Add(color(r, world, 0))
			}
			col = col.Div(float64(numSamples))
			// gamma correction
			col.X = math.Sqrt(col.X)
			col.Y = math.Sqrt(col.Y)
			col.Z = math.Sqrt(col.Z)
			ir := int32(255.99 * col.R())
			ig := int32(255.99 * col.G())
			ib := int32(255.99 * col.B())
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
