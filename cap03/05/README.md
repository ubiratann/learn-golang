Exercise description:

> Using the solution from the previous exercise:

1. In the package-level scope, using the var keyword, create a variable with the identifier "y". The type of this variable must be the underlying type of the type you created in the previous exercise.

2. In the main function:
    - This should already be done:
        - [x] Show the value of the variable "x"
        - [x] Demonstrate the type of the variable "x"
        - [x] Assign 42 to variable "x" using the operator "="
        - [x] Show the value of the variable "x"
    - Now do too:
        - Use cast to type the value of the variable "x" into its underlying type and, using the "=" operator, assign the value of "x" to "y"
        - Show the value of "y"
        - Show the type of "y"


> Code:
```go
package main

import "fmt"

type ubiratan int

var x ubiratan
var y int

func main() {
	fmt.Printf("%v\t%v\n", x, y)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%v %T\n", y, y)
}

```

> Output:
```console
0       0
42
42 int
```