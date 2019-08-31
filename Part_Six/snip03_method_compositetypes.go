package main

import (
	"fmt"
	"image/color"
	"math"
)

type Pointt struct {
	X, Y float64
}
type ColoredPoint struct {
	Pointt
	Color color.RGBA
}

func main() {
	var cp ColoredPoint

	cp.X, cp.Y = 0, 0
	cp.Color.R, cp.Color.G, cp.Color.B, cp.Color.A = 255, 255, 255, 0

	red := color.RGBA{255, 0, 0, 255}
	var cp2 = ColoredPoint{Pointt{0, 4}, red}

	// ColoredPoint	now shares the same method as 'Pointt' does
	// ColoredPoint still cannot be passed into the method directly (not the same type!)
	fmt.Println(cp2.Distance(cp.Pointt))

	fmt.Println(cp)
	fmt.Println(cp2)

	invokeWithoutAReceiver()
	invokeMethodJustLikeFunction()
}

func invokeWithoutAReceiver() {
	origin := Pointt{0, 0}
	q := Pointt{1, 1}

	distanceFromOrigin := origin.Distance
	fmt.Println(distanceFromOrigin(q))
}

func invokeMethodJustLikeFunction() {
	p := Pointt{1, 2}
	q := Pointt{4, 6}

	distance := Pointt.Distance
	fmt.Println(distance(p, q)) // type: func(Point, Point) float64

	scale := (*Pointt).ScaleBy
	scale(&p, 10)
	fmt.Println(p)
}

func (p Pointt) Distance(q Pointt) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + +dY*dY)
}

func (p *Pointt) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
