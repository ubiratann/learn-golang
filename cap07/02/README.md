Exercise description:

    1. Print the unicode point of all uppercase letters, 3 times each one.
> Code:
```go
package main

import "fmt"

func main() {
	for x := 65; x < 91; x++ {

		fmt.Println("\n", x)
		for y := 0; y < 3; y++ {
			fmt.Printf("%#U ", x)
		}
		fmt.Println("")
	}
}

```

> Output:
```console

 65
U+0041 'A' U+0041 'A' U+0041 'A' 

 66
U+0042 'B' U+0042 'B' U+0042 'B' 

 67
U+0043 'C' U+0043 'C' U+0043 'C'

...

 90
U+005A 'Z' U+005A 'Z' U+005A 'Z' 

 97
U+0061 'a' U+0061 'a' U+0061 'a' 

 98
U+0062 'b' U+0062 'b' U+0062 'b' 

 99
U+0063 'c' U+0063 'c' U+0063 'c' 

...

 122
U+007A 'z' U+007A 'z' U+007A 'z' 

```