package main

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i>0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1,2}
	q := Point{3,4}
	r := Point{1,5}
	path := Path{p,q,r}
	fmt.Println(p.Distance(q))
	fmt.Println(path.Distance())
}
