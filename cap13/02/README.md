Exercise description:

1. Create a function that receive a `int` variadic parameter and returns the sum of all `int`s receiveds
1. Pass a `int` slice as arguments of this function
1. Create other function that receives a slice of `int` and return the sum of all elements of the slice
1. Pass a `int` slice as argument of this function

> Code:
```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}

	x := variadic(5, 6, 7, 8)
	y := variadic(slice...)
	z := noVariadic(slice)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}

func variadic(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}

func noVariadic(x []int) int {
	return variadic(x...)
}

```

> Output:
```console
x: 26
y: 10
z: 10
```