package geom

// lerp returns the linear interpolation between a and b by amount t.
// The amount t is usually a value between 0 and 1. If t=0 a will be returned;
// if t=1 b will be returned.
func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

const epsilon = 1e-5
