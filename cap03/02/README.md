Exercise description:

1. Use `var` to declare three variables. They must be scoped at the package level. Do not assign values to these variables. Use the following identifiers and types for those.
    - "x" identifyer must be int
    - "y" identifyer must be string
    - "z" identifyer must be bool

    Then show these values on main function.

> Code :
```go
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}

```

> Output:
```console
0

false
```