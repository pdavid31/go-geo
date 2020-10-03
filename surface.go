package geo

// Surface interface describes functions on the Surface object type
type Surface interface {
	// Returns the area of this Surface, as measured in the spatial reference system of this Surface
	Area() float64
	// Returns the mathematical centroid for this Surface as a Point
	Centroid() Point
	// Returns A Point guaranteed to be on this Surface
	PointOnSurface() Point
}
