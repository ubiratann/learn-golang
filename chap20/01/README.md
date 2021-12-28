Exercise description:

1. In addition to the main goroutine, create two other goroutines.
1. Each additional goroutine must print separately.
1. Use waitgroups to make your goroutines terminate before the program ends.

> Code:
```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func1()
	go func2()

	wg.Wait()
}

func func1() {
	fmt.Println("Func 1")
	wg.Done()
}

func func2() {
	fmt.Println("Func 2")
	wg.Done()
}

```

> Output:
```console
Func 2
Func 1
```