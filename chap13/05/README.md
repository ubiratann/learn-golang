Exercise description:

1.Create a "square" type
1. Create a "circle" type
1. Create an "area" method for each type that calculates and returns the figure's area
1. Create a type "figure" that defines as an interface any type that has the method "area"
1. Create an "info" function that takes a "picture" type and returns the picture area
1. Create a value of type "square"
1. Create a value of type "circle"
1. Use the "info" function to demonstrate the "square" area
1. Use the "info" function to demonstrate the "circle" area

> Code:
```go
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

```

> Output:
```console
Square area = 144
Circle area = 31.41592653589793
```