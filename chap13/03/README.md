Exercise description:

1. Use a defer statement in a way that demonstrates that its execution only occurs at the end of the context to which it belongs.

> Code:
```go
package main

import "fmt"

func main() {

	fmt.Println("Executing with defer")
	withDefer()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func withDefer() {
	first()
	defer second()
	withoutDefer()
}

func withoutDefer() {
	first()
	second()
}

```

> Output:
```console
first
first
second
second
```