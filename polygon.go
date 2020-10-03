package geo

import (
	"reflect"
	"strings"
)

// Polygon type
type Polygon []LinearRing

/* CONSTRUCTOR */

// NewPolygon Polygon constructor
func NewPolygon(linearrings ...LinearRing) Polygon {
	var poly Polygon

	for _, lr := range linearrings {
		poly = append(poly, lr)
	}

	return poly
}

/* GEOMETRY */

// Lat function returns the latitude of the Geometry
func (p Polygon) Lat() float64 {
	panic("implement me")
}

// Lon function returns the longitude of the Geometry
func (p Polygon) Lon() float64 {
	panic("implement me")
}

// Z function returns the altitude of the Geometry
func (p Polygon) Z() float64 {
	panic("implement me")
}

// GeometryType returns GeometryType as a string
func (p Polygon) GeometryType() string {
	return "Polygon"
}

// SRID returns the SRID as an integer
func (p Polygon) SRID() int {
	return 4326
}

// Envelope returns the Envelope of the geometry
func (p Polygon) Envelope() Geometry {
	panic("implement me")
}

// AsText returns the string representation of the geometry
func (p Polygon) AsText() string {
	// examples:
	// POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10))
	// POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10), (20 30, 35 35, 30 20, 20 30))
	rep := strings.ToUpper(p.GeometryType()) + " "

	rep += p.ToString()

	return rep
}

// ToString returns the internal string representation of the geometry
func (p Polygon) ToString() string {
	rep := "("

	lastRing := len(p) - 1
	for i, lr := range p {
		rep += lr.ToString()

		if i != lastRing {
			rep += ", "
		}
	}

	rep += ")"
	return rep
}

// IsEmpty returns a bool indicating if the geometry is empty
func (p Polygon) IsEmpty() bool {
	return len(p) == 0
}

// Is3D returns a bool indicating if the geometry is three dimensional
func (p Polygon) Is3D() bool {
	for _, l := range p {
		if !l.Is3D() {
			return false
		}
	}

	return true
}

// Equals returns true if the geometry is equal to a given other
func (p Polygon) Equals(another Geometry) bool {
	// TODO: add case of unordered slices
	// check if given geometry is a point
	switch another.(type) {
	case Polygon:
		return reflect.DeepEqual(p, another)
	default:
		return false
	}
}

// Disjoint returns true if the geometry is disjoint to a given other
func (p Polygon) Disjoint(another Geometry) bool {
	panic("implement me")
}

// Intersects returns true if the geometry intersects another given geometry
func (p Polygon) Intersects(another Geometry) bool {
	panic("implement me")
}

// Touches returns true if the geometry touches another given geometry
func (p Polygon) Touches(another Geometry) bool {
	panic("implement me")
}

// Crosses returns true if the geometry crosses another given geometry
func (p Polygon) Crosses(another Geometry) bool {
	panic("implement me")
}

// Within returns true if the geometry is within another given geometry
func (p Polygon) Within(another Geometry) bool {
	panic("implement me")
}

// Contains returns true if the geometry contains another given geometry
func (p Polygon) Contains(another Geometry) bool {
	panic("implement me")
}

// Overlaps returns true if the geometry overlaps another given geometry
func (p Polygon) Overlaps(another Geometry) bool {
	panic("implement me")
}

// Distance calculates the distane to another given geometry
func (p Polygon) Distance(another Geometry) float64 {
	panic("implement me")
}

// Buffer returns the Buffer with the given distance around the geometry
func (p Polygon) Buffer(distance float64) Geometry {
	panic("implement me")
}

// ConvexHull returns the ConvexHull containing the geometry
func (p Polygon) ConvexHull() Geometry {
	panic("implement me")
}

// Intersection returns the Intersection of the geometry and another given geometry
func (p Polygon) Intersection(another Geometry) Geometry {
	panic("implement me")
}

// Union returns the Union of the geometry and another given geometry
func (p Polygon) Union(another Geometry) Geometry {
	panic("implement me")
}

// Difference returns the Difference of the geometry and another given geometry
func (p Polygon) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* SURFACE */

// Area returns the area of the surface
func (p Polygon) Area() float64 {
	panic("implement me")
}

// Centroid returns the centroid of the surface
func (p Polygon) Centroid() Point {
	panic("implement me")
}

// PointOnSurface returns a random point of the surface
func (p Polygon) PointOnSurface() Point {
	panic("implement me")
}

/* POLYGON */

// ExteriorRing returns the exterior ring of the polygon
func (p Polygon) ExteriorRing() LinearRing {
	panic("implement me")
}

// NumInteriorRing returns the amount of holes in the polygon
func (p Polygon) NumInteriorRing() int {
	panic("implement me")
}

// InteriorRingN return the interior ring n
func (p Polygon) InteriorRingN(n int) LinearRing {
	panic("implement me")
}

/* INTERFACE GUARD */
var (
	_ Geometry = (*Polygon)(nil)
	_ Surface  = (*Polygon)(nil)
)
