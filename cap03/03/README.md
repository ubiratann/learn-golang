Exercise description:

> Using the solution from the previous exercise:

1. In package-level scope, assign the following values to the variables:
    - for "x" assign 42
    - for "y" assign "James Bond"
    - for "z" set true
1. In the main function:
    - Use fmt.Sprintf to assign all these values to a single variable. Make this string type assignment to a variable named "s" using the short declaration operator.
    - Show the variable "s" value.
    
> Code :
```go
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\n%v\n%v", x, z, y)
	fmt.Println(s)
}

```

> Output:
```console
42
true
James Bond
```