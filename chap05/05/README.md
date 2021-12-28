Exercise description:

1. Create a variable of string type using a raw string literal.
2. Show this string.

> Code:
```go
package main

import "fmt"

func main() {
	raw := `string 
			in plain
					text`
	fmt.Printf("%s\n", raw)
}

```

> Output:
```console
string 
                        in plain
                                        text
```