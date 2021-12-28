package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type figure interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * 2 * math.Pi
}

func (s square) area() float64 {
	return s.side * s.side
}

func info(fig figure) float64 {
	return fig.area()
}

func main() {
	sq := square{12}
	cc := circle{5}

	fmt.Println("Square area =", info(sq))
	fmt.Println("Circle area =", info(cc))
}
