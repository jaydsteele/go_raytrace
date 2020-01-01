package geom

import "testing"

func TestVec3_R(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"R", fields{99.0, 0.0, 0.0}, 99.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.R(); got != tt.want {
				t.Errorf("Vec3.R() = %v, want %v", got, tt.want)
			}
		})
	}
}
