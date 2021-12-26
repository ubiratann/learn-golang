Exercise description:

1. Create a type. The underlying type must be int.
Create a variable for this type, with the identifier "x", using the var keyword.

1.	In the main function:
	- Show the value of the variable "x"
	- Show the type of the variable "x"
	- Assign 42 to variable "x" using the operator "="
	- Show the value of the variable "x"



> Code :
```go
package main

import "fmt"

type ubiratan int

var x ubiratan

func main() {

	fmt.Printf("value = %v\ntype = %T\n", x, x)
	x = 42
	fmt.Println(x)
}

```

> Output:
```console
value = 0
type = main.ubiratan
42
```