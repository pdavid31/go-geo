package geo

type Curve interface {
	// Returns the length of this Curve in its associated spatial reference
	Length() float64
	// Returns the start Point of this Curve
	StartPoint() Point
	// Returns the end Point of this Curve
	EndPoint() Point
	// Returns true if this Curve is closed [StartPoint ( ) = EndPoint ( )]
	IsClosed() bool
	// TODO: add IsRing()
}
