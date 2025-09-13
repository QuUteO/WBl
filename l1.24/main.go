package main

import (
	"fmt"
	"math"
)

type Pointer struct {
	x float64
	y float64
}

func NewPointer(x float64, y float64) *Pointer {
	return &Pointer{x: x, y: y}
}

func (p *Pointer) Distance(q *Pointer) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPointer(0, 0)
	point2 := NewPointer(3.0, 4.0)
	fmt.Println(point2.Distance(point1))
}
