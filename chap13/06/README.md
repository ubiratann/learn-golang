Exercise description:

1. Create and use an anonymous function

> Code:
```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function")
	}()
}

```

> Output:
```console
Anonymous function
```