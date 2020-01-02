package scene

import (
	"reflect"
	"testing"

	"github.com/jaydsteele/go_raytrace/geom"
)

func TestHitableList_Add(t *testing.T) {
	list := HitableList{}
	if len(list.items) != 0 {
		t.Errorf("HitableList should have length 0")
	}
	list.Add(Sphere{Center: geom.V3(0, 0, 0), Radius: 1})
	if len(list.items) != 1 {
		t.Errorf("HitableList should have length 1")
	}
	if !reflect.DeepEqual(list.items[0], Sphere{Center: geom.V3(0, 0, 0), Radius: 1}) {
		t.Errorf("Unexpected value in list")
	}
	list.Add(Sphere{Center: geom.V3(1, 2, 3), Radius: 99})
	if len(list.items) != 2 {
		t.Errorf("HitableList should have length 2")
	}
	if !reflect.DeepEqual(list.items[0], Sphere{Center: geom.V3(0, 0, 0), Radius: 1}) {
		t.Errorf("Unexpected value in list")
	}
	if !reflect.DeepEqual(list.items[1], Sphere{Center: geom.V3(1, 2, 3), Radius: 99}) {
		t.Errorf("Unexpected value in list")
	}

}
