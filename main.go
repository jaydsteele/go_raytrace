package main

import (
	"fmt"

	"github.com/jaydsteele/go_raytrace/geom"
)

func main() {
	nx := 200
	ny := 100
	fmt.Printf("P3\n%d %d\n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			v := geom.Vec3{
				X: float64(i) / float64(nx),
				Y: float64(j) / float64(ny),
				Z: 0.2,
			}
			ir := int32(255.99 * v.R())
			ig := int32(255.99 * v.G())
			ib := int32(255.99 * v.B())
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
