package geo

// MultiSurface interface describes functions on the MultiSurface object type
type MultiSurface interface {
	// Returns the area of this MultiSurface, as measured in the spatial reference system of this MultiSurface
	Area() float64
	// Returns the mathematical centroid for this MultiSurface
	Centroid() Point
	// Returns Point guaranteed to be on this MultiSurface
	PointOnSurface() Point
}
