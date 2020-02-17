package geo

type Polygon []LinearRing

/* CONSTRUCTOR */
func NewPolygon(linearrings ...LinearRing) Polygon {
	poly := Polygon{}

	for _, lr := range linearrings {
		poly = append(poly, lr)
	}

	return poly
}

/* GEOMETRY */
func (p Polygon) Lat() float64 {
	panic("implement me")
}

func (p Polygon) Lon() float64 {
	panic("implement me")
}

func (p Polygon) Z() float64 {
	panic("implement me")
}

func (p Polygon) GeometryType() string {
	return "Polygon"
}

func (p Polygon) SRID() int {
	return 4326
}

func (p Polygon) Envelope() Geometry {
	panic("implement me")
}

func (p Polygon) AsText() string {
	panic("implement me")
}

func (p Polygon) IsEmpty() bool {
	return len(p) == 0
}

func (p Polygon) Is3D() bool {
	for _, l := range p {
		if !l.Is3D() {
			return false
		}
	}

	return true
}

func (p Polygon) Equals(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Disjoint(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Intersects(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Touches(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Crosses(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Within(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Contains(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Overlaps(another Geometry) bool {
	panic("implement me")
}

func (p Polygon) Distance(another Geometry) float64 {
	panic("implement me")
}

func (p Polygon) Buffer(distance float64) Geometry {
	panic("implement me")
}

func (p Polygon) ConvexHull() Geometry {
	panic("implement me")
}

func (p Polygon) Intersection(another Geometry) Geometry {
	panic("implement me")
}

func (p Polygon) Union(another Geometry) Geometry {
	panic("implement me")
}

func (p Polygon) Difference(another Geometry) Geometry {
	panic("implement me")
}

/* SURFACE */
func (p Polygon) Area() float64 {
	panic("implement me")
}

func (p Polygon) Centroid() Point {
	panic("implement me")
}

func (p Polygon) PointOnSurface() Point {
	panic("implement me")
}

/* POLYGON */
func (p Polygon) ExteriorRing() LinearRing {
	panic("implement me")
}

func (p Polygon) NumInteriorRing() int {
	panic("implement me")
}

func (p Polygon) InteriorRingN(n int) LinearRing {
	panic("implement me")
}
