Exercise description: 

1. Using the short declaration operator (:=), assign the following values with identifiers "x", "y" and "z":
    1. 42
    1. "James Bond"
    1. true
    

2. Then, show the values with:
    1. One single print declaration
    1. Multiple prints declarations


> Code :
```go
package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Print("x=", x, " y=", y, " z=", z)
	fmt.Println("\nx= ", x)
	fmt.Println("y= ", y)
	fmt.Println("z= ", z)

}

```

> Output:
```console
x=42 y=James Bond z=true
x=  42
y=  James Bond
z=  true
```