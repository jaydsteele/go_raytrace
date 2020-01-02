package geom

// HitableList is a list of Hitable objects
type HitableList struct {
	items []Hitable
}

// Add a Hitable object to the HitableList
func (list *HitableList) Add(h Hitable) {
	if list.items == nil {
		list.items = make([]Hitable, 0)
	}
	list.items = append(list.items, h)
}

// Hit an item in the HitableList
func (list *HitableList) Hit(r Ray, tMin, tMax float64, rec *HitRecord) bool {
	tempRecord := HitRecord{}
	hitAnything := false
	closestSoFar := tMax
	for i := 0; i < len(list.items); i++ {
		if list.items[i].Hit(r, tMin, closestSoFar, &tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.t
			rec.Normal = tempRecord.Normal
			rec.p = tempRecord.p
			rec.t = tempRecord.t
		}
	}
	return hitAnything
}
