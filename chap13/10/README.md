Exercise description:

1. Create a function that returns another function, where this other function makes use of a variable other than its scope. 

> Code:
```go
package main

import "fmt"

func example() func() int {
	timesHere := 0

	return func() int {
		timesHere++
		return timesHere
	}
}

func main() {
	x := example()
	y := example()
	z := x
	x()
	x()
	x()
	fmt.Println("x:", x())
	fmt.Println("y:", y())
	fmt.Println("z:", z())

}

```

> Output:
```console
x: 4
y: 1
z: 5
```