package geo

import "fmt"

func pointToString(p Point) string {
	x := fmt.Sprintf("%f", p.x)
	y := fmt.Sprintf("%f", p.y)

	rep := "(" + x + " " + y

	if p.Is3D() {
		z := fmt.Sprintf("%f", p.z)
		rep += " " + z
	}

	if p.m != NullValue {
		m := fmt.Sprintf("%f", p.m)
		rep += " " + m
	}

	rep += ")"

	return rep
}
