package gopoints

// A simple set implementation as Go does not provide one. Uses map keys as the storage medium
// as they do not repeat and are unordered. The value used is a blank struct as it causes little
// to no overead due to its 0-byte size.

type PointSet struct {
	// The actual data-storage part
	data map[Point]struct{}
}

func (set *PointSet) Init() {
	// Only if the set is not already initialised in order to avoid accidentally removing data
	if set.data == nil {
		set.data = make(map[Point]struct{})
	}
}

func (set *PointSet) Add(point Point) {
	// Automatically init the set if it's not initialised yet (check handled by Init())
	set.Init()

	set.data[point] = struct{}{}
}

func (set *PointSet) AddArray(pointArray []Point) {
	// Automatically init the set if it's not initialised yet (check handled by Init())
	set.Init()

	for _, point := range pointArray {
		set.Add(point)
	}
}

func (set *PointSet) Remove(point Point) {
	// Automatically init the set if it's not initialised yet (check handled by Init())
	set.Init()

	delete(set.data, point)
}

func (set *PointSet) RemoveArray(pointArray []Point) {
	// Automatically init the set if it's not initialised yet (check handled by Init())
	set.Init()

	for _, point := range pointArray {
		set.Remove(point)
	}
}

func (set PointSet) CheckFor(point Point) bool {
	// Automatically init the set if it's not initialised yet (check handled by Init())
	set.Init()

	_, ok := set.data[point]
	return ok
}

func (set PointSet) AsArray() []Point {
	// Automatically init the set if it's not initialised yet (check handled by Init())
	set.Init()

	// Convert the set to an array for easy access
	arr := []Point{}
	for point := range set.data {
		arr = append(arr, point)
	}
	return arr
}
