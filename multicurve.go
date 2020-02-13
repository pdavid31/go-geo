package geo

type MultiCurve interface {
	// Returns the Length of this MultiCurve which is equal to the sum of the lengths of the element Curves
	Length() float64
	// if this MultiCurve is closed [StartPoint ( ) = EndPoint ( ) for each Curve in this MultiCurve]
	IsClosed() bool
}
