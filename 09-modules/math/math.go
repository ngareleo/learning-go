package math

import "math"

type Point struct {
	x, y   int
	parent *Point
	child  *Point
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p *Point) Attach(to *Point) {
	p.child = to
	to.parent = p
}

func (p *Point) Distance() float32 {
	xDiff := math.Pow(float64(p.x) - float64(p.child.x), 2)
	yDiff := math.Pow(float64(p.y) - float64(p.child.y) , 2)
	return float32(math.Sqrt(xDiff + yDiff))
}