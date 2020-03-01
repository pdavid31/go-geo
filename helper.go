package geo

import (
	"math"
)

// see geos c++ library

// see: geos::algorithm::Distance::pointToSegment
func distancePointToSegment(p, A, B Point) float64 {
	// if start == end, use pt distance
	if A.Equals(B) {
		return p.Distance(A)
	}

	/*
	   otherwise use comp.graphics.algorithms method:
	   (1)
	                   AC dot AB
	               r = ---------
	                   ||AB||^2

	   r has the following meaning:
	   r=0 P = A
	   r=1 P = B
	   r<0 P is on the backward extension of AB
	   r>1 P is on the forward extension of AB
	   0<r<1 P is interior to AB
	*/

	// TODO: check 3d space
	r := ((p.Lat()-A.Lat())*(B.Lat()-A.Lat()) + (p.Lon()-A.Lon())*(B.Lon()-A.Lon())) / ((B.Lat()-A.Lat())*(B.Lat()-A.Lat()) + (B.Lon()-A.Lon())*(B.Lon()-A.Lon()))

	if r <= 0 {
		return p.Distance(A)
	}

	if r >= 1 {
		return p.Distance(B)
	}

	/*
	   (2)
	           (Ay-Cy)(Bx-Ax)-(Ax-Cx)(By-Ay)
	       s = -----------------------------
	                       L^2

	   Then the distance from C to P = |s|*L.
	*/

	s := ((A.Lon()-p.Lon())*(B.Lat()-A.Lat()) - (A.Lat()-p.Lat())*(B.Lon()-A.Lon())) / ((B.Lat()-A.Lat())*(B.Lat()-A.Lat()) + (B.Lon()-A.Lon())*(B.Lon()-A.Lon()))

	return math.Abs(s) * math.Sqrt((B.Lat()-A.Lat())*(B.Lat()-A.Lat())+(B.Lon()-A.Lon())*(B.Lon()-A.Lon()))
}
