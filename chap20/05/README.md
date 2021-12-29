Exercise description:

1. Use atomic to fix the exercise 03

> Code:
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var gs = runtime.Gosched

func main() {

	maxGoRoutine := 1000
	var iterator uint64 = 0

	wg.Add(maxGoRoutine)
	for count := 0; count < maxGoRoutine; count++ {
		go func() {
			atomic.AddUint64(&iterator, 1)
			gs()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("\n\nFinal iterator: %v\nGoRoutines: %v\n\n\n", iterator, maxGoRoutine)
}

```

> Output:
```console


Final iterator: 1000
GoRoutines: 1000


```