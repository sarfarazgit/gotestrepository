package servicepackage

type Circle struct {
	Radius float64
}

func (c *Circle) AreaofCircle() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Shape interface {
	Area(a, b float64) float64
}

func Area(a, b float64) float64 {
	return a * b
}
