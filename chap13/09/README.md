Exercise description:

1. Callback: pass a function as argument fo another

> Code:
```go
package main

import "fmt"

func greet(x func() string) string {
	return x()
}

func sayHello() string {
	return "Hello"
}

func sayHi() string {
	return "Hi"
}

func main() {
	typeOf := 2
	compliment := ""

	switch typeOf {
	case 1:
		compliment = greet(sayHello)
	case 2:
		compliment = greet(sayHi)
	default:
		compliment = "..."
	}

	fmt.Println(compliment)
}

```

> Output:
```console
Hi
```