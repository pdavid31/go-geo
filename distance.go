package geo

import (
	"math"
)

// MinMax searches for minimum an maximum value in a float64 slice
func MinMax(a []float64) (float64, float64) {
	max := a[0]
	min := a[0]

	for _, v := range a {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}

	return min, max
}

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

// see: geos::algorithm::Distance::segmentToSegment
func distanceSegmentToSegment(A, B, C, D Point) float64 {
	// Check for zero-length segments
	if A.Equals(B) {
		return distancePointToSegment(A, C, D)
	}
	if C.Equals(D) {
		return distancePointToSegment(D, A, B)
	}

	/* AB and CD are line segments */
	/*
	   From comp.graphics.algo
	   Solving the above for r and s yields
	       (Ay-Cy)(Dx-Cx)-(Ax-Cx)(Dy-Cy)
	   r = ----------------------------- (eqn 1)
	       (Bx-Ax)(Dy-Cy)-(By-Ay)(Dx-Cx)
	       (Ay-Cy)(Bx-Ax)-(Ax-Cx)(By-Ay)
	   s = ----------------------------- (eqn 2)
	       (Bx-Ax)(Dy-Cy)-(By-Ay)(Dx-Cx)
	   Let P be the position vector of the intersection point, then
	       P=A+r(B-A) or
	       Px=Ax+r(Bx-Ax)
	       Py=Ay+r(By-Ay)
	   By examining the values of r & s, you can also determine some other
	   limiting conditions:
	   If 0<=r<=1 & 0<=s<=1, intersection exists;
	   If r<0 or r>1 or s<0 or s>1, line segments do not intersect;
	   If the denominator in eqn 1 is zero, AB & CD are parallel;
	   If the numerator in eqn 1 is also zero, AB & CD are collinear.
	*/

	noIntersection := false

	denom := (B.Lat()-A.Lat())*(D.Lon()-C.Lon()) - (B.Lon()-A.Lon())*(D.Lat()-C.Lat())

	if denom == 0 {
		noIntersection = true
	} else {
		rNum := (A.Lon()-C.Lon())*(D.Lat()-C.Lat()) - (A.Lat()-C.Lat())*(D.Lon()-C.Lon())
		sNum := (A.Lon()-C.Lon())*(B.Lat()-A.Lat()) - (A.Lat()-C.Lat())*(B.Lon()-A.Lon())

		s := sNum / denom
		r := rNum / denom

		if r < 0 || r > 1 || s < 0 || s > 1 {
			noIntersection = true
		}
	}

	if noIntersection {
		min, _ := MinMax([]float64{
			distancePointToSegment(A, C, D),
			distancePointToSegment(B, C, D),
			distancePointToSegment(C, A, B),
			distancePointToSegment(D, A, B),
		})

		return min
	}

	return 0.0
}
